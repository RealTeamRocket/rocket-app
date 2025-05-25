import 'base_api.dart';

class TrackingApi {
  static Future<void> addRun(String jwt, String lineString, String duration, double distance) async {
    final response = await BaseApi.post(
      '/api/v1/protected/runs',
      headers: {'Authorization': 'Bearer $jwt'},
      body: {
        'route': lineString,
        'duration': duration,
        'distance': distance,
      },
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to upload run');
    }
  }
}
