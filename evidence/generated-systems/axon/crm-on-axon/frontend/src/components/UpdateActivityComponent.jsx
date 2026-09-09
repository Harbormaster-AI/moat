import React, { Component } from 'react'
import ActivityService from '../services/ActivityService';

class UpdateActivityComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                subject: '',
                dueDate: '',
                startAt: '',
                endAt: '',
                location: '',
                activityType: '',
                status: '',
                priority: ''
        }
        this.updateActivity = this.updateActivity.bind(this);

        this.changesubjectHandler = this.changesubjectHandler.bind(this);
        this.changedueDateHandler = this.changedueDateHandler.bind(this);
        this.changestartAtHandler = this.changestartAtHandler.bind(this);
        this.changeendAtHandler = this.changeendAtHandler.bind(this);
        this.changelocationHandler = this.changelocationHandler.bind(this);
        this.changeActivityTypeHandler = this.changeActivityTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changePriorityHandler = this.changePriorityHandler.bind(this);
    }

    componentDidMount(){
        ActivityService.getActivityById(this.state.id).then( (res) =>{
            let activity = res.data;
            this.setState({
                subject: activity.subject,
                dueDate: activity.dueDate,
                startAt: activity.startAt,
                endAt: activity.endAt,
                location: activity.location,
                activityType: activity.activityType,
                status: activity.status,
                priority: activity.priority
            });
        });
    }

    updateActivity = (e) => {
        e.preventDefault();
        let activity = {
            activityId: this.state.id,
            subject: this.state.subject,
            dueDate: this.state.dueDate,
            startAt: this.state.startAt,
            endAt: this.state.endAt,
            location: this.state.location,
            activityType: this.state.activityType,
            status: this.state.status,
            priority: this.state.priority
        };
        console.log('activity => ' + JSON.stringify(activity));
        console.log('id => ' + JSON.stringify(this.state.id));
        ActivityService.updateActivity(activity).then( res => {
            this.props.history.push('/activitys');
        });
    }

    changesubjectHandler= (event) => {
        this.setState({subject: event.target.value});
    }
    changedueDateHandler= (event) => {
        this.setState({dueDate: event.target.value});
    }
    changestartAtHandler= (event) => {
        this.setState({startAt: event.target.value});
    }
    changeendAtHandler= (event) => {
        this.setState({endAt: event.target.value});
    }
    changelocationHandler= (event) => {
        this.setState({location: event.target.value});
    }
    changeActivityTypeHandler= (event) => {
        this.setState({activityType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changePriorityHandler= (event) => {
        this.setState({priority: event.target.value});
    }

    cancel(){
        this.props.history.push('/activitys');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Activity</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> subject: </label>
                                                <input placeholder="subject" name="subject" className="form-control" value={this.state.subject} onChange={this.changesubjectHandler}/>

                                            <label> dueDate: </label>
                                                <input type="date" placeholder="dueDate" name="dueDate" className="form-control" value={this.state.dueDate} onChange={this.changedueDateHandler}/>

                                            <label> startAt: </label>
                                                <input type="time" placeholder="startAt" name="startAt" className="form-control" value={this.state.startAt} onChange={this.changestartAtHandler}/>

                                            <label> endAt: </label>
                                                <input type="time" placeholder="endAt" name="endAt" className="form-control" value={this.state.endAt} onChange={this.changeendAtHandler}/>

                                            <label> location: </label>
                                                <input placeholder="location" name="location" className="form-control" value={this.state.location} onChange={this.changelocationHandler}/>

                                            <label> ActivityType: </label>
                                                <select value={this.state.activityType} onChange={this.changeActivityTypeHandler}>
                      <option name="ActivityType" className="form-control" >
                          Task
                      </option>
                      <option name="ActivityType" className="form-control" >
                          Call
                      </option>
                      <option name="ActivityType" className="form-control" >
                          Meeting
                      </option>
                      <option name="ActivityType" className="form-control" >
                          Demo
                      </option>
                      <option name="ActivityType" className="form-control" >
                          FollowUp
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          NotStarted
                      </option>
                      <option name="Status" className="form-control" >
                          InProgress
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          Deferred
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                            <label> Priority: </label>
                                                <select value={this.state.priority} onChange={this.changePriorityHandler}>
                      <option name="Priority" className="form-control" >
                          Low
                      </option>
                      <option name="Priority" className="form-control" >
                          Normal
                      </option>
                      <option name="Priority" className="form-control" >
                          High
                      </option>
                      <option name="Priority" className="form-control" >
                          Urgent
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateActivity}>Save</button>
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

export default UpdateActivityComponent
