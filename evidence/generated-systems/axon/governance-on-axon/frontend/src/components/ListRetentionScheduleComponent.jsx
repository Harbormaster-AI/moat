import React, { Component } from 'react'
import RetentionScheduleService from '../services/RetentionScheduleService'

class ListRetentionScheduleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                retentionSchedules: []
        }
        this.addRetentionSchedule = this.addRetentionSchedule.bind(this);
        this.editRetentionSchedule = this.editRetentionSchedule.bind(this);
        this.deleteRetentionSchedule = this.deleteRetentionSchedule.bind(this);
    }

    deleteRetentionSchedule(id){
        RetentionScheduleService.deleteRetentionSchedule(id).then( res => {
            this.setState({retentionSchedules: this.state.retentionSchedules.filter(retentionSchedule => retentionSchedule.retentionScheduleId !== id)});
        });
    }
    viewRetentionSchedule(id){
        this.props.history.push(`/view-retentionSchedule/${id}`);
    }
    editRetentionSchedule(id){
        this.props.history.push(`/add-retentionSchedule/${id}`);
    }

    componentDidMount(){
        RetentionScheduleService.getRetentionSchedules().then((res) => {
            this.setState({ retentionSchedules: res.data});
        });
    }

    addRetentionSchedule(){
        this.props.history.push('/add-retentionSchedule/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">RetentionSchedule List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addRetentionSchedule}> Add RetentionSchedule</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> RetentionPeriodMonths </th>
                                    <th> RetentionTrigger </th>
                                    <th> DispositionAction </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.retentionSchedules.map(
                                        retentionSchedule => 
                                        <tr key = {retentionSchedule.retentionScheduleId}>
                                             <td> { retentionSchedule.name } </td>
                                             <td> { retentionSchedule.retentionPeriodMonths } </td>
                                             <td> { retentionSchedule.retentionTrigger } </td>
                                             <td> { retentionSchedule.dispositionAction } </td>
                                             <td> { retentionSchedule.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editRetentionSchedule(retentionSchedule.retentionScheduleId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteRetentionSchedule(retentionSchedule.retentionScheduleId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewRetentionSchedule(retentionSchedule.retentionScheduleId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListRetentionScheduleComponent
