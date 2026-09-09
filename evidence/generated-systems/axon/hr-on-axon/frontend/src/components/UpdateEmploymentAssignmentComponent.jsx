import React, { Component } from 'react'
import EmploymentAssignmentService from '../services/EmploymentAssignmentService';

class UpdateEmploymentAssignmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                startDate: '',
                endDate: '',
                primary: '',
                assignmentType: '',
                status: ''
        }
        this.updateEmploymentAssignment = this.updateEmploymentAssignment.bind(this);

        this.changestartDateHandler = this.changestartDateHandler.bind(this);
        this.changeendDateHandler = this.changeendDateHandler.bind(this);
        this.changeprimaryHandler = this.changeprimaryHandler.bind(this);
        this.changeAssignmentTypeHandler = this.changeAssignmentTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        EmploymentAssignmentService.getEmploymentAssignmentById(this.state.id).then( (res) =>{
            let employmentAssignment = res.data;
            this.setState({
                startDate: employmentAssignment.startDate,
                endDate: employmentAssignment.endDate,
                primary: employmentAssignment.primary,
                assignmentType: employmentAssignment.assignmentType,
                status: employmentAssignment.status
            });
        });
    }

    updateEmploymentAssignment = (e) => {
        e.preventDefault();
        let employmentAssignment = {
            employmentAssignmentId: this.state.id,
            startDate: this.state.startDate,
            endDate: this.state.endDate,
            primary: this.state.primary,
            assignmentType: this.state.assignmentType,
            status: this.state.status
        };
        console.log('employmentAssignment => ' + JSON.stringify(employmentAssignment));
        console.log('id => ' + JSON.stringify(this.state.id));
        EmploymentAssignmentService.updateEmploymentAssignment(employmentAssignment).then( res => {
            this.props.history.push('/employmentAssignments');
        });
    }

    changestartDateHandler= (event) => {
        this.setState({startDate: event.target.value});
    }
    changeendDateHandler= (event) => {
        this.setState({endDate: event.target.value});
    }
    changeprimaryHandler= (event) => {
        this.setState({primary: event.target.value});
    }
    changeAssignmentTypeHandler= (event) => {
        this.setState({assignmentType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/employmentAssignments');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update EmploymentAssignment</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> startDate: </label>
                                                <input type="date" placeholder="startDate" name="startDate" className="form-control" value={this.state.startDate} onChange={this.changestartDateHandler}/>

                                            <label> endDate: </label>
                                                <input type="date" placeholder="endDate" name="endDate" className="form-control" value={this.state.endDate} onChange={this.changeendDateHandler}/>

                                            <label> primary: </label>
                                                <input type="checkbox" placeholder="primary" name="primary" className="form-control" value={this.state.primary} onChange={this.changeprimaryHandler}/>


                                            <label> AssignmentType: </label>
                                                <select value={this.state.assignmentType} onChange={this.changeAssignmentTypeHandler}>
                      <option name="AssignmentType" className="form-control" >
                          Primary
                      </option>
                      <option name="AssignmentType" className="form-control" >
                          Secondary
                      </option>
                      <option name="AssignmentType" className="form-control" >
                          Temporary
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Planned
                      </option>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateEmploymentAssignment}>Save</button>
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

export default UpdateEmploymentAssignmentComponent
