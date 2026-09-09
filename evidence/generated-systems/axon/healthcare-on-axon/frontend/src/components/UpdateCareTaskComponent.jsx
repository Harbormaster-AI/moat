import React, { Component } from 'react'
import CareTaskService from '../services/CareTaskService';

class UpdateCareTaskComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                description: '',
                dueDate: '',
                status: '',
                priority: ''
        }
        this.updateCareTask = this.updateCareTask.bind(this);

        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changedueDateHandler = this.changedueDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changePriorityHandler = this.changePriorityHandler.bind(this);
    }

    componentDidMount(){
        CareTaskService.getCareTaskById(this.state.id).then( (res) =>{
            let careTask = res.data;
            this.setState({
                description: careTask.description,
                dueDate: careTask.dueDate,
                status: careTask.status,
                priority: careTask.priority
            });
        });
    }

    updateCareTask = (e) => {
        e.preventDefault();
        let careTask = {
            careTaskId: this.state.id,
            description: this.state.description,
            dueDate: this.state.dueDate,
            status: this.state.status,
            priority: this.state.priority
        };
        console.log('careTask => ' + JSON.stringify(careTask));
        console.log('id => ' + JSON.stringify(this.state.id));
        CareTaskService.updateCareTask(careTask).then( res => {
            this.props.history.push('/careTasks');
        });
    }

    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }
    changedueDateHandler= (event) => {
        this.setState({dueDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changePriorityHandler= (event) => {
        this.setState({priority: event.target.value});
    }

    cancel(){
        this.props.history.push('/careTasks');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update CareTask</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> description: </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> dueDate: </label>
                                                <input type="date" placeholder="dueDate" name="dueDate" className="form-control" value={this.state.dueDate} onChange={this.changedueDateHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Requested
                      </option>
                      <option name="Status" className="form-control" >
                          Accepted
                      </option>
                      <option name="Status" className="form-control" >
                          InProgress
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                      <option name="Status" className="form-control" >
                          Failed
                      </option>
                    </select>

                                            <label> Priority: </label>
                                                <select value={this.state.priority} onChange={this.changePriorityHandler}>
                      <option name="Priority" className="form-control" >
                          Routine
                      </option>
                      <option name="Priority" className="form-control" >
                          Urgent
                      </option>
                      <option name="Priority" className="form-control" >
                          Stat
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCareTask}>Save</button>
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

export default UpdateCareTaskComponent
