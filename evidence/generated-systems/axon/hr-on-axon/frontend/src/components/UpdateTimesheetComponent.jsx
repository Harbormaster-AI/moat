import React, { Component } from 'react'
import TimesheetService from '../services/TimesheetService';

class UpdateTimesheetComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                periodStart: '',
                periodEnd: '',
                submissionDate: '',
                status: ''
        }
        this.updateTimesheet = this.updateTimesheet.bind(this);

        this.changeperiodStartHandler = this.changeperiodStartHandler.bind(this);
        this.changeperiodEndHandler = this.changeperiodEndHandler.bind(this);
        this.changesubmissionDateHandler = this.changesubmissionDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        TimesheetService.getTimesheetById(this.state.id).then( (res) =>{
            let timesheet = res.data;
            this.setState({
                periodStart: timesheet.periodStart,
                periodEnd: timesheet.periodEnd,
                submissionDate: timesheet.submissionDate,
                status: timesheet.status
            });
        });
    }

    updateTimesheet = (e) => {
        e.preventDefault();
        let timesheet = {
            timesheetId: this.state.id,
            periodStart: this.state.periodStart,
            periodEnd: this.state.periodEnd,
            submissionDate: this.state.submissionDate,
            status: this.state.status
        };
        console.log('timesheet => ' + JSON.stringify(timesheet));
        console.log('id => ' + JSON.stringify(this.state.id));
        TimesheetService.updateTimesheet(timesheet).then( res => {
            this.props.history.push('/timesheets');
        });
    }

    changeperiodStartHandler= (event) => {
        this.setState({periodStart: event.target.value});
    }
    changeperiodEndHandler= (event) => {
        this.setState({periodEnd: event.target.value});
    }
    changesubmissionDateHandler= (event) => {
        this.setState({submissionDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/timesheets');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Timesheet</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> periodStart: </label>
                                                <input type="date" placeholder="periodStart" name="periodStart" className="form-control" value={this.state.periodStart} onChange={this.changeperiodStartHandler}/>

                                            <label> periodEnd: </label>
                                                <input type="date" placeholder="periodEnd" name="periodEnd" className="form-control" value={this.state.periodEnd} onChange={this.changeperiodEndHandler}/>

                                            <label> submissionDate: </label>
                                                <input type="date" placeholder="submissionDate" name="submissionDate" className="form-control" value={this.state.submissionDate} onChange={this.changesubmissionDateHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Submitted
                      </option>
                      <option name="Status" className="form-control" >
                          Approved
                      </option>
                      <option name="Status" className="form-control" >
                          Rejected
                      </option>
                      <option name="Status" className="form-control" >
                          Processed
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateTimesheet}>Save</button>
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

export default UpdateTimesheetComponent
