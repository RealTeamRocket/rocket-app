import 'package:flutter/material.dart';
import '../constants/color_constants.dart';
import 'package:flutter_slidable/flutter_slidable.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

import 'package:mobile_app/utils/backend_api/friends_api.dart';
import 'package:mobile_app/utils/backend_api/user_api.dart';

class FriendlistPage extends StatefulWidget {
  const FriendlistPage({super.key, required this.title});

  final String title;

  @override
  State<FriendlistPage> createState() => _FriendlistPageState();
}

class _FriendlistPageState extends State<FriendlistPage> with TickerProviderStateMixin {
  final TextEditingController _searchController = TextEditingController();
  String _searchQuery = '';
  OverlayEntry? _overlayEntry;

  List<Friend> friends = [];
  List<UserWithImage> allUsers = [];

  @override
  void initState() {
    super.initState();
    loadData();
  }

  Future<void> loadData() async {
    final jwt = await const FlutterSecureStorage().read(key: 'jwt_token');
    if (jwt == null) return;
    var fetchedFriends = await FriendsApi.getAllFriends(jwt);
    final users = await FriendsApi.getAllUsers(jwt);
    setState(() {
      friends = fetchedFriends;
      allUsers = users;
    });
  }

  @override
  Widget build(BuildContext context) {
    final filteredFriends = friends.where((f) {
      return _searchQuery.isEmpty ||
          f['name']!.toLowerCase().contains(_searchQuery.toLowerCase());
    }).toList();

    return Container(
      color: ColorConstants.primaryColor,
      padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
      child: Column(
        children: [
          _buildSearchBar(),
          const SizedBox(height: 16),
          if (_searchQuery.isNotEmpty &&
              !friends.any((f) =>
              f['name']!.toLowerCase() == _searchQuery.toLowerCase()))
            _buildSearchResultCard(_searchQuery),
          const SizedBox(height: 8),
          Expanded(
            child: ListView.builder(
              itemCount: filteredFriends.length,
              itemBuilder: (context, index) {
                final friend = filteredFriends[index];
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
        hintText: 'Search for friends...',
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
        // TODO: sync with backend later/filtering for the right person => Call searchUsers API here to fetch real users
        setState(() {
          _searchQuery = value.trim();
        });
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
            child: const Icon(
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

  Widget _buildSearchResultCard(String name) {
    return Container(
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
          const CircleAvatar(
            radius: 28,
            backgroundImage:
            NetworkImage('https://i.pravatar.cc/150?img=65'),
          ),
          const SizedBox(width: 16),
          Expanded(
            child: Text(
              name,
              style: const TextStyle(
                fontSize: 16,
                fontWeight: FontWeight.bold,
                color: ColorConstants.white,
              ),
            ),
          ),
          ElevatedButton.icon(
            onPressed: () {
              // TODO: Send follow request to backend here (e.g. /api/follow)
              setState(() {
                friends.add({
                  'name': name,
                  'data': 'Newly followed',
                  'image': 'https://i.pravatar.cc/150?img=65',
                });
                _searchController.clear();
                _searchQuery = '';
                _showFollowOverlay(name);
              });
            },
            icon: const Icon(Icons.person_add, size: 18, color: Colors.white),
            label: const Text(
              'Follow',
              style: TextStyle(color: Colors.white),
            ),
            style: ElevatedButton.styleFrom(
              backgroundColor: ColorConstants.greenColor,
              padding:
              const EdgeInsets.symmetric(horizontal: 12, vertical: 8),
              shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(8),
              ),
            ),
          ),
        ],
      ),
    );
  }

  void _showFollowOverlay(String name) {
    final overlay = Overlay.of(context);

    _overlayEntry = OverlayEntry(
      builder: (context) => Center(
        child: Container(
          padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 12),
          decoration: BoxDecoration(
            color: ColorConstants.greenColor.withValues(alpha: 0.8),
            borderRadius: BorderRadius.circular(24),
          ),
          child: Text(
            'Followed\n$name',
            textAlign: TextAlign.center,
            style: const TextStyle(
              color: Colors.white,
              fontSize: 32,
              fontWeight: FontWeight.bold,
            ),
          ),
        ),
      ),
    );

    overlay.insert(_overlayEntry!);
    Future.delayed(const Duration(seconds: 2), () {
      _overlayEntry?.remove();
      _overlayEntry = null;
    });
  }
}
