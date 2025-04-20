import 'package:flutter/material.dart';
import '/constants/constants.dart';

class CustomMenuBar extends StatefulWidget {
  const CustomMenuBar({required this.onItemTapped, super.key});

  final void Function(int) onItemTapped;

  @override
  State<CustomMenuBar> createState() => _CustomMenuBarState();
}

class _CustomMenuBarState extends State<CustomMenuBar> {
  int _currentIndex = 0;

  @override
  Widget build(BuildContext context) {
    return BottomNavigationBar(
      type: BottomNavigationBarType.shifting,
      items: const <BottomNavigationBarItem>[
        BottomNavigationBarItem(
          icon: Icon(Icons.home),
          label: 'Home',
          backgroundColor: ColorConstants.secoundaryColor,
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.checklist),
          label: 'Search',
          backgroundColor: ColorConstants.secoundaryColor,
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.leaderboard),
          label: 'Leaderboards',
          backgroundColor: ColorConstants.secoundaryColor,
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.person),
          label: 'Profile',
          backgroundColor: ColorConstants.secoundaryColor,
        ),
      ],
      currentIndex: _currentIndex,
      selectedItemColor: ColorConstants.purpleColor,
      onTap: (index) {
        setState(() {
          _currentIndex = index;
        });
        widget.onItemTapped(index);
      },
    );
  }
}
