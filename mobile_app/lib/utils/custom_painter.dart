import 'dart:math';

import 'package:flutter/material.dart';

class CircularProgressPainter extends CustomPainter {
  final double progress; // Progress value between 0.0 and 1.0
  final Color color;
  final double strokeWidth;

  CircularProgressPainter({
    required this.progress,
    required this.color,
    required this.strokeWidth,
  });

  @override
  void paint(Canvas canvas, Size size) {
    final Paint paint = Paint()
      ..color = color
      ..strokeWidth = strokeWidth
      ..style = PaintingStyle.stroke
      ..strokeCap = StrokeCap.round;

    final double radius = (size.width / 2) - (strokeWidth / 2);
    final Offset center = Offset(size.width / 2, size.height / 2);
    final double startAngle = -pi / 2;
    final double sweepAngle = 2 * pi * progress;

    canvas.drawArc(
      Rect.fromCircle(center: center, radius: radius),
      startAngle,
      sweepAngle,
      false,
      paint,
    );
  }

  @override
  bool shouldRepaint(covariant CustomPainter oldDelegate) {
    return true;
  }
}
