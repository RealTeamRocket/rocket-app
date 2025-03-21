import 'package:cloud_firestore/cloud_firestore.dart';

void fetchUsers() async {
  final users = await FirebaseFirestore.instance.collection('test').get();

  for (final user in users.docs) {
    print(user.data());
  }
}
