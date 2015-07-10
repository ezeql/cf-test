import React from 'react';
import { Chart } from 'react-google-charts';
import TransactionStore from '../stores/TransactionStore';

export default class GeoMap extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            'BarChart': {
                data: [],
                chartType: "",
                options: {}
            }
        };
        this.change = this.change.bind(this);
    }

    render() {
        return (
            <div className="Geomap">
                <Chart  chartType={this.state.BarChart.chartType} 
                        width={"1000px"} height={"600px"} 
                        data={this.state.BarChart.data} 
                        options={this.state.BarChart.options} 
                        graph_id={this.state.BarChart.div_id} />
            </div>);
    }

    change() {
        var BarChartData = {
            dataArray: [
                ['Country', 'Transactions Count'],
            ],
            options: {
                magnifyingGlass: {
                    enable: true,
                    zoomFactor: 7.5
                },
                colorAxis: {
                    colors: ['lightgreen', 'cyan','yellow', '#e31b23'],
                    maxValue: 1000
                },
                enableRegionInteractivity: true
            }
        };

        var BarChart = {
            data: BarChartData.dataArray.concat(TransactionStore.countriesDataArray()),
            options: BarChartData.options,
            chartType: "GeoChart",
            div_id: "GeoChart"
        };

        this.setState({
            'BarChart': BarChart
        });
    }

    componentDidMount() {
        this.change();
        this.interval = setInterval(this.change, 5000);
    }

    componentWillUnmount() {
        clearInterval(this.interval);
    }
}