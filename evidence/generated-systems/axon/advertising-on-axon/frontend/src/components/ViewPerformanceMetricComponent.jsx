import React, { Component } from 'react'
import PerformanceMetricService from '../services/PerformanceMetricService'

class ViewPerformanceMetricComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            performanceMetric: {}
        }
    }

    componentDidMount(){
        PerformanceMetricService.getPerformanceMetricById(this.state.id).then( res => {
            this.setState({performanceMetric: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View PerformanceMetric Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> date:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.performanceMetric.date }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> value:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.performanceMetric.value }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> MetricType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.performanceMetric.metricType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPerformanceMetricComponent
