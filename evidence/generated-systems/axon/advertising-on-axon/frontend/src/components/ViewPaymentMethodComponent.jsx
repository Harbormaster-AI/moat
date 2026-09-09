import React, { Component } from 'react'
import PaymentMethodService from '../services/PaymentMethodService'

class ViewPaymentMethodComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            paymentMethod: {}
        }
    }

    componentDidMount(){
        PaymentMethodService.getPaymentMethodById(this.state.id).then( res => {
            this.setState({paymentMethod: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View PaymentMethod Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> last4:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.paymentMethod.last4 }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> cardholderName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.paymentMethod.cardholderName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> billingAddress:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.paymentMethod.billingAddress }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> MethodType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.paymentMethod.methodType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPaymentMethodComponent
