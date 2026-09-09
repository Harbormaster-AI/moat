import React, { Component } from 'react'
import ForecastService from '../services/ForecastService'

class ViewForecastComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            forecast: {}
        }
    }

    componentDidMount(){
        ForecastService.getForecastById(this.state.id).then( res => {
            this.setState({forecast: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Forecast Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> forecastNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.forecast.forecastNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> forecastHorizonStart:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.forecast.forecastHorizonStart }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> forecastHorizonEnd:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.forecast.forecastHorizonEnd }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Method:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.forecast.method }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewForecastComponent
