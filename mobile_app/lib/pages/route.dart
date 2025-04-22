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
  var distance = 0.0;
  var numOfGeoPoints = 0;
  var rocketPoints = 0;

  @override
  void initState() {
    super.initState();
    _setInitialPosition();
    _mapController = MapController(
      initPosition: GeoPoint(latitude: 48.61313, longitude: 9.45881),
    );
  }

  /// Here: Instead of the hardcoded points use the points from the list _routePoints
  Future<void> _setInitialPosition() async {
    try {
      setState(() {
        _initPosition = GeoPoint(latitude: 48.61313, longitude: 9.45881);
        _routePoints = [
          _initPosition!,
          GeoPoint(latitude: 48.6156, longitude: 9.45984),
          GeoPoint(latitude: 48.61651, longitude: 9.4549),
        ];
      });
    } catch (e) {
      // Handle the error accordingly
      ///Instead of print maybe don't show map and just show earned points and so on
      print(e);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text(widget.title)),
      body: Column(
        children: [
          SizedBox(
            height: 500,
            child:
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
                        ///Here draw the map and get the distance from point to point
                        for(var point in _routePoints){
                          /**
                           * Checks if all points have been already used
                           */
                          if(numOfGeoPoints == _routePoints.length){
                            break;
                          }
                          numOfGeoPoints++;
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
                            distance += roadInfo.distance ?? 0;
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
                            /**
                             * Adds the distance of the road to the total distance
                             */
                            distance += roadInfo.distance ?? 0;
                            setState(() {
                              distance = double.parse(distance.toStringAsFixed(2));
                            });
                          }
                        }
                        // Logik für die Karte
                      },
                      onGeoPointClicked: (geoPoint) {
                        // Behandlung des Klicks auf einen Geopunkt
                      },
                    ),
          ),
          Container(
            padding: const EdgeInsets.all(16.0),
            color: Colors.blueAccent,
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Text(
                      "Distanz: ${distance.toStringAsFixed(2)} km",
                      style: TextStyle(
                        fontSize: 16,
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                    Text(
                      "Dauer: 00:45:30",
                      style: TextStyle(
                        fontSize: 16,
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                  ],
                ),
                SizedBox(height: 8),
                Center(
                  child: Text(
                    "Punkte: ${distance.floor()}",
                    style: TextStyle(fontSize: 16, fontWeight: FontWeight.bold),
                  ),
                ),
                /**
                 * Here: Possible to add the challenges that are completed
                 */
                SizedBox(height: 16),
                Center(
                  child: ElevatedButton(
                    onPressed: () {
                      /**
                       * Here: Navigate back to the main app screen
                       */
                  },
                    child: Text("Back To App"),
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}

void main() {
  runApp(
    MaterialApp(
      title: 'Route Page',
      theme: ThemeData(primarySwatch: Colors.blue),
      home: RoutePage(title: 'Completed Run'),
    ),
  );
}
