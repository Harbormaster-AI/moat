import React, { Component } from 'react'
import ConsentService from '../services/ConsentService';

class UpdateConsentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                grantedAt: '',
                expiresAt: '',
                scope: '',
                consentType: '',
                status: ''
        }
        this.updateConsent = this.updateConsent.bind(this);

        this.changegrantedAtHandler = this.changegrantedAtHandler.bind(this);
        this.changeexpiresAtHandler = this.changeexpiresAtHandler.bind(this);
        this.changescopeHandler = this.changescopeHandler.bind(this);
        this.changeConsentTypeHandler = this.changeConsentTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ConsentService.getConsentById(this.state.id).then( (res) =>{
            let consent = res.data;
            this.setState({
                grantedAt: consent.grantedAt,
                expiresAt: consent.expiresAt,
                scope: consent.scope,
                consentType: consent.consentType,
                status: consent.status
            });
        });
    }

    updateConsent = (e) => {
        e.preventDefault();
        let consent = {
            consentId: this.state.id,
            grantedAt: this.state.grantedAt,
            expiresAt: this.state.expiresAt,
            scope: this.state.scope,
            consentType: this.state.consentType,
            status: this.state.status
        };
        console.log('consent => ' + JSON.stringify(consent));
        console.log('id => ' + JSON.stringify(this.state.id));
        ConsentService.updateConsent(consent).then( res => {
            this.props.history.push('/consents');
        });
    }

    changegrantedAtHandler= (event) => {
        this.setState({grantedAt: event.target.value});
    }
    changeexpiresAtHandler= (event) => {
        this.setState({expiresAt: event.target.value});
    }
    changescopeHandler= (event) => {
        this.setState({scope: event.target.value});
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
                                            <label> grantedAt: </label>
                                                <input type="time" placeholder="grantedAt" name="grantedAt" className="form-control" value={this.state.grantedAt} onChange={this.changegrantedAtHandler}/>

                                            <label> expiresAt: </label>
                                                <input type="time" placeholder="expiresAt" name="expiresAt" className="form-control" value={this.state.expiresAt} onChange={this.changeexpiresAtHandler}/>

                                            <label> scope: </label>
                                                <input placeholder="scope" name="scope" className="form-control" value={this.state.scope} onChange={this.changescopeHandler}/>

                                            <label> ConsentType: </label>
                                                <select value={this.state.consentType} onChange={this.changeConsentTypeHandler}>
                      <option name="ConsentType" className="form-control" >
                          DataAccess
                      </option>
                      <option name="ConsentType" className="form-control" >
                          PaymentInitiation
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Revoked
                      </option>
                      <option name="Status" className="form-control" >
                          Expired
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
