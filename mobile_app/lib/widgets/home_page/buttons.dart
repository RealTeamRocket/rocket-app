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
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 20.0),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          ElevatedButton(
            onPressed: () => onButtonPressed('Steps'),
            style: ElevatedButton.styleFrom(
              backgroundColor: selectedButton == 'Steps'
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
            child: Text("Steps"),
          ),
          SizedBox(width: 10.0),
          ElevatedButton(
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
            child: Text("Race"),
          ),
        ],
      ),
    );
  }
}
