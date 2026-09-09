import React, { Component } from 'react'
import MaintenancePlanService from '../services/MaintenancePlanService'

class ViewMaintenancePlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            maintenancePlan: {}
        }
    }

    componentDidMount(){
        MaintenancePlanService.getMaintenancePlanById(this.state.id).then( res => {
            this.setState({maintenancePlan: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View MaintenancePlan Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> planNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.maintenancePlan.planNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> interval:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.maintenancePlan.interval }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> lastServiceDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.maintenancePlan.lastServiceDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Strategy:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.maintenancePlan.strategy }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewMaintenancePlanComponent
