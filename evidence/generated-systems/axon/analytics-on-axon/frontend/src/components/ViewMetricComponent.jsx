import React, { Component } from 'react'
import MetricService from '../services/MetricService'

class ViewMetricComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            metric: {}
        }
    }

    componentDidMount(){
        MetricService.getMetricById(this.state.id).then( res => {
            this.setState({metric: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Metric Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.metric.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> expression:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.metric.expression }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> unit:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.metric.unit }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> MetricType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.metric.metricType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewMetricComponent
