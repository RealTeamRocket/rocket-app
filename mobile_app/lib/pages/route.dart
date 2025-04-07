import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:flutter_osm_plugin/flutter_osm_plugin.dart';
import 'package:latlong2/latlong.dart';
import 'package:geolocator/geolocator.dart';


class RoutePage extends StatefulWidget {
  const RoutePage({super.key, required this.title});

  final String title;

  @override
  State<RoutePage> createState() => _RoutePageState();
}

Future<Position> _determinePosition() async{
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

  @override
  void initState() {
    super.initState();
    _setInitialPosition();
  }
  Future<void> _setInitialPosition() async {
    try {
      Position position = await _determinePosition();
      setState(() {
        _initPosition = GeoPoint(latitude: position.latitude, longitude: position.longitude);
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
          : OSMViewer(controller: SimpleMapController
          (initPosition: _initPosition!,
            markerHome: const MarkerIcon(icon: Icon(Icons.location_on))),
          zoomOption: const ZoomOption(initZoom: 16, minZoomLevel: 11),
        )
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


