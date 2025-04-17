import 'dart:convert';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;

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
    final backendUrl = dotenv.get('BACKEND_URL', fallback: "http://10.0.2.2:8080");

    final response = await http.post(
      Uri.parse('$backendUrl/api/v1/protected/updateSteps'),
      headers: {
        'Authorization': 'Bearer $jwt',
      },
      body: jsonEncode({
        'steps': steps,
      }),
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
