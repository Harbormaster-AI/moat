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
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.forecast.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> horizon:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.forecast.horizon }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Granularity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.forecast.granularity }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewForecastComponent
