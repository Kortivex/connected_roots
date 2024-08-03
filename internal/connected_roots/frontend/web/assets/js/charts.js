function genCharts(daysOfWeek, chartTitles, chartLegend, toolbox, sensorData) {
    'use strict';

    const weekSunday = sensorData[0];
    const weekMonday = sensorData[1];
    const weekTuesday = sensorData[2];
    const weekWednesday = sensorData[3];
    const weekThursday = sensorData[4];
    const weekFriday = sensorData[5];
    const weekSaturday = sensorData[6];

    const charts = [];

    const createChart = (selector, options) => {
        const {getData} = window.phoenix.utils;
        const $chartEl = document.querySelector(selector);

        if ($chartEl) {
            const userOptions = getData($chartEl, 'echarts');
            const chart = window.echarts.init($chartEl);
            const getDefaultOptions = options;
            echartSetOption(chart, userOptions, getDefaultOptions);
            charts.push(chart);
        }
    };

    const echartSetOption = (
        chart,
        userOptions,
        getDefaultOptions,
        responsiveOptions
    ) => {
        const {breakpoints, resize} = window.phoenix.utils;
        const handleResize = options => {
            Object.keys(options).forEach(item => {
                if (window.innerWidth > breakpoints[item]) {
                    chart.setOption(options[item]);
                }
            });
        };

        const themeController = document.body;
        chart.setOption(window._.merge(getDefaultOptions(), userOptions));

        const navbarVerticalToggle = document.querySelector('.navbar-vertical-toggle');
        if (navbarVerticalToggle) {
            navbarVerticalToggle.addEventListener('navbar.vertical.toggle', () => {
                chart.resize();
                if (responsiveOptions) {
                    handleResize(responsiveOptions);
                }
            });
        }

        resize(() => {
            chart.resize();
            if (responsiveOptions) {
                handleResize(responsiveOptions);
            }
        });
        if (responsiveOptions) {
            handleResize(responsiveOptions);
        }

        themeController.addEventListener('clickControl', ({detail: {control}}) => {
            if (control === 'phoenixTheme') {
                chart.setOption(window._.merge(getDefaultOptions(), userOptions));
            }
            if (responsiveOptions) {
                handleResize(responsiveOptions);
            }
        });
    };

    const initCharts = () => {
        createChart('.temperature-chart', () => ({
            title: {
                text: chartTitles[0] + ' (°C)',
                textStyle: {
                    fontSize: 15
                },
            },
            tooltip: {
                trigger: 'axis',
                axisPointer: {
                    type: 'cross',
                    label: {
                        backgroundColor: '#6a7985'
                    }
                }
            },
            legend: {
                data: [chartLegend[0], chartLegend[1]],
                top: '9%'
            },
            toolbox: {
                feature: {
                    saveAsImage: {
                        title: toolbox[1],
                    },
                    dataView: {
                        title: toolbox[0],
                        readOnly: true,
                        lang: [toolbox[0], 'X']
                    },
                },
            },
            grid: {
                left: '5%',
                right: '5%',
                bottom: '5%',
                containLabel: true
            },
            xAxis: [
                {
                    type: 'category',
                    boundaryGap: false,
                    data: daysOfWeek
                }
            ],
            yAxis: [
                {
                    type: 'value'
                }
            ],
            series: [
                {
                    name: chartLegend[0],
                    type: 'line',
                    stack: 'Total',
                    areaStyle: {},
                    emphasis: {
                        focus: 'series'
                    },
                    data: [
                        weekMonday.avg_temperature_in,
                        weekTuesday.avg_temperature_in,
                        weekWednesday.avg_temperature_in,
                        weekThursday.avg_temperature_in,
                        weekFriday.avg_temperature_in,
                        weekSaturday.avg_temperature_in,
                        weekSunday.avg_temperature_in
                    ]
                },
                {
                    name: chartLegend[1],
                    type: 'line',
                    stack: 'Total',
                    areaStyle: {},
                    emphasis: {
                        focus: 'series'
                    },
                    data: [
                        weekMonday.avg_temperature_out,
                        weekTuesday.avg_temperature_out,
                        weekWednesday.avg_temperature_out,
                        weekThursday.avg_temperature_out,
                        weekFriday.avg_temperature_out,
                        weekSaturday.avg_temperature_out,
                        weekSunday.avg_temperature_out
                    ]
                },
            ],
            color: ['#F44336', '#f88980'] ,
        }));

        createChart('.humidity-chart', () => ({
            title: {
                text: chartTitles[1] + ' (%)',
                textStyle: {
                    fontSize: 15
                },
            },
            tooltip: {
                trigger: 'axis',
                axisPointer: {
                    type: 'cross',
                    label: {
                        backgroundColor: '#6a7985'
                    }
                }
            },
            legend: {
                data: [chartLegend[0], chartLegend[1]],
                top: '9%'
            },
                        toolbox: {
                feature: {
                    saveAsImage: {
                        title: toolbox[1],
                    },
                    dataView: {
                        title: toolbox[0],
                        readOnly: true,
                        lang: [toolbox[0], 'X']
                    },
                },
            },
            grid: {
                left: '5%',
                right: '5%',
                bottom: '5%',
                containLabel: true
            },
            xAxis: [
                {
                    type: 'category',
                    boundaryGap: false,
                    data: daysOfWeek
                }
            ],
            yAxis: [
                {
                    type: 'value'
                }
            ],
            series: [
                {
                    name: chartLegend[0],
                    type: 'line',
                    stack: 'Total',
                    areaStyle: {},
                    emphasis: {
                        focus: 'series'
                    },
                    data: [
                        weekMonday.avg_humidity_in,
                        weekTuesday.avg_humidity_in,
                        weekWednesday.avg_humidity_in,
                        weekThursday.avg_humidity_in,
                        weekFriday.avg_humidity_in,
                        weekSaturday.avg_humidity_in,
                        weekSunday.avg_humidity_in
                    ]
                },
                {
                    name: chartLegend[1],
                    type: 'line',
                    stack: 'Total',
                    areaStyle: {},
                    emphasis: {
                        focus: 'series'
                    },
                    data: [
                        weekMonday.avg_humidity_out,
                        weekTuesday.avg_humidity_out,
                        weekWednesday.avg_humidity_out,
                        weekThursday.avg_humidity_out,
                        weekFriday.avg_humidity_out,
                        weekSaturday.avg_humidity_out,
                        weekSunday.avg_humidity_out
                    ]
                },
            ],
            color: ['#00BCD4', '#89beca'],
        }));

        createChart('.salinity-chart', () => ({
            title: {
                text: chartTitles[3] + ' (%)',
                textStyle: {
                    fontSize: 15
                },
            },
            tooltip: {
                trigger: 'axis',
                axisPointer: {
                    type: 'cross',
                    label: {
                        backgroundColor: '#6a7985'
                    }
                }
            },
            legend: {
                data: [chartTitles[3]],
                top: '9%'
            },
                        toolbox: {
                feature: {
                    saveAsImage: {
                        title: toolbox[1],
                    },
                    dataView: {
                        title: toolbox[0],
                        readOnly: true,
                        lang: [toolbox[0], 'X']
                    },
                },
            },
            grid: {
                left: '5%',
                right: '5%',
                bottom: '5%',
                containLabel: true
            },
            xAxis: [
                {
                    type: 'category',
                    boundaryGap: false,
                    data: daysOfWeek
                }
            ],
            yAxis: [
                {
                    type: 'value'
                }
            ],
            series: [
                {
                    name: chartTitles[3],
                    type: 'line',
                    stack: 'Total',
                    areaStyle: {},
                    emphasis: {
                        focus: 'series'
                    },
                    data: [
                        weekMonday.avg_salt,
                        weekTuesday.avg_salt,
                        weekWednesday.avg_salt,
                        weekThursday.avg_salt,
                        weekFriday.avg_salt,
                        weekSaturday.avg_salt,
                        weekSunday.avg_salt
                    ]
                }
            ],
            color: ['#2196F3'],
        }));

        createChart('.soil-chart', () => ({
            title: {
                text: chartTitles[2] + ' (%)',
                textStyle: {
                    fontSize: 15
                },
            },
            tooltip: {
                trigger: 'axis',
                axisPointer: {
                    type: 'cross',
                    label: {
                        backgroundColor: '#6a7985'
                    }
                }
            },
            legend: {
                data: [chartTitles[2]],
                top: '9%'
            },
                        toolbox: {
                feature: {
                    saveAsImage: {
                        title: toolbox[1],
                    },
                    dataView: {
                        title: toolbox[0],
                        readOnly: true,
                        lang: [toolbox[0], 'X']
                    },
                },
            },
            grid: {
                left: '5%',
                right: '5%',
                bottom: '5%',
                containLabel: true
            },
            xAxis: [
                {
                    type: 'category',
                    boundaryGap: false,
                    data: daysOfWeek
                }
            ],
            yAxis: [
                {
                    type: 'value'
                }
            ],
            series: [
                {
                    name: chartTitles[2],
                    type: 'line',
                    stack: 'Total',
                    areaStyle: {},
                    emphasis: {
                        focus: 'series'
                    },
                    data: [
                        weekMonday.avg_soil,
                        weekTuesday.avg_soil,
                        weekWednesday.avg_soil,
                        weekThursday.avg_soil,
                        weekFriday.avg_soil,
                        weekSaturday.avg_soil,
                        weekSunday.avg_soil
                    ]
                }
            ],
            color: ['#795548'],
        }));

        createChart('.light-chart', () => ({
            title: {
                text: chartTitles[4] + ' (Lx)',
                textStyle: {
                    fontSize: 15
                },
            },
            tooltip: {
                trigger: 'axis',
                axisPointer: {
                    type: 'cross',
                    label: {
                        backgroundColor: '#6a7985'
                    }
                }
            },
            legend: {
                data: [chartTitles[4]],
                top: '9%'
            },
                        toolbox: {
                feature: {
                    saveAsImage: {
                        title: toolbox[1],
                    },
                    dataView: {
                        title: toolbox[0],
                        readOnly: true,
                        lang: [toolbox[0], 'X']
                    },
                },
            },
            grid: {
                left: '5%',
                right: '5%',
                bottom: '5%',
                containLabel: true
            },
            xAxis: [
                {
                    type: 'category',
                    boundaryGap: false,
                    data: daysOfWeek
                }
            ],
            yAxis: [
                {
                    type: 'value'
                }
            ],
            series: [
                {
                    name: chartTitles[4],
                    type: 'line',
                    stack: 'Total',
                    areaStyle: {},
                    emphasis: {
                        focus: 'series'
                    },
                    data: [
                        weekMonday.avg_light,
                        weekTuesday.avg_light,
                        weekWednesday.avg_light,
                        weekThursday.avg_light,
                        weekFriday.avg_light,
                        weekSaturday.avg_light,
                        weekSunday.avg_light
                    ]
                }
            ],
            color: ['#f5e87a'],
        }));

        createChart('.altitude-chart', () => ({
            title: {
                text: chartTitles[5] + ' (m)',
                textStyle: {
                    fontSize: 15
                },
            },
            tooltip: {
                trigger: 'axis',
                axisPointer: {
                    type: 'cross',
                    label: {
                        backgroundColor: '#6a7985'
                    }
                }
            },
            legend: {
                data: [chartTitles[5]],
                top: '9%'
            },
                        toolbox: {
                feature: {
                    saveAsImage: {
                        title: toolbox[1],
                    },
                    dataView: {
                        title: toolbox[0],
                        readOnly: true,
                        lang: [toolbox[0], 'X']
                    },
                },
            },
            grid: {
                left: '5%',
                right: '5%',
                bottom: '5%',
                containLabel: true
            },
            xAxis: [
                {
                    type: 'category',
                    boundaryGap: false,
                    data: daysOfWeek
                }
            ],
            yAxis: [
                {
                    type: 'value'
                }
            ],
            series: [
                {
                    name: chartTitles[5],
                    type: 'line',
                    stack: 'Total',
                    areaStyle: {},
                    emphasis: {
                        focus: 'series'
                    },
                    data: [
                        weekMonday.avg_altitude,
                        weekTuesday.avg_altitude,
                        weekWednesday.avg_altitude,
                        weekThursday.avg_altitude,
                        weekFriday.avg_altitude,
                        weekSaturday.avg_altitude,
                        weekSunday.avg_altitude
                    ]
                }
            ],
            color: ['#9C27B0'],
        }));

        createChart('.pressure-chart', () => ({
            title: {
                text: chartTitles[6] + ' (hPa)',
                textStyle: {
                    fontSize: 15
                },
            },
            tooltip: {
                trigger: 'axis',
                axisPointer: {
                    type: 'cross',
                    label: {
                        backgroundColor: '#6a7985'
                    }
                }
            },
            legend: {
                data: [chartTitles[6]],
                top: '9%'
            },
                        toolbox: {
                feature: {
                    saveAsImage: {
                        title: toolbox[1],
                    },
                    dataView: {
                        title: toolbox[0],
                        readOnly: true,
                        lang: [toolbox[0], 'X']
                    },
                },
            },
            grid: {
                left: '5%',
                right: '5%',
                bottom: '5%',
                containLabel: true
            },
            xAxis: [
                {
                    type: 'category',
                    boundaryGap: false,
                    data: daysOfWeek
                }
            ],
            yAxis: [
                {
                    type: 'value'
                }
            ],
            series: [
                {
                    name: chartTitles[6],
                    type: 'line',
                    stack: 'Total',
                    areaStyle: {},
                    emphasis: {
                        focus: 'series'
                    },
                    data: [
                        weekMonday.avg_pressure,
                        weekTuesday.avg_pressure,
                        weekWednesday.avg_pressure,
                        weekThursday.avg_pressure,
                        weekFriday.avg_pressure,
                        weekSaturday.avg_pressure,
                        weekSunday.avg_pressure
                    ]
                }
            ],
            color: ['#3F51B5'],
        }));

        window.addEventListener('resize', () => {
            charts.forEach(chart => {
                chart.resize();
            });
        });
    };

    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', initCharts);
    } else {
        initCharts();
    }
}

function processSensorData(sensorData) {
    const weekData = new Array(7).fill(null).map(() => ({
        weekday: 0,
        avg_voltage: 0,
        avg_battery: 0,
        avg_soil: 0,
        avg_salt: 0,
        avg_light: 0,
        avg_temperature_in: 0,
        avg_temperature_out: 0,
        avg_humidity_in: 0,
        avg_humidity_out: 0,
        avg_pressure: 0,
        avg_altitude: 0
    }));

    sensorData.forEach(day => {
        const index = day.weekday % 7;
        weekData[index] = day;
    });

    return weekData;
}
