import 'package:cloud_firestore/cloud_firestore.dart';

class User {
  String name;
  String nachName;

  User({required this.name, required this.nachName});

  Map<String, dynamic> toMap() {
    return {
      'Name': name,
      'Nachname': nachName,
    };
  }
}

void saveUser(User user) async {
  await FirebaseFirestore.instance.collection('test').add(user.toMap());
}
