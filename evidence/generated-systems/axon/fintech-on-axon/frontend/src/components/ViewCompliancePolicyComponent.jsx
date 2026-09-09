import React, { Component } from 'react'
import CompliancePolicyService from '../services/CompliancePolicyService'

class ViewCompliancePolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            compliancePolicy: {}
        }
    }

    componentDidMount(){
        CompliancePolicyService.getCompliancePolicyById(this.state.id).then( res => {
            this.setState({compliancePolicy: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View CompliancePolicy Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.compliancePolicy.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> policyCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.compliancePolicy.policyCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> description:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.compliancePolicy.description }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.compliancePolicy.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCompliancePolicyComponent
