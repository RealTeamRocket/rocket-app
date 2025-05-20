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
  late DateTime _lastResetDate;

  @override
  Future<void> onStart(DateTime timestamp, TaskStarter starter) async {
    final prefs = await SharedPreferences.getInstance();

    // 1) Lade oder initialisiere das Datum des letzten Resets
    final dateString = prefs.getString('lastResetDate');
    if (dateString != null) {
      _lastResetDate = DateTime.parse(dateString);
    } else {
      _lastResetDate = DateTime(timestamp.year, timestamp.month, timestamp.day);
      await prefs.setString(
        'lastResetDate',
        _lastResetDate.toIso8601String(),
      );
    }

    // 2) Starte den Pedometer-Stream
    _sub = Pedometer.stepCountStream.listen(
      _onStepCount,
      onError: (e) => debugPrint('Pedometer error in service: $e'),
    );
  }

  Future<void> _onStepCount(StepCount event) async {
    final now = event.timeStamp;
    final prefs = await SharedPreferences.getInstance();

    // 3) Tageswechsel prüfen
    if (!_isSameDay(now, _lastResetDate)) {
      // Reset am Tagesanfang
      _steps = 0;
      _lastResetDate = DateTime(now.year, now.month, now.day);
      await prefs.setString(
        'lastResetDate',
        _lastResetDate.toIso8601String(),
      );
    }

    // 4) Schritte aktualisieren (bei absoluten Zählern ggf. delta subtract)
    _steps = event.steps;

    // 5) Persistenz und Notification-Update
    await prefs.setInt('currentSteps', _steps);
    FlutterForegroundTask.updateService(
      notificationTitle: 'Schritte heute',
      notificationText: '$_steps',
    );

    // 6) Optional: Daten ans Main-Isolat senden
    FlutterForegroundTask.sendDataToMain(_steps);
  }

  @override
  Future<void> onRepeatEvent(DateTime timestamp) async {
    // Nicht benötigt, da der Stream alle Updates liefert
  }

  @override
  Future<void> onDestroy(DateTime timestamp, bool isTimeout) async {
    await _sub.cancel();
  }

  @override
  void onReceiveData(Object data) {
    // Eingehende Daten vom UI-Isolat (falls benötigt)
  }

  @override
  void onNotificationButtonPressed(String id) {
    // Falls du Notification-Buttons nutzen willst
  }

  @override
  void onNotificationPressed() {
    // Wenn die Notification angetippt wird
  }

  @override
  void onNotificationDismissed() {
    // Wenn die Notification dismissed wird
  }

  bool _isSameDay(DateTime a, DateTime b) {
    return a.year == b.year && a.month == b.month && a.day == b.day;
  }
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
