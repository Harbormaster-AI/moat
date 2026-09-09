import React, { Component } from 'react'
import AppointmentService from '../services/AppointmentService'

class ListAppointmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                appointments: []
        }
        this.addAppointment = this.addAppointment.bind(this);
        this.editAppointment = this.editAppointment.bind(this);
        this.deleteAppointment = this.deleteAppointment.bind(this);
    }

    deleteAppointment(id){
        AppointmentService.deleteAppointment(id).then( res => {
            this.setState({appointments: this.state.appointments.filter(appointment => appointment.appointmentId !== id)});
        });
    }
    viewAppointment(id){
        this.props.history.push(`/view-appointment/${id}`);
    }
    editAppointment(id){
        this.props.history.push(`/add-appointment/${id}`);
    }

    componentDidMount(){
        AppointmentService.getAppointments().then((res) => {
            this.setState({ appointments: res.data});
        });
    }

    addAppointment(){
        this.props.history.push('/add-appointment/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Appointment List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAppointment}> Add Appointment</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> AppointmentDate </th>
                                    <th> Reason </th>
                                    <th> Status </th>
                                    <th> Priority </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.appointments.map(
                                        appointment => 
                                        <tr key = {appointment.appointmentId}>
                                             <td> { appointment.appointmentDate } </td>
                                             <td> { appointment.reason } </td>
                                             <td> { appointment.status } </td>
                                             <td> { appointment.priority } </td>
                                             <td>
                                                 <button onClick={ () => this.editAppointment(appointment.appointmentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAppointment(appointment.appointmentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAppointment(appointment.appointmentId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAppointmentComponent
