import 'package:flutter/material.dart';
import '/widgets/widgets.dart';
import '/constants/constants.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:flutter_foreground_task/flutter_foreground_task.dart';

import '../utils/backend_api/rocketpoints_api.dart';

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
  int? rocketPoints;

  final FlutterSecureStorage _secureStorage = const FlutterSecureStorage();

  late void Function(Object) _taskDataCallback;

  @override
  void initState() {
    super.initState();
    _loadRocketPoints();

    _taskDataCallback = (Object data) {
      if (data is int) {
        setState(() {
          currentSteps = data;
        });
      }
    };

    FlutterForegroundTask.addTaskDataCallback(_taskDataCallback);

    // Actively request the current steps from the foreground task
    FlutterForegroundTask.sendDataToTask('getCurrentSteps');
  }

  @override
  void dispose() {
    FlutterForegroundTask.removeTaskDataCallback(_taskDataCallback);
    super.dispose();
  }

  Future<void> _loadRocketPoints() async {
    try {
      final jwt = await _secureStorage.read(key: 'jwt_token');
      if (jwt == null) return;
      final response = await RocketPointsApi.fetchRocketPoints(jwt);
      setState(() {
        rocketPoints = response.rocketPoints;
      });
    } catch (e) {
      setState(() {
        rocketPoints = null;
      });
    }
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
                                color: ColorConstants.purpleColor.withOpacity(0.3),
                                width: 2.5,
                              ),
                              boxShadow: [
                                BoxShadow(
                                  color: ColorConstants.secoundaryColor.withOpacity(0.2),
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
                                rocketPoints != null
                                    ? '🚀 $rocketPoints RPs'
                                    : '🚀 ... RPs',
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
            StepCounterWidget(currentSteps: currentSteps, dailyGoal: dailyGoal),
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
