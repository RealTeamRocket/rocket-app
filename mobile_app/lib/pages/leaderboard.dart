import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:mobile_app/widgets/menu_bar.dart';
import 'package:flutter/material.dart';

class LeaderboardPage extends StatefulWidget {
  const LeaderboardPage({Key? key, required this.title}) : super(key: key);
    final String title;
  @override
  State<LeaderboardPage> createState() => _LeaderboardPageState();
}

class _LeaderboardPageState extends State<LeaderboardPage> with SingleTickerProviderStateMixin{

  List<User> users = [
    User(name: 'John Doe', rocketpoints: 100, isFriend: true),
    User(name: 'Jane Smith', rocketpoints: 200, isFriend: false),
    User(name: 'Alice Johnson', rocketpoints: 150, isFriend: true),
    User(name: 'Bob Brown', rocketpoints: 50, isFriend: false),
    User(name: 'Charlie Davis', rocketpoints: 300, isFriend: true),
    User(name: 'Diana Prince', rocketpoints: 250, isFriend: false),
    User(name: 'Ethan Hunt', rocketpoints: 400, isFriend: true),
    User(name: 'Fiona Apple', rocketpoints: 350, isFriend: false),
    User(name: 'George Clooney', rocketpoints: 450, isFriend: true),
    User(name: 'Hannah Montana', rocketpoints: 500, isFriend: false),
    User(name: 'Ian Somerhalder', rocketpoints: 600, isFriend: true),
    User(name: 'Jessica Alba', rocketpoints: 550, isFriend: false),
    User(name: 'Kevin Spacey', rocketpoints: 700, isFriend: true),
    User(name: 'Liam Neeson', rocketpoints: 650, isFriend: false),
    User(name: 'Megan Fox', rocketpoints: 800, isFriend: true),
    User(name: 'Nina Dobrev', rocketpoints: 750, isFriend: false),
  ];



  late AnimationController _controller;
  late Animation<double> _animation;
  int selectedToggle = 1;
  bool onlyFriends = false;
  List<User> sortedUsers = [];
  void getList() {
      if (!onlyFriends) {
        sortedUsers = List.from(users)
          ..sort((a, b) => b.rocketpoints.compareTo(a.rocketpoints));
      } else {
        sortedUsers = users
            .where((user) => user.isFriend)
            .toList()
          ..sort((a, b) => b.rocketpoints.compareTo(a.rocketpoints));
      }
  }
  int _currentIndex = 0;


  /// TODO: Needs to be implemented
  void _onItemTapped(int index) {
    setState(() {
      _currentIndex = index;
    });

    switch (index) {
      case 0:
        Navigator.pushReplacementNamed(context, '/home');
        break;
      case 1:
        Navigator.pushReplacementNamed(context, '/search');
        break;
      case 2:
        Navigator.pushReplacementNamed(context, '/leaderboards');
        break;
      case 3:
        Navigator.pushReplacementNamed(context, '/profile');
        break;
    }

  }

  void addFriend(User user) {
    print('Friend ${user.name} has been added!'); // TODO: Implement add friend functionality
  }




  @override
  void initState() {
    super.initState();
    getList();
    _controller = AnimationController(
      vsync: this,
      duration: const Duration(milliseconds: 300),
    );
    _animation = Tween<double>(begin: 0, end: 1).animate(_controller)
      ..addListener(() {
        setState(() {});
      });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
        backgroundColor: Colors.blueGrey[100],
      ),
      backgroundColor: Colors.blueGrey[100],
      body: Column(
        children: [

          /**
           * Here: Buttons to filter the leaderboard
           */
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
                      alignment: onlyFriends ? Alignment.centerRight : Alignment.centerLeft,
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
                          onTap: () {
                            setState(() {
                              onlyFriends = false;
                            });
                            getList();
                          },
                          child: Container(
                            width: 98, // Breite anpassen
                            alignment: Alignment.center,
                            child: Text(
                              "All",
                              style: TextStyle(
                                color: !onlyFriends ? Colors.white : Colors.black,
                                fontSize: 15,
                              ),
                            ),
                          ),
                        ),
                        InkWell(
                          onTap: () {
                            setState(() {
                              onlyFriends = true;
                            });
                            getList();
                          },
                          child: Container(
                            width: 98, // Breite anpassen
                            alignment: Alignment.center,
                            child: Text(
                              "Friends",
                              style: TextStyle(
                                color: onlyFriends ? Colors.white : Colors.black,
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
          /**
           * Here: Top part of the page (Podium)
           */
          /**
           * TODO: Instead of a print Mehod with the podium open maybe a popup to add as a friend or see profile if is friend
           *
           * TODO: Max. Length of text inside the container/podium needs to be implemented
           */

          if (sortedUsers.length >= 3) ...[
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
                      Column(
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          GestureDetector(
                            onTap: () {
                              showDialog(
                                context: context,
                                builder: (BuildContext context) {
                                  return AlertDialog(
                                    title: Text(sortedUsers[0].name),
                                    content: Column(
                                      mainAxisSize: MainAxisSize.min,
                                      children: [
                                        Text('Rocketpoints: ${sortedUsers[0].rocketpoints}'),
                                        Text('Signed up: 01.01.2023'), // Beispiel-Datum
                                        if(sortedUsers[0].isFriend)...[
                                          Text('Status: Friend'),
                                         ]else...[
                                          Text('Status: Not a friend'),
                                          IconButton(
                                            icon: Icon(Icons.person_add_alt, color: Colors.black),
                                            onPressed: () {
                                              addFriend(sortedUsers[0]);
                                            },
                                          )
                                        ]
                                      ],
                                    ),
                                    actions: [
                                      TextButton(
                                        onPressed: () {
                                          Navigator.of(context).pop(); // Schließt das Pop-up
                                        },
                                        child: const Text('Okay'),
                                      ),
                                    ],
                                  );
                                },
                              );
                            },
                            child: Container(
                            width: 80, // Breite des Containers
                            height: 80, // Höhe des Containers
                            decoration: BoxDecoration(
                              border: Border.all(color: Colors.black),
                              color: Color(0xFFFFD700),
                              shape: BoxShape.circle,
                            ),
                            child: Column(
                              mainAxisAlignment: MainAxisAlignment.center,
                              children: [
                                  Icon(FontAwesomeIcons.award, color: Colors.white, size: 24),
                                const SizedBox(height: 5), // Abstand zwischen Icon und Text
                                Text(
                                  '${sortedUsers[0].name.length > 12 ? sortedUsers[0].name.substring(0, 12): sortedUsers[0].name}\n${sortedUsers[0].rocketpoints} RP',
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
                          ),
                        ],
                      ),
                      Container(),
                      Container(),
                      Container(),
                      Column(
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          GestureDetector(
                            onTap: () {
                              showDialog(
                                context: context,
                                builder: (BuildContext context) {
                                  return AlertDialog(
                                    title: Text(sortedUsers[1].name),
                                    content: Column(
                                      mainAxisSize: MainAxisSize.min,
                                      children: [
                                        Text('Rocketpoints: ${sortedUsers[1].rocketpoints}'),
                                        Text('Signed up: 01.01.2021'), // Beispiel-Datum
                                        if(sortedUsers[1].isFriend)...[
                                          Text('Status: Friend'),
                                        ]else...[
                                          Text('Status: Not a friend'),
                                          IconButton(
                                            icon: Icon(Icons.person_add_alt, color: Colors.black),
                                            onPressed: () {
                                              addFriend(sortedUsers[1]);
                                            },
                                          )
                                        ]
                                      ],
                                    ),
                                    actions: [
                                      TextButton(
                                        onPressed: () {
                                          Navigator.of(context).pop(); // Schließt das Pop-up
                                        },
                                        child: const Text('Okay'),
                                      ),
                                    ],
                                  );
                                },
                              );
                            },
                            child: Container(
                            width: 80, // Breite des Containers
                            height: 80, // Höhe des Containers
                            decoration: BoxDecoration(
                              border: Border.all(color: Colors.black),
                              color: Color(0xFFC0C0C0),
                              shape: BoxShape.circle,
                            ),
                            child: Column(
                              mainAxisAlignment: MainAxisAlignment.center,
                              children: [
                                Icon(FontAwesomeIcons.medal, color: Colors.white, size: 24),
                                const SizedBox(height: 5), // Abstand zwischen Icon und Text
                                Text(
                                  '${sortedUsers[1].name.length > 12 ? sortedUsers[1].name.substring(0, 12): sortedUsers[1].name}\n${sortedUsers[1].rocketpoints} RP',
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
                          ),
                        ],
                      ),
                      Container(),
                      Column(
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          GestureDetector(
                            onTap: () {
                              showDialog(
                                context: context,
                                builder: (BuildContext context) {
                                  return AlertDialog(
                                    title: Text(sortedUsers[2].name),
                                    content: Column(
                                      mainAxisSize: MainAxisSize.min,
                                      children: [
                                        Text('Rocketpoints: ${sortedUsers[2].rocketpoints}'),
                                        Text('Signed up: 01.01.2022'), // Beispiel-Datum
                                        if(sortedUsers[2].isFriend)...[
                                          Text('Status: Friend'),
                                        ]else...[
                                          Text('Status: Not a friend'),
                                          IconButton(
                                            icon: Icon(Icons.person_add_alt, color: Colors.black),
                                            onPressed: () {
                                              addFriend(sortedUsers[2]);
                                            },
                                          )
                                        ]
                                      ],
                                    ),
                                    actions: [
                                      TextButton(
                                        onPressed: () {
                                          Navigator.of(context).pop(); // Schließt das Pop-up
                                        },
                                        child: const Text('Okay'),
                                      ),
                                    ],
                                  );
                                },
                              );
                            },
                            child: Container(
                            width: 80, // Breite des Containers
                            height: 80, // Höhe des Containers
                            decoration: BoxDecoration(
                              border: Border.all(color: Colors.black),
                              color: Color(0xFFCD7F32),
                              shape: BoxShape.circle,
                            ),
                            child: Column(
                              mainAxisAlignment: MainAxisAlignment.center,
                              children: [
                                Icon(FontAwesomeIcons.trophy, color: Colors.white, size: 24),
                                const SizedBox(height: 5), // Abstand zwischen Icon und Text
                                Text(
                                  '${sortedUsers[2].name.length > 12 ? sortedUsers[2].name.substring(0, 12): sortedUsers[2].name}\n${sortedUsers[2].rocketpoints} RP',
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
                          ),
                        ],
                      ),
                    ],
                  ),
                ),
                )
              ],
            ),
          ] else
            const Center(child: Text('Nicht genügend Benutzer für das Podium')),

          /**
           * Here: Bottom part of the page
           */

          const SizedBox(height: 30),
          Expanded(
            child: ListView.builder(
              /// With .clamp(0, 5) we can limit the number of items to 5 or to any number we want
              /// Or leave it out to show all items
              itemCount: sortedUsers.length > 3 ? (sortedUsers.length - 3) : 0,
              itemBuilder: (context, index) {
                final user = sortedUsers[index + 3];
                return ListTile(
                  leading: Text(
                    '${index + 4}.',
                    style: const TextStyle(fontSize: 16, fontWeight: FontWeight.bold, color: Colors.black), // Größere Schriftgröße
                  ),
                  title: Text(
                    user.name,
                    style: const TextStyle(fontSize: 14, color: Colors.black), // Größere Schriftgröße
                  ),
                  subtitle: Text(
                    'Rocketpoints: ${user.rocketpoints}',
                    style: const TextStyle(fontSize: 12), // Größere Schriftgröße
                  ),
                  trailing: ElevatedButton(
                    onPressed: () {
                      addFriend(user);
                      // TODO: Add friend functionality
                    },
                    style: ElevatedButton.styleFrom(
                      backgroundColor: user.isFriend ? Colors.green : Colors.grey,
                    ),
                    child: Icon(
                      user.isFriend ? Icons.person : Icons.person_add_alt,
                      color: Colors.white,
                      size: 24, // Größeres Icon
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
}

class User {
  final String name;
  final int rocketpoints;
  final bool isFriend;

  User({required this.name, required this.rocketpoints, required this.isFriend});
}
