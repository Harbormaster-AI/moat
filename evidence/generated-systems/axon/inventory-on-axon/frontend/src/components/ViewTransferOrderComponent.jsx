import React, { Component } from 'react'
import TransferOrderService from '../services/TransferOrderService'

class ViewTransferOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            transferOrder: {}
        }
    }

    componentDidMount(){
        TransferOrderService.getTransferOrderById(this.state.id).then( res => {
            this.setState({transferOrder: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View TransferOrder Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> orderNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.transferOrder.orderNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> requestedShipDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.transferOrder.requestedShipDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> requestedReceiveDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.transferOrder.requestedReceiveDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> shippedDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.transferOrder.shippedDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> receivedDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.transferOrder.receivedDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.transferOrder.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewTransferOrderComponent
