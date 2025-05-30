import 'base_api.dart';
import 'dart:convert';

class RunModel {
  final String id;
  final String route;
  final String duration;
  final double distance;
  final String createdAt;

  RunModel({
    required this.id,
    required this.route,
    required this.duration,
    required this.distance,
    required this.createdAt,
  });

  factory RunModel.fromJson(Map<String, dynamic> json) {
    return RunModel(
      id: json['id'] as String,
      route: json['route'] as String,
      duration: json['duration'] as String,
      distance: (json['distance'] as num).toDouble(),
      createdAt: json['created_at'] as String,
    );
  }
}

class TrackingApi {
  static Future<void> addRun(
    String jwt,
    String lineString,
    String duration,
    double distance,
  ) async {
    final response = await BaseApi.post(
      '/api/v1/protected/runs',
      headers: {'Authorization': 'Bearer $jwt'},
      body: {'route': lineString, 'duration': duration, 'distance': distance},
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to upload run');
    }
  }

  static Future<List<RunModel>> getAllRuns(String jwt) async {
    final response = await BaseApi.get(
      '/api/v1/protected/runs',
      headers: {'Authorization': 'Bearer $jwt'},
    );
    if (response.statusCode != 200) {
      throw Exception('Failed to fetch runs');
    }
    final List<dynamic> data = jsonDecode(response.body);
    return data.map((json) => RunModel.fromJson(json)).toList();
  }

  static Future<void> deleteRun(String jwt, String runId) async {
    final response = await BaseApi.delete(
      '/api/v1/protected/runs/$runId',
      headers: {'Authorization': 'Bearer $jwt'},
    );
    if (response.statusCode != 200) {
      throw Exception('Failed to delete run');
    }
  }
}
