import React, { Component } from 'react'
import ClinicianService from '../services/ClinicianService';

class UpdateClinicianComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                firstName: '',
                lastName: '',
                licenseNumber: '',
                clinicianType: '',
                specialty: ''
        }
        this.updateClinician = this.updateClinician.bind(this);

        this.changefirstNameHandler = this.changefirstNameHandler.bind(this);
        this.changelastNameHandler = this.changelastNameHandler.bind(this);
        this.changelicenseNumberHandler = this.changelicenseNumberHandler.bind(this);
        this.changeClinicianTypeHandler = this.changeClinicianTypeHandler.bind(this);
        this.changeSpecialtyHandler = this.changeSpecialtyHandler.bind(this);
    }

    componentDidMount(){
        ClinicianService.getClinicianById(this.state.id).then( (res) =>{
            let clinician = res.data;
            this.setState({
                firstName: clinician.firstName,
                lastName: clinician.lastName,
                licenseNumber: clinician.licenseNumber,
                clinicianType: clinician.clinicianType,
                specialty: clinician.specialty
            });
        });
    }

    updateClinician = (e) => {
        e.preventDefault();
        let clinician = {
            clinicianId: this.state.id,
            firstName: this.state.firstName,
            lastName: this.state.lastName,
            licenseNumber: this.state.licenseNumber,
            clinicianType: this.state.clinicianType,
            specialty: this.state.specialty
        };
        console.log('clinician => ' + JSON.stringify(clinician));
        console.log('id => ' + JSON.stringify(this.state.id));
        ClinicianService.updateClinician(clinician).then( res => {
            this.props.history.push('/clinicians');
        });
    }

    changefirstNameHandler= (event) => {
        this.setState({firstName: event.target.value});
    }
    changelastNameHandler= (event) => {
        this.setState({lastName: event.target.value});
    }
    changelicenseNumberHandler= (event) => {
        this.setState({licenseNumber: event.target.value});
    }
    changeClinicianTypeHandler= (event) => {
        this.setState({clinicianType: event.target.value});
    }
    changeSpecialtyHandler= (event) => {
        this.setState({specialty: event.target.value});
    }

    cancel(){
        this.props.history.push('/clinicians');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Clinician</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> firstName: </label>
                                                <input placeholder="firstName" name="firstName" className="form-control" value={this.state.firstName} onChange={this.changefirstNameHandler}/>

                                            <label> lastName: </label>
                                                <input placeholder="lastName" name="lastName" className="form-control" value={this.state.lastName} onChange={this.changelastNameHandler}/>

                                            <label> licenseNumber: </label>
                                                <input placeholder="licenseNumber" name="licenseNumber" className="form-control" value={this.state.licenseNumber} onChange={this.changelicenseNumberHandler}/>

                                            <label> ClinicianType: </label>
                                                <select value={this.state.clinicianType} onChange={this.changeClinicianTypeHandler}>
                      <option name="ClinicianType" className="form-control" >
                          Physician
                      </option>
                      <option name="ClinicianType" className="form-control" >
                          NursePractitioner
                      </option>
                      <option name="ClinicianType" className="form-control" >
                          PhysicianAssistant
                      </option>
                      <option name="ClinicianType" className="form-control" >
                          RegisteredNurse
                      </option>
                      <option name="ClinicianType" className="form-control" >
                          Pharmacist
                      </option>
                      <option name="ClinicianType" className="form-control" >
                          Therapist
                      </option>
                      <option name="ClinicianType" className="form-control" >
                          Technician
                      </option>
                    </select>

                                            <label> Specialty: </label>
                                                <select value={this.state.specialty} onChange={this.changeSpecialtyHandler}>
                      <option name="Specialty" className="form-control" >
                          InternalMedicine
                      </option>
                      <option name="Specialty" className="form-control" >
                          FamilyMedicine
                      </option>
                      <option name="Specialty" className="form-control" >
                          Cardiology
                      </option>
                      <option name="Specialty" className="form-control" >
                          Oncology
                      </option>
                      <option name="Specialty" className="form-control" >
                          Orthopedics
                      </option>
                      <option name="Specialty" className="form-control" >
                          Pediatrics
                      </option>
                      <option name="Specialty" className="form-control" >
                          Radiology
                      </option>
                      <option name="Specialty" className="form-control" >
                          Pathology
                      </option>
                      <option name="Specialty" className="form-control" >
                          Anesthesiology
                      </option>
                      <option name="Specialty" className="form-control" >
                          Surgery
                      </option>
                      <option name="Specialty" className="form-control" >
                          Psychiatry
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateClinician}>Save</button>
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

export default UpdateClinicianComponent
