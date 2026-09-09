import React, { Component } from 'react'
import PaymentOrderService from '../services/PaymentOrderService'

class ViewPaymentOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            paymentOrder: {}
        }
    }

    componentDidMount(){
        PaymentOrderService.getPaymentOrderById(this.state.id).then( res => {
            this.setState({paymentOrder: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View PaymentOrder Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> orderReference:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.paymentOrder.orderReference }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> requestedExecutionDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.paymentOrder.requestedExecutionDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> purpose:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.paymentOrder.purpose }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> PaymentMethod:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.paymentOrder.paymentMethod }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.paymentOrder.status }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Priority:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.paymentOrder.priority }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPaymentOrderComponent
