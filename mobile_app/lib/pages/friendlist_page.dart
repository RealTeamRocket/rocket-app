import 'package:flutter/material.dart';
import '../constants/color_constants.dart';
import 'package:flutter_slidable/flutter_slidable.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

import 'package:mobile_app/utils/backend_api/friends_api.dart';

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

  @override
  void dispose() {
    _searchController.dispose();
    super.dispose();
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
    // 1) Filter followed friends that match the text
    final filteredFriends = friends.where((f) {
      return _searchQuery.isEmpty || f.username.toLowerCase().contains(_searchQuery.toLowerCase());
    }).toList();

    // 2) Filter all Users, that are not friends and match the searched text
    final searchResults = allUsers.where((u) {
      final matchesQuery = _searchQuery.isNotEmpty && u.username.toLowerCase().contains(_searchQuery.toLowerCase());
      final notFriendYet = !friends.any((f) => f.id == u.id);
      return matchesQuery && notFriendYet;
    }).toList();

    return Container(
      color: ColorConstants.primaryColor,
      padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
      child: Column(
        children: [
          _buildSearchBar(),
          const SizedBox(height: 16),

          // 3) show filtered list
          if (searchResults.isNotEmpty) ...[
            Text('Add new friends:', style: TextStyle(color: Colors.white70)),
            const SizedBox(height: 8),
            SizedBox(
              height: searchResults.length * 72.0, // adjust itemHeight
              child: ListView.builder(
                physics: const NeverScrollableScrollPhysics(),
                itemCount: searchResults.length,
                itemBuilder: (ctx, i) => _buildSearchResultCard(searchResults[i]),
              ),
            ),
            const SizedBox(height: 16),
          ],

          // 4) followed friends
          Expanded(
            child: ListView.builder(
              itemCount: filteredFriends.length,
              itemBuilder: (ctx, i) => _buildFriendCard(filteredFriends[i]),
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
        // filtered in build() afterwards
        setState(() {
          _searchQuery = value.trim();
        });
      },
    );
  }

  Widget _buildFriendCard(Friend friend) {
    return Slidable(
      key: ValueKey(friend.id),
      endActionPane: ActionPane(
        motion: const StretchMotion(),
        extentRatio: 0.25,
        children: [
          CustomSlidableAction(
            onPressed: (_) async {
              final jwt = await const FlutterSecureStorage().read(key: 'jwt_token');
              if (jwt == null) return;

              try {
                // 1) Unfollow
                await FriendsApi.deleteFriend(jwt, friend.username);
                // 2) refresh lists
                final freshFriends = await FriendsApi.getAllFriends(jwt);
                final freshUsers = await FriendsApi.getAllUsers(jwt);
                setState(() {
                  friends = freshFriends;
                  allUsers = freshUsers;
                });
              } catch (e) {
                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(content: Text('Failed unfollowing: $e')),
                );
              }
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
              backgroundImage: friend.imageData != null
                  ? MemoryImage(friend.imageData!)
                  : const AssetImage('assets/avatar_placeholder.png') as ImageProvider,
            ),
            const SizedBox(width: 16),
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    friend.username,
                    style: const TextStyle(
                      fontSize: 16,
                      fontWeight: FontWeight.bold,
                      color: ColorConstants.white,
                    ),
                  ),
                  const SizedBox(height: 4),
                  Text(
                    friend.email, // maybe some other stats/data?
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

  Widget _buildSearchResultCard(UserWithImage user) {
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
          CircleAvatar(
            radius: 28,
            backgroundImage: user.imageData != null
                ? MemoryImage(user.imageData!)
                : const AssetImage('assets/avatar_placeholder.png') as ImageProvider,
          ),
          const SizedBox(width: 16),
          Expanded(
            child: Text(
              user.username,
              style: const TextStyle(
                fontSize: 16,
                fontWeight: FontWeight.bold,
                color: ColorConstants.white,
              ),
            ),
          ),
          ElevatedButton.icon(
            onPressed: () async {
              final jwt = await const FlutterSecureStorage().read(key:'jwt_token');
              if(jwt!=null){
                try{
                  // follow
                  await FriendsApi.addFriend(jwt, user.username);
                  // refresh lists
                  final freshFriends = await FriendsApi.getAllFriends(jwt);
                  final freshUsers   = await FriendsApi.getAllUsers(jwt);
                  setState(() {
                    friends  = freshFriends;
                    allUsers = freshUsers;
                    _searchQuery = '';
                    _searchController.clear();
                  });
                  // visual feedback
                  _showFollowOverlay(user.username);
                } catch (e) {
                  ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(content: Text('Error: $e')),
                  );
                }
              }
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
