import 'dart:async';
import 'package:flutter/material.dart';
import '/widgets/widgets.dart';
import '/constants/constants.dart';

class RunPage extends StatefulWidget {
  const RunPage({super.key, required this.title});

  final String title;

  @override
  State<RunPage> createState() => _RunPageState();
}

class _RunPageState extends State<RunPage> {
  final int dailyGoal = 2000;
  int currentSteps = 1000;
  String selectedButton = 'Steps';
  Timer? _timer;

  @override
  void initState() {
    super.initState();
    // Counter fills RAM quickly - even when testing other pages
    // Uncomment the line below to start the step counter
    // _startStepCounter();
  }

  @override
  void dispose() {
    _timer?.cancel();
    super.dispose();
  }

  void _startStepCounter() {
    _timer = Timer.periodic(Duration(seconds: 1), (timer) {
      setState(() {
        currentSteps += 10;
      });
    });
  }

  void _onButtonPressed(String button) {
    setState(() {
      selectedButton = button;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: ColorConstants.deepBlue,
      body: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: <Widget>[
          if (selectedButton == 'Steps') ...[
            StepCounterWidget(currentSteps: currentSteps, dailyGoal: dailyGoal),
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
    );
  }
}
