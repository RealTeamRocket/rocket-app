import 'dart:async';
import 'package:flutter/material.dart';
import 'package:permission_handler/permission_handler.dart';
import 'package:pedometer/pedometer.dart';
import '/widgets/widgets.dart';
import '/constants/constants.dart';

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key, required this.title});

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  final int dailyGoal = 2000;
  int currentSteps = 0;
  String selectedButton = 'Steps';
  Stream<StepCount>? _stepCountStream;

  @override
  void initState() {
    super.initState();
    _requestActivityRecognitionPermission();
  }

  Future<void> _requestActivityRecognitionPermission() async {
    var status = await Permission.activityRecognition.status;
    if (!status.isGranted) {
      status = await Permission.activityRecognition.request();
    }
    if (status.isGranted) {
      _startStepCounter();
    } else {
      _showPermissionDeniedDialog();
    }
  }

  @override
  void dispose() {
    super.dispose();
  }

  void _startStepCounter() {
    _stepCountStream = Pedometer.stepCountStream;
    _stepCountStream?.listen(onStepCount, onError: onStepCountError);
  }

  void onStepCount(StepCount event) {
    setState(() {
      currentSteps = event.steps;
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
