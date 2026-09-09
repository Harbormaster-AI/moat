import React, { Component } from 'react'
import AppointmentService from '../services/AppointmentService'

class ViewAppointmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            appointment: {}
        }
    }

    componentDidMount(){
        AppointmentService.getAppointmentById(this.state.id).then( res => {
            this.setState({appointment: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Appointment Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> appointmentDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.appointment.appointmentDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> reason:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.appointment.reason }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.appointment.status }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Priority:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.appointment.priority }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAppointmentComponent
