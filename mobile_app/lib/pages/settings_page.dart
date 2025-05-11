// filepath: /Users/ron/dev/rocket-app/mobile_app/lib/pages/settings_page.dart
import 'package:flutter/material.dart';
import 'package:mobile_app/constants/constants.dart';
import 'package:mobile_app/utils/backend_api/settings_api.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class SettingsPage extends StatefulWidget {
  const SettingsPage({super.key});

  @override
  State<SettingsPage> createState() => _SettingsPageState();
}

class _SettingsPageState extends State<SettingsPage> {
  final TextEditingController _stepGoalController = TextEditingController();
  bool _isLoading = false;
  String? _errorMessage;
  final _storage = FlutterSecureStorage();

  @override
  void initState() {
    super.initState();
    _fetchSettings();
  }

  Future<void> _fetchSettings() async {
    setState(() {
      _isLoading = true;
      _errorMessage = null;
    });

    try {
      final jwt = await _storage.read(key: 'jwt_token');
      if (jwt == null) {
        throw Exception('JWT token not found');
      }
      final settings = await SettingsApi.getSettings(jwt!);
      _stepGoalController.text = settings.stepGoal.toString();
    } catch (e) {
      setState(() {
        _errorMessage = 'Failed to load settings: $e';
      });
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  }

  Future<void> _updateSettings() async {
    setState(() {
      _isLoading = true;
      _errorMessage = null;
    });

    try {
      final jwt = await _storage.read(key: 'jwt_token');
      final stepGoal = int.tryParse(_stepGoalController.text) ?? 0;
      await SettingsApi.updateSettings(jwt!, stepGoal);
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Settings updated successfully')),
      );
    } catch (e) {
      setState(() {
        _errorMessage = 'Failed to update settings: $e';
      });
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Settings'),
        backgroundColor: ColorConstants.secoundaryColor,
      ),
      backgroundColor: ColorConstants.primaryColor,
      body: _isLoading
          ? const Center(child: CircularProgressIndicator())
          : Padding(
              padding: const EdgeInsets.all(16.0),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  if (_errorMessage != null)
                    Padding(
                      padding: const EdgeInsets.only(bottom: 16.0),
                      child: Text(
                        _errorMessage!,
                        style: const TextStyle(color: Colors.red),
                      ),
                    ),
                  const Text(
                    'Step Goal',
                    style: TextStyle(
                      fontSize: 18.0,
                      fontWeight: FontWeight.bold,
                      color: Colors.white,
                    ),
                  ),
                  const SizedBox(height: 8.0),
                  TextField(
                    controller: _stepGoalController,
                    keyboardType: TextInputType.number,
                    style: const TextStyle(color: Colors.white),
                    decoration: InputDecoration(
                      filled: true,
                      fillColor: ColorConstants.secoundaryColor,
                      border: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(12),
                        borderSide: BorderSide.none,
                      ),
                      hintText: 'Enter your step goal',
                      hintStyle: const TextStyle(color: Colors.white54),
                    ),
                  ),
                  const SizedBox(height: 16.0),
                  ElevatedButton(
                    style: ElevatedButton.styleFrom(
                      backgroundColor: ColorConstants.greenColor,
                      padding: const EdgeInsets.symmetric(vertical: 16.0),
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(12.0),
                      ),
                    ),
                    onPressed: _updateSettings,
                    child: const Text(
                      'Save',
                      style: TextStyle(
                        fontSize: 18.0,
                        fontWeight: FontWeight.bold,
                        color: Colors.white,
                      ),
                    ),
                  ),
                ],
              ),
            ),
    );
  }

  @override
  void dispose() {
    _stepGoalController.dispose();
    super.dispose();
  }
}