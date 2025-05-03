import 'package:flutter/material.dart';
import '../constants/color_constants.dart';
import 'package:flutter_slidable/flutter_slidable.dart';

class FriendlistPage extends StatefulWidget {
  const FriendlistPage({super.key, required this.title});

  final String title;

  @override
  State<FriendlistPage> createState() => _FriendlistPageState();
}

class _FriendlistPageState extends State<FriendlistPage> {
  final TextEditingController _searchController = TextEditingController();

  List<Map<String, String>> friends = [
    {
      'name': 'Julianne Hough',
      'data': 'Placeholder for something',
      'image': 'https://i.pravatar.cc/150?img=1',
    },
    {
      'name': 'Jessica Lowndes',
      'data': 'Placeholder for something',
      'image': 'https://i.pravatar.cc/150?img=2',
    },
    {
      'name': 'Kristen Bell',
      'data': 'Placeholder for something',
      'image': 'https://i.pravatar.cc/150?img=3',
    },
    {
      'name': 'Mila Kunis',
      'data': 'Placeholder for something',
      'image': 'https://i.pravatar.cc/150?img=4',
    },
    {
      'name': 'Rosie Jones',
      'data': 'Placeholder for something',
      'image': 'https://i.pravatar.cc/150?img=5',
    },
    {
      'name': 'Alexandra Pierce',
      'data': 'Placeholder for something',
      'image': 'https://i.pravatar.cc/150?img=6',
    },
    {
      'name': 'Daniela Hartmann',
      'data': 'Placeholder for something',
      'image': 'https://i.pravatar.cc/150?img=7',
    },
    {
      'name': 'Nina Lopez',
      'data': 'Placeholder for something',
      'image': 'https://i.pravatar.cc/150?img=8',
    },
    {
      'name': 'Tessa Morgan',
      'data': 'Placeholder for something',
      'image': 'https://i.pravatar.cc/150?img=9',
    },
    {
      'name': 'Carla Bergström',
      'data': 'Placeholder for something',
      'image': 'https://i.pravatar.cc/150?img=10',
    },
  ];

  @override
  Widget build(BuildContext context) {
    return Container(
      color: ColorConstants.primaryColor,
      padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
      child: Column(
        children: [
          _buildSearchBar(),
          const SizedBox(height: 16),
          Expanded(
            child: ListView.builder(
              itemCount: friends.length,
              itemBuilder: (context, index) {
                final friend = friends[index];
                return _buildFriendCard(friend);
              },
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildSearchBar() {
    return TextField(
      controller: _searchController,
      style: const TextStyle(color: ColorConstants.white),
      decoration: InputDecoration(
        hintText: 'Search for new friends...',
        hintStyle: const TextStyle(color: Colors.white54),
        prefixIcon: const Icon(Icons.search, color: Colors.white),
        filled: true,
        fillColor: ColorConstants.secoundaryColor,
        border: OutlineInputBorder(
          borderRadius: BorderRadius.circular(12),
          borderSide: BorderSide.none,
        ),
        contentPadding: const EdgeInsets.symmetric(vertical: 0, horizontal: 16),
      ),
      onChanged: (value) {
        // TODO: sync with backend later/filtering for the right person
      },
    );
  }

  Widget _buildFriendCard(Map<String, String> friend) {
    return Slidable(
      key: ValueKey(friend['name']),
      endActionPane: ActionPane(
        motion: const StretchMotion(),
        extentRatio: 0.25,
        children: [
          CustomSlidableAction(
            onPressed: (_) {
              // TODO: Unfollow logic with API
              setState(() {
                friends.remove(friend);
              });
            },
            backgroundColor: const Color(0xFFB5544D),
            borderRadius: BorderRadius.circular(16),
            child: Icon(
              Icons.person_remove,
              size: 30,
              color: Colors.white,
            ),
          ),
        ],
      ),
      child: Container(
        margin: const EdgeInsets.symmetric(vertical: 8),
        padding: const EdgeInsets.all(12),
        decoration: BoxDecoration(
          color: ColorConstants.secoundaryColor,
          borderRadius: BorderRadius.circular(16),
          boxShadow: [
            BoxShadow(
              color: Colors.black.withValues(alpha: 0.05),
              blurRadius: 4,
              offset: const Offset(0, 2),
            )
          ],
        ),
        child: Row(
          children: [
            CircleAvatar(
              radius: 28,
              backgroundImage: NetworkImage(friend['image']!),
            ),
            const SizedBox(width: 16),
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    friend['name']!,
                    style: const TextStyle(
                      fontSize: 16,
                      fontWeight: FontWeight.bold,
                      color: ColorConstants.white,
                    ),
                  ),
                  const SizedBox(height: 4),
                  Text(
                    friend['data']!,
                    style: const TextStyle(color: Colors.white70),
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }

}
