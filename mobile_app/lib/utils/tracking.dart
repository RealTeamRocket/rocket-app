import 'package:flutter/material.dart';
import 'package:flutter_osm_plugin/flutter_osm_plugin.dart';
import 'package:geolocator/geolocator.dart';


class Tracking extends MapController{
  List<GeoPoint> _routePoints = [];
  late MapController _mapController;

  void startTracking() {
    _mapController.enableTracking();
  }

  void stopTracking() {
    _mapController.disabledTracking();
    print("Tracking gestoppt. Aufgezeichnete Punkte: $_routePoints");
  }

  void addPoint(GeoPoint point) {
    _routePoints.add(point);
    print("Punkt hinzugefügt: $point");
  }
  void getCurrentPosition() async {
    Position position = await Geolocator.getCurrentPosition();
    GeoPoint currentPoint = GeoPoint(latitude: position.latitude, longitude: position.longitude);
    addPoint(currentPoint);
  }
}


/**
 * location 8.0.0 Maybe need to use that instead of geolocator for background location
 * https://pub.dev/packages/location
 *
 * Also this class will send the List into the Route class so there the points can be used to create a route
 * How do I sent them into the database from here?
 * How do I start and stop the tracking?
 * Where will the button send the request to start and stop the tracking?
 */
