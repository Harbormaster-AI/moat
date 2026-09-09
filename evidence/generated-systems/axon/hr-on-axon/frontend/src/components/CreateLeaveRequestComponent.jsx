import React, { Component } from 'react'
import LeaveRequestService from '../services/LeaveRequestService';

class CreateLeaveRequestComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                requestNumber: '',
                startDate: '',
                endDate: '',
                reason: '',
                hours: '',
                status: ''
        }
        this.changerequestNumberHandler = this.changerequestNumberHandler.bind(this);
        this.changestartDateHandler = this.changestartDateHandler.bind(this);
        this.changeendDateHandler = this.changeendDateHandler.bind(this);
        this.changereasonHandler = this.changereasonHandler.bind(this);
        this.changehoursHandler = this.changehoursHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateLeaveRequest = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            leaveRequest.leaveRequestId=''
            LeaveRequestService.createLeaveRequest(leaveRequest).then(res =>{
                this.props.history.push('/leaveRequests');
            });
        }else{
            LeaveRequestService.updateLeaveRequest(leaveRequest).then( res => {
                this.props.history.push('/leaveRequests');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add LeaveRequest</h3>
        }else{
            return <h3 className="text-center">Update LeaveRequest</h3>
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
                                            <label> requestNumber:&emsp; </label>
                                                <input placeholder="requestNumber" name="requestNumber" className="form-control" value={this.state.requestNumber} onChange={this.changerequestNumberHandler}/>

                                            <label> startDate:&emsp; </label>
                                                <input type="date" placeholder="startDate" name="startDate" className="form-control" value={this.state.startDate} onChange={this.changestartDateHandler}/>

                                            <label> endDate:&emsp; </label>
                                                <input type="date" placeholder="endDate" name="endDate" className="form-control" value={this.state.endDate} onChange={this.changeendDateHandler}/>

                                            <label> reason:&emsp; </label>
                                                <input placeholder="reason" name="reason" className="form-control" value={this.state.reason} onChange={this.changereasonHandler}/>

                                            <label> hours:&emsp; </label>
                                                <input placeholder="hours" name="hours" className="form-control" value={this.state.hours} onChange={this.changehoursHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateLeaveRequest}>Save</button>
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

export default CreateLeaveRequestComponent
