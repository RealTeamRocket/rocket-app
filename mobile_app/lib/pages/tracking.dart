import 'dart:async';
import 'dart:math';
import 'package:flutter/material.dart';
import 'package:geolocator/geolocator.dart';
import 'package:flutter_osm_plugin/flutter_osm_plugin.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:mobile_app/pages/pages.dart';

import '../constants/color_constants.dart';
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

  // For displaying runs
  List<RunModel> _runs = [];
  bool _isLoadingRuns = false;

  @override
  void initState() {
    super.initState();
    _fetchRuns();
  }

  Future<void> _fetchRuns() async {
    setState(() {
      _isLoadingRuns = true;
    });
    try {
      final storage = FlutterSecureStorage();
      final jwt = await storage.read(key: 'jwt_token');
      if (jwt != null) {
        final runs = await TrackingApi.getAllRuns(jwt);
        setState(() {
          _runs = runs;
        });
      }
    } catch (e) {
      // Optionally show error
    } finally {
      setState(() {
        _isLoadingRuns = false;
      });
    }
  }

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
      ScaffoldMessenger.of(
        context,
      ).showSnackBar(SnackBar(content: Text("No tracking data to save.")));
      return;
    }

    try {
      await TrackingApi.addRun(jwt, lineString, duration, distance);
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text("Tracking-Daten erfolgreich gespeichert.")),
      );
      await _fetchRuns(); // Refresh runs after saving
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

  // Helper to parse WKT LineString to List<GeoPoint>
  List<GeoPoint> parseLineString(String lineString) {
    final regex = RegExp(r'LINESTRING\((.*)\)');
    final match = regex.firstMatch(lineString);
    if (match == null) return [];
    final coords = match.group(1)!.split(',');
    List<GeoPoint> points = [];
    for (var c in coords) {
      var parts = c.trim().split(RegExp(r'\s+'));
      if (parts.length != 2) continue;
      try {
        double lon = double.parse(parts[0]);
        double lat = double.parse(parts[1]);
        points.add(GeoPoint(latitude: lat, longitude: lon));
      } catch (e) {
        // Skip invalid points
        continue;
      }
    }
    return points;
  }

  String _formatDate(String isoString) {
    try {
      final date = DateTime.parse(isoString);
      return "${date.day.toString().padLeft(2, '0')}.${date.month.toString().padLeft(2, '0')}.${date.year}, "
          "${date.hour.toString().padLeft(2, '0')}:${date.minute.toString().padLeft(2, '0')}";
    } catch (_) {
      return isoString;
    }
  }

  @override
  Widget build(BuildContext context) {
    Widget content;

    if (_isTracking && _wasStarted) {
      content = ElevatedButton(
          onPressed: _stopTracking,
          style: ElevatedButton.styleFrom(
            backgroundColor: ColorConstants.purpleColor,
            foregroundColor: ColorConstants.white,
          ),
          child: Text("Stop")
      );
    } else if (!_isTracking && _wasStarted) {
      content = Row(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          ElevatedButton(
            onPressed: () {
              Navigator.pushAndRemoveUntil(
                context,
                MaterialPageRoute(
                  builder:
                      (context) =>
                          AppNavigator(title: 'Rocket App', initialIndex: 2),
                ),
                (route) => false,
              );
            },
            style: ElevatedButton.styleFrom(
              backgroundColor: ColorConstants.purpleColor,
              foregroundColor: ColorConstants.white,
            ),
            child: Text("Home"),
          ),
          SizedBox(width: 20),
          ElevatedButton(
            onPressed: () {
              Navigator.pushAndRemoveUntil(
                context,
                MaterialPageRoute(
                  builder:
                      (context) => RoutePage(
                        title: "Completed Run",
                        routePoints: _lastRoutePoints,
                        elapsedTime: _formattedTime,
                      ),
                ),
                (route) => false,
              );
            },
            style: ElevatedButton.styleFrom(
              backgroundColor: ColorConstants.purpleColor,
              foregroundColor: ColorConstants.white,
            ),
            child: Text("Show Route"),
          ),
        ],
      );
    } else {
      content = ElevatedButton(
          onPressed: _startTracking,
          style: ElevatedButton.styleFrom(
            backgroundColor: ColorConstants.purpleColor,
            foregroundColor: ColorConstants.white,
          ),
          child: Text("Start"),
      );
    }

    return Scaffold(
      appBar: AppBar(
        iconTheme: IconThemeData(color: ColorConstants.white),
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
      backgroundColor: ColorConstants.primaryColor,
      body: Padding(
        padding: EdgeInsets.all(16),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            if (_stopwatch.elapsed.inMinutes >= 10)
              Icon(Icons.star, color: Colors.yellow, size: 48),
            Text(
              "Timer: $_formattedTime",
              style: TextStyle(
                fontSize: 32,
                color: ColorConstants.white,
              ),
              textAlign: TextAlign.center,
            ),
            SizedBox(height: 30),
            Center(child: content),
            SizedBox(height: 30),
            Divider(),
            Text(
              "Past Runs",
              style: TextStyle(
                  fontSize: 20,
                  color: ColorConstants.white,
                  fontWeight: FontWeight.bold
              ),
            ),
            SizedBox(height: 10),
            _isLoadingRuns
                ? CircularProgressIndicator()
                : Expanded(
                  child:
                      _runs.isEmpty
                          ? Text(
                              "No runs found.",
                              style: TextStyle(
                              color: ColorConstants.white,
                            ),
                          )
                          : ListView.builder(
                            itemCount: _runs.length,
                            itemBuilder: (context, index) {
                              final run = _runs[index];
                              return Card(
                                margin: const EdgeInsets.symmetric(
                                  vertical: 6,
                                  horizontal: 2,
                                ),
                                elevation: 2,
                                child: ListTile(
                                  contentPadding: const EdgeInsets.symmetric(
                                    vertical: 8,
                                    horizontal: 16,
                                  ),
                                  leading: Icon(
                                    Icons.directions_run,
                                    color: Colors.blueAccent,
                                    size: 32,
                                  ),
                                  title: Row(
                                    children: [
                                      Icon(
                                        Icons.route,
                                        size: 18,
                                        color: Colors.grey[700],
                                      ),
                                      SizedBox(width: 6),
                                      Text(
                                        "${run.distance.toStringAsFixed(2)} km",
                                      ),
                                    ],
                                  ),
                                  subtitle: Column(
                                    crossAxisAlignment:
                                        CrossAxisAlignment.start,
                                    children: [
                                      Row(
                                        children: [
                                          Icon(
                                            Icons.calendar_today,
                                            size: 16,
                                            color: Colors.grey[600],
                                          ),
                                          SizedBox(width: 4),
                                          Text(_formatDate(run.createdAt)),
                                        ],
                                      ),
                                      SizedBox(height: 2),
                                      Row(
                                        children: [
                                          Icon(
                                            Icons.timer,
                                            size: 18,
                                            color: Colors.grey[700],
                                          ),
                                          SizedBox(width: 6),
                                          Text(run.duration),
                                        ],
                                      ),
                                    ],
                                  ),
                                  trailing: Icon(
                                    Icons.chevron_right,
                                    color: Colors.grey[600],
                                  ),
                                  onTap: () {
                                    final geoPoints = parseLineString(
                                      run.route,
                                    );
                                    Navigator.push(
                                      context,
                                      MaterialPageRoute(
                                        builder:
                                            (context) => RoutePage(
                                              title: _formatDate(run.createdAt),
                                              routePoints: geoPoints,
                                              elapsedTime: run.duration,
                                            ),
                                      ),
                                    );
                                  },
                                ),
                              );
                            },
                          ),
                ),
          ],
        ),
      ),
    );
  }
}
