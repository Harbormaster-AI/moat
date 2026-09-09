import React, { Component } from 'react'
import ShiftAssignmentService from '../services/ShiftAssignmentService';

class CreateShiftAssignmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                assignmentDate: ''
        }
        this.changeassignmentDateHandler = this.changeassignmentDateHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ShiftAssignmentService.getShiftAssignmentById(this.state.id).then( (res) =>{
                let shiftAssignment = res.data;
                this.setState({
                    assignmentDate: shiftAssignment.assignmentDate
                });
            });
        }        
    }
    saveOrUpdateShiftAssignment = (e) => {
        e.preventDefault();
        let shiftAssignment = {
                shiftAssignmentId: this.state.id,
                assignmentDate: this.state.assignmentDate
            };
        console.log('shiftAssignment => ' + JSON.stringify(shiftAssignment));

        // step 5
        if(this.state.id === '_add'){
            shiftAssignment.shiftAssignmentId=''
            ShiftAssignmentService.createShiftAssignment(shiftAssignment).then(res =>{
                this.props.history.push('/shiftAssignments');
            });
        }else{
            ShiftAssignmentService.updateShiftAssignment(shiftAssignment).then( res => {
                this.props.history.push('/shiftAssignments');
            });
        }
    }
    
    changeassignmentDateHandler= (event) => {
        this.setState({assignmentDate: event.target.value});
    }

    cancel(){
        this.props.history.push('/shiftAssignments');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ShiftAssignment</h3>
        }else{
            return <h3 className="text-center">Update ShiftAssignment</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> assignmentDate:&emsp; </label>
                                                <input type="date" placeholder="assignmentDate" name="assignmentDate" className="form-control" value={this.state.assignmentDate} onChange={this.changeassignmentDateHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateShiftAssignment}>Save</button>
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

export default CreateShiftAssignmentComponent
