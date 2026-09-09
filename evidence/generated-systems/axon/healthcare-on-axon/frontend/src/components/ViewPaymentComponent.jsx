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
                            <div className = "col" style={{textAlign:"right"}}><label> paymentDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payment.paymentDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Method:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payment.method }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPaymentComponent
