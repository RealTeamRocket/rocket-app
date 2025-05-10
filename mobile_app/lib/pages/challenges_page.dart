import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import '../constants/color_constants.dart';
import 'package:flutter_slidable/flutter_slidable.dart';

import '../utils/backend_api/daily_challenges_api.dart';

class ChallengePage extends StatefulWidget {
  const ChallengePage({super.key, required this.title});

  final String title;

  @override
  State<ChallengePage> createState() => _ChallengePageState();
}

class _ChallengePageState extends State<ChallengePage>
    with SingleTickerProviderStateMixin {
  List<Challenge> _challenges = [];

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
      debugPrint(challenges.toString());
      setState(() {
        _challenges = challenges;
      });
    } catch (e) {
      debugPrint('Error loading challenges: $e');
    }
  }

  @override
  Widget build(BuildContext context) {
    const int totalChallenges = 5;
    double progressValue =
        totalChallenges == 0
            ? 0
            : ((totalChallenges - _challenges.length) / totalChallenges).clamp(
              0.0,
              1.0,
            );

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
                          HapticFeedback.mediumImpact();
                          final challengeToMark = challenge;
                          setState(() {
                            _challenges.remove(challengeToMark);
                          });
                          showPointsOverlay(challengeToMark.points.toString());
                          markChallengeAsDone(challengeToMark);
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
                          HapticFeedback.mediumImpact();
                          final challengeToMark = challenge;
                          setState(() {
                            _challenges.remove(challengeToMark);
                          });
                          showPointsOverlay(challengeToMark.points.toString());
                          markChallengeAsDone(challengeToMark);
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
      // await _loadChallenges(); // only needed if we want to use only backend to sync/show challenges
      if (!mounted) return;
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
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 20.0, vertical: 16.0),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Expanded(
                child: Text(
                  challenge.text,
                  style: const TextStyle(
                    fontSize: 18.0,
                    fontWeight: FontWeight.w600,
                    color: ColorConstants.white,
                  ),
                ),
              ),
              Text(
                '+${challenge.points} RP',
                style: const TextStyle(
                  fontSize: 16.0,
                  fontWeight: FontWeight.bold,
                  color: ColorConstants.greenColor,
                ),
              ),
            ],
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
            border: Border.all(color: ColorConstants.white, width: 2.0),
          ),
          child: ClipRRect(
            borderRadius: BorderRadius.circular(20.0),
            child: LinearProgressIndicator(
              value: progressValue,
              minHeight: 12.0,
              backgroundColor: ColorConstants.primaryColor.withValues(
                alpha: 0.3,
              ),
              valueColor: const AlwaysStoppedAnimation<Color>(
                ColorConstants.greenColor,
              ),
            ),
          ),
        ),
      ],
    );
  }

  void showPointsOverlay(String points) {
    final overlay = Overlay.of(context);
    late OverlayEntry overlayEntry;

    final controller = AnimationController(
      vsync: Navigator.of(context),
      duration: const Duration(milliseconds: 800),
    );

    final opacityAnimation = Tween<double>(
      begin: 1.0,
      end: 0.0,
    ).animate(CurvedAnimation(parent: controller, curve: Curves.easeOut));

    final slideAnimation = Tween<Offset>(
      begin: Offset.zero,
      end: const Offset(0, -0.5),
    ).animate(CurvedAnimation(parent: controller, curve: Curves.easeOut));

    overlayEntry = OverlayEntry(
      builder: (context) {
        return Positioned.fill(
          child: IgnorePointer(
            child: Center(
              child: AnimatedBuilder(
                animation: controller,
                builder:
                    (context, child) => Opacity(
                      opacity: opacityAnimation.value,
                      child: Transform.translate(
                        offset:
                            slideAnimation.value * 100, // Moves ~50px upward
                        child: child,
                      ),
                    ),
                child: Text(
                  '+$points RP',
                  style: const TextStyle(
                    fontSize: 32,
                    fontWeight: FontWeight.bold,
                    color: ColorConstants.greenColor,
                    shadows: [
                      Shadow(
                        blurRadius: 6.0,
                        color: Colors.black38,
                        offset: Offset(2, 2),
                      ),
                    ],
                  ),
                ),
              ),
            ),
          ),
        );
      },
    );

    overlay.insert(overlayEntry);
    controller.forward();

    Future.delayed(const Duration(milliseconds: 900), () {
      controller.dispose();
      overlayEntry.remove();
    });
  }
}
