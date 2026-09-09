import React, { Component } from 'react'
import CarePlanService from '../services/CarePlanService'

class ViewCarePlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            carePlan: {}
        }
    }

    componentDidMount(){
        CarePlanService.getCarePlanById(this.state.id).then( res => {
            this.setState({carePlan: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View CarePlan Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> planNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.carePlan.planNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> goalSummary:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.carePlan.goalSummary }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.carePlan.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCarePlanComponent
