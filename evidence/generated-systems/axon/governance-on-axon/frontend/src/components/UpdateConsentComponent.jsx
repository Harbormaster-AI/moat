import React, { Component } from 'react'
import ConsentService from '../services/ConsentService';

class UpdateConsentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                subjectIdentifier: '',
                captureDate: '',
                expiryDate: '',
                consentType: '',
                status: ''
        }
        this.updateConsent = this.updateConsent.bind(this);

        this.changesubjectIdentifierHandler = this.changesubjectIdentifierHandler.bind(this);
        this.changecaptureDateHandler = this.changecaptureDateHandler.bind(this);
        this.changeexpiryDateHandler = this.changeexpiryDateHandler.bind(this);
        this.changeConsentTypeHandler = this.changeConsentTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ConsentService.getConsentById(this.state.id).then( (res) =>{
            let consent = res.data;
            this.setState({
                subjectIdentifier: consent.subjectIdentifier,
                captureDate: consent.captureDate,
                expiryDate: consent.expiryDate,
                consentType: consent.consentType,
                status: consent.status
            });
        });
    }

    updateConsent = (e) => {
        e.preventDefault();
        let consent = {
            consentId: this.state.id,
            subjectIdentifier: this.state.subjectIdentifier,
            captureDate: this.state.captureDate,
            expiryDate: this.state.expiryDate,
            consentType: this.state.consentType,
            status: this.state.status
        };
        console.log('consent => ' + JSON.stringify(consent));
        console.log('id => ' + JSON.stringify(this.state.id));
        ConsentService.updateConsent(consent).then( res => {
            this.props.history.push('/consents');
        });
    }

    changesubjectIdentifierHandler= (event) => {
        this.setState({subjectIdentifier: event.target.value});
    }
    changecaptureDateHandler= (event) => {
        this.setState({captureDate: event.target.value});
    }
    changeexpiryDateHandler= (event) => {
        this.setState({expiryDate: event.target.value});
    }
    changeConsentTypeHandler= (event) => {
        this.setState({consentType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/consents');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Consent</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> subjectIdentifier: </label>
                                                <input placeholder="subjectIdentifier" name="subjectIdentifier" className="form-control" value={this.state.subjectIdentifier} onChange={this.changesubjectIdentifierHandler}/>

                                            <label> captureDate: </label>
                                                <input type="date" placeholder="captureDate" name="captureDate" className="form-control" value={this.state.captureDate} onChange={this.changecaptureDateHandler}/>

                                            <label> expiryDate: </label>
                                                <input type="date" placeholder="expiryDate" name="expiryDate" className="form-control" value={this.state.expiryDate} onChange={this.changeexpiryDateHandler}/>

                                            <label> ConsentType: </label>
                                                <select value={this.state.consentType} onChange={this.changeConsentTypeHandler}>
                      <option name="ConsentType" className="form-control" >
                          Marketing
                      </option>
                      <option name="ConsentType" className="form-control" >
                          Profiling
                      </option>
                      <option name="ConsentType" className="form-control" >
                          Cookies
                      </option>
                      <option name="ConsentType" className="form-control" >
                          Location
                      </option>
                      <option name="ConsentType" className="form-control" >
                          Biometric
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Granted
                      </option>
                      <option name="Status" className="form-control" >
                          Withdrawn
                      </option>
                      <option name="Status" className="form-control" >
                          Expired
                      </option>
                      <option name="Status" className="form-control" >
                          NotRequired
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateConsent}>Save</button>
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

export default UpdateConsentComponent
