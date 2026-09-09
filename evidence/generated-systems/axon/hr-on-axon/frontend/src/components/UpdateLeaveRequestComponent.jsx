import React, { Component } from 'react'
import LeaveRequestService from '../services/LeaveRequestService';

class UpdateLeaveRequestComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                requestNumber: '',
                startDate: '',
                endDate: '',
                reason: '',
                hours: '',
                status: ''
        }
        this.updateLeaveRequest = this.updateLeaveRequest.bind(this);

        this.changerequestNumberHandler = this.changerequestNumberHandler.bind(this);
        this.changestartDateHandler = this.changestartDateHandler.bind(this);
        this.changeendDateHandler = this.changeendDateHandler.bind(this);
        this.changereasonHandler = this.changereasonHandler.bind(this);
        this.changehoursHandler = this.changehoursHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        LeaveRequestService.getLeaveRequestById(this.state.id).then( (res) =>{
            let leaveRequest = res.data;
            this.setState({
                requestNumber: leaveRequest.requestNumber,
                startDate: leaveRequest.startDate,
                endDate: leaveRequest.endDate,
                reason: leaveRequest.reason,
                hours: leaveRequest.hours,
                status: leaveRequest.status
            });
        });
    }

    updateLeaveRequest = (e) => {
        e.preventDefault();
        let leaveRequest = {
            leaveRequestId: this.state.id,
            requestNumber: this.state.requestNumber,
            startDate: this.state.startDate,
            endDate: this.state.endDate,
            reason: this.state.reason,
            hours: this.state.hours,
            status: this.state.status
        };
        console.log('leaveRequest => ' + JSON.stringify(leaveRequest));
        console.log('id => ' + JSON.stringify(this.state.id));
        LeaveRequestService.updateLeaveRequest(leaveRequest).then( res => {
            this.props.history.push('/leaveRequests');
        });
    }

    changerequestNumberHandler= (event) => {
        this.setState({requestNumber: event.target.value});
    }
    changestartDateHandler= (event) => {
        this.setState({startDate: event.target.value});
    }
    changeendDateHandler= (event) => {
        this.setState({endDate: event.target.value});
    }
    changereasonHandler= (event) => {
        this.setState({reason: event.target.value});
    }
    changehoursHandler= (event) => {
        this.setState({hours: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/leaveRequests');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update LeaveRequest</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> requestNumber: </label>
                                                <input placeholder="requestNumber" name="requestNumber" className="form-control" value={this.state.requestNumber} onChange={this.changerequestNumberHandler}/>

                                            <label> startDate: </label>
                                                <input type="date" placeholder="startDate" name="startDate" className="form-control" value={this.state.startDate} onChange={this.changestartDateHandler}/>

                                            <label> endDate: </label>
                                                <input type="date" placeholder="endDate" name="endDate" className="form-control" value={this.state.endDate} onChange={this.changeendDateHandler}/>

                                            <label> reason: </label>
                                                <input placeholder="reason" name="reason" className="form-control" value={this.state.reason} onChange={this.changereasonHandler}/>

                                            <label> hours: </label>
                                                <input placeholder="hours" name="hours" className="form-control" value={this.state.hours} onChange={this.changehoursHandler}/>

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
                          Cancelled
                      </option>
                      <option name="Status" className="form-control" >
                          Taken
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateLeaveRequest}>Save</button>
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

export default UpdateLeaveRequestComponent
