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
import 'pages/pages.dart' as pages;
import 'package:permission_handler/permission_handler.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  try {
    await dotenv.load(fileName: ".env");
  } catch (e) {
    debugPrint("Error loading .env file using fallback: $e");
  }

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
              ? const pages.AppNavigator(title: 'Rocket App')
              : const pages.WelcomePage(),
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
  int _stepsToday = 0;
  int? _baselineSteps;
  late DateTime _lastResetDate;

  @override
  Future<void> onStart(DateTime timestamp, TaskStarter starter) async {
    await dotenv.load(fileName: ".env");
    final prefs = await SharedPreferences.getInstance();

    final dateString = prefs.getString('lastResetDate');
    if (dateString != null) {
      _lastResetDate = DateTime.parse(dateString);
    } else {
      _lastResetDate = DateTime(timestamp.year, timestamp.month, timestamp.day);
      await prefs.setString('lastResetDate', _lastResetDate.toIso8601String());
    }

    _baselineSteps = prefs.getInt('baselineSteps');

    int? lastSteps = prefs.getInt('lastPedometerSteps');
    if (_baselineSteps != null && lastSteps != null) {
      _stepsToday = lastSteps - _baselineSteps!;
    } else {
      _stepsToday = 0;
    }

    FlutterForegroundTask.sendDataToMain(_stepsToday);

    _sub = Pedometer.stepCountStream.listen(
      _onStepCount,
      onError: (e) => debugPrint('Pedometer error in service: $e'),
    );
  }

  Future<void> _onStepCount(StepCount event) async {
    debugPrint('Step event received: ${event.steps} at ${event.timeStamp}');
    final now = event.timeStamp;
    final prefs = await SharedPreferences.getInstance();

    await prefs.setInt('lastPedometerSteps', event.steps);

    if (!_isSameDay(now, _lastResetDate)) {
      _baselineSteps = event.steps;
      _stepsToday = 0;
      _lastResetDate = DateTime(now.year, now.month, now.day);
      await prefs.setString('lastResetDate', _lastResetDate.toIso8601String());
      await prefs.setInt('baselineSteps', _baselineSteps!);
    }

    if (_baselineSteps == null) {
      _baselineSteps = event.steps;
      await prefs.setInt('baselineSteps', _baselineSteps!);
    }
    _stepsToday = event.steps - _baselineSteps!;

    await prefs.setInt('currentSteps', _stepsToday);
    FlutterForegroundTask.updateService(
      notificationTitle: 'Steps Today',
      notificationText: '$_stepsToday',
    );

    FlutterForegroundTask.sendDataToMain(_stepsToday);
  }

  @override
  Future<void> onRepeatEvent(DateTime timestamp) async {
    try {
      final prefs = await SharedPreferences.getInstance();
      final steps = prefs.getInt('currentSteps') ?? 0;
      final jwt = prefs.getString('jwt_token');
      debugPrint("These values get send to backend: $jwt, $steps");
      if (jwt != null && jwt.isNotEmpty) {
        await api.DailyStepsApi.sendDailySteps(steps, jwt);
        debugPrint('Sent steps to backend: $steps');
      } else {
        debugPrint('JWT not found in SharedPreferences, skipping step sync.');
      }
    } catch (e) {
      debugPrint('Error sending steps to backend: $e');
    }
  }

  @override
  Future<void> onDestroy(DateTime timestamp, bool isTimeout) async {
    await _sub.cancel();
  }

  @override
  void onReceiveData(Object data) async {
    if (data is String && data == 'getCurrentSteps') {
      FlutterForegroundTask.sendDataToMain(_stepsToday);
    } else if (data is Map && data.containsKey('jwt_token')) {
      debugPrint("this is the data from the task ${data['jwt_token']}");
      final prefs = await SharedPreferences.getInstance();
      await prefs.setString('jwt_token', data['jwt_token']);
      debugPrint('JWT saved to SharedPreferences in background isolate.');
    }
  }

  @override
  void onNotificationButtonPressed(String id) {}

  @override
  void onNotificationPressed() {}

  @override
  void onNotificationDismissed() {}

  bool _isSameDay(DateTime a, DateTime b) {
    return a.year == b.year && a.month == b.month && a.day == b.day;
  }
}

Future<void> _requestPermissions() async {
  final NotificationPermission notificationPermission =
      await FlutterForegroundTask.checkNotificationPermission();
  if (notificationPermission != NotificationPermission.granted) {
    await FlutterForegroundTask.requestNotificationPermission();
  }

  if (Platform.isAndroid) {
    if (await Permission.activityRecognition.isDenied) {
      await Permission.activityRecognition.request();
    }
    if (await Permission.activityRecognition.isPermanentlyDenied) {
      await openAppSettings();
    }
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
      eventAction: ForegroundTaskEventAction.repeat(900000), // 15 minutes in ms
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
