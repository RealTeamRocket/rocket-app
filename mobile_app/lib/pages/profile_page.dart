import 'package:flutter/material.dart';
import '/constants/constants.dart';
// import '/widgets/widgets.dart';

class Profile extends StatelessWidget {
  const Profile({super.key});

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          CircleAvatar(
            radius: 50,
            backgroundImage: AssetImage('assets/images/profile_picture.png'),
          ),
          const SizedBox(height: 16),
          const Text(
            'Profile Name',
            style: TextStyle(
              color: ColorConstants.white,
              fontSize: 24,
              fontWeight: FontWeight.bold,
            ),
          ),
        ],
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
        width: double.infinity,
        color: ColorConstants.greenColor,
        child: const Center(
          child: Text(
            'History',
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