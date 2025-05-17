import 'package:flutter/material.dart';
import '/utils/utils.dart';
import '/widgets/widgets.dart';
import '/constants/constants.dart';

class RunPage extends StatefulWidget {
  const RunPage({super.key, required this.title});
  final String title;

  @override
  State<RunPage> createState() => _RunPageState();
}

class _RunPageState extends State<RunPage> {
  final int dailyGoal = 10000;
  int currentSteps = 0;
  String selectedButton = 'Steps';

  late PedometerService _pedometerService;

  @override
  void initState() {
    super.initState();
    _pedometerService = PedometerService();
    _pedometerService.onStepsUpdated = (steps) {
      setState(() {
        currentSteps = steps;
      });
    };

    _pedometerService.onError = (msg) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text(msg)),
        );
      }
    };

    _pedometerService.init();
  }

  void _onButtonPressed(String button) {
    setState(() {
      selectedButton = button;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: ColorConstants.primaryColor,
      body: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: <Widget>[
          if (selectedButton == 'Steps') ...[
            /// RocketPoints-Card
            Column(
              children: [
                const SizedBox(height: 20.0),
                SizedBox(
                  height: 100,
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Center(
                        child: IntrinsicWidth(
                          child: Container(
                            decoration: BoxDecoration(
                              color: ColorConstants.secoundaryColor,
                              borderRadius: BorderRadius.circular(16.0),
                              border: Border.all(
                                color: ColorConstants.purpleColor.withValues(alpha: 0.3),
                                width: 2.5,
                              ),
                              boxShadow: [
                                BoxShadow(
                                  color: ColorConstants.secoundaryColor.withValues(alpha: 0.2),
                                  blurRadius: 6.0,
                                  offset: const Offset(0, 3),
                                ),
                              ],
                            ),
                            padding: const EdgeInsets.symmetric(
                              vertical: 16.0,
                              horizontal: 16.0,
                            ),
                            child: Center(
                              child: Text(
                                '🚀 100 RPs',
                                style: const TextStyle(
                                  fontSize: 28.0,
                                  fontWeight: FontWeight.bold,
                                  color: ColorConstants.greenColor,
                                ),
                              ),
                            ),
                          ),
                        ),
                      ),
                    ],
                  ),
                ),
                const SizedBox(height: 12.0),
              ],
            ),

            /// Progressbar
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
                color: ColorConstants.blackColor,
              ),
            ),
            const SizedBox(height: 20.0),
          ],

          /// Buttons
          ButtonsWidget(
            selectedButton: selectedButton,
            onButtonPressed: _onButtonPressed,
          ),
        ],
      ),
    );
  }
}
