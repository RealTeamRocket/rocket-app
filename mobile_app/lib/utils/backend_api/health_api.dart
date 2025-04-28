import 'dart:convert';
import 'base_api.dart';

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
}

class HealthApi {
  static Future<HealthStats> fetchHealth() async {
    final response = await BaseApi.get('/api/v1/health');

    if (response.statusCode == 200) {
      return HealthStats.fromJson(
        jsonDecode(response.body) as Map<String, dynamic>,
      );
    } else {
      throw Exception('Failed to fetch stats');
    }
  }
}
