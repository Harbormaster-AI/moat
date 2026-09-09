import React, { Component } from 'react'
import EmploymentAssignmentService from '../services/EmploymentAssignmentService'

class ViewEmploymentAssignmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            employmentAssignment: {}
        }
    }

    componentDidMount(){
        EmploymentAssignmentService.getEmploymentAssignmentById(this.state.id).then( res => {
            this.setState({employmentAssignment: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View EmploymentAssignment Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> startDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.employmentAssignment.startDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> endDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.employmentAssignment.endDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> primary:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.employmentAssignment.primary }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> AssignmentType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.employmentAssignment.assignmentType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.employmentAssignment.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewEmploymentAssignmentComponent
