import 'dart:convert';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;

class LoginResponse {
  final String token;

  const LoginResponse({
    required this.token,
  });

  factory LoginResponse.fromJson(Map<String, dynamic> json) {
    return LoginResponse(
      token: json['token'],
    );
  }
}

class LoginApi {
  static Future<LoginResponse> login(String email, String password) async {
    final backendUrl = dotenv.get('BACKEND_URL', fallback: "http://10.0.2.2:8080");
    final response = await http.post(
      Uri.parse('$backendUrl/api/v1/login'),
      headers: {
        'Content-Type': 'application/json',
      },
      body: jsonEncode({
        'email': email,
        'password': password,
      }),
    );

    if (response.statusCode == 200) {
      return LoginResponse.fromJson(
        jsonDecode(response.body) as Map<String, dynamic>,
      );
    } else if (response.statusCode == 401) {
      throw Exception('Incorrect email or password');
    } else {
      throw Exception('Failed to login');
    }
  }
}
