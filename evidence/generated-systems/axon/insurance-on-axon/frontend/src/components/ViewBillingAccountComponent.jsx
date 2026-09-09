import React, { Component } from 'react'
import BillingAccountService from '../services/BillingAccountService'

class ViewBillingAccountComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            billingAccount: {}
        }
    }

    componentDidMount(){
        BillingAccountService.getBillingAccountById(this.state.id).then( res => {
            this.setState({billingAccount: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View BillingAccount Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> accountNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.billingAccount.accountNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> balance:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.billingAccount.balance }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.billingAccount.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewBillingAccountComponent
