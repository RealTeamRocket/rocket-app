import 'package:background_fetch/background_fetch.dart';
import 'package:flutter/material.dart';
import 'package:mobile_app/utils/scheduler/step_scheduler.dart';
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
  runApp(MyApp());
  BackgroundFetch.registerHeadlessTask(backgroundFetchHeadlessTask);
  StepScheduler.initialize();
}

class MyApp extends StatefulWidget {
  const MyApp({super.key});

  @override
  _MyAppState createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
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

      final authStatus = await api.AuthStatus.fetchAuthStatus(jwt);
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
    );
  }
}
