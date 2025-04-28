import 'dart:convert';
import 'base_api.dart';

class StepsResponse {
  final bool success;

  const StepsResponse({
    required this.success,
  });

  factory StepsResponse.fromJson(Map<String, dynamic> json) {
    return StepsResponse(
      success: json['success'] ?? false,
    );
  }

  @override
  String toString() => "success: $success";
}

class DailyStepsApi {
  static Future<StepsResponse> sendDailySteps(int steps, String jwt) async {
    final response = await BaseApi.post(
      '/api/v1/protected/updateSteps',
      headers: {
        'Authorization': 'Bearer $jwt',
      },
      body: {
        'steps': steps,
      },
    );

    if (response.statusCode == 200) {
      return StepsResponse.fromJson(
        jsonDecode(response.body) as Map<String, dynamic>,
      );
    } else {
      throw Exception('Failed to send daily steps: ${response.statusCode}');
    }
  }
}
