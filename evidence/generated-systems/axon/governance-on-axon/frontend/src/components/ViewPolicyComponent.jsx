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
                            <div className = "col" style={{textAlign:"right"}}><label> title:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.policy.title }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> versionLabel:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.policy.versionLabel }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> approvalDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.policy.approvalDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> nextReviewDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.policy.nextReviewDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> documentUrl:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.policy.documentUrl }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> PolicyType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.policy.policyType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.policy.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPolicyComponent
