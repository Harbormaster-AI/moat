import React, { Component } from 'react'
import FlightHealthEventService from '../services/FlightHealthEventService'

class ViewFlightHealthEventComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            flightHealthEvent: {}
        }
    }

    componentDidMount(){
        FlightHealthEventService.getFlightHealthEventById(this.state.id).then( res => {
            this.setState({flightHealthEvent: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View FlightHealthEvent Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> eventCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.flightHealthEvent.eventCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Severity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.flightHealthEvent.severity }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewFlightHealthEventComponent
