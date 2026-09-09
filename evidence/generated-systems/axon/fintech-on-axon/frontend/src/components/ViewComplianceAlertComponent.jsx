import React, { Component } from 'react'
import ComplianceAlertService from '../services/ComplianceAlertService'

class ViewComplianceAlertComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            complianceAlert: {}
        }
    }

    componentDidMount(){
        ComplianceAlertService.getComplianceAlertById(this.state.id).then( res => {
            this.setState({complianceAlert: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ComplianceAlert Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> alertCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.complianceAlert.alertCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> raisedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.complianceAlert.raisedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> notes:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.complianceAlert.notes }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Severity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.complianceAlert.severity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.complianceAlert.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewComplianceAlertComponent
