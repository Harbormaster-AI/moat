import React, { Component } from 'react'
import MaintenanceAppointmentService from '../services/MaintenanceAppointmentService';

class UpdateMaintenanceAppointmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                appointmentDate: '',
                status: ''
        }
        this.updateMaintenanceAppointment = this.updateMaintenanceAppointment.bind(this);

        this.changeappointmentDateHandler = this.changeappointmentDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        MaintenanceAppointmentService.getMaintenanceAppointmentById(this.state.id).then( (res) =>{
            let maintenanceAppointment = res.data;
            this.setState({
                appointmentDate: maintenanceAppointment.appointmentDate,
                status: maintenanceAppointment.status
            });
        });
    }

    updateMaintenanceAppointment = (e) => {
        e.preventDefault();
        let maintenanceAppointment = {
            maintenanceAppointmentId: this.state.id,
            appointmentDate: this.state.appointmentDate,
            status: this.state.status
        };
        console.log('maintenanceAppointment => ' + JSON.stringify(maintenanceAppointment));
        console.log('id => ' + JSON.stringify(this.state.id));
        MaintenanceAppointmentService.updateMaintenanceAppointment(maintenanceAppointment).then( res => {
            this.props.history.push('/maintenanceAppointments');
        });
    }

    changeappointmentDateHandler= (event) => {
        this.setState({appointmentDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/maintenanceAppointments');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update MaintenanceAppointment</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> appointmentDate: </label>
                                                <input type="date" placeholder="appointmentDate" name="appointmentDate" className="form-control" value={this.state.appointmentDate} onChange={this.changeappointmentDateHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Scheduled
                      </option>
                      <option name="Status" className="form-control" >
                          InProgress
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                      <option name="Status" className="form-control" >
                          Deferred
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateMaintenanceAppointment}>Save</button>
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

export default UpdateMaintenanceAppointmentComponent
