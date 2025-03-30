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

  Color _getProgressColor(double progress) {
    return ColorConstants.purpleColor;
  }

  @override
  Widget build(BuildContext context) {
    double progress = currentSteps / dailyGoal;

    return Scaffold(
      appBar: AppBar(
        backgroundColor: ColorConstants.deepBlue,
        title: Text(
          widget.title,
          style: TextStyle(color: ColorConstants.white),
        ),
      ),
      backgroundColor: ColorConstants.deepBlue,
      body: Column(
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
          const SizedBox(height: 20.0),
          Container(
            padding: const EdgeInsets.symmetric(horizontal: 20.0),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                ElevatedButton(
                  onPressed: () {},
                  style: ElevatedButton.styleFrom(
                    padding: EdgeInsets.symmetric(
                      horizontal: 30.0,
                      vertical: 15.0,
                    ),
                    textStyle: TextStyle(
                      fontSize: 20.0,
                      fontWeight: FontWeight.bold,
                    ),
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(7.0),
                    ),
                  ),
                  child: Text("Steps"),
                ),
                SizedBox(width: 10.0),
                ElevatedButton(
                  onPressed: () {},
                  style: ElevatedButton.styleFrom(
                    padding: EdgeInsets.symmetric(
                      horizontal: 30.0,
                      vertical: 15.0,
                    ),
                    textStyle: TextStyle(
                      fontSize: 20.0,
                      fontWeight: FontWeight.bold,
                    ),
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(7.0),
                    ),
                  ),
                  child: Text("Race"),
                ),
              ],
            ),
          ),
        ],
      ),
      bottomNavigationBar: const CustomMenuBar(),
    );
  }
}
