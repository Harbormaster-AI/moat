import React, { Component } from 'react'
import ClaimPaymentService from '../services/ClaimPaymentService'

class ViewClaimPaymentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            claimPayment: {}
        }
    }

    componentDidMount(){
        ClaimPaymentService.getClaimPaymentById(this.state.id).then( res => {
            this.setState({claimPayment: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ClaimPayment Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> paymentNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.claimPayment.paymentNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> amount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.claimPayment.amount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> paymentDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.claimPayment.paymentDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> PayeeType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.claimPayment.payeeType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Method:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.claimPayment.method }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.claimPayment.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewClaimPaymentComponent
