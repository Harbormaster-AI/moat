import React, { Component } from 'react'
import AppointmentService from '../services/AppointmentService';

class UpdateAppointmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                appointmentDate: '',
                reason: '',
                status: '',
                priority: ''
        }
        this.updateAppointment = this.updateAppointment.bind(this);

        this.changeappointmentDateHandler = this.changeappointmentDateHandler.bind(this);
        this.changereasonHandler = this.changereasonHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changePriorityHandler = this.changePriorityHandler.bind(this);
    }

    componentDidMount(){
        AppointmentService.getAppointmentById(this.state.id).then( (res) =>{
            let appointment = res.data;
            this.setState({
                appointmentDate: appointment.appointmentDate,
                reason: appointment.reason,
                status: appointment.status,
                priority: appointment.priority
            });
        });
    }

    updateAppointment = (e) => {
        e.preventDefault();
        let appointment = {
            appointmentId: this.state.id,
            appointmentDate: this.state.appointmentDate,
            reason: this.state.reason,
            status: this.state.status,
            priority: this.state.priority
        };
        console.log('appointment => ' + JSON.stringify(appointment));
        console.log('id => ' + JSON.stringify(this.state.id));
        AppointmentService.updateAppointment(appointment).then( res => {
            this.props.history.push('/appointments');
        });
    }

    changeappointmentDateHandler= (event) => {
        this.setState({appointmentDate: event.target.value});
    }
    changereasonHandler= (event) => {
        this.setState({reason: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changePriorityHandler= (event) => {
        this.setState({priority: event.target.value});
    }

    cancel(){
        this.props.history.push('/appointments');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Appointment</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> appointmentDate: </label>
                                                <input type="time" placeholder="appointmentDate" name="appointmentDate" className="form-control" value={this.state.appointmentDate} onChange={this.changeappointmentDateHandler}/>

                                            <label> reason: </label>
                                                <input placeholder="reason" name="reason" className="form-control" value={this.state.reason} onChange={this.changereasonHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Proposed
                      </option>
                      <option name="Status" className="form-control" >
                          Booked
                      </option>
                      <option name="Status" className="form-control" >
                          Arrived
                      </option>
                      <option name="Status" className="form-control" >
                          Fulfilled
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                      <option name="Status" className="form-control" >
                          NoShow
                      </option>
                      <option name="Status" className="form-control" >
                          EnteredInError
                      </option>
                    </select>

                                            <label> Priority: </label>
                                                <select value={this.state.priority} onChange={this.changePriorityHandler}>
                      <option name="Priority" className="form-control" >
                          Routine
                      </option>
                      <option name="Priority" className="form-control" >
                          Urgent
                      </option>
                      <option name="Priority" className="form-control" >
                          Stat
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateAppointment}>Save</button>
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

export default UpdateAppointmentComponent
