import React, { Component } from 'react'
import GoalService from '../services/GoalService';

class CreateGoalComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                title: '',
                description: '',
                targetDate: '',
                weight: '',
                status: ''
        }
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changetargetDateHandler = this.changetargetDateHandler.bind(this);
        this.changeweightHandler = this.changeweightHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateGoal = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            goal.goalId=''
            GoalService.createGoal(goal).then(res =>{
                this.props.history.push('/goals');
            });
        }else{
            GoalService.updateGoal(goal).then( res => {
                this.props.history.push('/goals');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Goal</h3>
        }else{
            return <h3 className="text-center">Update Goal</h3>
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
                                            <label> title:&emsp; </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> description:&emsp; </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> targetDate:&emsp; </label>
                                                <input type="date" placeholder="targetDate" name="targetDate" className="form-control" value={this.state.targetDate} onChange={this.changetargetDateHandler}/>

                                            <label> weight:&emsp; </label>
                                                <input placeholder="weight" name="weight" className="form-control" value={this.state.weight} onChange={this.changeweightHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateGoal}>Save</button>
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

export default CreateGoalComponent
