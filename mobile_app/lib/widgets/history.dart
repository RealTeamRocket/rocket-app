import 'package:flutter/material.dart';
import 'package:fl_chart/fl_chart.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import '/constants/constants.dart';
import '/utils/backend_api/user_api.dart';

class History extends StatefulWidget {
  const History({super.key});

  @override
  State<History> createState() => _HistoryState();
}

class _HistoryState extends State<History> {
  late Future<List<UserStatistics>> _userStatisticsFuture = Future.value([]);

  @override
  void initState() {
    super.initState();

    final storage = FlutterSecureStorage();
    storage.read(key: 'jwt_token').then((jwt) {
      if (jwt != null) {
        setState(() {
          _userStatisticsFuture = UserApi.fetchUserStatistics(jwt);
        });
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const Text(
            'History',
            style: TextStyle(
              color: ColorConstants.white,
              fontSize: 28,
              fontWeight: FontWeight.bold,
            ),
          ),
          const SizedBox(height: 16),
          Expanded(
            child: FutureBuilder<List<UserStatistics>>(
              future: _userStatisticsFuture,
              builder: (context, snapshot) {
                if (snapshot.connectionState == ConnectionState.waiting) {
                  return const Center(child: CircularProgressIndicator());
                } else if (snapshot.hasError) {
                  return Center(
                    child: Text(
                      'Error: ${snapshot.error}',
                      style: const TextStyle(color: Colors.red),
                    ),
                  );
                } else if (!snapshot.hasData || snapshot.data!.isEmpty) {
                  return const Center(
                    child: Text(
                      'No data available',
                      style: TextStyle(color: Colors.white),
                    ),
                  );
                }

                final userStatistics = snapshot.data!;
                final stepsData = userStatistics.map((e) => e.steps.toDouble()).toList();
                final averageSteps =
                    stepsData.reduce((a, b) => a + b) / stepsData.length;

                return Column(
                  children: [
                    Expanded(
                      child: LineChart(
                        LineChartData(
                          gridData: FlGridData(show: true),
                          titlesData: FlTitlesData(
                            leftTitles: AxisTitles(
                              sideTitles: SideTitles(showTitles: true),
                            ),
                            bottomTitles: AxisTitles(
                              sideTitles: SideTitles(
                                showTitles: true,
                                getTitlesWidget: (value, meta) {
                                  final days = userStatistics
                                      .map((e) => e.day)
                                      .toList();
                                  return Text(
                                    days[value.toInt() % days.length],
                                    style: const TextStyle(
                                      color: Colors.white,
                                      fontSize: 12,
                                    ),
                                  );
                                },
                              ),
                            ),
                          ),
                          borderData: FlBorderData(show: true),
                          lineBarsData: [
                            LineChartBarData(
                              spots: List.generate(
                                stepsData.length,
                                (index) => FlSpot(index.toDouble(), stepsData[index]),
                              ),
                              isCurved: true,
                              color: ColorConstants.greenColor,
                              barWidth: 4,
                              isStrokeCapRound: true,
                              belowBarData: BarAreaData(show: false),
                            ),
                          ],
                        ),
                      ),
                    ),
                    const SizedBox(height: 16),
                    Text(
                      '7-Day Average: ${averageSteps.toStringAsFixed(0)} steps',
                      style: const TextStyle(
                        color: ColorConstants.white,
                        fontSize: 18,
                      ),
                    ),
                  ],
                );
              },
            ),
          ),
        ],
      ),
    );
  }
}
