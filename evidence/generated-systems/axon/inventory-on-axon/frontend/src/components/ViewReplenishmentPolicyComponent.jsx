import React, { Component } from 'react'
import ReplenishmentPolicyService from '../services/ReplenishmentPolicyService'

class ViewReplenishmentPolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            replenishmentPolicy: {}
        }
    }

    componentDidMount(){
        ReplenishmentPolicyService.getReplenishmentPolicyById(this.state.id).then( res => {
            this.setState({replenishmentPolicy: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ReplenishmentPolicy Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> minLevel:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.replenishmentPolicy.minLevel }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> maxLevel:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.replenishmentPolicy.maxLevel }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> reorderPoint:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.replenishmentPolicy.reorderPoint }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> reorderQuantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.replenishmentPolicy.reorderQuantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> leadTimeDays:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.replenishmentPolicy.leadTimeDays }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> reviewPeriodDays:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.replenishmentPolicy.reviewPeriodDays }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> PolicyType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.replenishmentPolicy.policyType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewReplenishmentPolicyComponent
