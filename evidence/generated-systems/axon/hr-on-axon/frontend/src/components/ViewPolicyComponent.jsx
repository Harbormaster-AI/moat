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
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.policy.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> effectiveDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.policy.effectiveDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> description:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.policy.description }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPolicyComponent
