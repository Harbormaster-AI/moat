import React, { Component } from 'react'
import WorkScheduleService from '../services/WorkScheduleService'

class ListWorkScheduleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                workSchedules: []
        }
        this.addWorkSchedule = this.addWorkSchedule.bind(this);
        this.editWorkSchedule = this.editWorkSchedule.bind(this);
        this.deleteWorkSchedule = this.deleteWorkSchedule.bind(this);
    }

    deleteWorkSchedule(id){
        WorkScheduleService.deleteWorkSchedule(id).then( res => {
            this.setState({workSchedules: this.state.workSchedules.filter(workSchedule => workSchedule.workScheduleId !== id)});
        });
    }
    viewWorkSchedule(id){
        this.props.history.push(`/view-workSchedule/${id}`);
    }
    editWorkSchedule(id){
        this.props.history.push(`/add-workSchedule/${id}`);
    }

    componentDidMount(){
        WorkScheduleService.getWorkSchedules().then((res) => {
            this.setState({ workSchedules: res.data});
        });
    }

    addWorkSchedule(){
        this.props.history.push('/add-workSchedule/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">WorkSchedule List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addWorkSchedule}> Add WorkSchedule</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> StandardHoursPerWeek </th>
                                    <th> ScheduleType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.workSchedules.map(
                                        workSchedule => 
                                        <tr key = {workSchedule.workScheduleId}>
                                             <td> { workSchedule.name } </td>
                                             <td> { workSchedule.standardHoursPerWeek } </td>
                                             <td> { workSchedule.scheduleType } </td>
                                             <td>
                                                 <button onClick={ () => this.editWorkSchedule(workSchedule.workScheduleId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteWorkSchedule(workSchedule.workScheduleId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewWorkSchedule(workSchedule.workScheduleId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListWorkScheduleComponent
