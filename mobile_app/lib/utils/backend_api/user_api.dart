import 'dart:convert';
import 'package:flutter/foundation.dart';
import 'package:mobile_app/utils/backend_api/base_api.dart';

class UserStatistics {
  final int steps;
  final String day;

  const UserStatistics({required this.steps, required this.day});

  factory UserStatistics.fromJson(Map<String, dynamic> json) {
    return UserStatistics(
      steps: json['steps'] as int,
      day: json['day'] as String,
    );
  }
}

class UserImage {
  final String username;
  final String name;
  final String mimeType;
  final String data;

  const UserImage({
    required this.username,
    required this.name,
    required this.mimeType,
    required this.data,
  });

  factory UserImage.fromJson(Map<String, dynamic> json) {
    return UserImage(
      username: json['username'] as String,
      name: json['name'] as String,
      mimeType: json['mime_type'] as String,
      data: json['data'] as String,
    );
  }
}

class UserApi {
  /// Fetch user statistics
  static Future<List<UserStatistics>> fetchUserStatistics(String jwt) async {
    final response = await BaseApi.post(
      '/api/v1/protected/user/statistics',
      headers: {'Authorization': 'Bearer $jwt'},
    );

    if (response.statusCode != 200) {
      throw Exception(
        'Failed to fetch user statistics: ${response.statusCode}',
      );
    }

    final List<dynamic> jsonList = jsonDecode(response.body) as List<dynamic>;
    return jsonList
        .map((json) => UserStatistics.fromJson(json as Map<String, dynamic>))
        .toList();
  }

  /// Fetch user image
  static Future<UserImage> fetchUserImage(String jwt, {String? userId}) async {
    // Prepare the request body
    final body = userId != null ? {'user_id': userId} : null;

    final response = await BaseApi.post(
      '/api/v1/protected/user/image',
      headers: {'Authorization': 'Bearer $jwt'},
      body: body,
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to fetch user image: ${response.statusCode}');
    }

    return UserImage.fromJson(
      jsonDecode(response.body) as Map<String, dynamic>,
    );
  }
}
