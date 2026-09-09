import React, { Component } from 'react'
import MaintenanceAppointmentService from '../services/MaintenanceAppointmentService'

class ListMaintenanceAppointmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                maintenanceAppointments: []
        }
        this.addMaintenanceAppointment = this.addMaintenanceAppointment.bind(this);
        this.editMaintenanceAppointment = this.editMaintenanceAppointment.bind(this);
        this.deleteMaintenanceAppointment = this.deleteMaintenanceAppointment.bind(this);
    }

    deleteMaintenanceAppointment(id){
        MaintenanceAppointmentService.deleteMaintenanceAppointment(id).then( res => {
            this.setState({maintenanceAppointments: this.state.maintenanceAppointments.filter(maintenanceAppointment => maintenanceAppointment.maintenanceAppointmentId !== id)});
        });
    }
    viewMaintenanceAppointment(id){
        this.props.history.push(`/view-maintenanceAppointment/${id}`);
    }
    editMaintenanceAppointment(id){
        this.props.history.push(`/add-maintenanceAppointment/${id}`);
    }

    componentDidMount(){
        MaintenanceAppointmentService.getMaintenanceAppointments().then((res) => {
            this.setState({ maintenanceAppointments: res.data});
        });
    }

    addMaintenanceAppointment(){
        this.props.history.push('/add-maintenanceAppointment/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">MaintenanceAppointment List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addMaintenanceAppointment}> Add MaintenanceAppointment</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> AppointmentDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.maintenanceAppointments.map(
                                        maintenanceAppointment => 
                                        <tr key = {maintenanceAppointment.maintenanceAppointmentId}>
                                             <td> { maintenanceAppointment.appointmentDate } </td>
                                             <td> { maintenanceAppointment.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editMaintenanceAppointment(maintenanceAppointment.maintenanceAppointmentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteMaintenanceAppointment(maintenanceAppointment.maintenanceAppointmentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewMaintenanceAppointment(maintenanceAppointment.maintenanceAppointmentId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListMaintenanceAppointmentComponent
