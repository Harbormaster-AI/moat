import React, { Component } from 'react'
import MaintenanceAppointmentService from '../services/MaintenanceAppointmentService';

class CreateMaintenanceAppointmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                appointmentDate: '',
                status: ''
        }
        this.changeappointmentDateHandler = this.changeappointmentDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            MaintenanceAppointmentService.getMaintenanceAppointmentById(this.state.id).then( (res) =>{
                let maintenanceAppointment = res.data;
                this.setState({
                    appointmentDate: maintenanceAppointment.appointmentDate,
                    status: maintenanceAppointment.status
                });
            });
        }        
    }
    saveOrUpdateMaintenanceAppointment = (e) => {
        e.preventDefault();
        let maintenanceAppointment = {
                maintenanceAppointmentId: this.state.id,
                appointmentDate: this.state.appointmentDate,
                status: this.state.status
            };
        console.log('maintenanceAppointment => ' + JSON.stringify(maintenanceAppointment));

        // step 5
        if(this.state.id === '_add'){
            maintenanceAppointment.maintenanceAppointmentId=''
            MaintenanceAppointmentService.createMaintenanceAppointment(maintenanceAppointment).then(res =>{
                this.props.history.push('/maintenanceAppointments');
            });
        }else{
            MaintenanceAppointmentService.updateMaintenanceAppointment(maintenanceAppointment).then( res => {
                this.props.history.push('/maintenanceAppointments');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add MaintenanceAppointment</h3>
        }else{
            return <h3 className="text-center">Update MaintenanceAppointment</h3>
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
                                                <input type="date" placeholder="appointmentDate" name="appointmentDate" className="form-control" value={this.state.appointmentDate} onChange={this.changeappointmentDateHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateMaintenanceAppointment}>Save</button>
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

export default CreateMaintenanceAppointmentComponent
