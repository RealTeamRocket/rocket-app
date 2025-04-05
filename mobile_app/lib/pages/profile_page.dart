import 'package:flutter/material.dart';
import '/constants/constants.dart';
import '/widgets/widgets.dart';

class ProfilePage extends StatefulWidget {
  const ProfilePage({super.key});

  @override
  State<ProfilePage> createState() => _ProfilePageState();
}

class _ProfilePageState extends State<ProfilePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: ColorConstants.deepBlue,
      body: Column(
        children: [
          Expanded(
            child: Profile(),
          ),
          Expanded(
            child: History(),
          ),
        ],
      ),
    );
  }
}