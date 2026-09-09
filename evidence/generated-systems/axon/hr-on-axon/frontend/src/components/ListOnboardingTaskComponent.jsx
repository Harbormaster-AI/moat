import React, { Component } from 'react'
import OnboardingTaskService from '../services/OnboardingTaskService'

class ListOnboardingTaskComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                onboardingTasks: []
        }
        this.addOnboardingTask = this.addOnboardingTask.bind(this);
        this.editOnboardingTask = this.editOnboardingTask.bind(this);
        this.deleteOnboardingTask = this.deleteOnboardingTask.bind(this);
    }

    deleteOnboardingTask(id){
        OnboardingTaskService.deleteOnboardingTask(id).then( res => {
            this.setState({onboardingTasks: this.state.onboardingTasks.filter(onboardingTask => onboardingTask.onboardingTaskId !== id)});
        });
    }
    viewOnboardingTask(id){
        this.props.history.push(`/view-onboardingTask/${id}`);
    }
    editOnboardingTask(id){
        this.props.history.push(`/add-onboardingTask/${id}`);
    }

    componentDidMount(){
        OnboardingTaskService.getOnboardingTasks().then((res) => {
            this.setState({ onboardingTasks: res.data});
        });
    }

    addOnboardingTask(){
        this.props.history.push('/add-onboardingTask/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">OnboardingTask List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addOnboardingTask}> Add OnboardingTask</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> TaskNumber </th>
                                    <th> Name </th>
                                    <th> DueDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.onboardingTasks.map(
                                        onboardingTask => 
                                        <tr key = {onboardingTask.onboardingTaskId}>
                                             <td> { onboardingTask.taskNumber } </td>
                                             <td> { onboardingTask.name } </td>
                                             <td> { onboardingTask.dueDate } </td>
                                             <td> { onboardingTask.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editOnboardingTask(onboardingTask.onboardingTaskId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteOnboardingTask(onboardingTask.onboardingTaskId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewOnboardingTask(onboardingTask.onboardingTaskId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListOnboardingTaskComponent
