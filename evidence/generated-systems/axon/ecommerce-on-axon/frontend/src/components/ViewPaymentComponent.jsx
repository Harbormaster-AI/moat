import React, { Component } from 'react'
import PaymentService from '../services/PaymentService'

class ViewPaymentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            payment: {}
        }
    }

    componentDidMount(){
        PaymentService.getPaymentById(this.state.id).then( res => {
            this.setState({payment: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Payment Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> paymentNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payment.paymentNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> amount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payment.amount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> transactionId:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payment.transactionId }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> authorizedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payment.authorizedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> capturedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payment.capturedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payment.status }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> PaymentMethod:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payment.paymentMethod }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPaymentComponent
