import 'dart:convert';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;

class AuthStatus {
  final bool authenticated;

  const AuthStatus({
    required this.authenticated,
  });

  factory AuthStatus.fromJson(Map<String, dynamic> json) {
    return AuthStatus(
      authenticated: json['authenticated'] == 'true',
    );
  }

  static Future<AuthStatus> fetchAuthStatus(String jwt) async {
    final backendUrl = dotenv.get('BACKEND_URL', fallback: "http://10.0.2.2:8080");
    final response = await http.get(
      Uri.parse('$backendUrl/api/v1/protected/'),
      headers: {
        'Authorization': 'Bearer $jwt',
      },
    );
    if (response.statusCode == 200) {
      return AuthStatus.fromJson(
        jsonDecode(response.body) as Map<String, dynamic>,
      );
    } else {
      throw Exception('Failed to fetch auth status');
    }
  }
  @override
  String toString() {
    return """
      authenticated: $authenticated,
    """;
  }
}
