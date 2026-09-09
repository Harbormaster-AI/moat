import React, { Component } from 'react'
import AgreementService from '../services/AgreementService';

class CreateAgreementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                agreementNumber: '',
                effectiveDate: '',
                agreementType: '',
                status: ''
        }
        this.changeagreementNumberHandler = this.changeagreementNumberHandler.bind(this);
        this.changeeffectiveDateHandler = this.changeeffectiveDateHandler.bind(this);
        this.changeAgreementTypeHandler = this.changeAgreementTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateAgreement = (e) => {
        e.preventDefault();
        let agreement = {
                agreementId: this.state.id,
                agreementNumber: this.state.agreementNumber,
                effectiveDate: this.state.effectiveDate,
                agreementType: this.state.agreementType,
                status: this.state.status
            };
        console.log('agreement => ' + JSON.stringify(agreement));

        // step 5
        if(this.state.id === '_add'){
            agreement.agreementId=''
            AgreementService.createAgreement(agreement).then(res =>{
                this.props.history.push('/agreements');
            });
        }else{
            AgreementService.updateAgreement(agreement).then( res => {
                this.props.history.push('/agreements');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Agreement</h3>
        }else{
            return <h3 className="text-center">Update Agreement</h3>
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
                                            <label> agreementNumber:&emsp; </label>
                                                <input placeholder="agreementNumber" name="agreementNumber" className="form-control" value={this.state.agreementNumber} onChange={this.changeagreementNumberHandler}/>

                                            <label> effectiveDate:&emsp; </label>
                                                <input type="date" placeholder="effectiveDate" name="effectiveDate" className="form-control" value={this.state.effectiveDate} onChange={this.changeeffectiveDateHandler}/>

                                            <label> AgreementType:&emsp; </label>
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

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAgreement}>Save</button>
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

export default CreateAgreementComponent
