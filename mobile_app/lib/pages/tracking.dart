import 'dart:async';
import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:geolocator/geolocator.dart';
import 'package:http/http.dart' as http;
import 'package:flutter_osm_plugin/flutter_osm_plugin.dart';
import 'package:http/http.dart' as dotenv;
import 'package:http/http.dart' as storage;
import 'package:mobile_app/utils/backend_api/backend_api.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'dart:convert';
import 'dart:io';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;
import 'package:path/path.dart';

import '../utils/backend_api/tracking_api.dart';

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
        ScaffoldMessenger.of(context as BuildContext).showSnackBar(
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
    String? userId;

    setState(() {
      _isTracking = false;
    });

    try {
      final storage = FlutterSecureStorage();
      userId = await storage.read(key: 'userId');
      if (userId == null) {
        debugPrint("Benutzer-ID ist null");
        return;
      }

      debugPrint("Benutzer-ID erfolgreich abgerufen: $userId");

      // Hier kannst du die userId weiterverwenden, z. B. für API-Aufrufe
    } catch (e) {
      debugPrint("Fehler beim Abrufen der Benutzer-ID: $e");
    }

    if (userId == null) {
      ScaffoldMessenger.of(context as BuildContext).showSnackBar(
        SnackBar(
          content: Text("Benutzer-ID nicht gefunden. Anmeldung erforderlich."),
        ),
      );
      return;
    }

    final lineString = geoPointsToLineString(_geoPoints);

    if (lineString.isEmpty) {
      ScaffoldMessenger.of(context as BuildContext).showSnackBar(
        SnackBar(content: Text("Keine Tracking-Daten vorhanden.")),
      );
      return;
    }

    try {
      await TrackingApi.addRun(userId, lineString);
      ScaffoldMessenger.of(context as BuildContext).showSnackBar(
        SnackBar(content: Text("Tracking-Daten erfolgreich gespeichert.")),
      );
    } catch (e) {
      ScaffoldMessenger.of(context as BuildContext).showSnackBar(
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

  @override
  Widget build(BuildContext context) {
    Widget content;

    if (_isTracking && _wasStarted) {
      content = ElevatedButton(onPressed: _stopTracking, child: Text("Stop"));
    } else if (!_isTracking && _wasStarted) {
      content = Row(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          ElevatedButton(onPressed: () {}, child: Text("Home")),
          SizedBox(width: 20),
          ElevatedButton(onPressed: () {}, child: Text("Run anzeigen")),
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
            /**
             * TODO
             */
            ///Hier noch Icons verändern oder nach 20 min 2 Stars und 30 3 stars.
            if (_stopwatch.elapsed.inMinutes >= 10) // Icon anzeigen, wenn länger als 10 Minuten
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

void main() {
  runApp(MaterialApp(home: TrackingPage(title: 'Tracking App')));
}
