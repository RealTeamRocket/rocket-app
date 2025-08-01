import 'package:flutter/material.dart';
import '/constants/constants.dart';

class ButtonsWidget extends StatelessWidget {
  final String selectedButton;
  final Function(String) onButtonPressed;

  const ButtonsWidget({
    super.key,
    required this.selectedButton,
    required this.onButtonPressed,
  });

  @override
  Widget build(BuildContext context) {
    return Center(
      child: ElevatedButton(
        onPressed: () => onButtonPressed('Race'),
        style: ElevatedButton.styleFrom(
          backgroundColor: selectedButton == 'Race'
              ? ColorConstants.greenColor
              : ColorConstants.blueColor,
          foregroundColor: ColorConstants.white,
          padding: EdgeInsets.symmetric(
            horizontal: 30.0,
            vertical: 15.0,
          ),
          textStyle: TextStyle(
            fontSize: 20.0,
            fontWeight: FontWeight.bold,
          ),
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(7.0),
          ),
        ),
        child: Row(
          mainAxisSize: MainAxisSize.min,
          children: [
            Icon(Icons.directions_run, size: 24.0),
            SizedBox(width: 8.0),
            Text("Race"),
          ],
        ),
      ),
    );
  }
}
