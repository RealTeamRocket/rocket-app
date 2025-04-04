import 'package:flutter/material.dart';
import '/constants/constants.dart';
import '/widgets/widgets.dart';

class Profile extends StatelessWidget {
  const Profile({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 300, // Höhe hinzugefügt
      width: double.infinity,
      color: ColorConstants.deepBlue, // Farbanpassung
      child: const Center(
        child: Text(
          'Profile',
          style: TextStyle(color: ColorConstants.white, fontSize: 28), // Größere Schrift
        ),
      ),
    );
  }
}

class History extends StatelessWidget {
  const History({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 300, // Höhe hinzugefügt
      width: double.infinity,
      color: ColorConstants.deepBlue, // Farbanpassung
      child: const Center(
        child: Text(
          'History',
          style: TextStyle(color: ColorConstants.white, fontSize: 28), // Größere Schrift
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
      appBar: AppBar(
        title: const Text('Profile Page'),
        backgroundColor: ColorConstants.deepBlue, // Farbanpassung
      ),
      body: SingleChildScrollView( // Scrollbarkeit hinzugefügt
        child: Column(
          children: const [
            Profile(),
            SizedBox(height: 16), // Abstand zwischen den Widgets
            History(),
          ],
        ),
      ),
    );
  }
}