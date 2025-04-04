import 'package:flutter/material.dart';
import '/constants/constants.dart';
import '/widgets/widgets.dart';

class Profile extends StatelessWidget {
  const Profile({super.key});

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Container(
        //height: 300,
        width: double.infinity,
        color: ColorConstants.greenColor,
        child: const Center(
          child: Text(
            'Profile',
            style: TextStyle(color: ColorConstants.white, fontSize: 28),
          ),
        ),
      ),
    );
  }
}

class History extends StatelessWidget {
  const History({super.key});

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Container(
        //height: 300,
        width: double.infinity,
        color: ColorConstants.deepBlue,
        child: const Center(
          child: Text(
            'Histroy',
            style: TextStyle(color: ColorConstants.white, fontSize: 28),
          ),
        ),
      ),
    );
  }
}

class ProfilePage extends StatefulWidget {
  const ProfilePage({super.key});

  @override
  State<ProfilePage> createState() => _ProfilePageState();
}

class _ProfilePageState extends State<ProfilePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        children: [
          Expanded(
            child: Profile(),
          ),
          const SizedBox(height: 10),
          Expanded(
            child: History(),
          ),
        ],
      ),
    );
  }
}