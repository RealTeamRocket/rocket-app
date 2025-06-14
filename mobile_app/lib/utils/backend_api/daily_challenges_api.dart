import 'dart:convert';
import 'package:mobile_app/utils/backend_api/base_api.dart';

/// Model for single challenge
class Challenge {
  final String id;
  final String text;
  final int points;

  const Challenge({required this.id, required this.text, required this.points});

  factory Challenge.fromJson(Map<String, dynamic> json) => Challenge(
    id: json['id'] as String,
    text: json['text'] as String,
    points: json['points'] as int,
  );
}

/// Model for daily challenge progress
class DailyChallengeProgress {
  final int completed;
  final int total;

  const DailyChallengeProgress({required this.completed, required this.total});

  factory DailyChallengeProgress.fromJson(Map<String, dynamic> json) =>
      DailyChallengeProgress(
        completed: json['completed'] as int,
        total: json['total'] as int,
      );
}

/// API-Service: Returning list of challenges
class ChallengesApi {
  static Future<List<Challenge>> fetchChallenges(String jwt) async {
    final response = await BaseApi.get(
      '/api/v1/protected/challenges/new',
      headers: {'Authorization': 'Bearer $jwt'},
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to fetch challenges: ${response.statusCode}');
    }

    final items = jsonDecode(response.body) as List<dynamic>;

    /// map into list of challenges
    return items
        .map((e) => Challenge.fromJson(e as Map<String, dynamic>))
        .toList();
  }

  static Future<void> markAsDone(
    String jwt,
    String challengeId,
    int rocketPoints,
  ) async {
    final response = await BaseApi.post(
      '/api/v1/protected/challenges/complete',
      headers: {
        'Authorization': 'Bearer $jwt',
        'Content-Type': 'application/json',
      },
      body: {'challenge_id': challengeId, 'rocket_points': rocketPoints},
    );

    if (response.statusCode != 200) {
      throw Exception(
        'Failed to mark challenge as done: ${response.statusCode}',
      );
    }
  }

  static Future<DailyChallengeProgress> fetchChallengeProgress(
    String jwt,
  ) async {
    final response = await BaseApi.get(
      '/api/v1/protected/challenges/progress',
      headers: {'Authorization': 'Bearer $jwt'},
    );

    if (response.statusCode != 200) {
      throw Exception(
        'Failed to fetch challenge progress: ${response.statusCode}',
      );
    }

    return DailyChallengeProgress.fromJson(
      jsonDecode(response.body) as Map<String, dynamic>,
    );
  }
}
