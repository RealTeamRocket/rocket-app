import 'package:flutter/material.dart';
import '../constants/color_constants.dart';
import '/widgets/widgets.dart';
import 'pages.dart';

class HomePage extends StatefulWidget {
  const HomePage({super.key, required this.title});

  final String title;

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  int _selectedIndex = 0;

  final List<Widget> _pages = <Widget>[
    const RunPage(title: 'Run'),
    Center(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: const [Text("Search"), Icon(Icons.checklist)],
      ),
    ),
    Center(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: const [Text("Leaderboards"), Icon(Icons.leaderboard)],
      ),
    ),
    const ProfilePage(),
  ];

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: ColorConstants.secoundaryColor,
        title: Center(
          child: Text(
            widget.title,
            style: TextStyle(color: ColorConstants.white),
          ),
        ),
      ),
      body: Center(child: _pages.elementAt(_selectedIndex)),
      bottomNavigationBar: CustomMenuBar(onItemTapped: _onItemTapped),
    );
  }
}
