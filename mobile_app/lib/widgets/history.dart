import 'package:flutter/material.dart';
import 'package:fl_chart/fl_chart.dart';
import '/constants/constants.dart';

class History extends StatelessWidget {
  const History({super.key});

  @override
  Widget build(BuildContext context) {
    final List<double> stepsData = [5000, 7000, 8000, 6000, 7500, 9000, 8500];
    final double averageSteps = stepsData.reduce((a, b) => a + b) / stepsData.length;

    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const Text(
            'History',
            style: TextStyle(color: ColorConstants.white, fontSize: 28),
          ),
          const SizedBox(height: 16),
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
                        const days = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'];
                        return Text(
                          days[value.toInt() % days.length],
                          style: const TextStyle(color: Colors.white, fontSize: 12),
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
            style: const TextStyle(color: ColorConstants.white, fontSize: 18),
          ),
        ],
      ),
    );
  }
}