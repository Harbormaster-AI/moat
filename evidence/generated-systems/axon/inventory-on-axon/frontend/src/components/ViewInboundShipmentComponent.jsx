import React, { Component } from 'react'
import InboundShipmentService from '../services/InboundShipmentService'

class ViewInboundShipmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            inboundShipment: {}
        }
    }

    componentDidMount(){
        InboundShipmentService.getInboundShipmentById(this.state.id).then( res => {
            this.setState({inboundShipment: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View InboundShipment Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> shipmentNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inboundShipment.shipmentNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> expectedArrivalDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inboundShipment.expectedArrivalDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> arrivalDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inboundShipment.arrivalDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> carrierName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inboundShipment.carrierName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inboundShipment.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewInboundShipmentComponent
