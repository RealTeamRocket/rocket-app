import 'dart:async';
import 'package:flutter/material.dart';
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
  int currentSteps = 1000;
  String selectedButton = 'Steps';
  Timer? _timer;

  @override
  void initState() {
    super.initState();
    _startStepCounter();
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
      appBar: AppBar(
        backgroundColor: ColorConstants.greyColor,
        title: Center(
          child: Text(
            widget.title,
            style: TextStyle(color: ColorConstants.white),
          ),
        ),
      ),
      backgroundColor: ColorConstants.white,
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
