import 'dart:convert';
import 'dart:io';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;
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

  static Future<void> updateStepGoal(String jwt, int stepGoal) async {
    // Ensure stepGoal is valid
    final response = await BaseApi.post(
      '/api/v1/protected/settings/step-goal',
      headers: {'Authorization': 'Bearer $jwt'},
      body: jsonEncode({'stepGoal': stepGoal}),
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to update step goal');
    }
  }

  static Future<void> updateImage(String jwt, File imageFile) async {
    final String backendUrl = dotenv.get('BACKEND_URL', fallback: "http://10.0.2.2:8080");
    final uri = Uri.parse('$backendUrl/api/v1/protected/settings/image');
    String apiKey = dotenv.get('API_KEY', fallback: "default-api-key");
    final request = http.MultipartRequest('POST', uri)
      ..headers['Authorization'] = 'Bearer $jwt'
      ..headers['X-API-KEY'] = apiKey
      ..files.add(await http.MultipartFile.fromPath('image', imageFile.path));

    final streamedResponse = await request.send();
    final response = await http.Response.fromStream(streamedResponse);

    if (response.statusCode != 200) {
      print('Upload fehlgeschlagen: ${response.statusCode}');
      print('Fehlerantwort: ${response.body}');
      throw Exception('Image upload failed');
    }
  }
}
