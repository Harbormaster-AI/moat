import React, { Component } from 'react'
import ShiftAssignmentService from '../services/ShiftAssignmentService';

class UpdateShiftAssignmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                assignmentDate: ''
        }
        this.updateShiftAssignment = this.updateShiftAssignment.bind(this);

        this.changeassignmentDateHandler = this.changeassignmentDateHandler.bind(this);
    }

    componentDidMount(){
        ShiftAssignmentService.getShiftAssignmentById(this.state.id).then( (res) =>{
            let shiftAssignment = res.data;
            this.setState({
                assignmentDate: shiftAssignment.assignmentDate
            });
        });
    }

    updateShiftAssignment = (e) => {
        e.preventDefault();
        let shiftAssignment = {
            shiftAssignmentId: this.state.id,
            assignmentDate: this.state.assignmentDate
        };
        console.log('shiftAssignment => ' + JSON.stringify(shiftAssignment));
        console.log('id => ' + JSON.stringify(this.state.id));
        ShiftAssignmentService.updateShiftAssignment(shiftAssignment).then( res => {
            this.props.history.push('/shiftAssignments');
        });
    }

    changeassignmentDateHandler= (event) => {
        this.setState({assignmentDate: event.target.value});
    }

    cancel(){
        this.props.history.push('/shiftAssignments');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ShiftAssignment</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> assignmentDate: </label>
                                                <input type="date" placeholder="assignmentDate" name="assignmentDate" className="form-control" value={this.state.assignmentDate} onChange={this.changeassignmentDateHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateShiftAssignment}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateShiftAssignmentComponent
