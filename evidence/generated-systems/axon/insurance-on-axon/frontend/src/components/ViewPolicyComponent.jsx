import React, { Component } from 'react'
import PolicyService from '../services/PolicyService'

class ViewPolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            policy: {}
        }
    }

    componentDidMount(){
        PolicyService.getPolicyById(this.state.id).then( res => {
            this.setState({policy: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Policy Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> policyNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.policy.policyNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> effectivePeriod:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.policy.effectivePeriod }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> totalPremium:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.policy.totalPremium }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.policy.status }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> PaymentPlan:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.policy.paymentPlan }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPolicyComponent
