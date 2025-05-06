import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:mobile_app/utils/backend_api/base_api.dart';

/// Model for a ranked user
class RankedUser {
  final String id;
  final String username;
  final String email;
  final int rocketPoints;

  const RankedUser({
    required this.id,
    required this.username,
    required this.email,
    required this.rocketPoints,
  });

  factory RankedUser.fromJson(Map<String, dynamic> json) {
    return RankedUser(
      id: json['id'] as String,
      username: json['username'] as String,
      email: json['email'] as String,
      rocketPoints: json['rocket_points'] as int,
    );
  }
}

class RankingApi {
  /// Fetch rankings for all users
  static Future<List<RankedUser>> fetchUserRankings(String jwt) async {
    final response = await BaseApi.get(
      '/api/v1/protected/ranking/users',
      headers: {'Authorization': 'Bearer $jwt'},
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to fetch user rankings: ${response.statusCode}');
    }

    debugPrint("Response body (users): ${response.body}");

    if (response.body.isEmpty) {
      debugPrint("Response body is null or empty for users");
      return [];
    }

    final decoded = jsonDecode(response.body);

    // Ensure the decoded response is a list
    if (decoded is! List) {
      throw Exception('Invalid response format for users: Expected a list');
    }

    return decoded
        .map((e) => RankedUser.fromJson(e as Map<String, dynamic>))
        .toList();
  }

  /// Fetch rankings for friends
  static Future<List<RankedUser>> fetchFriendRankings(String jwt) async {
    final response = await BaseApi.get(
      '/api/v1/protected/ranking/friends',
      headers: {'Authorization': 'Bearer $jwt'},
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to fetch friend rankings: ${response.statusCode}');
    }

    debugPrint("Response body (friends): ${response.body}");

    if (response.body.isEmpty) {
      debugPrint("Response body is null or empty for friends");
      return [];
    }

    final decoded = jsonDecode(response.body);

    if (decoded is! List) {
      debugPrint("Invalid response format for friends: ${response.body}");
      return [];
    }

    return decoded
        .map((e) => RankedUser.fromJson(e as Map<String, dynamic>))
        .toList();
  }
}
