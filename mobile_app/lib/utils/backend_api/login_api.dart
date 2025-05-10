import 'dart:convert';
import 'base_api.dart';

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
    final response = await BaseApi.post(
      '/api/v1/login',
      body: {
        'email': email,
        'password': password,
      },
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
