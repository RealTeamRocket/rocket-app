import 'package:flutter/material.dart';
import '../constants/color_constants.dart';

class ChallengesPage extends StatefulWidget{
  const ChallengesPage({super.key, required this.title});

  final String title;

  @override
  State<ChallengesPage> createState() => _ChallengesPageState();

}
class _ChallengesPageState extends State<ChallengesPage> {

  /// daily challenges
  final List<String> _challenges = const [
    'Mache 10.000 Schritte',
    'Trinke 2 Liter Wasser',
    '30 Minuten Yoga',
    '8 Stunden Schlaf',
  ];

  @override
  Widget build(BuildContext context) {

    final double progressValue = 0.5;

    return Scaffold(
      appBar: AppBar(
        title: const Text('Tägliche Challenges'),
        backgroundColor: ColorConstants.deepBlue,
      ),
      backgroundColor: ColorConstants.white,
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          children: [
            // Liste der Challenges
            Expanded(
              child: ListView.builder(
                itemCount: _challenges.length,
                itemBuilder: (context, index) {
                  return _buildChallengeCard(_challenges[index]);
                },
              ),
            ),
            const SizedBox(height: 24.0),
            _buildProgressSection(progressValue),
          ],
        ),
      ),
    );

  }

  Widget _buildChallengeCard(String challengeText) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 8.0),
      child: Container(
        decoration: BoxDecoration(
          color: ColorConstants.white,
          borderRadius: BorderRadius.circular(16.0),
          border: Border.all(
            color: ColorConstants.greyColor.withOpacity(0.3),
          ),
          boxShadow: [
            BoxShadow(
              color: ColorConstants.greyColor.withOpacity(0.2),
              blurRadius: 6.0,
              offset: const Offset(0, 3),
            ),
          ],
        ),
        child: ListTile(
          contentPadding: const EdgeInsets.symmetric(
            horizontal: 20.0,
            vertical: 16.0,
          ),
          title: Text(
            challengeText,
            style: const TextStyle(
              fontSize: 18.0,
              fontWeight: FontWeight.w600,
              color: ColorConstants.deepBlue,
            ),
          ),
          trailing: const Icon(
            Icons.check_circle_outline,
            color: ColorConstants.purpleColor,
            size: 32.0,
          ),
          onTap: () {
            // switch state of challenge per tap or other way
          },
        ),
      ),
    );
  }

  /// Progressbar
  Widget _buildProgressSection(double progressValue) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          'Fortschritt: ${(progressValue * 100).toInt()}%',
          style: const TextStyle(
            fontSize: 16,
            fontWeight: FontWeight.bold,
            color: ColorConstants.deepBlue,
          ),
        ),
        const SizedBox(height: 8.0),
        ClipRRect(
          borderRadius: BorderRadius.circular(20.0),
          child: LinearProgressIndicator(
            value: progressValue,
            minHeight: 12.0,
            backgroundColor: ColorConstants.greyColor.withOpacity(0.3),
            valueColor: const AlwaysStoppedAnimation<Color>(
              ColorConstants.greenColor,
            ),
          ),
        ),
      ],
    );
  }
}


