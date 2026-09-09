import React, { Component } from 'react'
import LeavePolicyService from '../services/LeavePolicyService'

class ViewLeavePolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            leavePolicy: {}
        }
    }

    componentDidMount(){
        LeavePolicyService.getLeavePolicyById(this.state.id).then( res => {
            this.setState({leavePolicy: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View LeavePolicy Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.leavePolicy.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> accrualRate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.leavePolicy.accrualRate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> carryoverAllowed:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.leavePolicy.carryoverAllowed }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> maxBalance:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.leavePolicy.maxBalance }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> LeaveCategory:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.leavePolicy.leaveCategory }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> AccrualUnit:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.leavePolicy.accrualUnit }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewLeavePolicyComponent
