import React, { Component } from 'react'
import PatientService from '../services/PatientService';

class CreatePatientComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
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
        this.changefirstNameHandler = this.changefirstNameHandler.bind(this);
        this.changelastNameHandler = this.changelastNameHandler.bind(this);
        this.changemrnHandler = this.changemrnHandler.bind(this);
        this.changedateOfBirthHandler = this.changedateOfBirthHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
        this.changeprimaryLanguageHandler = this.changeprimaryLanguageHandler.bind(this);
        this.changeSexAtBirthHandler = this.changeSexAtBirthHandler.bind(this);
        this.changeBloodTypeHandler = this.changeBloodTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdatePatient = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            patient.patientId=''
            PatientService.createPatient(patient).then(res =>{
                this.props.history.push('/patients');
            });
        }else{
            PatientService.updatePatient(patient).then( res => {
                this.props.history.push('/patients');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Patient</h3>
        }else{
            return <h3 className="text-center">Update Patient</h3>
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
                                            <label> firstName:&emsp; </label>
                                                <input placeholder="firstName" name="firstName" className="form-control" value={this.state.firstName} onChange={this.changefirstNameHandler}/>

                                            <label> lastName:&emsp; </label>
                                                <input placeholder="lastName" name="lastName" className="form-control" value={this.state.lastName} onChange={this.changelastNameHandler}/>

                                            <label> mrn:&emsp; </label>
                                                <input placeholder="mrn" name="mrn" className="form-control" value={this.state.mrn} onChange={this.changemrnHandler}/>

                                            <label> dateOfBirth:&emsp; </label>
                                                <input type="date" placeholder="dateOfBirth" name="dateOfBirth" className="form-control" value={this.state.dateOfBirth} onChange={this.changedateOfBirthHandler}/>

                                            <label> address:&emsp; </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                            <label> primaryLanguage:&emsp; </label>
                                                <input placeholder="primaryLanguage" name="primaryLanguage" className="form-control" value={this.state.primaryLanguage} onChange={this.changeprimaryLanguageHandler}/>

                                            <label> SexAtBirth:&emsp; </label>
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

                                            <label> BloodType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePatient}>Save</button>
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

export default CreatePatientComponent
