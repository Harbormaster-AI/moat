import React, { Component } from 'react'
import ClinicianService from '../services/ClinicianService'

class ListClinicianComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                clinicians: []
        }
        this.addClinician = this.addClinician.bind(this);
        this.editClinician = this.editClinician.bind(this);
        this.deleteClinician = this.deleteClinician.bind(this);
    }

    deleteClinician(id){
        ClinicianService.deleteClinician(id).then( res => {
            this.setState({clinicians: this.state.clinicians.filter(clinician => clinician.clinicianId !== id)});
        });
    }
    viewClinician(id){
        this.props.history.push(`/view-clinician/${id}`);
    }
    editClinician(id){
        this.props.history.push(`/add-clinician/${id}`);
    }

    componentDidMount(){
        ClinicianService.getClinicians().then((res) => {
            this.setState({ clinicians: res.data});
        });
    }

    addClinician(){
        this.props.history.push('/add-clinician/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Clinician List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addClinician}> Add Clinician</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> FirstName </th>
                                    <th> LastName </th>
                                    <th> LicenseNumber </th>
                                    <th> ClinicianType </th>
                                    <th> Specialty </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.clinicians.map(
                                        clinician => 
                                        <tr key = {clinician.clinicianId}>
                                             <td> { clinician.firstName } </td>
                                             <td> { clinician.lastName } </td>
                                             <td> { clinician.licenseNumber } </td>
                                             <td> { clinician.clinicianType } </td>
                                             <td> { clinician.specialty } </td>
                                             <td>
                                                 <button onClick={ () => this.editClinician(clinician.clinicianId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteClinician(clinician.clinicianId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewClinician(clinician.clinicianId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListClinicianComponent
