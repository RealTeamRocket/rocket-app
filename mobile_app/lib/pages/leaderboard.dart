
import 'package:mobile_app/widgets/menu_bar.dart';
import 'package:flutter/material.dart';

class LeaderboardPage extends StatefulWidget {
  const LeaderboardPage({Key? key, required this.title}) : super(key: key);
    final String title;
  @override
  State<LeaderboardPage> createState() => _LeaderboardPageState();
}

class _LeaderboardPageState extends State<LeaderboardPage>{

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





  @override
  void initState() {
    super.initState();
    getList(); //
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
          Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              OutlinedButton(
                onPressed: () {
                  setState(() {
                    onlyFriends = false;
                  });
                  getList();
                },
                style: ElevatedButton.styleFrom(
                  backgroundColor: !onlyFriends ? Colors.blue : Colors.grey,
                ),
                child: const Text('All'),
              ),
              const SizedBox(width: 10),
              OutlinedButton(
                onPressed: () {
                  setState(() {
                    onlyFriends = true;
                  });
                  getList();
                },
                style: ElevatedButton.styleFrom(
                  backgroundColor: onlyFriends ? Colors.blue : Colors.grey,
                ),
                child: const Text('Friends'),
              ),
            ],
          ),
          if (sortedUsers.length >= 3) ...[
            const SizedBox(height: 20),
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Expanded(
                  child: GridView.count(
                    crossAxisCount: 3,
                    mainAxisSpacing: 50,
                    crossAxisSpacing: 50,
                    shrinkWrap: true,
                    children: [
                      Container(),
                      Container(
                        decoration: BoxDecoration(
                          border: Border.all(color: Colors.black),
                          color: Colors.blue,
                        ),
                        child: Center(
                          child: Text(
                            '${sortedUsers[0].name} - ${sortedUsers[0].rocketpoints}',
                            style: const TextStyle(color: Colors.white),
                          ),
                        ),
                      ),
                    ],
                  ),
                ),
              ],
            ),
            const SizedBox(height: 20),
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                Expanded(
                  child: GridView.count(
                    crossAxisCount: 4,
                    mainAxisSpacing: 10,
                    crossAxisSpacing: 10,
                    shrinkWrap: true,
                    children: [
                      Container(),
                      Container(
                        decoration: BoxDecoration(
                          border: Border.all(color: Colors.black),
                          color: Colors.blue,
                        ),
                        child: Center(
                          child: Text(
                            '${sortedUsers[1].name} - ${sortedUsers[1].rocketpoints}',
                            style: const TextStyle(color: Colors.white),
                          ),
                        ),
                      ),
                      Container(
                        decoration: BoxDecoration(
                          border: Border.all(color: Colors.black),
                          color: Colors.blue,
                        ),
                        child: Center(
                          child: Text(
                            '${sortedUsers[2].name} - ${sortedUsers[2].rocketpoints}',
                            style: const TextStyle(color: Colors.white),
                          ),
                        ),
                      ),
                    ],
                  ),
                ),
              ],
            ),
          ] else
            const Center(child: Text('Nicht genügend Benutzer für das Podium')),
          const SizedBox(height: 20),
          Expanded(
            child: ListView.builder(

              itemCount: sortedUsers.length > 3 ? (sortedUsers.length - 3).clamp(0, 5) : 0,
              itemBuilder: (context, index) {
                final user = sortedUsers[index + 3];
                return ListTile(
                  title: Text(user.name),
                  subtitle: Text('Rocketpoints: ${user.rocketpoints}'),
                  trailing: Icon(
                    user.isFriend ? Icons.person : Icons.person_outline,
                    color: user.isFriend ? Colors.green : Colors.grey,
                  ),
                );
              },
            ),
          ),
        ],
      ),
      bottomNavigationBar: CustomMenuBar(
        onItemTapped: _onItemTapped,
      ),
    );
  }

}

void main() {
  runApp(MaterialApp(
    home: LeaderboardPage(title: 'Leaderboard'),
  ));
}

class User {
  final String name;
  final int rocketpoints;
  final bool isFriend;

  User({required this.name, required this.rocketpoints, required this.isFriend});
}