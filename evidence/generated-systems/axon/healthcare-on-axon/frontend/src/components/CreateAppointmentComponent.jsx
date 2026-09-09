import React, { Component } from 'react'
import AppointmentService from '../services/AppointmentService';

class CreateAppointmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                appointmentDate: '',
                reason: '',
                status: '',
                priority: ''
        }
        this.changeappointmentDateHandler = this.changeappointmentDateHandler.bind(this);
        this.changereasonHandler = this.changereasonHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changePriorityHandler = this.changePriorityHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateAppointment = (e) => {
        e.preventDefault();
        let appointment = {
                appointmentId: this.state.id,
                appointmentDate: this.state.appointmentDate,
                reason: this.state.reason,
                status: this.state.status,
                priority: this.state.priority
            };
        console.log('appointment => ' + JSON.stringify(appointment));

        // step 5
        if(this.state.id === '_add'){
            appointment.appointmentId=''
            AppointmentService.createAppointment(appointment).then(res =>{
                this.props.history.push('/appointments');
            });
        }else{
            AppointmentService.updateAppointment(appointment).then( res => {
                this.props.history.push('/appointments');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Appointment</h3>
        }else{
            return <h3 className="text-center">Update Appointment</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> appointmentDate:&emsp; </label>
                                                <input type="time" placeholder="appointmentDate" name="appointmentDate" className="form-control" value={this.state.appointmentDate} onChange={this.changeappointmentDateHandler}/>

                                            <label> reason:&emsp; </label>
                                                <input placeholder="reason" name="reason" className="form-control" value={this.state.reason} onChange={this.changereasonHandler}/>

                                            <label> Status:&emsp; </label>
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

                                            <label> Priority:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAppointment}>Save</button>
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

export default CreateAppointmentComponent
