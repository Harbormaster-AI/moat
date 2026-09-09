import React, { Component } from 'react'
import ForecastLineService from '../services/ForecastLineService'

class ViewForecastLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            forecastLine: {}
        }
    }

    componentDidMount(){
        ForecastLineService.getForecastLineById(this.state.id).then( res => {
            this.setState({forecastLine: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ForecastLine Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> period:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.forecastLine.period }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.forecastLine.quantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> confidence:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.forecastLine.confidence }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewForecastLineComponent
