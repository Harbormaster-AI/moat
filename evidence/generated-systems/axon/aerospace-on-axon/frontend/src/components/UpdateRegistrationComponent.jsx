import React, { Component } from 'react'
import RegistrationService from '../services/RegistrationService';

class UpdateRegistrationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                tailNumber: '',
                registryCountry: ''
        }
        this.updateRegistration = this.updateRegistration.bind(this);

        this.changetailNumberHandler = this.changetailNumberHandler.bind(this);
        this.changeregistryCountryHandler = this.changeregistryCountryHandler.bind(this);
    }

    componentDidMount(){
        RegistrationService.getRegistrationById(this.state.id).then( (res) =>{
            let registration = res.data;
            this.setState({
                tailNumber: registration.tailNumber,
                registryCountry: registration.registryCountry
            });
        });
    }

    updateRegistration = (e) => {
        e.preventDefault();
        let registration = {
            registrationId: this.state.id,
            tailNumber: this.state.tailNumber,
            registryCountry: this.state.registryCountry
        };
        console.log('registration => ' + JSON.stringify(registration));
        console.log('id => ' + JSON.stringify(this.state.id));
        RegistrationService.updateRegistration(registration).then( res => {
            this.props.history.push('/registrations');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Registration</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> tailNumber: </label>
                                                <input placeholder="tailNumber" name="tailNumber" className="form-control" value={this.state.tailNumber} onChange={this.changetailNumberHandler}/>

                                            <label> registryCountry: </label>
                                                <input placeholder="registryCountry" name="registryCountry" className="form-control" value={this.state.registryCountry} onChange={this.changeregistryCountryHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateRegistration}>Save</button>
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

export default UpdateRegistrationComponent
