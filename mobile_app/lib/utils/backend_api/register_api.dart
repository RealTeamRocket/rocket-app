import 'dart:convert';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;

class RegisterResponse {
  final String message;

  const RegisterResponse({
    required this.message,
  });

  factory RegisterResponse.fromJson(Map<String, dynamic> json) {
    return RegisterResponse(
      message: json['message'],
    );
  }
}

class RegisterApi {
  static Future<RegisterResponse> register(String email, String username, String password) async {
    final backendUrl = dotenv.get('BACKEND_URL', fallback: "http://10.0.2.2:8080");
    final response = await http.post(
      Uri.parse('$backendUrl/api/v1/register'),
      headers: {
        'Content-Type': 'application/json',
      },
      body: jsonEncode({
        'email': email,
        'username': username,
        'password': password,
      }),
    );

    if (response.statusCode == 200) {
      return RegisterResponse.fromJson(
        jsonDecode(response.body) as Map<String, dynamic>,
      );
    } else if (response.statusCode == 400) {
      throw Exception('Email already exists');
    } else {
      throw Exception('Failed to register');
    }
  }
}
