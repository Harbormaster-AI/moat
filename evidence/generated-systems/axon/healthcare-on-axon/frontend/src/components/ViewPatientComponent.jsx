import React, { Component } from 'react'
import PatientService from '../services/PatientService'

class ViewPatientComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            patient: {}
        }
    }

    componentDidMount(){
        PatientService.getPatientById(this.state.id).then( res => {
            this.setState({patient: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Patient Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> firstName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.patient.firstName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> lastName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.patient.lastName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> mrn:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.patient.mrn }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> dateOfBirth:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.patient.dateOfBirth }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> address:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.patient.address }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> primaryLanguage:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.patient.primaryLanguage }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> SexAtBirth:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.patient.sexAtBirth }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> BloodType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.patient.bloodType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPatientComponent
