import React, { Component } from 'react'
import BuildScheduleService from '../services/BuildScheduleService'

class ListBuildScheduleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                buildSchedules: []
        }
        this.addBuildSchedule = this.addBuildSchedule.bind(this);
        this.editBuildSchedule = this.editBuildSchedule.bind(this);
        this.deleteBuildSchedule = this.deleteBuildSchedule.bind(this);
    }

    deleteBuildSchedule(id){
        BuildScheduleService.deleteBuildSchedule(id).then( res => {
            this.setState({buildSchedules: this.state.buildSchedules.filter(buildSchedule => buildSchedule.buildScheduleId !== id)});
        });
    }
    viewBuildSchedule(id){
        this.props.history.push(`/view-buildSchedule/${id}`);
    }
    editBuildSchedule(id){
        this.props.history.push(`/add-buildSchedule/${id}`);
    }

    componentDidMount(){
        BuildScheduleService.getBuildSchedules().then((res) => {
            this.setState({ buildSchedules: res.data});
        });
    }

    addBuildSchedule(){
        this.props.history.push('/add-buildSchedule/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">BuildSchedule List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addBuildSchedule}> Add BuildSchedule</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ScheduleNumber </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.buildSchedules.map(
                                        buildSchedule => 
                                        <tr key = {buildSchedule.buildScheduleId}>
                                             <td> { buildSchedule.scheduleNumber } </td>
                                             <td> { buildSchedule.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editBuildSchedule(buildSchedule.buildScheduleId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteBuildSchedule(buildSchedule.buildScheduleId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewBuildSchedule(buildSchedule.buildScheduleId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListBuildScheduleComponent
