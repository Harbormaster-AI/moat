import React, { Component } from 'react'
import PaymentContractService from '../services/PaymentContractService'

class ViewPaymentContractComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            paymentContract: {}
        }
    }

    componentDidMount(){
        PaymentContractService.getPaymentContractById(this.state.id).then( res => {
            this.setState({paymentContract: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View PaymentContract Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> contractNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.paymentContract.contractNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> pricingPlanCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.paymentContract.pricingPlanCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.paymentContract.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPaymentContractComponent
