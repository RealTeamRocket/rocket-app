import 'package:flutter/material.dart';
import 'utils/backend_api/backend_api.dart' as api;
import 'package:flutter_dotenv/flutter_dotenv.dart';

import 'pages/pages.dart';

void main() async {
  await dotenv.load(fileName: ".env");
  try {
    final healtStats = await api.HealthStats.fetchHealth();
    debugPrint("$healtStats");
  } catch (e) {
    debugPrint(e.toString());
  }
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Rocket App',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.blueGrey),
      ),
      home: const MyHomePage(title: 'Step Counter'),
    );
  }
}
