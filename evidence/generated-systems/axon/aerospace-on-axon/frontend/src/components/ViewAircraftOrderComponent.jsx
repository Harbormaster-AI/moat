import React, { Component } from 'react'
import AircraftOrderService from '../services/AircraftOrderService'

class ViewAircraftOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            aircraftOrder: {}
        }
    }

    componentDidMount(){
        AircraftOrderService.getAircraftOrderById(this.state.id).then( res => {
            this.setState({aircraftOrder: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View AircraftOrder Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> orderNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aircraftOrder.orderNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> totalAmount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aircraftOrder.totalAmount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aircraftOrder.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAircraftOrderComponent
