import 'dart:convert';
import 'dart:io';


import 'base_api.dart';

class SettingsResponse {
  final String id;
  final String userId;
  final String imageId;
  final int stepGoal;

  const SettingsResponse({
    required this.id,
    required this.userId,
    required this.imageId,
    required this.stepGoal,
  });

  factory SettingsResponse.fromJson(Map<String, dynamic> json) {
    return SettingsResponse(
      id: json['id'],
      userId: json['user_id'],
      imageId: json['image_id'],
      stepGoal: json['step_goal'],
    );
  }
}

class SettingsApi {
  static Future<SettingsResponse> getSettings(String jwt) async {
    final response = await BaseApi.get(
      '/api/v1/protected/settings',
      headers: {'Authorization': 'Bearer $jwt'},
    );

    if (response.statusCode == 200) {
      return SettingsResponse.fromJson(
        jsonDecode(response.body) as Map<String, dynamic>,
      );
    } else if (response.statusCode == 404) {
      throw Exception('Settings not found');
    } else {
      throw Exception('Failed to fetch settings');
    }
  }

  static Future<void> updateSettings(String jwt, int stepGoal, {File? imageFile,}) async {
    // Prepare the settings JSON as a form field
    final settingsJson = jsonEncode({'stepGoal': stepGoal});

    // Call the BaseApi's postMultipart method
    final response = await BaseApi.postMultipart(
      '/api/v1/protected/settings/update',
      headers: {'Authorization': 'Bearer $jwt'},
      fields: {'settings': settingsJson},
      file: imageFile,
      fileFieldName: 'image',
    );

    if (response.statusCode != 200) {
      final responseBody = await response.stream.bytesToString();
      throw Exception('Failed to update settings: $responseBody');
    }
  }

static Future<File> getUserImage(String jwt) async {
  final response = await BaseApi.get(
    '/api/v1/protected/image',
    headers: {'Authorization': 'Bearer $jwt'},
  );

  if (response.statusCode == 200) {
    final bytes = response.bodyBytes;
    final tempDir = Directory.systemTemp;
    final file = File('${tempDir.path}/user_image.jpg');
    await file.writeAsBytes(bytes);
    return file;
  } else if (response.statusCode == 404) {
    throw Exception('User image not found');
  } else {
    throw Exception('Failed to fetch user image');
  }
}
}