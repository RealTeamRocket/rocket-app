import 'dart:async';
import 'package:flutter/material.dart';
import '/widgets/widgets.dart';
import '/utils/utils.dart';
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
  late Timer _timer;

  @override
  void initState() {
    super.initState();
    _startStepCounter();
  }

  @override
  void dispose() {
    _timer.cancel();
    super.dispose();
  }

  void _startStepCounter() {
    _timer = Timer.periodic(Duration(seconds: 1), (timer) {
      setState(() {
        currentSteps += 10;
      });
    });
  }

  Color _getProgressColor(double progress) {
    if (progress >= 1.0) {
      return Colors.green;
    } else if (progress >= 0.5) {
      return Colors.orange;
    } else {
      return Colors.red;
    }
  }

  @override
  Widget build(BuildContext context) {
    double progress = currentSteps / dailyGoal;

    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text(widget.title),
      ),
      backgroundColor: ColorConstants.backgroundColor,
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Stack(
              alignment: Alignment.center,
              children: [
                SizedBox(
                  width: 300.0,
                  height: 300.0,
                  child: CustomPaint(
                    painter: CircularProgressPainter(
                      progress: progress,
                      color: _getProgressColor(progress),
                      strokeWidth: 20.0,
                    ),
                  ),
                ),
                Text(
                  '$currentSteps',
                  style: TextStyle(
                    fontSize: 50.0,
                    fontWeight: FontWeight.bold,
                    color: ColorConstants.white,
                  ),
                ),
              ],
            ),
          ],
        ),
      ),
      bottomNavigationBar: const CustomMenuBar(),
    );
  }
}
