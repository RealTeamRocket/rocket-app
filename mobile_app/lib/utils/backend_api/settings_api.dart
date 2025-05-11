import 'base_api.dart';
import 'dart:convert';

class SettingsResponse {
  final String id;
  final String userId;
  final String? imageId;
  final int stepGoal;

  const SettingsResponse({
    required this.id,
    required this.userId,
    this.imageId,
    required this.stepGoal,
  });

  factory SettingsResponse.fromJson(Map<String, dynamic> json) {
    return SettingsResponse(
      id: json['id'],
      userId: json['userId'],
      imageId: json['imageId'],
      stepGoal: json['stepGoal'],
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

  static Future<void> updateSettings(String jwt, int stepGoal, {String? imageId}) async {
    final body = {
      'stepGoal': stepGoal,
      if (imageId != null) 'imageId': imageId,
    };

    final response = await BaseApi.post(
      '/api/v1/protected/settings/update',
      headers: {'Authorization': 'Bearer $jwt'},
      body: body,
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to update settings');
    }
  }

  static Future<void> createSettings(String jwt, int stepGoal, {String? imageId}) async {
    final body = {
      'stepGoal': stepGoal,
      if (imageId != null) 'imageId': imageId,
    };

    final response = await BaseApi.post(
      '/api/v1/protected/settings/create',
      headers: {'Authorization': 'Bearer $jwt'},
      body: body,
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to create settings');
    }
  }
}