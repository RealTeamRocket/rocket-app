import 'package:flutter/material.dart';
import '../constants/color_constants.dart';
import '/widgets/widgets.dart';
import 'leaderboard.dart';
import 'pages.dart';

class HomePage extends StatefulWidget {
  const HomePage({super.key, required this.title});

  final String title;

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  int _selectedIndex = 2; /// set position of starting page to homepage in the middle

  final List<Widget> _pages = <Widget>[
    const ChallengePage(title: 'Challenges'),
    const LeaderboardPage(title: 'Leaderboard'),
    const RunPage(title: 'Run'),
    const FriendlistPage(title: 'friends'),
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
        title: SizedBox(
          height: kToolbarHeight,
          child: Stack(
            children: [
              Align(
                alignment: Alignment.center,
                child: Text(
                  widget.title,
                  style: TextStyle(color: ColorConstants.white),
                ),
              ),
              if (_selectedIndex == 4)
                Align(
                  alignment: Alignment.centerRight,
                  child: IconButton(
                    icon: Icon(Icons.settings, color: ColorConstants.white),
                    onPressed: () {
                      Navigator.pushNamed(context, '/settings');
                    },
                  ),
                ),
            ],
          ),
        ),
        centerTitle: true,
      ),
      body: Center(child: _pages.elementAt(_selectedIndex)),
      bottomNavigationBar: CustomMenuBar(onItemTapped: _onItemTapped),
    );
  }
}
