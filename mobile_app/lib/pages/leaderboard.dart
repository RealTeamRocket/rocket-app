import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:flutter/material.dart';
import 'package:mobile_app/utils/backend_api/ranking_api.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class LeaderboardPage extends StatefulWidget {
  const LeaderboardPage({Key? key, required this.title}) : super(key: key);
  final String title;

  @override
  State<LeaderboardPage> createState() => _LeaderboardPageState();
}

class _LeaderboardPageState extends State<LeaderboardPage> with SingleTickerProviderStateMixin {
  late AnimationController _controller;

  List<RankedUser> allUsers = [];
  List<RankedUser> friends = [];
  List<RankedUser> displayedUsers = [];
  bool isLoading = true;
  int selectedTab = 0; // 0 for "All", 1 for "Friends"

  @override
  void initState() {
    super.initState();
    _controller = AnimationController(
      vsync: this,
      duration: const Duration(milliseconds: 300),
    );
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

  void addFriend(RankedUser user) {
    print('Friend ${user.username} has been added!'); // TODO: Implement add friend functionality
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
        backgroundColor: Colors.blueGrey[100],
      ),
      backgroundColor: Colors.blueGrey[100],
      body: isLoading
          ? const Center(child: CircularProgressIndicator())
          : Column(
              children: [
                // Toggle Buttons for "All" and "Friends"
                const SizedBox(height: 7),
                Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Container(
                      height: 44,
                      width: 200,
                      decoration: BoxDecoration(
                        color: Colors.grey,
                        borderRadius: BorderRadius.circular(100),
                        border: Border.all(color: Colors.black, width: 2),
                      ),
                      child: Stack(
                        alignment: Alignment.center,
                        clipBehavior: Clip.none,
                        children: [
                          AnimatedAlign(
                            alignment: selectedTab == 1 ? Alignment.centerRight : Alignment.centerLeft,
                            duration: const Duration(milliseconds: 300),
                            child: Container(
                              width: 100,
                              height: 44,
                              decoration: BoxDecoration(
                                color: const Color(0xff9ba0fc),
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
                                      color: selectedTab == 0 ? Colors.white : Colors.black,
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
                                      color: selectedTab == 1 ? Colors.white : Colors.black,
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

                // Podium for Top 3 Users
                if (displayedUsers.length >= 3) ...[
                  const SizedBox(height: 30),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Expanded(
                        child: SizedBox(
                          height: 200,
                          child: GridView.count(
                            crossAxisCount: 5,
                            mainAxisSpacing: 1,
                            crossAxisSpacing: 1,
                            shrinkWrap: true,
                            children: [
                              Container(),
                              Container(),
                              buildPodiumUser(displayedUsers[0], Colors.yellow, FontAwesomeIcons.award),
                              Container(),
                              Container(),
                              Container(),
                              buildPodiumUser(displayedUsers[1], Colors.grey, FontAwesomeIcons.medal),
                              Container(),
                              buildPodiumUser(displayedUsers[2], Colors.brown, FontAwesomeIcons.trophy),
                            ],
                          ),
                        ),
                      )
                    ],
                  ),
                ] else ...[
                    const Center(child: Text('Not enough users for the podium')),
                    const SizedBox(height: 10),
                    Expanded(
                      child: ListView.builder(
                        itemCount: displayedUsers.length,
                        itemBuilder: (context, index) {
                          final user = displayedUsers[index];
                          return ListTile(
                            leading: Text(
                              '${index + 1}.',
                              style: const TextStyle(fontSize: 16, fontWeight: FontWeight.bold, color: Colors.black),
                            ),
                            title: Text(
                              user.username,
                              style: const TextStyle(fontSize: 14, color: Colors.black),
                            ),
                            subtitle: Text(
                              'Rocketpoints: ${user.rocketPoints}',
                              style: const TextStyle(fontSize: 12),
                            ),
                          );
                        },
                      ),
                    ),
                ],

                // List of Remaining Users
                const SizedBox(height: 30),
                Expanded(
                  child: ListView.builder(
                    itemCount: displayedUsers.length > 3 ? (displayedUsers.length - 3) : 0,
                    itemBuilder: (context, index) {
                      final user = displayedUsers[index + 3];
                      return ListTile(
                        leading: Text(
                          '${index + 4}.',
                          style: const TextStyle(fontSize: 16, fontWeight: FontWeight.bold, color: Colors.black),
                        ),
                        title: Text(
                          user.username,
                          style: const TextStyle(fontSize: 14, color: Colors.black),
                        ),
                        subtitle: Text(
                          'Rocketpoints: ${user.rocketPoints}',
                          style: const TextStyle(fontSize: 12),
                        ),
                        trailing: ElevatedButton(
                          onPressed: () {
                            addFriend(user);
                          },
                          style: ElevatedButton.styleFrom(
                            backgroundColor: Colors.grey,
                          ),
                          child: const Icon(
                            Icons.person_add_alt,
                            color: Colors.white,
                            size: 24,
                          ),
                        ),
                      );
                    },
                  ),
                ),
              ],
            ),
    );
  }

  Widget buildPodiumUser(RankedUser user, Color color, IconData icon) {
    return GestureDetector(
      onTap: () {
        showDialog(
          context: context,
          builder: (BuildContext context) {
            return AlertDialog(
              title: Text(user.username),
              content: Column(
                mainAxisSize: MainAxisSize.min,
                children: [
                  Text('Rocketpoints: ${user.rocketPoints}'),
                  Text('Signed up: 01.01.2023'), // Example date
                  Text('Status: Not a friend'),
                  IconButton(
                    icon: const Icon(Icons.person_add_alt, color: Colors.black),
                    onPressed: () {
                      addFriend(user);
                    },
                  ),
                ],
              ),
              actions: [
                TextButton(
                  onPressed: () {
                    Navigator.of(context).pop();
                  },
                  child: const Text('Okay'),
                ),
              ],
            );
          },
        );
      },
      child: Container(
        width: 80,
        height: 80,
        decoration: BoxDecoration(
          border: Border.all(color: Colors.black),
          color: color,
          shape: BoxShape.circle,
        ),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Icon(icon, color: Colors.white, size: 24),
            const SizedBox(height: 5),
            Text(
              '${user.username.length > 12 ? user.username.substring(0, 12) : user.username}\n${user.rocketPoints} RP',
              style: const TextStyle(
                color: Colors.black,
                fontSize: 12,
                fontWeight: FontWeight.bold,
              ),
              textAlign: TextAlign.center,
            ),
          ],
        ),
      ),
    );
  }
}
