import 'dart:convert';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;

/// Model for single challenge
class Challenge {
  final String id;
  final String text;
  final int points;

  const Challenge({
    required this.id,
    required this.text,
    required this.points,
  });

  factory Challenge.fromJson(Map<String, dynamic> json) => Challenge(
    id: json['id'] as String,
    text: json['text'] as String,
    points: json['points'] as int,
  );
}

/// API-Service: Returning list of challenges
class ChallengesApi {
  static Future<List<Challenge>> fetchChallenges(String jwt) async {
    final backendUrl = dotenv.get('BACKEND_URL', fallback: 'http://10.0.2.2:8080');
    final uri = Uri.parse('$backendUrl/api/v1/protected/dailies');
    final response = await http.get(
      uri,
      headers: {'Authorization': 'Bearer $jwt'},
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to fetch challenges: ${response.statusCode}');
    }

    /// JSON parsing
    final body = jsonDecode(response.body) as Map<String, dynamic>;
    final items = body['challenges'] as List<dynamic>? ?? [];

    /// map into list of challenges
    return items
        .map((e) => Challenge.fromJson(e as Map<String, dynamic>))
        .toList();
  }
}
