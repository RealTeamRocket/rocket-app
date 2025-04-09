import 'package:flutter/material.dart';
import '../constants/color_constants.dart';

class LeaderboardsPage extends StatefulWidget{
  const LeaderboardsPage({super.key, required this.title});

  final String title;

  @override
  State<LeaderboardsPage> createState() => _LeaderboardsPageState();

}
class _LeaderboardsPageState extends State<LeaderboardsPage> {

  /// daily challenges
  final List<String> _challenges = const [
    'Do 100 Push-ups',
    'Drink 2 liter of water',
    'Do 30 minutes of Yoga',
    'Sleep 8 hours',
    'Do 100 Push-ups',
    'Drink 2 liter of water',
    'Do 30 minutes of Yoga',
    'Sleep 8 hours',
  ];

  List<bool> _completed = [false, false, true, false, false, false, true, false];

  @override
  Widget build(BuildContext context) {

    final double progressValue = _completed.where((c) => c).length / _completed.length;

    return Container(
      color: ColorConstants.white,
      padding: const EdgeInsets.all(16.0),
      child: Column(
        children: [
          Expanded(
            child: ListView.builder(
              itemCount: _challenges.length,
              itemBuilder: (context, index) {
                return _buildChallengeCard(_challenges[index], index);
              },
            ),
          ),
          const SizedBox(height: 24.0),
          _buildProgressSection(progressValue),
        ],
      ),
    );


  }

  Widget _buildChallengeCard(String challengeText, int index) {
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
          trailing: Icon(
            _completed[index]
                ? Icons.check_circle
                : Icons.radio_button_unchecked,
            color: _completed[index]
                ? ColorConstants.greenColor
                : ColorConstants.greyColor,
            size: 32.0,
          ),
          onTap: () {
            setState(() {
              _completed[index] = !_completed[index];
            });
          },
        ),
      ),
    );
  }

  /// Progressbar
  Widget _buildProgressSection(double progressValue) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.center,
      children: [
        Text(
          'Progress: ${(progressValue * 100).toInt()}%',
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


