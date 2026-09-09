import React, { Component } from 'react'
import InspectionPlanService from '../services/InspectionPlanService'

class ViewInspectionPlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            inspectionPlan: {}
        }
    }

    componentDidMount(){
        InspectionPlanService.getInspectionPlanById(this.state.id).then( res => {
            this.setState({inspectionPlan: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View InspectionPlan Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> planNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inspectionPlan.planNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> revision:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inspectionPlan.revision }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> SamplingPlan:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inspectionPlan.samplingPlan }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inspectionPlan.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewInspectionPlanComponent
