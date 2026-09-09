import React, { Component } from 'react'
import ShipmentService from '../services/ShipmentService'

class ViewShipmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            shipment: {}
        }
    }

    componentDidMount(){
        ShipmentService.getShipmentById(this.state.id).then( res => {
            this.setState({shipment: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Shipment Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> shipmentNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.shipment.shipmentNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> shippedDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.shipment.shippedDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> deliveredDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.shipment.deliveredDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> trackingNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.shipment.trackingNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> shippingAddress:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.shipment.shippingAddress }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.shipment.status }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Carrier:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.shipment.carrier }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewShipmentComponent
