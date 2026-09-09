import React, { Component } from 'react'
import AgreementService from '../services/AgreementService';

class UpdateAgreementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                agreementNumber: '',
                effectiveDate: '',
                agreementType: '',
                status: ''
        }
        this.updateAgreement = this.updateAgreement.bind(this);

        this.changeagreementNumberHandler = this.changeagreementNumberHandler.bind(this);
        this.changeeffectiveDateHandler = this.changeeffectiveDateHandler.bind(this);
        this.changeAgreementTypeHandler = this.changeAgreementTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        AgreementService.getAgreementById(this.state.id).then( (res) =>{
            let agreement = res.data;
            this.setState({
                agreementNumber: agreement.agreementNumber,
                effectiveDate: agreement.effectiveDate,
                agreementType: agreement.agreementType,
                status: agreement.status
            });
        });
    }

    updateAgreement = (e) => {
        e.preventDefault();
        let agreement = {
            agreementId: this.state.id,
            agreementNumber: this.state.agreementNumber,
            effectiveDate: this.state.effectiveDate,
            agreementType: this.state.agreementType,
            status: this.state.status
        };
        console.log('agreement => ' + JSON.stringify(agreement));
        console.log('id => ' + JSON.stringify(this.state.id));
        AgreementService.updateAgreement(agreement).then( res => {
            this.props.history.push('/agreements');
        });
    }

    changeagreementNumberHandler= (event) => {
        this.setState({agreementNumber: event.target.value});
    }
    changeeffectiveDateHandler= (event) => {
        this.setState({effectiveDate: event.target.value});
    }
    changeAgreementTypeHandler= (event) => {
        this.setState({agreementType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/agreements');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Agreement</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> agreementNumber: </label>
                                                <input placeholder="agreementNumber" name="agreementNumber" className="form-control" value={this.state.agreementNumber} onChange={this.changeagreementNumberHandler}/>

                                            <label> effectiveDate: </label>
                                                <input type="date" placeholder="effectiveDate" name="effectiveDate" className="form-control" value={this.state.effectiveDate} onChange={this.changeeffectiveDateHandler}/>

                                            <label> AgreementType: </label>
                                                <select value={this.state.agreementType} onChange={this.changeAgreementTypeHandler}>
                      <option name="AgreementType" className="form-control" >
                          TermsOfService
                      </option>
                      <option name="AgreementType" className="form-control" >
                          PrivacyPolicy
                      </option>
                      <option name="AgreementType" className="form-control" >
                          LoanAgreement
                      </option>
                      <option name="AgreementType" className="form-control" >
                          AccountAgreement
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
                          Terminated
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateAgreement}>Save</button>
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

export default UpdateAgreementComponent
