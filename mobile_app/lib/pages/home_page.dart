import 'package:flutter/material.dart';
import 'package:pedometer/pedometer.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:permission_handler/permission_handler.dart';
import '/widgets/widgets.dart';
import '/constants/constants.dart';

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key, required this.title});
  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  final int dailyGoal = 10000;
  int currentSteps = 0;
  String selectedButton = 'Steps';
  Stream<StepCount>? _stepCountStream;

  int? _initialStepCount;
  DateTime? _initialStepDate;

  @override
  void initState() {
    super.initState();
    _requestActivityRecognitionPermission();
  }

  /// request permission to track steps
  Future<void> _requestActivityRecognitionPermission() async {
    var status = await Permission.activityRecognition.status;
    if (!status.isGranted) {
      status = await Permission.activityRecognition.request();
    }
    if (status.isGranted) {
      // Load initial step data from persistent storage
      await _loadInitialStepData();
      _startStepCounter();
    } else {
      _showPermissionDeniedDialog();
    }
  }

  Future<void> _loadInitialStepData() async {
    final prefs = await SharedPreferences.getInstance();
    final savedInitialStep = prefs.getInt('initialStepCount');
    final savedDateString = prefs.getString('initialStepDate');
    if (savedDateString != null) {
      _initialStepDate = DateTime.tryParse(savedDateString);
    }
    // If the saved date is not today, reset the initial step count
    if (_initialStepDate == null || !_isSameDay(_initialStepDate!, DateTime.now())) {
      _initialStepCount = null;
      _initialStepDate = DateTime.now();
      await prefs.remove('initialStepCount');
      await prefs.setString('initialStepDate', _initialStepDate!.toIso8601String());
    } else {
      _initialStepCount = savedInitialStep;
    }
  }

  bool _isSameDay(DateTime d1, DateTime d2) {
    return d1.year == d2.year && d1.month == d2.month && d1.day == d2.day;
  }

  void _startStepCounter() {
    _stepCountStream = Pedometer.stepCountStream;
    _stepCountStream?.listen(onStepCount, onError: onStepCountError);
  }

  void onStepCount(StepCount event) async {
    DateTime eventTime = event.timeStamp ?? DateTime.now();

    if (_initialStepDate == null || !_isSameDay(_initialStepDate!, eventTime)) {
      _initialStepCount = event.steps;
      _initialStepDate = eventTime;
      final prefs = await SharedPreferences.getInstance();
      await prefs.setInt('initialStepCount', _initialStepCount!);
      await prefs.setString('initialStepDate', _initialStepDate!.toIso8601String());
    }

    /// For the very first reading on app launch
    if (_initialStepCount == null) {
      _initialStepCount = event.steps;
      final prefs = await SharedPreferences.getInstance();
      await prefs.setInt('initialStepCount', _initialStepCount!);
    }

    setState(() {
      currentSteps = event.steps - _initialStepCount!;
    });
  }

  void onStepCountError(error) {
    debugPrint("Pedometer Error: $error");
  }

  void _showPermissionDeniedDialog() {
    showDialog(
      context: context,
      builder: (context) {
        return AlertDialog(
          title: Text("Permission Denied"),
          content: Text("Activity recognition permission is required to count your steps. Please grant the permission in settings."),
          actions: [
            TextButton(
              onPressed: () => Navigator.of(context).pop(),
              child: Text("OK"),
            ),
          ],
        );
      },
    );
  }

  void _onButtonPressed(String button) {
    setState(() {
      selectedButton = button;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: ColorConstants.deepBlue,
        title: Center(
          child: Text(
            widget.title,
            style: TextStyle(color: ColorConstants.white),
          ),
        ),
      ),
      backgroundColor: ColorConstants.deepBlue,
      body: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: <Widget>[
          if (selectedButton == 'Steps') ...[
            StepCounterWidget(
              currentSteps: currentSteps,
              dailyGoal: dailyGoal,
            ),
            const SizedBox(height: 20.0),
          ] else if (selectedButton == 'Race') ...[
            Text(
              'To be implemented',
              style: TextStyle(
                fontSize: 30.0,
                fontWeight: FontWeight.bold,
                color: ColorConstants.white,
              ),
            ),
            const SizedBox(height: 20.0),
          ],
          ButtonsWidget(
            selectedButton: selectedButton,
            onButtonPressed: _onButtonPressed,
          ),
        ],
      ),
      bottomNavigationBar: const CustomMenuBar(),
    );
  }
}
