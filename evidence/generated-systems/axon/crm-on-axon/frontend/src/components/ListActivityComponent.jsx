import React, { Component } from 'react'
import ActivityService from '../services/ActivityService'

class ListActivityComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                activitys: []
        }
        this.addActivity = this.addActivity.bind(this);
        this.editActivity = this.editActivity.bind(this);
        this.deleteActivity = this.deleteActivity.bind(this);
    }

    deleteActivity(id){
        ActivityService.deleteActivity(id).then( res => {
            this.setState({activitys: this.state.activitys.filter(activity => activity.activityId !== id)});
        });
    }
    viewActivity(id){
        this.props.history.push(`/view-activity/${id}`);
    }
    editActivity(id){
        this.props.history.push(`/add-activity/${id}`);
    }

    componentDidMount(){
        ActivityService.getActivitys().then((res) => {
            this.setState({ activitys: res.data});
        });
    }

    addActivity(){
        this.props.history.push('/add-activity/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Activity List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addActivity}> Add Activity</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Subject </th>
                                    <th> DueDate </th>
                                    <th> StartAt </th>
                                    <th> EndAt </th>
                                    <th> Location </th>
                                    <th> ActivityType </th>
                                    <th> Status </th>
                                    <th> Priority </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.activitys.map(
                                        activity => 
                                        <tr key = {activity.activityId}>
                                             <td> { activity.subject } </td>
                                             <td> { activity.dueDate } </td>
                                             <td> { activity.startAt } </td>
                                             <td> { activity.endAt } </td>
                                             <td> { activity.location } </td>
                                             <td> { activity.activityType } </td>
                                             <td> { activity.status } </td>
                                             <td> { activity.priority } </td>
                                             <td>
                                                 <button onClick={ () => this.editActivity(activity.activityId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteActivity(activity.activityId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewActivity(activity.activityId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListActivityComponent
