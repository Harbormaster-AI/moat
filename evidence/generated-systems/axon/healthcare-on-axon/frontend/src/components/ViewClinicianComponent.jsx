import React, { Component } from 'react'
import ClinicianService from '../services/ClinicianService'

class ViewClinicianComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            clinician: {}
        }
    }

    componentDidMount(){
        ClinicianService.getClinicianById(this.state.id).then( res => {
            this.setState({clinician: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Clinician Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> firstName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.clinician.firstName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> lastName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.clinician.lastName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> licenseNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.clinician.licenseNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ClinicianType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.clinician.clinicianType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Specialty:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.clinician.specialty }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewClinicianComponent
