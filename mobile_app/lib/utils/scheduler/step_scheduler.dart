import 'package:background_fetch/background_fetch.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:shared_preferences/shared_preferences.dart';

import '../backend_api/daily_steps_api.dart';

// HEADLESS TASK
@pragma('vm:entry-point')
void backgroundFetchHeadlessTask(HeadlessTask task) async {
  String taskId = task.taskId;
  bool isTimeout = task.timeout;
  try {
    await dotenv.load(fileName: ".env");
  } catch (e) {
    debugPrint("Error loading .env file using fallback: $e");
  }
  // Enable integration testing with the Flutter Driver extension.
  if (isTimeout) {
    // This task has exceeded its allowed running-time.
    // You must stop what you're doing and immediately .finish(taskId)
    debugPrint("[BackgroundFetch] Headless task timed-out: $taskId");
    BackgroundFetch.finish(taskId);
    return;
  }
  debugPrint('[BackgroundFetch] Headless event received.');
  try {
    await sendData();
  } catch (e) {
    debugPrint("BackgroundFetch: Error in headless task: $e");
  }
  BackgroundFetch.finish(taskId);
}

class StepScheduler {
  static const String taskId = "stepSchedulerTask";

  static Future<void> initialize() async {
    // Configure Background Fetch
    try {
      var status = await BackgroundFetch.configure(
        BackgroundFetchConfig(
          minimumFetchInterval: 15,
          forceAlarmManager: false,
          stopOnTerminate: false,
          startOnBoot: true,
          enableHeadless: true,
          requiresBatteryNotLow: false,
          requiresCharging: false,
          requiresStorageNotLow: false,
          requiresDeviceIdle: false,
          requiredNetworkType: NetworkType.NONE,
        ),
        _onBackgroundFetch,
        _onBackgroundFetchTimeout,
      );
      debugPrint('[BackgroundFetch] configure success: $status');
    } on Exception catch (e) {
      debugPrint("[BackgroundFetch] configure ERROR: $e");
    }
  }
}

void _onBackgroundFetch(String taskId) async {
  debugPrint("[BackgroundFetch] Event received: $taskId at ${DateTime.now()}");
  if (taskId == "flutter_background_fetch") {
    try {
      await sendData();
    } catch (e) {
      debugPrint("BackgroundFetch: Error in task: $e");
    }
  }
  BackgroundFetch.finish(taskId);
}

void _onBackgroundFetchTimeout(String taskId) {
  // <-- Task timeout handler.
  // This task has exceeded its allowed running-time.  You must stop what you're doing and immediately .finish(taskId)
  debugPrint("[BackgroundFetch] TASK TIMEOUT taskId: $taskId");
  BackgroundFetch.finish(taskId);
}

Future<void> sendData() async {
  final prefs = await SharedPreferences.getInstance();
  final steps = prefs.getInt('currentSteps') ?? 0;
  final storage = FlutterSecureStorage();
  final jwt = await storage.read(key: 'jwt_token');

  try {
    await DailyStepsApi.sendDailySteps(steps, jwt ?? "");
    debugPrint("BackgroundFetch: Steps uploaded successfully: $steps steps.");
  } catch (e) {
    debugPrint("BackgroundFetch: Error uploading steps: $e");
  }
}
