import 'package:flutter/material.dart';
import '../constants/color_constants.dart';
import 'package:flutter_slidable/flutter_slidable.dart';

class LeaderboardsPage extends StatefulWidget {
  const LeaderboardsPage({super.key, required this.title});

  final String title;

  @override
  State<LeaderboardsPage> createState() => _LeaderboardsPageState();
}

class _LeaderboardsPageState extends State<LeaderboardsPage> {
  final List<String> _challenges = [
    'Do 100 Push-ups',
    'Drink 2 liter of water',
    'Do 30 minutes of Yoga',
    'Sleep 8 hours',
    'Do 100 Push-ups',
    'Drink 2 liter of water',
    'Do 30 minutes of Yoga',
    'Sleep 8 hours',
  ];

  int _completedCount = 0;
  int _initialChallengeCount = 0;

  @override
  void initState() {
    super.initState();
    _initialChallengeCount = _challenges.length;
  }

  @override
  Widget build(BuildContext context) {
    double progressValue = _initialChallengeCount == 0 ? 0 : _completedCount / _initialChallengeCount;

    return Container(
      color: ColorConstants.primaryColor,
      padding: const EdgeInsets.all(16.0),
      child: Column(
        children: [
          Expanded(
            child: ListView.builder(
              itemCount: _challenges.length,
              itemBuilder: (context, index) {
                return Slidable(
                  key: Key(_challenges[index]),

                  /// swipe to the right
                  startActionPane: ActionPane(
                    motion: const StretchMotion(),
                    extentRatio: 0.25,
                    children: [
                      SlidableAction(
                        onPressed: (_) {
                          setState(() {
                            _completedCount++;
                            _challenges.removeAt(index);
                          });
                        },
                        backgroundColor: ColorConstants.greenColor,
                        foregroundColor: Colors.white,
                        icon: Icons.check,
                        label: 'Done',
                        borderRadius: BorderRadius.circular(16),
                      ),
                    ],
                  ),

                  /// swipe to the left
                  endActionPane: ActionPane(
                    motion: const StretchMotion(),
                    extentRatio: 0.25,
                    children: [
                      SlidableAction(
                        onPressed: (_) {
                          setState(() {
                            _completedCount++;
                            _challenges.removeAt(index);
                          });
                        },
                        backgroundColor: ColorConstants.greenColor,
                        foregroundColor: Colors.white,
                        icon: Icons.check,
                        label: 'Done',
                        borderRadius: BorderRadius.circular(16),
                      ),
                    ],
                  ),

                  child: _buildChallengeCard(_challenges[index]),
                );
              },
            ),
          ),
          const SizedBox(height: 24.0),
          _buildProgressSection(progressValue),
        ],
      ),
    );
  }

  /// challenge Cards
  Widget _buildChallengeCard(String challengeText) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 8.0),
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
              color: ColorConstants.white,
            ),
          ),
        ),
      ),
    );
  }

  Widget _buildProgressSection(double progressValue) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.center,
      children: [
        Text(
          'Progress: ${(progressValue * 100).toInt()}%',
          style: const TextStyle(
            fontSize: 16,
            fontWeight: FontWeight.bold,
            color: ColorConstants.white,
          ),
        ),
        const SizedBox(height: 8.0),
        Container(
          decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(20.0),
            border: Border.all(
              color: ColorConstants.white,
              width: 2.0,
            ),
          ),
          child: ClipRRect(
            borderRadius: BorderRadius.circular(20.0),
            child: LinearProgressIndicator(
              value: progressValue,
              minHeight: 12.0,
              backgroundColor: ColorConstants.primaryColor.withValues(alpha: 0.3),
              valueColor: const AlwaysStoppedAnimation<Color>(
                ColorConstants.greenColor,
              ),
            ),
          ),
        ),
      ],
    );
  }
}