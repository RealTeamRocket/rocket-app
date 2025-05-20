import 'dart:async';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter_foreground_task/flutter_foreground_task.dart';
import 'package:mobile_app/pages/settings_page.dart';
import 'package:pedometer/pedometer.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'utils/backend_api/backend_api.dart' as api;
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:jwt_decoder/jwt_decoder.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'pages/pages.dart';

void main() async {
  try {
    await dotenv.load(fileName: ".env");
  } catch (e) {
    debugPrint("Error loading .env file using fallback: $e");
  }
  // BackgroundFetch.registerHeadlessTask(backgroundFetchHeadlessTask);
  // StepScheduler.initialize();
  FlutterForegroundTask.initCommunicationPort();
  await _requestPermissions();
  _initService();
  await _startService();

  runApp(MyApp());
}

class MyApp extends StatefulWidget {
  const MyApp({super.key});

  @override
  MyAppState createState() => MyAppState();
}

class MyAppState extends State<MyApp> {
  bool? isAuthenticated;
  final _storage = FlutterSecureStorage();

  @override
  void initState() {
    super.initState();
    checkIfAuthenticated();
  }

  Future<void> checkIfAuthenticated() async {
    try {
      final jwt = await _storage.read(key: 'jwt_token');
      if (jwt == null) {
        setState(() {
          isAuthenticated = false;
        });
        return;
      }

      if (JwtDecoder.isExpired(jwt)) {
        await _storage.delete(key: 'jwt_token');
        setState(() {
          isAuthenticated = false;
        });
        return;
      }

      final authStatus = await api.AuthApi.fetchAuthStatus(jwt);
      setState(() {
        isAuthenticated = authStatus.authenticated;
      });
    } catch (e) {
      debugPrint('Error fetching auth status: $e');
      setState(() {
        isAuthenticated = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    if (isAuthenticated == null) {
      return MaterialApp(
        home: Scaffold(body: Center(child: CircularProgressIndicator())),
      );
    }

    return MaterialApp(
      title: 'Rocket App',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.blueGrey),
      ),
      home:
          isAuthenticated!
              ? const HomePage(title: 'Rocket App')
              : const WelcomePage(),
      routes: {'/settings': (context) => SettingsPage()},
    );
  }
}

@pragma('vm:entry-point')
void startCallback() {
  FlutterForegroundTask.setTaskHandler(MyTaskHandler());
}

class MyTaskHandler extends TaskHandler {
  late StreamSubscription<StepCount> _sub;
  int _steps = 0;

  @override
  Future<void> onStart(DateTime timestamp, TaskStarter starter) async {
    // Erlaubnis prüfen, SharedPreferences laden etc.
    _sub = Pedometer.stepCountStream.listen((event) async {
      _steps = event.steps;
      final prefs = await SharedPreferences.getInstance();
      await prefs.setInt('currentSteps', _steps);

      FlutterForegroundTask.updateService(
        notificationTitle: 'Schritte',
        notificationText: '$_steps',
      );
    }, onError: (e) {
      debugPrint('Pedometer error in service: $e');
    });
  }

  @override
  Future<void> onDestroy(DateTime timestamp, bool isTimeout) async {
    await _sub.cancel();
  }

  Future<void> _incrementCount() async {
    // get pedometer data
    final prefs = await SharedPreferences.getInstance();
    _steps = prefs.getInt('currentSteps') ?? 0;
    debugPrint("Steps: $_steps");

    // Update notification content.
    FlutterForegroundTask.updateService(
      notificationTitle: 'Foreground Service Running',
      notificationText: 'Steps: $_steps',
    );

    // Send data to main isolate.
    FlutterForegroundTask.sendDataToMain(_steps);
  }

  @override
  Future<void> onRepeatEvent(DateTime timestamp) async {
    await _incrementCount();
  }

  @override
  void onReceiveData(Object data) {}

  @override
  void onNotificationButtonPressed(String id) {}

  @override
  void onNotificationPressed() {}

  @override
  void onNotificationDismissed() {}
}

// Permission and service helpers
Future<void> _requestPermissions() async {
  final NotificationPermission notificationPermission =
      await FlutterForegroundTask.checkNotificationPermission();
  if (notificationPermission != NotificationPermission.granted) {
    await FlutterForegroundTask.requestNotificationPermission();
  }

  if (Platform.isAndroid) {
    if (!await FlutterForegroundTask.isIgnoringBatteryOptimizations) {
      await FlutterForegroundTask.requestIgnoreBatteryOptimization();
    }
    if (!await FlutterForegroundTask.canScheduleExactAlarms) {
      await FlutterForegroundTask.openAlarmsAndRemindersSettings();
    }
  }
}

void _initService() {
  FlutterForegroundTask.init(
    androidNotificationOptions: AndroidNotificationOptions(
      channelId: 'foreground_service',
      channelName: 'Foreground Service Notification',
      channelDescription:
          'This notification appears when the foreground service is running.',
      onlyAlertOnce: true,
    ),
    iosNotificationOptions: const IOSNotificationOptions(
      showNotification: false,
      playSound: false,
    ),
    foregroundTaskOptions: ForegroundTaskOptions(
      eventAction: ForegroundTaskEventAction.repeat(1000), // 1 second
      autoRunOnBoot: true,
      autoRunOnMyPackageReplaced: true,
      allowWakeLock: true,
      allowWifiLock: true,
    ),
  );
}

Future<void> _startService() async {
  if (await FlutterForegroundTask.isRunningService) {
    await FlutterForegroundTask.restartService();
  } else {
    await FlutterForegroundTask.startService(
      serviceId: 256,
      notificationTitle: 'Foreground Service Running',
      notificationText: 'Steps: 0',
      notificationIcon: null,
      callback: startCallback,
    );
  }
}
