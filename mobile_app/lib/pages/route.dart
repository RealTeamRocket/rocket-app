import 'package:flutter/material.dart';
import 'package:flutter_osm_plugin/flutter_osm_plugin.dart';
import 'package:mobile_app/constants/color_constants.dart';
import 'dart:math';

import 'package:mobile_app/pages/pages.dart';

class RoutePage extends StatefulWidget {
  const RoutePage({
    super.key,
    required this.title,
    required this.routePoints,
    required this.elapsedTime,
  });

  final String title;
  final List<GeoPoint> routePoints;
  final String elapsedTime;

  @override
  State<RoutePage> createState() => _RoutePageState();
}

class _RoutePageState extends State<RoutePage> {
  GeoPoint? _initPosition;
  late List<GeoPoint> _routePoints;
  late MapController _mapController;
  double distance = 0.0;

  @override
  void initState() {
    super.initState();
    _routePoints = widget.routePoints;
    _initPosition = _routePoints.isNotEmpty
        ? _routePoints.first
        : GeoPoint(latitude: 48.61313, longitude: 9.45881);
    _mapController = MapController(
      initPosition: _initPosition!,
    );
    distance = _calculateTotalDistance(_routePoints);
  }

  double _calculateTotalDistance(List<GeoPoint> points) {
    if (points.length < 2) return 0.0;
    double total = 0.0;
    for (int i = 0; i < points.length - 1; i++) {
      total += _calculateDistanceBetween(points[i], points[i + 1]);
    }
    return total;
  }

  // Haversine formula for distance in km
  double _calculateDistanceBetween(GeoPoint a, GeoPoint b) {
    const double R = 6371; // Earth's radius in km
    double dLat = _deg2rad(b.latitude - a.latitude);
    double dLon = _deg2rad(b.longitude - a.longitude);
    double lat1 = _deg2rad(a.latitude);
    double lat2 = _deg2rad(b.latitude);

    double hav = sin(dLat / 2) * sin(dLat / 2) +
        sin(dLon / 2) * sin(dLon / 2) * cos(lat1) * cos(lat2);
    double c = 2 * atan2(sqrt(hav), sqrt(1 - hav));
    return R * c;
  }

  double _deg2rad(double deg) => deg * (pi / 180);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        iconTheme: IconThemeData(
          color: ColorConstants.white,
        ),
        title: Text(
          widget.title,
          style: TextStyle(
            color: ColorConstants.white,
            fontWeight: FontWeight.bold,
          ),
        ),
        centerTitle: true,
        backgroundColor: ColorConstants.secoundaryColor,
      ),
      backgroundColor: ColorConstants.secoundaryColor,
      body: Column(
        children: [
          SizedBox(
            height: 500,
            child: _initPosition == null || _routePoints.isEmpty
                ? Center(
                  child: Text(
                      "No route available",
                      style: TextStyle(
                        color: ColorConstants.white,
                        fontWeight: FontWeight.bold,
                      ),
                    )
                  )
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
                      if (_routePoints.length > 1) {
                        for (int i = 0; i < _routePoints.length - 1; i++) {
                          await _mapController.drawRoad(
                            _routePoints[i],
                            _routePoints[i + 1],
                            roadType: RoadType.foot,
                            roadOption: RoadOption(
                              roadColor: Colors.red,
                              roadWidth: 8,
                            ),
                          );
                        }
                      }
                    },
                    onGeoPointClicked: (geoPoint) {},
                  ),
          ),
          Container(
            padding: const EdgeInsets.all(16.0),
            color: ColorConstants.secoundaryColor,
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Text(
                      "Distance: ${distance.toStringAsFixed(2)} km",
                      style: TextStyle(
                        fontSize: 16,
                        fontWeight: FontWeight.bold,
                        color: ColorConstants.white
                      ),
                    ),
                    Text(
                      "Time: ${widget.elapsedTime}",
                      style: TextStyle(
                        fontSize: 16,
                        fontWeight: FontWeight.bold,
                        color: ColorConstants.white
                      ),
                    ),
                  ],
                ),
                SizedBox(height: 16),
                Center(
                  child: ElevatedButton(
                    onPressed: () {
                      Navigator.pushAndRemoveUntil(
                        context,
                        MaterialPageRoute(
                          builder: (context) => AppNavigator(title: 'Rocket App', initialIndex: 2),
                        ),
                        (route) => false,
                      );
                    },
                    style: ElevatedButton.styleFrom(
                      backgroundColor: ColorConstants.purpleColor,
                      foregroundColor: ColorConstants.white,
                    ),
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
