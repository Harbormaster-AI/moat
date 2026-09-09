import React, { Component } from 'react'
import BillingProfileService from '../services/BillingProfileService'

class ViewBillingProfileComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            billingProfile: {}
        }
    }

    componentDidMount(){
        BillingProfileService.getBillingProfileById(this.state.id).then( res => {
            this.setState({billingProfile: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View BillingProfile Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> billingName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.billingProfile.billingName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> taxId:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.billingProfile.taxId }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> billingAddress:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.billingProfile.billingAddress }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> PaymentTerms:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.billingProfile.paymentTerms }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewBillingProfileComponent
