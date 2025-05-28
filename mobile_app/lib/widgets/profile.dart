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
  late Future<UserImage?> _userImageFuture;

  @override
  void initState() {
    super.initState();

    _userImageFuture = Future.value(null);

    final storage = FlutterSecureStorage();
    storage.read(key: 'jwt_token').then((jwt) {
      if (jwt != null) {
        setState(() {
          _userImageFuture = _fetchUserImage(jwt);
        });
      }
    });
  }

  Future<UserImage?> _fetchUserImage(String jwt) async {
    try {
      final userImage = await UserApi.fetchUserImage(jwt);
      return userImage;
    } catch (e) {
      debugPrint('Error fetching user image: $e');
      return null;
    }
  }

  void _showImageDialog(Uint8List imageData) {
    showDialog(
      context: context,
      builder: (context) => Dialog(
        backgroundColor: Colors.transparent,
        child: GestureDetector(
          onTap: () => Navigator.of(context).pop(),
          child: Container(
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(16),
              color: Colors.black.withOpacity(0.8),
            ),
            padding: const EdgeInsets.all(16),
            child: Image.memory(
              imageData,
              fit: BoxFit.contain,
            ),
          ),
        ),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          FutureBuilder<UserImage?>(
            future: _userImageFuture,
            builder: (context, snapshot) {
              if (snapshot.connectionState == ConnectionState.waiting) {
                return const CircleAvatar(
                  radius: 50,
                  backgroundColor: Colors.grey,
                  child: CircularProgressIndicator(),
                );
              } else if (snapshot.hasError || snapshot.data == null || snapshot.data!.data == null) {
                return const CircleAvatar(
                  radius: 50,
                  backgroundColor: Colors.grey,
                );
              } else {
                final userImage = snapshot.data!;
                final imageData = base64Decode(userImage.data!);
                return GestureDetector(
                  onTap: () => _showImageDialog(Uint8List.fromList(imageData)),
                  child: CircleAvatar(
                    radius: 50,
                    backgroundImage: MemoryImage(Uint8List.fromList(imageData)),
                  ),
                );
              }
            },
          ),
          const SizedBox(height: 16),
          FutureBuilder<UserImage?>(
            future: _userImageFuture,
            builder: (context, snapshot) {
              if (snapshot.connectionState == ConnectionState.waiting) {
                return const Text(
                  'Loading...',
                  style: TextStyle(
                    color: ColorConstants.white,
                    fontSize: 24,
                    fontWeight: FontWeight.bold,
                  ),
                );
              } else if (snapshot.hasError || snapshot.data == null) {
                return const Text(
                  'Error loading username',
                  style: TextStyle(
                    color: ColorConstants.white,
                    fontSize: 24,
                    fontWeight: FontWeight.bold,
                  ),
                );
              } else {
                final username = snapshot.data!.username;
                return Text(
                  username,
                  style: const TextStyle(
                    color: ColorConstants.white,
                    fontSize: 24,
                    fontWeight: FontWeight.bold,
                  ),
                );
              }
            },
          ),
        ],
      ),
    );
  }
}
