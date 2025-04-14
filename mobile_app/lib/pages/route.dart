import 'package:flutter/material.dart';
import 'package:flutter_osm_plugin/flutter_osm_plugin.dart';
import 'package:geolocator/geolocator.dart';

import '../utils/tracking.dart';


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
        'Location permissions are permanently denied, we cannot request permissions.');
  }
  return await Geolocator.getCurrentPosition();
}

class _RoutePageState extends State<RoutePage> {
  GeoPoint? _initPosition;
  final Tracking _tracking = Tracking();
  List<GeoPoint> get _routePoints => _tracking.routePoints;
  late MapController _mapController;


  @override
  void initState() {
    _tracking.startTracking();
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

      });
    } catch (e) {
      // Handle the error accordingly
      print(e);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: _initPosition == null
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
          if (isReady && _routePoints.length >= 2) {
            _mapController.drawRoad(
              _routePoints.first, // Startpunkt
              _routePoints.last,  // Endpunkt
              roadType: RoadType.foot,
              intersectPoint: _routePoints.sublist(1, _routePoints.length - 1), // Zwischenpunkte
              roadOption: RoadOption(
                roadColor: Colors.yellow,
                roadWidth: 10.0,
              ),
            );
          }
        },
        onGeoPointClicked: (geoPoint) {
          // Behandlung des Klicks auf einen Geopunkt, falls erforderlich
        },
      ),
    );
  }
}

void main() {
  runApp(MaterialApp(
    title: 'Route Page',
    theme: ThemeData(
      primarySwatch: Colors.blue,
    ),
    home: RoutePage(title: 'Route Page'),
  ));
}