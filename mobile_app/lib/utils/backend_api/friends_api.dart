import 'dart:convert';
import 'package:mobile_app/utils/backend_api/base_api.dart';

/// Model for a Friend
class Friend {
  final String id;
  final String username;
  final String email;
  final int rocketPoints;

  const Friend({required this.id, required this.username, required this.email, required this.rocketPoints });

  factory Friend.fromJson(Map<String, dynamic> json) {
    return Friend(
      id: json['id'] as String,
      username: json['username'] as String,
      email: json['email'] as String,
      rocketPoints: json['rocket_points'] as int,
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
    final response = await BaseApi.post(
      '/api/v1/protected/friends/add',
      headers: {
        'Authorization': 'Bearer $jwt',
      },
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
      headers: {
        'Authorization': 'Bearer $jwt',
      },
      body: {'friend_name': friendName},
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to delete friend: ${response.statusCode}');
    }
  }
}
