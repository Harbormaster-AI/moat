import React, { Component } from 'react'
import OnboardingTaskService from '../services/OnboardingTaskService';

class CreateOnboardingTaskComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                taskNumber: '',
                name: '',
                dueDate: '',
                status: ''
        }
        this.changetaskNumberHandler = this.changetaskNumberHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changedueDateHandler = this.changedueDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateOnboardingTask = (e) => {
        e.preventDefault();
        let onboardingTask = {
                onboardingTaskId: this.state.id,
                taskNumber: this.state.taskNumber,
                name: this.state.name,
                dueDate: this.state.dueDate,
                status: this.state.status
            };
        console.log('onboardingTask => ' + JSON.stringify(onboardingTask));

        // step 5
        if(this.state.id === '_add'){
            onboardingTask.onboardingTaskId=''
            OnboardingTaskService.createOnboardingTask(onboardingTask).then(res =>{
                this.props.history.push('/onboardingTasks');
            });
        }else{
            OnboardingTaskService.updateOnboardingTask(onboardingTask).then( res => {
                this.props.history.push('/onboardingTasks');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add OnboardingTask</h3>
        }else{
            return <h3 className="text-center">Update OnboardingTask</h3>
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
                                            <label> taskNumber:&emsp; </label>
                                                <input placeholder="taskNumber" name="taskNumber" className="form-control" value={this.state.taskNumber} onChange={this.changetaskNumberHandler}/>

                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> dueDate:&emsp; </label>
                                                <input type="date" placeholder="dueDate" name="dueDate" className="form-control" value={this.state.dueDate} onChange={this.changedueDateHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateOnboardingTask}>Save</button>
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

export default CreateOnboardingTaskComponent
