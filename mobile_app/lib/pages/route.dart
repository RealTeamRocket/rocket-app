import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:flutter_osm_plugin/flutter_osm_plugin.dart';
import 'package:latlong2/latlong.dart';



class RoutePage extends StatefulWidget {
  const RoutePage({super.key, required this.title});

  final String title;

  @override
  State<RoutePage> createState() => _RoutePageState();
}



class _RoutePageState extends State<RoutePage> {

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: Text(widget.title),
        ),
        body: OSMViewer(controller: SimpleMapController
          (initPosition: GeoPoint(latitude: 48.783333, longitude: 9.183333),
            markerHome: const MarkerIcon(icon: Icon(Icons.home))),
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


