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
  int currentSteps = 1000;
  late Timer _timer;
  bool isRunning = false;
  Duration elapsedTime = Duration.zero;

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
        if (isRunning) {
          elapsedTime += Duration(seconds: 1);
        }
      });
    });
  }

  void _toggleRun() {
    setState(() {
      if (isRunning) {
        isRunning = false;
      } else {
        isRunning = true;
        elapsedTime = Duration.zero;
      }
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

  String _formatDuration(Duration duration) {
    if (duration == Duration.zero && !isRunning) {
      return "";
    }
    String twoDigits(int n) => n.toString().padLeft(2, '0');
    String twoDigitMinutes = twoDigits(duration.inMinutes.remainder(60));
    String twoDigitSeconds = twoDigits(duration.inSeconds.remainder(60));
    return "${twoDigits(duration.inHours)}:$twoDigitMinutes:$twoDigitSeconds";
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
      body: Column(
        children: <Widget>[
          Expanded(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: <Widget>[
                Padding(
                  padding: const EdgeInsets.all(10.0),
                  child: Text(
                    _formatDuration(elapsedTime),
                    style: TextStyle(
                      fontSize: 30.0,
                      fontWeight: FontWeight.bold,
                      color: ColorConstants.white,
                    ),
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.all(10.0),
                  child: Container(
                    decoration: BoxDecoration(
                      gradient: LinearGradient(
                        colors: [Colors.blue, Colors.purple],
                      ),
                      borderRadius: BorderRadius.circular(10),
                    ),
                    child: SizedBox(
                      width: 200,
                      height: 60,
                      child: ElevatedButton(
                        style: ElevatedButton.styleFrom(
                          backgroundColor: Colors.transparent,
                          shadowColor: Colors.transparent,
                        ),
                        onPressed: _toggleRun,
                        child: Text(
                          isRunning ? 'Stop Run' : 'Start Run',
                          style: TextStyle(
                            color: ColorConstants.white,
                            fontSize: 20,
                          ),
                        ),
                      ),
                    ),
                  ),
                ),
              ],
            ),
          ),
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
          Expanded(
            child: Container(),
          ),
        ],
      ),
      bottomNavigationBar: const CustomMenuBar(),
    );
  }
}
