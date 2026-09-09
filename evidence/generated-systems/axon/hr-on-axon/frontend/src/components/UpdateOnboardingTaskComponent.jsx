import React, { Component } from 'react'
import OnboardingTaskService from '../services/OnboardingTaskService';

class UpdateOnboardingTaskComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                taskNumber: '',
                name: '',
                dueDate: '',
                status: ''
        }
        this.updateOnboardingTask = this.updateOnboardingTask.bind(this);

        this.changetaskNumberHandler = this.changetaskNumberHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changedueDateHandler = this.changedueDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        OnboardingTaskService.getOnboardingTaskById(this.state.id).then( (res) =>{
            let onboardingTask = res.data;
            this.setState({
                taskNumber: onboardingTask.taskNumber,
                name: onboardingTask.name,
                dueDate: onboardingTask.dueDate,
                status: onboardingTask.status
            });
        });
    }

    updateOnboardingTask = (e) => {
        e.preventDefault();
        let onboardingTask = {
            onboardingTaskId: this.state.id,
            taskNumber: this.state.taskNumber,
            name: this.state.name,
            dueDate: this.state.dueDate,
            status: this.state.status
        };
        console.log('onboardingTask => ' + JSON.stringify(onboardingTask));
        console.log('id => ' + JSON.stringify(this.state.id));
        OnboardingTaskService.updateOnboardingTask(onboardingTask).then( res => {
            this.props.history.push('/onboardingTasks');
        });
    }

    changetaskNumberHandler= (event) => {
        this.setState({taskNumber: event.target.value});
    }
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changedueDateHandler= (event) => {
        this.setState({dueDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/onboardingTasks');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update OnboardingTask</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> taskNumber: </label>
                                                <input placeholder="taskNumber" name="taskNumber" className="form-control" value={this.state.taskNumber} onChange={this.changetaskNumberHandler}/>

                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> dueDate: </label>
                                                <input type="date" placeholder="dueDate" name="dueDate" className="form-control" value={this.state.dueDate} onChange={this.changedueDateHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          NotStarted
                      </option>
                      <option name="Status" className="form-control" >
                          InProgress
                      </option>
                      <option name="Status" className="form-control" >
                          Blocked
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateOnboardingTask}>Save</button>
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

export default UpdateOnboardingTaskComponent
