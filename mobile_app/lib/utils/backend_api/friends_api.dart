import 'dart:convert';
import 'dart:typed_data';
import 'package:flutter/material.dart';
import 'package:mobile_app/utils/backend_api/base_api.dart';

/// Model for a Friend
class Friend {
  final String id;
  final String username;
  final String email;
  final int rocketPoints;
  final Uint8List? imageData;

  const Friend({
    required this.id,
    required this.username,
    required this.email,
    required this.rocketPoints,
    this.imageData,
  });

  factory Friend.fromJson(Map<String, dynamic> json) {
    Uint8List? data;
    if (json['image_data'] != null) {
      final base64Str = json['image_data'] as String;
      data = base64Decode(base64Str);
    }

    return Friend(
      id: json['id'] as String,
      username: json['username'] as String,
      email: json['email'] as String,
      rocketPoints: json['rocket_points'] as int,
      imageData: data,
    );
  }
}

/// Extended model for all users including image data
class UserWithImage {
  final String id;
  final String username;
  final String email;
  final int rocketPoints;
  final String imageName;
  final Uint8List? imageData;

  const UserWithImage({
    required this.id,
    required this.username,
    required this.email,
    required this.rocketPoints,
    required this.imageName,
    this.imageData,
  });

  factory UserWithImage.fromJson(Map<String, dynamic> json) {
    Uint8List? data;
    if (json['image_data'] != null) {
      final base64Str = json['image_data'] as String;
      data = base64Decode(base64Str);
    }
    return UserWithImage(
      id: json['id'] as String,
      username: json['username'] as String,
      email: json['email'] as String,
      rocketPoints: json['rocket_points'] as int,
      imageName: json['image_name'] as String,
      imageData: data,
    );
  }
}

class FriendsApi {
  /// Fetch all friends of the authenticated user
  static Future<List<Friend>> getAllFriends(String jwt) async {
    final response = await BaseApi.get(
      '/api/v1/protected/friends',
      headers: {'Authorization': 'Bearer $jwt'},
    );

    if (response.statusCode == 404) {
      return [];
    }

    if (response.statusCode != 200) {
      throw Exception('Failed to fetch friends: ${response.statusCode}');
    }

    final decoded = jsonDecode(response.body);

    if (decoded is! List) {
      throw Exception('Invalid response format for friends: Expected a list');
    }

    return decoded
        .map((e) => Friend.fromJson(e as Map<String, dynamic>))
        .toList();
  }

  /// Add a friend by username
  static Future<void> addFriend(String jwt, String friendName) async {
    debugPrint(friendName);
    final response = await BaseApi.post(
      '/api/v1/protected/friends/add',
      headers: {'Authorization': 'Bearer $jwt'},
      body: {'friend_name': friendName},
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to add friend: ${response.statusCode}');
    }
  }

  /// Delete a friend by username
  static Future<void> deleteFriend(String jwt, String friendName) async {
    final response = await BaseApi.delete(
      '/api/v1/protected/friends/delete',
      headers: {'Authorization': 'Bearer $jwt'},
      body: {'friend_name': friendName},
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to delete friend: ${response.statusCode}');
    }
  }

  /// Fetch all users (for searching new friends)
  static Future<List<UserWithImage>> getAllUsers(String jwt) async {
    final response = await BaseApi.get(
      '/api/v1/protected/users',
      headers: {'Authorization': 'Bearer $jwt'},
    );
    debugPrint(response.body);

    if (response.statusCode == 404) {
      return [];
    }

    if (response.statusCode != 200) {
      throw Exception('Failed to fetch users: \${response.statusCode}');
    }

    final decoded = jsonDecode(response.body);

    if (decoded is! List) {
      throw Exception('Invalid response format for users: Expected a list');
    }

    return decoded
        .map((e) => UserWithImage.fromJson(e as Map<String, dynamic>))
        .toList();
  }
}
