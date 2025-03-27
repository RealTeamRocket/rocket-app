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
  final int totalSteps = 2000;
  final int currentSteps = 1000;

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
            Stack(
              alignment: Alignment.center,
              children: [
                Container(
                  width: 300.0,
                  height: 300.0,
                  child: CustomPaint(
                    painter: CircularProgressPainter(
                      progress: progress,
                      color: ColorConstants.redColor,
                      strokeWidth: 10.0,
                    ),
                  ),
                ),
                Text(
                  '$currentSteps',
                  style: TextStyle(
                    fontSize: 42.0,
                    fontWeight: FontWeight.bold,
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
