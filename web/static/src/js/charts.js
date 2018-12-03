if ($(".report").length) {

    window.MessagePie = function (data) {


        Highcharts.chart('message-pie', {
            chart: {
                plotBackgroundColor: null,
                plotBorderWidth: null,
                plotShadow: false,
                type: 'pie'
            },
            title: {
                text: 'Messages Sent'
            },
            tooltip: {
                pointFormat: '{series.name}: <b>{point.y} Messages ({point.percentage:.0f}%)</b>'
            },
            plotOptions: {
                pie: {
                    allowPointSelect: true,
                    cursor: 'pointer',
                    dataLabels: {
                        enabled: true,
                        format: '<b>{point.name}</b>',
                        style: {
                            // color: (Highcharts.theme && Highcharts.theme.contrastTextColor) || 'black'
                        }
                    }
                }
            },
            series: [{
                name: 'Messages',
                colorByPoint: true,
                data: data,
            }]
        });
    }

     window.WordPie = function (data) {
        Highcharts.chart('words-pie', {
            chart: {
                plotBackgroundColor: null,
                plotBorderWidth: null,
                plotShadow: false,
                type: 'pie'
            },
            title: {
                text: 'Words Sent'
            },
            tooltip: {
                pointFormat: '{series.name}: <b>{point.y} Words ({point.percentage:.0f}%)</b>'
            },
            plotOptions: {
                pie: {
                    allowPointSelect: true,
                    cursor: 'pointer',
                    dataLabels: {
                        enabled: true,
                        format: '<b>{point.name}</b>',
                        style: {
                            // color: (Highcharts.theme && Highcharts.theme.contrastTextColor) || 'black'
                        }
                    }
                }
            },
            series: [{
                name: 'Words',
                colorByPoint: true,
                data: data,
            }]
        });
    }


    window.MonthlyMessages = function (months, data) {
        Highcharts.chart('monthly-messages', {
            chart: {
                type: 'column'
            },
            title: {
                text: 'Messages sent per month'
            },
            xAxis: {
                categories: months,
                crosshair: true
            },
            yAxis: {
                min: 0,
                title: {
                    text: 'Messages Sent'
                }
            },
            tooltip: {
                headerFormat: '<span style="font-size:10px">{point.key}</span><table>',
                pointFormat: '<tr><td style="color:{series.color};padding:0">{series.name}: </td>' +
                    '<td style="padding:0"><b>{point.y} messages</b></td></tr>',
                footerFormat: '</table>',
                shared: true,
                useHTML: true
            },
            plotOptions: {
                column: {
                    pointPadding: 0.2,
                    borderWidth: 0
                }
            },
            series: data,
            // series: [{
            //     name: 'Tokyo',
            //     // data: [49.9, 71.5, 106.4, 129.2, 144.0, 176.0, 135.6, 148.5, 216.4, 194.1, 95.6, 54.4]
            //
            // }, {
            //     name: 'New York',
            //     // data: [83.6, 78.8, 98.5, 93.4, 106.0, 84.5, 105.0, 104.3, 91.2, 83.5, 106.6, 92.3]
            //
            // }, {
            //     name: 'London',
            //     data: [42, 33, 34, 39, 52, 75, 57, 60, 47, 39, 46]
            //
            // }, {
            //     name: 'Mendy',
            //     data: [["Aug 2017",42]]
            // }]
        });
    }


}