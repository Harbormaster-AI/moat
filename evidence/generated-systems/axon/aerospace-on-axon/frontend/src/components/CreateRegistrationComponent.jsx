import React, { Component } from 'react'
import RegistrationService from '../services/RegistrationService';

class CreateRegistrationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                tailNumber: '',
                registryCountry: ''
        }
        this.changetailNumberHandler = this.changetailNumberHandler.bind(this);
        this.changeregistryCountryHandler = this.changeregistryCountryHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            RegistrationService.getRegistrationById(this.state.id).then( (res) =>{
                let registration = res.data;
                this.setState({
                    tailNumber: registration.tailNumber,
                    registryCountry: registration.registryCountry
                });
            });
        }        
    }
    saveOrUpdateRegistration = (e) => {
        e.preventDefault();
        let registration = {
                registrationId: this.state.id,
                tailNumber: this.state.tailNumber,
                registryCountry: this.state.registryCountry
            };
        console.log('registration => ' + JSON.stringify(registration));

        // step 5
        if(this.state.id === '_add'){
            registration.registrationId=''
            RegistrationService.createRegistration(registration).then(res =>{
                this.props.history.push('/registrations');
            });
        }else{
            RegistrationService.updateRegistration(registration).then( res => {
                this.props.history.push('/registrations');
            });
        }
    }
    
    changetailNumberHandler= (event) => {
        this.setState({tailNumber: event.target.value});
    }
    changeregistryCountryHandler= (event) => {
        this.setState({registryCountry: event.target.value});
    }

    cancel(){
        this.props.history.push('/registrations');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Registration</h3>
        }else{
            return <h3 className="text-center">Update Registration</h3>
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
                                            <label> tailNumber:&emsp; </label>
                                                <input placeholder="tailNumber" name="tailNumber" className="form-control" value={this.state.tailNumber} onChange={this.changetailNumberHandler}/>

                                            <label> registryCountry:&emsp; </label>
                                                <input placeholder="registryCountry" name="registryCountry" className="form-control" value={this.state.registryCountry} onChange={this.changeregistryCountryHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateRegistration}>Save</button>
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

export default CreateRegistrationComponent
