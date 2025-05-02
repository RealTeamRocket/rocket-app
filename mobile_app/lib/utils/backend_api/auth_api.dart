import 'dart:convert';
import 'base_api.dart';

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
}

class AuthApi  {
  static Future<AuthStatus> fetchAuthStatus(String jwt) async {
    final response = await BaseApi.get('/api/v1/protected/', headers: {
      'Authorization': 'Bearer $jwt',
    });

    if (response.statusCode == 200) {
      return AuthStatus.fromJson(
        jsonDecode(response.body) as Map<String, dynamic>,
      );
    } else {
      throw Exception('Failed to fetch auth status');
    }
  }
}
