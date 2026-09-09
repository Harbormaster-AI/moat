import React, { Component } from 'react'
import GoalService from '../services/GoalService'

class ListGoalComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                goals: []
        }
        this.addGoal = this.addGoal.bind(this);
        this.editGoal = this.editGoal.bind(this);
        this.deleteGoal = this.deleteGoal.bind(this);
    }

    deleteGoal(id){
        GoalService.deleteGoal(id).then( res => {
            this.setState({goals: this.state.goals.filter(goal => goal.goalId !== id)});
        });
    }
    viewGoal(id){
        this.props.history.push(`/view-goal/${id}`);
    }
    editGoal(id){
        this.props.history.push(`/add-goal/${id}`);
    }

    componentDidMount(){
        GoalService.getGoals().then((res) => {
            this.setState({ goals: res.data});
        });
    }

    addGoal(){
        this.props.history.push('/add-goal/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Goal List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addGoal}> Add Goal</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Title </th>
                                    <th> Description </th>
                                    <th> TargetDate </th>
                                    <th> Weight </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.goals.map(
                                        goal => 
                                        <tr key = {goal.goalId}>
                                             <td> { goal.title } </td>
                                             <td> { goal.description } </td>
                                             <td> { goal.targetDate } </td>
                                             <td> { goal.weight } </td>
                                             <td> { goal.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editGoal(goal.goalId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteGoal(goal.goalId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewGoal(goal.goalId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListGoalComponent
