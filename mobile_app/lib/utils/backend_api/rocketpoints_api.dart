import 'dart:convert';
import 'base_api.dart';

class RocketPointsResponse {
  final int rocketPoints;

  const RocketPointsResponse({required this.rocketPoints});

  factory RocketPointsResponse.fromJson(Map<String, dynamic> json) {
    return RocketPointsResponse(rocketPoints: json['rocket_points'] ?? 0);
  }

  @override
  String toString() => "rocketPoints: $rocketPoints";
}

class RocketPointsApi {
  static Future<RocketPointsResponse> fetchRocketPoints(String jwt) async {
    final response = await BaseApi.get(
      '/api/v1/protected/user/rocketpoints',
      headers: {'Authorization': 'Bearer $jwt'},
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to fetch rocket points: ${response.statusCode}');
    } else {
      return RocketPointsResponse.fromJson(
        jsonDecode(response.body) as Map<String, dynamic>,
      );
    }
  }
}
