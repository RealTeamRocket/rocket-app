import 'dart:convert';
import 'dart:io';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;
import 'package:path/path.dart';

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
      'X-API-KEY': apiKey,
      'Content-Type': 'application/json',
      ...?headers,
    };
    final requestBody = body is String ? body : jsonEncode(body);
    return await http.post(uri, headers: defaultHeaders, body: requestBody);
  }

  static Future<http.Response> delete(String endpoint, {Map<String, String>? headers, Object? body}) async {
    final uri = Uri.parse('$backendUrl$endpoint');
    final defaultHeaders = {
      'X-API-KEY': apiKey,
      'Content-Type': 'application/json',
      ...?headers,
    };
    return await http.delete(uri, headers: defaultHeaders, body: jsonEncode(body));
  }

  static Future<http.StreamedResponse> postMultipart(String endpoint, {
    Map<String, String>? headers,
    Map<String, String>? fields,
    File? file,
    String? fileFieldName = 'file',}) async {
    final uri = Uri.parse('$backendUrl$endpoint');
    final request = http.MultipartRequest('POST', uri);

    request.headers.addAll({
      'X-API-KEY': apiKey,
      ...?headers,
    });

    if (fields != null) {
      request.fields.addAll(fields);
    }

    // Add file if provided
    if (file != null) {
      final fileName = basename(file.path);
      final fileStream = http.ByteStream(file.openRead());
      final fileLength = await file.length();

      request.files.add(
        http.MultipartFile(
          fileFieldName ?? 'file',
          fileStream,
          fileLength,
          filename: fileName,
        ),
      );
    }

    // Send the request
    return await request.send();
  }
}
