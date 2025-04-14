// ignore_for_file: no_leading_underscores_for_local_identifiers

import 'package:flutter_osm_plugin/flutter_osm_plugin.dart';
import 'package:location/location.dart';

class Tracking{
  List<GeoPoint> routePoints = [];
  final Location location = Location();

  void startTracking() async {

    /**
     * This are 3 test points to show how the route will look like
     */
    routePoints.add(GeoPoint(latitude: 48.61313, longitude: 9.45881));
    routePoints.add(GeoPoint(latitude: 48.6156, longitude: 9.45984));
    routePoints.add(GeoPoint(latitude: 48.61651, longitude: 9.4549));

    bool _serviceEnabled;
    PermissionStatus _permissionGranted;

    _serviceEnabled = await location.serviceEnabled();
    if (!_serviceEnabled) {
      _serviceEnabled = await location.requestService();
      if (!_serviceEnabled) {
        return;
      }
    }

    _permissionGranted = await location.hasPermission();

    if(_permissionGranted == PermissionStatus.denied) {
      _permissionGranted = await location.requestPermission();
      if (_permissionGranted != PermissionStatus.granted) {
        return;
      }
    }
    location.enableBackgroundMode(enable: true);
    /**
     * The following code should be in one .dart where all background activities are handled
     *
     */
    location.onLocationChanged.listen((LocationData currentLocation) {
      if (currentLocation.longitude == null || currentLocation.latitude == null) {
        return;
      }
      GeoPoint point = GeoPoint(
        latitude: currentLocation.latitude!,
        longitude: currentLocation.longitude!,
      );
      routePoints.add(point);
    });
  }
}


/**
 * location 8.0.0 Maybe need to use that instead of 'Geo-locator' for background location
 * https://pub.dev/packages/location
 *
 * Also this class will send the List into the Route class so there the points can be used to create a route
 * How do I sent them into the database from here?
 * How do I start and stop the tracking?
 * Where will the button send the request to start and stop the tracking?
 */
