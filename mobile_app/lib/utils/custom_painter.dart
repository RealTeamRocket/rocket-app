import 'dart:math';
import 'package:flutter/material.dart';
import 'package:mobile_app/constants/constants.dart';

class CircularProgressPainter extends CustomPainter {
  final double progress;
  final Color color;
  final double strokeWidth;

  CircularProgressPainter({
    required this.progress,
    required this.color,
    required this.strokeWidth,
  });

  @override
  void paint(Canvas canvas, Size size) {
    final Paint backgroundPaint =
        Paint()
          ..color = ColorConstants.greyColor
          ..strokeWidth = strokeWidth
          ..style = PaintingStyle.stroke
          ..strokeCap = StrokeCap.round;

    final Paint progressPaint =
        Paint()
          ..color = color
          ..strokeWidth = strokeWidth
          ..style = PaintingStyle.stroke
          ..strokeCap = StrokeCap.round;

    final double radius = (size.width / 2) - (strokeWidth / 2);
    final Offset center = Offset(size.width / 2, size.height / 2);
    final double startAngle = 3 * pi / 4;
    final double sweepAngle = 3 * pi / 2 * progress;

    // Draw the background arc
    if (progress <= 1.00) {
      canvas.drawArc(
        Rect.fromCircle(center: center, radius: radius),
        startAngle,
        3 * pi / 2,
        false,
        backgroundPaint,
      );
    }

    // Draw the progress arc

    if (progress <= 1.00) {
      canvas.drawArc(
        Rect.fromCircle(center: center, radius: radius),
        startAngle,
        sweepAngle,
        false,
        progressPaint,
      );
    } else {
      canvas.drawArc(
        Rect.fromCircle(center: center, radius: radius),
        startAngle,
        3 * pi / 2,
        false,
        progressPaint,
      );
    }

    // Draw the text "Steps" at the bottom center
    final TextPainter textPainter = TextPainter(
      text: TextSpan(
        text: 'Steps',
        style: TextStyle(
          color: ColorConstants.white,
          fontSize: 40.0,
          fontWeight: FontWeight.bold,
        ),
      ),
      textDirection: TextDirection.ltr,
    );
    textPainter.layout();
    final double textX = center.dx - textPainter.width / 2;
    final double textY = center.dy + radius - textPainter.height / 2 - 30;
    textPainter.paint(canvas, Offset(textX, textY));
  }

  @override
  bool shouldRepaint(covariant CustomPainter oldDelegate) {
    return true;
  }
}
