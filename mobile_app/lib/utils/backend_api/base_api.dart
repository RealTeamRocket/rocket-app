import 'dart:convert';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;

class BaseApi {
  static final String backendUrl = dotenv.get('BACKEND_URL', fallback: "http://10.0.2.2:8080");
  static final String apiKey = dotenv.get('API_KEY', fallback: "default-api-key");

  static Future<http.Response> get(String endpoint, {Map<String, String>? headers}) async {
    final uri = Uri.parse('$backendUrl$endpoint');
    final defaultHeaders = {
      'X-API-KEY': apiKey,
      ...?headers,
    };
    return await http.get(uri, headers: defaultHeaders);
  }

  static Future<http.Response> post(String endpoint, {Map<String, String>? headers, Object? body}) async {
    final uri = Uri.parse('$backendUrl$endpoint');
    final defaultHeaders = {
      'x-api-key': apiKey,
      'Content-Type': 'application/json',
      ...?headers,
    };
    return await http.post(uri, headers: defaultHeaders, body: jsonEncode(body));
  }
}
