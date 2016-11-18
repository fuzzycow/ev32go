$(function () {
    var series1 = {
        label: "Series 1",
        values: [{time: 0, y: 0}, ]
    }
    var series2 = {
        label: "Series 1",
        values: [{time: 0, y: 0}, ]
    }

    var lineChartData = [series1, series2];

    $('#lineChart').epoch({
        type: 'time.line',
        data: lineChartData
    });

    for (var y = 0; y <= 1000; y += 1) {
        lineChartData.push({time: y, y: Math.random()})
    }

    $('#gaugeChart').epoch({
        type: 'time.gauge',
        value: 0.5,
        ticks: 10
    });

});
