import React, { Component } from 'react'
import PatientService from '../services/PatientService'

class ListPatientComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                patients: []
        }
        this.addPatient = this.addPatient.bind(this);
        this.editPatient = this.editPatient.bind(this);
        this.deletePatient = this.deletePatient.bind(this);
    }

    deletePatient(id){
        PatientService.deletePatient(id).then( res => {
            this.setState({patients: this.state.patients.filter(patient => patient.patientId !== id)});
        });
    }
    viewPatient(id){
        this.props.history.push(`/view-patient/${id}`);
    }
    editPatient(id){
        this.props.history.push(`/add-patient/${id}`);
    }

    componentDidMount(){
        PatientService.getPatients().then((res) => {
            this.setState({ patients: res.data});
        });
    }

    addPatient(){
        this.props.history.push('/add-patient/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Patient List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPatient}> Add Patient</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> FirstName </th>
                                    <th> LastName </th>
                                    <th> Mrn </th>
                                    <th> DateOfBirth </th>
                                    <th> Address </th>
                                    <th> PrimaryLanguage </th>
                                    <th> SexAtBirth </th>
                                    <th> BloodType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.patients.map(
                                        patient => 
                                        <tr key = {patient.patientId}>
                                             <td> { patient.firstName } </td>
                                             <td> { patient.lastName } </td>
                                             <td> { patient.mrn } </td>
                                             <td> { patient.dateOfBirth } </td>
                                             <td> { patient.address } </td>
                                             <td> { patient.primaryLanguage } </td>
                                             <td> { patient.sexAtBirth } </td>
                                             <td> { patient.bloodType } </td>
                                             <td>
                                                 <button onClick={ () => this.editPatient(patient.patientId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePatient(patient.patientId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPatient(patient.patientId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPatientComponent
