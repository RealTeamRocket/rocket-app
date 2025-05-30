import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:flutter/material.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:mobile_app/utils/backend_api/backend_api.dart';

import '../constants/color_constants.dart';

class LeaderboardPage extends StatefulWidget {
  const LeaderboardPage({super.key, required this.title});
  final String title;

  @override
  State<LeaderboardPage> createState() => _LeaderboardPageState();
}

class _LeaderboardPageState extends State<LeaderboardPage> {
  List<RankedUser> allUsers = [];
  List<RankedUser> friends = [];
  List<RankedUser> displayedUsers = [];
  bool isLoading = true;
  int selectedTab = 0; // 0 for "All", 1 for "Friends"

  @override
  void initState() {
    super.initState();
    fetchRankings();
  }

  Future<void> fetchRankings() async {
    setState(() {
      isLoading = true;
    });

    try {
      final storage = FlutterSecureStorage();
      final jwt = await storage.read(key: 'jwt_token');
      if (jwt == null) {
        debugPrint("JWT token is null");
        return;
      }

      // Fetch all users and friends rankings
      final fetchedAllUsers = await RankingApi.fetchUserRankings(jwt);
      final fetchedFriends = await RankingApi.fetchFriendRankings(jwt);

      setState(() {
        allUsers = fetchedAllUsers;
        friends = fetchedFriends;
        displayedUsers = allUsers;
      });
    } catch (e) {
      debugPrint("Error fetching rankings: $e");
    } finally {
      setState(() {
        isLoading = false;
      });
    }
  }

  void switchTab(int tabIndex) {
    setState(() {
      selectedTab = tabIndex;
      displayedUsers = tabIndex == 0 ? allUsers : friends;
    });
  }

  void addFriend(RankedUser user) async {
    final storage = FlutterSecureStorage();
    final jwt = await storage.read(key: 'jwt_token');
    if (jwt == null) {
      debugPrint("JWT token is null");
      return;
    }
    try {
      debugPrint("Adding friend: ${user.username}");
      await FriendsApi.addFriend(jwt, user.username);
      debugPrint("Friend added successfully");
      setState(() {
        // Add to friends list if not already present
        if (!friends.any((f) => f.username == user.username)) {
          friends.add(user);
        }
      });
    } catch (e) {
      debugPrint("Error adding friend: $e");
    }
  }

  bool isFriend(RankedUser user) {
    return friends.any((f) => f.username == user.username);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(
            widget.title,
            style: TextStyle(
              color: ColorConstants.white,
              fontSize: 20,
              fontWeight: FontWeight.bold,
            ),
        ),
        centerTitle: true,
        backgroundColor: ColorConstants.primaryColor,
      ),
      backgroundColor: ColorConstants.primaryColor,
      body: isLoading
          ? const Center(child: CircularProgressIndicator())
          : SafeArea(
              child: SingleChildScrollView(
                padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.center,
                  children: [
                    // Toggle Buttons for "All" and "Friends"
                    Row(
                      mainAxisAlignment: MainAxisAlignment.center,
                      children: [
                        Container(
                          height: 44,
                          width: 200,
                          decoration: BoxDecoration(
                            color: ColorConstants.secoundaryColor,
                            borderRadius: BorderRadius.circular(100),
                            border: Border.all(color: ColorConstants.purpleColor.withValues(alpha: 0.3), width: 2),
                          ),
                          child: Stack(
                            alignment: Alignment.center,
                            clipBehavior: Clip.none,
                            children: [
                              AnimatedAlign(
                                alignment: selectedTab == 1
                                    ? Alignment.centerRight
                                    : Alignment.centerLeft,
                                duration: const Duration(milliseconds: 300),
                                child: Container(
                                  width: 100,
                                  height: 44,
                                  decoration: BoxDecoration(
                                    color: ColorConstants.purpleColor,
                                    borderRadius: BorderRadius.circular(100),
                                  ),
                                ),
                              ),
                              Row(
                                mainAxisAlignment: MainAxisAlignment.spaceAround,
                                children: [
                                  InkWell(
                                    onTap: () => switchTab(0),
                                    child: Container(
                                      width: 98,
                                      alignment: Alignment.center,
                                      child: Text(
                                        "All",
                                        style: TextStyle(
                                          color: ColorConstants.white,
                                          fontWeight: FontWeight.w600,
                                          fontSize: 15,
                                        ),
                                      ),
                                    ),
                                  ),
                                  InkWell(
                                    onTap: () => switchTab(1),
                                    child: Container(
                                      width: 98,
                                      alignment: Alignment.center,
                                      child: Text(
                                        "Friends",
                                        style: TextStyle(
                                          color: ColorConstants.white,
                                          fontWeight: FontWeight.w600,
                                          fontSize: 15,
                                        ),
                                      ),
                                    ),
                                  ),
                                ],
                              ),
                            ],
                          ),
                        ),
                      ],
                    ),
                    const SizedBox(height: 24),

                    // Podium for Top 3 Users (only if 3 or more)
                    if (displayedUsers.length >= 3) ...[
                      Center(
                        child: SizedBox(
                          width: 320,
                          child: Row(
                            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                            crossAxisAlignment: CrossAxisAlignment.end,
                            children: [
                              // 2nd Place
                              _PodiumSpot(
                                child: buildPodiumUserCircle(
                                  displayedUsers[1],
                                  Colors.grey,
                                  FontAwesomeIcons.medal,
                                  size: 65,
                                ),
                                place: "2nd",
                                points: displayedUsers[1].rocketPoints,
                                height: 110,
                              ),
                              // 1st Place (center, bigger)
                              _PodiumSpot(
                                child: buildPodiumUserCircle(
                                  displayedUsers[0],
                                  Colors.yellow[700]!,
                                  FontAwesomeIcons.award,
                                  size: 85,
                                ),
                                place: "1st",
                                points: displayedUsers[0].rocketPoints,
                                height: 140,
                              ),
                              // 3rd Place
                              _PodiumSpot(
                                child: buildPodiumUserCircle(
                                  displayedUsers[2],
                                  Colors.brown,
                                  FontAwesomeIcons.trophy,
                                  size: 55,
                                ),
                                place: "3rd",
                                points: displayedUsers[2].rocketPoints,
                                height: 90,
                              ),
                            ],
                          ),
                        ),
                      ),
                      const SizedBox(height: 32),
                    ],

                    // List of Users
                    if (displayedUsers.isNotEmpty)
                      ListView.separated(
                        physics: const NeverScrollableScrollPhysics(),
                        shrinkWrap: true,
                        itemCount: displayedUsers.length < 3
                            ? displayedUsers.length
                            : displayedUsers.length - 3,
                        separatorBuilder: (_, __) => const Divider(height: 1),
                        itemBuilder: (context, index) {
                          final user = displayedUsers.length < 3
                              ? displayedUsers[index]
                              : displayedUsers[index + 3];
                          return ListTile(
                            leading: Text(
                              '${displayedUsers.length < 3 ? index + 1 : index + 4}.',
                              style: const TextStyle(
                                fontSize: 18,
                                fontWeight: FontWeight.bold,
                                color: ColorConstants.white,
                              ),
                            ),
                            title: Text(
                              user.username,
                              style: const TextStyle(
                                fontSize: 16,
                                fontWeight: FontWeight.bold,
                                color: ColorConstants.white,
                              ),
                            ),
                            subtitle: Text(
                              'Rocketpoints: ${user.rocketPoints}',
                              style: const TextStyle(
                                  fontSize: 14,
                                  color: ColorConstants.purpleColor
                              ),
                            ),
                            trailing: selectedTab == 0
                                ? ElevatedButton(
                                    onPressed: isFriend(user)
                                        ? null
                                        : () {
                                            addFriend(user);
                                          },
                                    style: ElevatedButton.styleFrom(
                                      backgroundColor: isFriend(user)
                                          ? ColorConstants.greenColor
                                          : Colors.grey.withValues(alpha: 0.3),
                                      shape: const CircleBorder(),
                                      padding: const EdgeInsets.all(8),
                                    ),
                                    child: Icon(
                                      isFriend(user)
                                          ? Icons.check
                                          : Icons.person_add_alt,
                                      color: isFriend(user)
                                          ? ColorConstants.greenColor
                                          : ColorConstants.white,
                                      size: 22,
                                    ),
                                  )
                                : null,
                          );
                        },
                      ),
                  ],
                ),
              ),
            ),
    );
  }

  // Only icon and username in the circle
  Widget buildPodiumUserCircle(RankedUser user, Color color, IconData icon, {double size = 80}) {
    return GestureDetector(
      onTap: () {
        showDialog(
          context: context,
          builder: (BuildContext context) {
            return AlertDialog(
              backgroundColor: ColorConstants.secoundaryColor.withValues(alpha: 0.9),
              title: Center(
                child: Text(
                  user.username,
                  style: const TextStyle(color: ColorConstants.white),
                ),
              ),
              content: Column(
                mainAxisSize: MainAxisSize.min,
                children: [
                  Text(
                      'Rocketpoints: ${user.rocketPoints}',
                      style: const TextStyle(color: ColorConstants.white),
                  ),
                  Text(
                      'Signed up: 01.01.2023',
                      style: const TextStyle(color: ColorConstants.white),
                  ), // Example date
                  Text(
                      'Status: Not a friend',
                      style: const TextStyle(color: ColorConstants.white),
                  ),
                  IconButton(
                    icon: const Icon(Icons.person_add_alt, color: ColorConstants.white),
                    onPressed: () {
                      addFriend(user);
                    },
                  ),
                ],
              ),
              actions: [
                ElevatedButton(
                  onPressed: () {
                    Navigator.of(context).pop();
                  },
                  style: ElevatedButton.styleFrom(
                    backgroundColor: ColorConstants.purpleColor,
                    foregroundColor: Colors.white,
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(8),
                    ),
                  ),
                  child: const Text('Okay'),
                ),
              ],
            );
          },
        );
      },
      child: Container(
        width: size,
        height: size,
        decoration: BoxDecoration(
          border: Border.all(color: ColorConstants.purpleColor.withValues(alpha: 0.3)),
          color: color,
          shape: BoxShape.circle,
          boxShadow: [
            BoxShadow(
              color: Colors.black.withValues(alpha: 0.08),
              blurRadius: 6,
              offset: const Offset(0, 3),
            ),
          ],
        ),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Icon(icon, color: Colors.white, size: size * 0.4),
            const SizedBox(height: 4),
            Text(
              user.username.length > 10
                  ? user.username.substring(0, 10)
                  : user.username,
              style: const TextStyle(
                color: Colors.black,
                fontSize: 14,
                fontWeight: FontWeight.bold,
              ),
              textAlign: TextAlign.center,
              overflow: TextOverflow.ellipsis,
            ),
          ],
        ),
      ),
    );
  }
}

class _PodiumSpot extends StatelessWidget {
  final Widget child;
  final String place;
  final int points;
  final double height;

  const _PodiumSpot({
    required this.child,
    required this.place,
    required this.points,
    required this.height,
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.end,
      children: [
        SizedBox(height: height, child: child),
        const SizedBox(height: 8),
        Text(
          place,
          style: const TextStyle(
              fontWeight: FontWeight.bold,
              fontSize: 18,
              color: ColorConstants.white,
          ),
        ),
        Text(
          '$points RP',
          style: const TextStyle(
            color: ColorConstants.purpleColor,
            fontSize: 15,
            fontWeight: FontWeight.w500,
          ),
        ),
      ],
    );
  }
}
