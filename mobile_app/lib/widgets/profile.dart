import 'dart:convert';
import 'dart:typed_data';
import 'package:flutter/material.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import '/constants/constants.dart';
import '/utils/backend_api/user_api.dart';

class Profile extends StatefulWidget {
  const Profile({super.key});

  @override
  State<Profile> createState() => _ProfileState();
}

class _ProfileState extends State<Profile> {
  late Future<ImageProvider?> _userImageFuture;

  @override
  void initState() {
    super.initState();

    _userImageFuture = Future.value(null);

    final storage = FlutterSecureStorage();
    storage.read(key: 'jwt_token').then((jwt) {
      if (jwt != null) {
        setState(() {
          _userImageFuture = _fetchUserImage(jwt, 'b69d9b37-5b38-4a28-b1b2-8edaf4ea8673');
        });
      }
    });
  }

  Future<ImageProvider?> _fetchUserImage(String jwt, String userId) async {
    try {
      final userImage = await UserApi.fetchUserImage(jwt, userId);

      final imageData = base64Decode(userImage.data);

      return MemoryImage(Uint8List.fromList(imageData));
    } catch (e) {
      debugPrint('Error fetching user image: $e');
      return null;
    }
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          FutureBuilder<ImageProvider?>(
            future: _userImageFuture,
            builder: (context, snapshot) {
              if (snapshot.connectionState == ConnectionState.waiting) {
                return const CircleAvatar(
                  radius: 50,
                  backgroundColor: Colors.grey,
                  child: CircularProgressIndicator(),
                );
              } else if (snapshot.hasError || snapshot.data == null) {
                return const CircleAvatar(
                  radius: 50,
                  backgroundColor: Colors.grey,
                );
              } else {
                // If the image is successfully fetched, display it
                return CircleAvatar(
                  radius: 50,
                  backgroundImage: snapshot.data,
                );
              }
            },
          ),
          const SizedBox(height: 16),
          const Text(
            'You',
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
