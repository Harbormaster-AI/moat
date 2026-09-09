import React, { Component } from 'react'
import GoalService from '../services/GoalService';

class UpdateGoalComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                title: '',
                description: '',
                targetDate: '',
                weight: '',
                status: ''
        }
        this.updateGoal = this.updateGoal.bind(this);

        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changetargetDateHandler = this.changetargetDateHandler.bind(this);
        this.changeweightHandler = this.changeweightHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        GoalService.getGoalById(this.state.id).then( (res) =>{
            let goal = res.data;
            this.setState({
                title: goal.title,
                description: goal.description,
                targetDate: goal.targetDate,
                weight: goal.weight,
                status: goal.status
            });
        });
    }

    updateGoal = (e) => {
        e.preventDefault();
        let goal = {
            goalId: this.state.id,
            title: this.state.title,
            description: this.state.description,
            targetDate: this.state.targetDate,
            weight: this.state.weight,
            status: this.state.status
        };
        console.log('goal => ' + JSON.stringify(goal));
        console.log('id => ' + JSON.stringify(this.state.id));
        GoalService.updateGoal(goal).then( res => {
            this.props.history.push('/goals');
        });
    }

    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }
    changetargetDateHandler= (event) => {
        this.setState({targetDate: event.target.value});
    }
    changeweightHandler= (event) => {
        this.setState({weight: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/goals');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Goal</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> title: </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> description: </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> targetDate: </label>
                                                <input type="date" placeholder="targetDate" name="targetDate" className="form-control" value={this.state.targetDate} onChange={this.changetargetDateHandler}/>

                                            <label> weight: </label>
                                                <input placeholder="weight" name="weight" className="form-control" value={this.state.weight} onChange={this.changeweightHandler}/>

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

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateGoal}>Save</button>
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

export default UpdateGoalComponent
