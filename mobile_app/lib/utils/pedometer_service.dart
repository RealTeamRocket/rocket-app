import 'package:pedometer/pedometer.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:permission_handler/permission_handler.dart';

class PedometerService {
  int? _initialStepCount;
  DateTime? _initialStepDate;
  Stream<StepCount>? _stepCountStream;
  Function(int)? onStepsUpdated;
  Function(String)? onError;

  Future<void> init() async {
    await _requestActivityRecognitionPermission();
    await _loadInitialStepData();
    _startStepCounter();
  }

  Future<bool> _requestActivityRecognitionPermission() async {
    var status = await Permission.activityRecognition.status;

    if (!status.isGranted) {
      status = await Permission.activityRecognition.request();
    }

    if (!status.isGranted) {
      if (onError != null) {
        onError!("Permission denied. Step tracking is disabled.");
      }
      return false;
    }

    return true;
  }

  Future<void> _loadInitialStepData() async {
    final prefs = await SharedPreferences.getInstance();
    final savedInitialStep = prefs.getInt('initialStepCount');
    final savedDateString = prefs.getString('initialStepDate');

    if (savedDateString != null) {
      _initialStepDate = DateTime.tryParse(savedDateString);
    }

    if (_initialStepDate == null || !_isSameDay(_initialStepDate!, DateTime.now())) {
      _initialStepCount = null;
      _initialStepDate = DateTime.now();
      await prefs.remove('initialStepCount');
      await prefs.setString('initialStepDate', _initialStepDate!.toIso8601String());
    } else {
      _initialStepCount = savedInitialStep;
    }
  }

  void _startStepCounter() {
    _stepCountStream = Pedometer.stepCountStream;
    _stepCountStream?.listen(_onStepCount, onError: _onStepCountError);
  }

  void _onStepCount(StepCount event) async {
    DateTime eventTime = event.timeStamp;

    /// New day check
    if (_initialStepDate == null || !_isSameDay(_initialStepDate!, eventTime)) {
      _initialStepCount = event.steps;
      _initialStepDate = eventTime;
      final prefs = await SharedPreferences.getInstance();
      await prefs.setInt('initialStepCount', _initialStepCount!);
      await prefs.setString('initialStepDate', _initialStepDate!.toIso8601String());
    }

    /// First reading of the day if needed
    if (_initialStepCount == null) {
      _initialStepCount = event.steps;
      final prefs = await SharedPreferences.getInstance();
      await prefs.setInt('initialStepCount', _initialStepCount!);
    }

    final currentSteps = event.steps - _initialStepCount!;

    if (onStepsUpdated != null) {
      onStepsUpdated!(currentSteps);
    }
  }

  void _onStepCountError(error) {
    final errorMessage = "Step tracking failed: $error Please check permissions or restart the app.";
    print(errorMessage);

    if (onError != null) {
      onError!(errorMessage);
    }
  }

  bool _isSameDay(DateTime d1, DateTime d2) {
    return d1.year == d2.year && d1.month == d2.month && d1.day == d2.day;
  }
}
