import 'package:flutter/material.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import '../constants/color_constants.dart';
import 'package:flutter_slidable/flutter_slidable.dart';

import '../utils/backend_api/daily_challenges_api.dart';

class LeaderboardsPage extends StatefulWidget {
  const LeaderboardsPage({super.key, required this.title});

  final String title;

  @override
  State<LeaderboardsPage> createState() => _LeaderboardsPageState();
}

class _LeaderboardsPageState extends State<LeaderboardsPage> {
  List<Challenge> _challenges = [];
  int _completedCount = 0;
  int _initialChallengeCount = 0;


  @override
  void initState() {
    super.initState();
    _loadChallenges();
  }

  Future<void> _loadChallenges() async {
    try {
      final jwt = await FlutterSecureStorage().read(key: 'jwt_token');
      if (jwt == null) {
        throw Exception('JWT not found');
      }

      final challenges = await ChallengesApi.fetchChallenges(jwt);
      setState(() {
        _challenges = challenges;
        _initialChallengeCount = challenges.length;
      });
    } catch (e) {
      debugPrint('Error loading challenges: $e');
    }
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
                final challenge = _challenges[index];
                return Slidable(
                  key: Key(challenge.id),

                  /// swipe to the right
                  startActionPane: ActionPane(
                    motion: const StretchMotion(),
                    extentRatio: 0.25,
                    children: [
                      SlidableAction(
                        onPressed: (_) {
                          setState(() {
                            _completedCount++;
                            markChallengeAsDone(challenge);
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
                            markChallengeAsDone(challenge);
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

                  child: _buildChallengeCard(challenge),
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

  Future<void> markChallengeAsDone(Challenge challenge) async {
    final jwt = await FlutterSecureStorage().read(key: 'jwt_token');
    if (jwt == null) return;

    try {
      await ChallengesApi.markAsDone(jwt, challenge.id, challenge.points);
      await _loadChallenges();
      if (!mounted) return;
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          content: Text('Challenge done!'),
          duration: const Duration(seconds: 2),
        ),
      );
    } catch (e) {
      debugPrint('Failure at completing challenge: $e');
    }
  }


  /// challenge Cards
  Widget _buildChallengeCard(Challenge challenge) {
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
            challenge.text,
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