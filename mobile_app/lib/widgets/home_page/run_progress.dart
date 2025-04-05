import 'package:flutter/material.dart';

import '../../pages/pages.dart';
import '../menu_bar.dart';
import '../painters/painters.dart';

class MyHomePageState extends State<MyHomePage> {
  final int totalSteps = 2000;
  final int currentSteps = 900;

  @override
  Widget build(BuildContext context) {
    double progress = currentSteps / totalSteps;

    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text(widget.title),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            SizedBox(
              width: 300.0,
              height: 300.0,
              child: CustomPaint(
                painter: CircularProgressPainter(
                  progress: progress,
                  color: Colors.red,
                  strokeWidth: 6.0,
                ),
              ),
            ),
          ],
        ),
      ),
      bottomNavigationBar: const CustomMenuBar(),
    );
  }
}
