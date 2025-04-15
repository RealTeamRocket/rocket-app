import 'package:flutter/material.dart';
import 'package:flutter_osm_plugin/flutter_osm_plugin.dart';
import 'package:geolocator/geolocator.dart';

class RoutePage extends StatefulWidget {
  const RoutePage({super.key, required this.title});

  final String title;

  @override
  State<RoutePage> createState() => _RoutePageState();
}

Future<Position> _determinePosition() async {
  bool serviceEnabled;
  LocationPermission locationPermission;
  serviceEnabled = await Geolocator.isLocationServiceEnabled();
  if (!serviceEnabled) {
    return Future.error('Location services are disabled.');
  }
  locationPermission = await Geolocator.checkPermission();
  if (locationPermission == LocationPermission.denied) {
    locationPermission = await Geolocator.requestPermission();
    if (locationPermission == LocationPermission.denied) {
      return Future.error('Location permissions are denied');
    }
  }
  if (locationPermission == LocationPermission.deniedForever) {
    return Future.error(
      'Location permissions are permanently denied, we cannot request permissions.',
    );
  }
  return await Geolocator.getCurrentPosition();
}

class _RoutePageState extends State<RoutePage> {
  GeoPoint? _initPosition;
  List<GeoPoint> _routePoints = [];
  late MapController _mapController;
  var _distance = 0.0;
  var numOfPoints = 0;

  @override
  void initState() {
    super.initState();
    _setInitialPosition();
    _mapController = MapController(
      initPosition: GeoPoint(latitude: 48.61313, longitude: 9.45881),
    );
  }

  Future<void> _setInitialPosition() async {
    try {
      Position position = await _determinePosition();
      setState(() {
        _initPosition = GeoPoint(latitude: 48.61313, longitude: 9.45881);
        _routePoints = [
          _initPosition!,
          GeoPoint(latitude: 48.6156, longitude: 9.45984),
          GeoPoint(latitude: 48.61651, longitude: 9.4549), // Beispielzielpunkt
        ];
      });
    } catch (e) {
      // Handle the error accordingly
      print(e);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text(widget.title)),
      body:
          _initPosition == null
              ? Center(child: CircularProgressIndicator())
              : OSMFlutter(
                controller: _mapController,
                osmOption: OSMOption(
                  zoomOption: ZoomOption(
                    minZoomLevel: 5,
                    maxZoomLevel: 18,
                    initZoom: 16,
                  ),
                  userLocationMarker: UserLocationMaker(
                    personMarker: MarkerIcon(
                      icon: Icon(
                        Icons.location_on,
                        color: Colors.red,
                        size: 48,
                      ),
                    ),
                    directionArrowMarker: MarkerIcon(
                      icon: Icon(
                        Icons.arrow_forward,
                        color: Colors.blue,
                        size: 48,
                      ),
                    ),
                  ),
                ),
                mapIsLoading: Center(child: CircularProgressIndicator()),
                onMapIsReady: (isReady) async {
                  RoadInfo roadInfo;
                  for(var point in _routePoints){
                    /**
                     * Checks if all points have been already used
                     */
                    if(numOfPoints == _routePoints.length){
                      break;
                    }
                    numOfPoints++;
                    /**
                     * Draws the road between the last and first point
                     */
                    if(_routePoints.indexOf(point) == _routePoints.length-1){
                        roadInfo = await _mapController.drawRoad(
                        point,
                        _routePoints[0],
                        roadType: RoadType.foot,
                        roadOption: RoadOption(
                          roadColor: Colors.yellow,
                          roadWidth: 8,
                        ),
                      );
                       print("Rückweg: ${roadInfo.distance?.toStringAsFixed(4)} km");
                      _distance += roadInfo.distance ?? 0;
                      /**
                       * Draws road between each other point 0->1, 1->2, ...
                       */
                    }else{
                      roadInfo = await _mapController.drawRoad(
                        point,
                        _routePoints[_routePoints.indexOf(point) + 1],
                        roadType: RoadType.foot,
                        roadOption: RoadOption(
                          roadColor: Colors.yellow,
                          roadWidth: 8,
                        ),
                      );
                      print(" ${roadInfo.distance?.toStringAsFixed(4)} km");
                      /**
                       * Adds the distance of the road to the total distance
                       */
                      _distance += roadInfo.distance ?? 0;
                    }
                  }
                  /**
                   * Here the distance will be saved into the database but there is still an inaccuracy
                   * The whole distance gets printed two times.
                   */
                  print("Gesamte Distanz: ${_distance.toStringAsFixed(4)} km");
                },
                onGeoPointClicked: (geoPoint) {
                  // Behandlung des Klicks auf einen Geopunkt, falls erforderlich
                },
              ),

    );
  }
}

void main() {
  runApp(
    MaterialApp(
      title: 'Route Page',
      theme: ThemeData(primarySwatch: Colors.blue),
      home: RoutePage(title: 'Route Page'),
    ),
  );
}
