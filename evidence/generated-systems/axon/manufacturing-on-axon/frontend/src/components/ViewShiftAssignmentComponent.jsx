import React, { Component } from 'react'
import ShiftAssignmentService from '../services/ShiftAssignmentService'

class ViewShiftAssignmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            shiftAssignment: {}
        }
    }

    componentDidMount(){
        ShiftAssignmentService.getShiftAssignmentById(this.state.id).then( res => {
            this.setState({shiftAssignment: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ShiftAssignment Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> assignmentDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.shiftAssignment.assignmentDate }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewShiftAssignmentComponent
