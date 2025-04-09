import 'package:flutter/material.dart';
import '../painters/painters.dart';
import '/constants/constants.dart';

class StepCounterWidget extends StatelessWidget {
  final int currentSteps;
  final int dailyGoal;

  const StepCounterWidget({
    super.key,
    required this.currentSteps,
    required this.dailyGoal,
  });

  Color _getProgressColor(double progress) {
    return ColorConstants.purpleColor;
  }

  @override
  Widget build(BuildContext context) {
    double progress = currentSteps / dailyGoal;

    return Stack(
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
    );
  }
}
