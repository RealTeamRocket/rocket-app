import 'dart:convert';
import 'package:flutter_dotenv/flutter_dotenv.dart';

import 'package:http/http.dart' as http;

class HealthStats {
  final int idle;
  final int inUse;
  final int maxIdleClosed;
  final int maxLifetimeClosed;
  final String message;
  final int openConnections;
  final String status;
  final int waitCount;
  final String waitDuration;

  const HealthStats({
    required this.idle,
    required this.inUse,
    required this.maxIdleClosed,
    required this.maxLifetimeClosed,
    required this.message,
    required this.openConnections,
    required this.status,
    required this.waitCount,
    required this.waitDuration,
  });

  factory HealthStats.fromJson(Map<String, dynamic> json) {
    return HealthStats(
      idle: int.parse(json['idle']),
      inUse: int.parse(json['in_use']),
      maxIdleClosed: int.parse(json['max_idle_closed']),
      maxLifetimeClosed: int.parse(json['max_lifetime_closed']),
      message: json['message'],
      openConnections: int.parse(json['open_connections']),
      status: json['status'],
      waitCount: int.parse(json['wait_count']),
      waitDuration: json['wait_duration'],
    );
  }

  static Future<HealthStats> fetchHealth() async {
    final backendUrl = dotenv.get('BACKEND_URL', fallback: "http://10.0.2.2:8080");
    final response = await http.get(
      Uri.parse('$backendUrl/api/v1/health'),
    );
    if (response.statusCode == 200) {
      return HealthStats.fromJson(
        jsonDecode(response.body) as Map<String, dynamic>,
      );
    } else {
      throw Exception('Failed to fetch stats');
    }
  }

  @override
  String toString() {
    return """
      idle: $idle,
      inUse: $inUse,
      maxIdleClosed: $maxIdleClosed,
      maxLifetimeClosed: $maxLifetimeClosed,
      message: $message,
      openConnections: $openConnections,
      status: $status,
      waitCount: $waitCount,
      waitDuration: $waitDuration,
    """;
  }
}
