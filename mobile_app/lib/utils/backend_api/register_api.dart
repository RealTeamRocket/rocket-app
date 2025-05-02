import 'dart:convert';
import 'base_api.dart';

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
    final response = await BaseApi.post(
      '/api/v1/register',
      body: {
        'email': email,
        'username': username,
        'password': password,
      },
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
