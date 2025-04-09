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
      backgroundColor: ColorConstants.primaryColor,
      body: Column(
        children: [
          Flexible(flex: 2, child: Profile()),
          Flexible(flex: 3, child: History()),
        ],
      ),
    );
  }
}
