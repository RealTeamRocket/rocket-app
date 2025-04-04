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
      items: const <BottomNavigationBarItem>[
        BottomNavigationBarItem(
          icon: Icon(Icons.home),
          label: 'Home',
          backgroundColor: ColorConstants.greyColor,
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.checklist),
          label: 'Search',
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.leaderboard),
          label: 'Leaderboards',
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.person),
          label: 'Profile',
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
