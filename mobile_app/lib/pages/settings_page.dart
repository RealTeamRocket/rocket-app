import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';

class SettingsPage extends StatefulWidget {
  @override
  _SettingsPageState createState() => _SettingsPageState();
}

class _SettingsPageState extends State<SettingsPage> {
  final TextEditingController _usernameController = TextEditingController();
  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _stepGoalController = TextEditingController();

  @override
  void initState() {
    super.initState();
    _loadSettings();
  }

  Future<void> _loadSettings() async {
    final prefs = await SharedPreferences.getInstance();
    setState(() {
      _usernameController.text = prefs.getString('username') ?? '';
      _emailController.text = prefs.getString('email') ?? '';
      _stepGoalController.text = prefs.getInt('stepGoal')?.toString() ?? '10000';
    });
  }

  Future<void> _saveSettings() async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.setString('username', _usernameController.text);
    await prefs.setString('email', _emailController.text);
    await prefs.setInt('stepGoal', int.parse(_stepGoalController.text));
    ScaffoldMessenger.of(context).showSnackBar(
      SnackBar(content: Text('Einstellungen gespeichert!')),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Einstellungen'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          children: [
            TextField(
              controller: _usernameController,
              decoration: InputDecoration(labelText: 'Nutzername'),
            ),
            TextField(
              controller: _emailController,
              decoration: InputDecoration(labelText: 'E-Mail'),
            ),
            TextField(
              controller: _stepGoalController,
              decoration: InputDecoration(labelText: 'Schrittziel'),
              keyboardType: TextInputType.number,
            ),
            SizedBox(height: 20),
            ElevatedButton(
              onPressed: _saveSettings,
              child: Text('Speichern'),
            ),
          ],
        ),
      ),
    );
  }
}