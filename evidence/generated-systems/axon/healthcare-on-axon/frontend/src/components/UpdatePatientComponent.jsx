import React, { Component } from 'react'
import PatientService from '../services/PatientService';

class UpdatePatientComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                firstName: '',
                lastName: '',
                mrn: '',
                dateOfBirth: '',
                address: '',
                primaryLanguage: '',
                sexAtBirth: '',
                bloodType: ''
        }
        this.updatePatient = this.updatePatient.bind(this);

        this.changefirstNameHandler = this.changefirstNameHandler.bind(this);
        this.changelastNameHandler = this.changelastNameHandler.bind(this);
        this.changemrnHandler = this.changemrnHandler.bind(this);
        this.changedateOfBirthHandler = this.changedateOfBirthHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
        this.changeprimaryLanguageHandler = this.changeprimaryLanguageHandler.bind(this);
        this.changeSexAtBirthHandler = this.changeSexAtBirthHandler.bind(this);
        this.changeBloodTypeHandler = this.changeBloodTypeHandler.bind(this);
    }

    componentDidMount(){
        PatientService.getPatientById(this.state.id).then( (res) =>{
            let patient = res.data;
            this.setState({
                firstName: patient.firstName,
                lastName: patient.lastName,
                mrn: patient.mrn,
                dateOfBirth: patient.dateOfBirth,
                address: patient.address,
                primaryLanguage: patient.primaryLanguage,
                sexAtBirth: patient.sexAtBirth,
                bloodType: patient.bloodType
            });
        });
    }

    updatePatient = (e) => {
        e.preventDefault();
        let patient = {
            patientId: this.state.id,
            firstName: this.state.firstName,
            lastName: this.state.lastName,
            mrn: this.state.mrn,
            dateOfBirth: this.state.dateOfBirth,
            address: this.state.address,
            primaryLanguage: this.state.primaryLanguage,
            sexAtBirth: this.state.sexAtBirth,
            bloodType: this.state.bloodType
        };
        console.log('patient => ' + JSON.stringify(patient));
        console.log('id => ' + JSON.stringify(this.state.id));
        PatientService.updatePatient(patient).then( res => {
            this.props.history.push('/patients');
        });
    }

    changefirstNameHandler= (event) => {
        this.setState({firstName: event.target.value});
    }
    changelastNameHandler= (event) => {
        this.setState({lastName: event.target.value});
    }
    changemrnHandler= (event) => {
        this.setState({mrn: event.target.value});
    }
    changedateOfBirthHandler= (event) => {
        this.setState({dateOfBirth: event.target.value});
    }
    changeaddressHandler= (event) => {
        this.setState({address: event.target.value});
    }
    changeprimaryLanguageHandler= (event) => {
        this.setState({primaryLanguage: event.target.value});
    }
    changeSexAtBirthHandler= (event) => {
        this.setState({sexAtBirth: event.target.value});
    }
    changeBloodTypeHandler= (event) => {
        this.setState({bloodType: event.target.value});
    }

    cancel(){
        this.props.history.push('/patients');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Patient</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> firstName: </label>
                                                <input placeholder="firstName" name="firstName" className="form-control" value={this.state.firstName} onChange={this.changefirstNameHandler}/>

                                            <label> lastName: </label>
                                                <input placeholder="lastName" name="lastName" className="form-control" value={this.state.lastName} onChange={this.changelastNameHandler}/>

                                            <label> mrn: </label>
                                                <input placeholder="mrn" name="mrn" className="form-control" value={this.state.mrn} onChange={this.changemrnHandler}/>

                                            <label> dateOfBirth: </label>
                                                <input type="date" placeholder="dateOfBirth" name="dateOfBirth" className="form-control" value={this.state.dateOfBirth} onChange={this.changedateOfBirthHandler}/>

                                            <label> address: </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                            <label> primaryLanguage: </label>
                                                <input placeholder="primaryLanguage" name="primaryLanguage" className="form-control" value={this.state.primaryLanguage} onChange={this.changeprimaryLanguageHandler}/>

                                            <label> SexAtBirth: </label>
                                                <select value={this.state.sexAtBirth} onChange={this.changeSexAtBirthHandler}>
                      <option name="SexAtBirth" className="form-control" >
                          Male
                      </option>
                      <option name="SexAtBirth" className="form-control" >
                          Female
                      </option>
                      <option name="SexAtBirth" className="form-control" >
                          Unknown
                      </option>
                    </select>

                                            <label> BloodType: </label>
                                                <select value={this.state.bloodType} onChange={this.changeBloodTypeHandler}>
                      <option name="BloodType" className="form-control" >
                          APositive
                      </option>
                      <option name="BloodType" className="form-control" >
                          ANegative
                      </option>
                      <option name="BloodType" className="form-control" >
                          BPositive
                      </option>
                      <option name="BloodType" className="form-control" >
                          BNegative
                      </option>
                      <option name="BloodType" className="form-control" >
                          ABPositive
                      </option>
                      <option name="BloodType" className="form-control" >
                          ABNegative
                      </option>
                      <option name="BloodType" className="form-control" >
                          OPositive
                      </option>
                      <option name="BloodType" className="form-control" >
                          ONegative
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePatient}>Save</button>
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

export default UpdatePatientComponent
