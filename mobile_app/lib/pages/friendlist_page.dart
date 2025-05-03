import 'package:flutter/material.dart';
import '../constants/color_constants.dart';

class FriendlistPage extends StatefulWidget {
  const FriendlistPage({super.key, required this.title});

  final String title;

  @override
  State<FriendlistPage> createState() => _FriendlistPageState();
}

class _FriendlistPageState extends State<FriendlistPage> {
  @override
  Widget build(BuildContext context) {
    return Container(
      color: ColorConstants.primaryColor,
      child: const Center(
        child: Text(
          'Implement friendlist here',
          style: TextStyle(
            fontSize: 20,
            fontWeight: FontWeight.bold,
            color: ColorConstants.white,
          ),
        ),
      ),
    );
  }
}
