import 'dart:async';
import 'dart:math';
import 'package:flutter/material.dart';
import 'package:geolocator/geolocator.dart';
import 'package:flutter_osm_plugin/flutter_osm_plugin.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:mobile_app/pages/pages.dart';

import '../utils/backend_api/tracking_api.dart';
import 'route.dart';

class TrackingPage extends StatefulWidget {
  const TrackingPage({super.key, required this.title});

  final String title;

  @override
  State<TrackingPage> createState() => _TrackingPageState();
}

class _TrackingPageState extends State<TrackingPage> {
  bool _isTracking = false;
  bool _wasStarted = false;
  List<GeoPoint> _geoPoints = [];
  List<GeoPoint> _lastRoutePoints = [];
  Stopwatch _stopwatch = Stopwatch();
  late Timer _timer;
  String _formattedTime = "00:00:00";
  StreamSubscription<Position>? _positionStream;

  Future<void> _startTracking() async {
    bool serviceEnabled = await Geolocator.isLocationServiceEnabled();
    LocationPermission permission = await Geolocator.checkPermission();

    if (!serviceEnabled || permission == LocationPermission.denied) {
      permission = await Geolocator.requestPermission();
      if (permission == LocationPermission.denied ||
          permission == LocationPermission.deniedForever) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text("Standortberechtigung benötigt.")),
        );
        return;
      }
    }

    setState(() {
      _wasStarted = true;
      _isTracking = true;
      _geoPoints.clear();
      _stopwatch.reset();
      _stopwatch.start();
    });

    _timer = Timer.periodic(Duration(milliseconds: 10), (timer) {
      setState(() {
        final elapsed = _stopwatch.elapsed;
        final minutes = elapsed.inMinutes
            .remainder(60)
            .toString()
            .padLeft(2, '0');
        final seconds = elapsed.inSeconds
            .remainder(60)
            .toString()
            .padLeft(2, '0');
        final milliseconds = (elapsed.inMilliseconds.remainder(1000) ~/ 10)
            .toString()
            .padLeft(2, '0');
        _formattedTime = "$minutes:$seconds:$milliseconds";
      });
    });

    _positionStream = Geolocator.getPositionStream(
      locationSettings: LocationSettings(
        accuracy: LocationAccuracy.high,
        distanceFilter: 5,
      ),
    ).listen((Position position) {
      setState(() {
        _geoPoints.add(
          GeoPoint(latitude: position.latitude, longitude: position.longitude),
        );
      });
    });
  }

  void _stopTracking() async {
    _positionStream?.cancel();
    _stopwatch.stop();
    _timer.cancel();

    setState(() {
      _isTracking = false;
      _lastRoutePoints = List<GeoPoint>.from(_geoPoints);
    });

    final storage = FlutterSecureStorage();
    final jwt = await storage.read(key: 'jwt_token');
    if (jwt == null) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text("JWT nicht gefunden. Anmeldung erforderlich.")),
      );
      return;
    }

    final lineString = geoPointsToLineString(_geoPoints);
    final duration = _formattedTime;
    final distance = _calculateTotalDistance(_geoPoints);

    if (lineString.isEmpty) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text("Keine Tracking-Daten vorhanden.")),
      );
      return;
    }

    try {
      await TrackingApi.addRun(jwt, lineString, duration, distance);
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text("Tracking-Daten erfolgreich gespeichert.")),
      );
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text("Fehler beim Speichern der Tracking-Daten: $e")),
      );
    }
  }

  String geoPointsToLineString(List<GeoPoint> points) {
    if (points.isEmpty) return "";
    final coordinates = points
        .map((p) => "${p.longitude} ${p.latitude}")
        .join(", ");
    return "LINESTRING($coordinates)";
  }

  double _calculateTotalDistance(List<GeoPoint> points) {
    if (points.length < 2) return 0.0;
    double total = 0.0;
    for (int i = 0; i < points.length - 1; i++) {
      total += _calculateDistanceBetween(points[i], points[i + 1]);
    }
    return total;
  }

  double _calculateDistanceBetween(GeoPoint a, GeoPoint b) {
    const double R = 6371; // Earth's radius in km
    double dLat = _deg2rad(b.latitude - a.latitude);
    double dLon = _deg2rad(b.longitude - a.longitude);
    double lat1 = _deg2rad(a.latitude);
    double lat2 = _deg2rad(b.latitude);

    double hav =
        (sin(dLat / 2) * sin(dLat / 2)) +
        (sin(dLon / 2) * sin(dLon / 2) * cos(lat1) * cos(lat2));
    double c = 2 * atan2(sqrt(hav), sqrt(1 - hav));
    return R * c;
  }

  double _deg2rad(double deg) => deg * (pi / 180);

  @override
  Widget build(BuildContext context) {
    Widget content;

    if (_isTracking && _wasStarted) {
      content = ElevatedButton(onPressed: _stopTracking, child: Text("Stop"));
    } else if (!_isTracking && _wasStarted) {
      content = Row(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          ElevatedButton(
            onPressed: () {
              Navigator.pushAndRemoveUntil(
                context,
                MaterialPageRoute(
                  builder: (context) => AppNavigator(title: 'Rocket App', initialIndex: 2),
                ),
                (route) => false,
              );
            },
            child: Text("Home"),
          ),
          SizedBox(width: 20),
          ElevatedButton(
            onPressed: () {
              Navigator.pushAndRemoveUntil(
                context,
                MaterialPageRoute(
                  builder: (context) => RoutePage(
                    title: "Completed Run",
                    routePoints: _lastRoutePoints,
                    elapsedTime: _formattedTime,
                  ),
                ),
                (route) => false,
              );
            },
            child: Text("Run anzeigen"),
          ),
        ],
      );
    } else {
      content = ElevatedButton(onPressed: _startTracking, child: Text("Start"));
    }

    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
        backgroundColor: Colors.blueGrey[100],
      ),
      backgroundColor: Colors.blueGrey[100],
      body: Padding(
        padding: EdgeInsets.all(16),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            if (_stopwatch.elapsed.inMinutes >= 10)
              Icon(
                Icons.star,
                color: Colors.yellow,
                size: 48,
              ),
            Text(
              "Timer: $_formattedTime",
              style: TextStyle(fontSize: 32),
              textAlign: TextAlign.center,
            ),
            SizedBox(height: 30),
            Center(child: content),
          ],
        ),
      ),
    );
  }
}
