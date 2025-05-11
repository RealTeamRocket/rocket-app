import 'package:flutter/material.dart';
import '/constants/constants.dart';
import 'dart:io';
import '/utils/backend_api/settings_api.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class Profile extends StatefulWidget {
  const Profile({super.key});

  @override
  State<Profile> createState() => _ProfileState();
}

class _ProfileState extends State<Profile> {
  String? imagePath;
  String? username;
  bool isLoading = true;
  final _storage = const FlutterSecureStorage();

  @override
  void initState() {
    super.initState();
    _fetchProfileData();
  }

  Future<void> _fetchProfileData() async {
    try {
      final jwt = await _storage.read(key: 'jwt_token');
      if (jwt == null) throw Exception('JWT token not found');

      // Fetch settings
      final settings = await SettingsApi.getSettings(jwt);
      print('Settings: $settings');
      // Fetch image
      final imageFile = await SettingsApi.getImage(jwt, settings.imageId);

      setState(() {
        username = settings.userId; // Assuming userId is the username
        imagePath = imageFile.path;
        isLoading = false;
      });
    } catch (e) {
      setState(() {
        isLoading = false;
      });
      // Handle error (e.g., show a snackbar or log the error)
    }
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          isLoading
              ? const CircularProgressIndicator()
              : CircleAvatar(
                  radius: 50,
                  backgroundImage: imagePath != null
                      ? FileImage(File(imagePath!))
                      : const AssetImage('assets/images/profile_picture.png')
                          as ImageProvider,
                ),
          const SizedBox(height: 16),
          Text(
            username ?? 'Unknown User',
            style: const TextStyle(
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