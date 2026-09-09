import React, { Component } from 'react'
import MaintenanceAppointmentService from '../services/MaintenanceAppointmentService'

class ViewMaintenanceAppointmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            maintenanceAppointment: {}
        }
    }

    componentDidMount(){
        MaintenanceAppointmentService.getMaintenanceAppointmentById(this.state.id).then( res => {
            this.setState({maintenanceAppointment: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View MaintenanceAppointment Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> appointmentDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.maintenanceAppointment.appointmentDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.maintenanceAppointment.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewMaintenanceAppointmentComponent
