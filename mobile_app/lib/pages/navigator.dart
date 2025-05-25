import 'package:flutter/material.dart';
import '../constants/color_constants.dart';
import '/widgets/widgets.dart';
import 'leaderboard.dart';
import 'pages.dart';

class AppNavigator extends StatefulWidget {
  const AppNavigator({
    super.key,
    required this.title,
    this.initialIndex = 2, // Default to "Run" tab in the middle
  });

  final String title;
  final int initialIndex;

  @override
  State<AppNavigator> createState() => _AppNavigatorState();
}

class _AppNavigatorState extends State<AppNavigator> {
  late int _selectedIndex;

  final List<Widget> _pages = <Widget>[
    const ChallengePage(title: 'Challenges'),
    const LeaderboardPage(title: 'Leaderboard'),
    const RunPage(title: 'Run'),
    const FriendlistPage(title: 'friends'),
    const ProfilePage(),
  ];

  @override
  void initState() {
    super.initState();
    _selectedIndex = widget.initialIndex;
  }

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
