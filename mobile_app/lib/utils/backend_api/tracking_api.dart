import 'base_api.dart';
import 'dart:convert';


class TrackingResponse {
  final String status;
  const TrackingResponse({
    required this.status,
  });

  factory TrackingResponse.fromJson(Map<String, dynamic> json) {
    return TrackingResponse(
      status: json['status'],
    );
  }
}

class TrackingApi {
  // AddRun-Funktion
  static Future<TrackingResponse> addRun(String userId, String lineString) async {
    final response = await BaseApi.post(
      '/api/v1/runs',
      body: {
        'user_id': userId,
        'linestring': lineString,
      },
    );

    if (response.statusCode == 200) {
      return TrackingResponse.fromJson(
        jsonDecode(response.body) as Map<String, dynamic>,
      );
    } else {
      throw Exception('Failed to add run: ${response.body}');
    }
  }

  // GetLatestRoute-Funktion
  static Future<TrackingResponse> getLatestRoute(String userId) async {
    final response = await BaseApi.get(
      '/api/v1/runs/latest?user_id=$userId',
    );

    if (response.statusCode == 200) {
      return TrackingResponse.fromJson(
        jsonDecode(response.body) as Map<String, dynamic>,
      );
    } else {
      throw Exception('Failed to get latest route: ${response.body}');
    }
  }

  // GetRouteByID-Funktion
  static Future<TrackingResponse> getRouteByID(String routeId, String userId) async {
    final response = await BaseApi.get(
      '/api/v1/runs/$routeId?user_id=$userId',
    );

    if (response.statusCode == 200) {
      return TrackingResponse.fromJson(
        jsonDecode(response.body) as Map<String, dynamic>,
      );
    } else {
      throw Exception('Failed to get route by ID: ${response.body}');
    }
  }

  // DeleteRunByID-Funktion
  static Future<void> deleteRunByID(String runId, String userId) async {
    final response = await BaseApi.delete(
      '/api/v1/runs/$runId',
      body: {'user_id': userId},
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to delete run: ${response.body}');
    }
  }
}
