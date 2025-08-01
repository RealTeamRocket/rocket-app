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
        crossAxisAlignment: CrossAxisAlignment.center,
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
                }
                if (snapshot.hasError) {
                  return Center(
                    child: Text(
                      'Error:  ${snapshot.error}',
                      style: const TextStyle(color: Colors.red),
                    ),
                  );
                }
                if (!snapshot.hasData || snapshot.data!.isEmpty) {
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

                return Center(
                  child: SizedBox(
                    width: 340,
                    child: Column(
                      mainAxisSize: MainAxisSize.min,
                      crossAxisAlignment: CrossAxisAlignment.center,
                      children: [
                        SizedBox(
                          height: 220,
                          child: LineChart(
                            LineChartData(
                              gridData: FlGridData(
                                show: true,
                                drawVerticalLine: true,
                                horizontalInterval: 1000,
                                getDrawingHorizontalLine: (value) => FlLine(
                                  color: Colors.white24,
                                  strokeWidth: 1,
                                ),
                              ),
                              titlesData: FlTitlesData(
                                topTitles: AxisTitles(
                                  sideTitles: SideTitles(showTitles: false),
                                ),
                                leftTitles: AxisTitles(
                                  sideTitles: SideTitles(showTitles: false),
                                ),
                                rightTitles: AxisTitles(
                                  sideTitles: SideTitles(
                                    showTitles: true,
                                    getTitlesWidget: (value, meta) {
                                      return Padding(
                                        padding: const EdgeInsets.only(left: 8),
                                        child: Text(
                                          value.toInt().toString(),
                                          style: const TextStyle(
                                            color: Colors.white,
                                            fontSize: 12,
                                          ),
                                        ),
                                      );
                                    },
                                    reservedSize: 44,
                                  ),
                                ),
                                bottomTitles: AxisTitles(
                                  sideTitles: SideTitles(
                                    showTitles: true,
                                    interval: 1,
                                    getTitlesWidget: (value, meta) {
                                      final int index = value.round();
                                      if (index < 0 || index >= userStatistics.length) return const SizedBox.shrink();
                                      return Padding(
                                        padding: const EdgeInsets.only(top: 8),
                                        child: Text(
                                          userStatistics[index].day,
                                          style: const TextStyle(
                                            color: Colors.white,
                                            fontSize: 12,
                                          ),
                                        ),
                                      );
                                    },
                                    reservedSize: 32,
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
                                  isCurved: false,
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
                            fontWeight: FontWeight.bold,
                          ),
                        ),
                      ],
                    ),
                  ),
                );
              },
            ),
          ),
        ],
      ),
    );
  }
}
