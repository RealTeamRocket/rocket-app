import 'package:flutter/material.dart';
import 'package:mobile_app/utils/backend_api/backend_api.dart' as api;
import '/widgets/widgets.dart';
import '/constants/constants.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:flutter_foreground_task/flutter_foreground_task.dart';

import 'tracking.dart';

class RunPage extends StatefulWidget {
  const RunPage({super.key, required this.title});
  final String title;

  @override
  State<RunPage> createState() => _RunPageState();
}

class _RunPageState extends State<RunPage> {
  int? dailyGoal;
  int currentSteps = 0;
  int? rocketPoints;

  final FlutterSecureStorage _secureStorage = const FlutterSecureStorage();

  late void Function(Object) _taskDataCallback;

  @override
  void initState() {
    super.initState();
    _loadRocketPoints();
    _loadStepGoal();

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
      final response = await api.RocketPointsApi.fetchRocketPoints(jwt);
      setState(() {
        rocketPoints = response.rocketPoints;
      });
    } catch (e) {
      setState(() {
        rocketPoints = null;
      });
    }
  }

  Future<void> _loadStepGoal() async {
    try {
      final jwt = await _secureStorage.read(key: 'jwt_token');
      if (jwt == null) return;
      final settings = await api.SettingsApi.getSettings(jwt);
      setState(() {
        dailyGoal = settings.stepGoal;
      });
    } catch (e) {
      setState(() {
        dailyGoal = null;
      });
    }
  }

  void _onRacePressed() {
    Navigator.push(
      context,
      MaterialPageRoute(builder: (context) => TrackingPage(title: 'Tracking')),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: ColorConstants.primaryColor,
      body: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: <Widget>[
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
          dailyGoal != null
              ? StepCounterWidget(currentSteps: currentSteps, dailyGoal: dailyGoal!)
              : const CircularProgressIndicator(),
          const SizedBox(height: 20.0),

          /// Centered Race Button
          Center(
            child: ElevatedButton(
              onPressed: _onRacePressed,
              style: ElevatedButton.styleFrom(
                backgroundColor: ColorConstants.greenColor,
                foregroundColor: ColorConstants.white,
                padding: const EdgeInsets.symmetric(
                  horizontal: 30.0,
                  vertical: 15.0,
                ),
                textStyle: const TextStyle(
                  fontSize: 20.0,
                  fontWeight: FontWeight.bold,
                ),
                shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(7.0),
                ),
              ),
              child: Row(
                mainAxisSize: MainAxisSize.min,
                children: const [
                  Icon(Icons.directions_run, size: 24.0),
                  SizedBox(width: 8.0),
                  Text("Run"),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}
