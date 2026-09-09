import React, { Component } from 'react'
import DirectDebitMandateService from '../services/DirectDebitMandateService';

class UpdateDirectDebitMandateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                mandateId: '',
                signedAt: '',
                scheme: '',
                status: ''
        }
        this.updateDirectDebitMandate = this.updateDirectDebitMandate.bind(this);

        this.changemandateIdHandler = this.changemandateIdHandler.bind(this);
        this.changesignedAtHandler = this.changesignedAtHandler.bind(this);
        this.changeSchemeHandler = this.changeSchemeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        DirectDebitMandateService.getDirectDebitMandateById(this.state.id).then( (res) =>{
            let directDebitMandate = res.data;
            this.setState({
                mandateId: directDebitMandate.mandateId,
                signedAt: directDebitMandate.signedAt,
                scheme: directDebitMandate.scheme,
                status: directDebitMandate.status
            });
        });
    }

    updateDirectDebitMandate = (e) => {
        e.preventDefault();
        let directDebitMandate = {
            directDebitMandateId: this.state.id,
            mandateId: this.state.mandateId,
            signedAt: this.state.signedAt,
            scheme: this.state.scheme,
            status: this.state.status
        };
        console.log('directDebitMandate => ' + JSON.stringify(directDebitMandate));
        console.log('id => ' + JSON.stringify(this.state.id));
        DirectDebitMandateService.updateDirectDebitMandate(directDebitMandate).then( res => {
            this.props.history.push('/directDebitMandates');
        });
    }

    changemandateIdHandler= (event) => {
        this.setState({mandateId: event.target.value});
    }
    changesignedAtHandler= (event) => {
        this.setState({signedAt: event.target.value});
    }
    changeSchemeHandler= (event) => {
        this.setState({scheme: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/directDebitMandates');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update DirectDebitMandate</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> mandateId: </label>
                                                <input placeholder="mandateId" name="mandateId" className="form-control" value={this.state.mandateId} onChange={this.changemandateIdHandler}/>

                                            <label> signedAt: </label>
                                                <input type="time" placeholder="signedAt" name="signedAt" className="form-control" value={this.state.signedAt} onChange={this.changesignedAtHandler}/>

                                            <label> Scheme: </label>
                                                <select value={this.state.scheme} onChange={this.changeSchemeHandler}>
                      <option name="Scheme" className="form-control" >
                          SEPA
                      </option>
                      <option name="Scheme" className="form-control" >
                          ACH
                      </option>
                      <option name="Scheme" className="form-control" >
                          BACS
                      </option>
                      <option name="Scheme" className="form-control" >
                          BECS
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Suspended
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                      <option name="Status" className="form-control" >
                          Expired
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateDirectDebitMandate}>Save</button>
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

export default UpdateDirectDebitMandateComponent
