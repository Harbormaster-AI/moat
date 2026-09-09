import React, { Component } from 'react'
import LoanApplicationService from '../services/LoanApplicationService';

class CreateLoanApplicationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                applicationNumber: '',
                amountRequested: '',
                termMonths: '',
                submittedAt: '',
                product: '',
                purpose: '',
                status: ''
        }
        this.changeapplicationNumberHandler = this.changeapplicationNumberHandler.bind(this);
        this.changeamountRequestedHandler = this.changeamountRequestedHandler.bind(this);
        this.changetermMonthsHandler = this.changetermMonthsHandler.bind(this);
        this.changesubmittedAtHandler = this.changesubmittedAtHandler.bind(this);
        this.changeProductHandler = this.changeProductHandler.bind(this);
        this.changePurposeHandler = this.changePurposeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            LoanApplicationService.getLoanApplicationById(this.state.id).then( (res) =>{
                let loanApplication = res.data;
                this.setState({
                    applicationNumber: loanApplication.applicationNumber,
                    amountRequested: loanApplication.amountRequested,
                    termMonths: loanApplication.termMonths,
                    submittedAt: loanApplication.submittedAt,
                    product: loanApplication.product,
                    purpose: loanApplication.purpose,
                    status: loanApplication.status
                });
            });
        }        
    }
    saveOrUpdateLoanApplication = (e) => {
        e.preventDefault();
        let loanApplication = {
                loanApplicationId: this.state.id,
                applicationNumber: this.state.applicationNumber,
                amountRequested: this.state.amountRequested,
                termMonths: this.state.termMonths,
                submittedAt: this.state.submittedAt,
                product: this.state.product,
                purpose: this.state.purpose,
                status: this.state.status
            };
        console.log('loanApplication => ' + JSON.stringify(loanApplication));

        // step 5
        if(this.state.id === '_add'){
            loanApplication.loanApplicationId=''
            LoanApplicationService.createLoanApplication(loanApplication).then(res =>{
                this.props.history.push('/loanApplications');
            });
        }else{
            LoanApplicationService.updateLoanApplication(loanApplication).then( res => {
                this.props.history.push('/loanApplications');
            });
        }
    }
    
    changeapplicationNumberHandler= (event) => {
        this.setState({applicationNumber: event.target.value});
    }
    changeamountRequestedHandler= (event) => {
        this.setState({amountRequested: event.target.value});
    }
    changetermMonthsHandler= (event) => {
        this.setState({termMonths: event.target.value});
    }
    changesubmittedAtHandler= (event) => {
        this.setState({submittedAt: event.target.value});
    }
    changeProductHandler= (event) => {
        this.setState({product: event.target.value});
    }
    changePurposeHandler= (event) => {
        this.setState({purpose: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/loanApplications');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add LoanApplication</h3>
        }else{
            return <h3 className="text-center">Update LoanApplication</h3>
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
                                            <label> applicationNumber:&emsp; </label>
                                                <input placeholder="applicationNumber" name="applicationNumber" className="form-control" value={this.state.applicationNumber} onChange={this.changeapplicationNumberHandler}/>

                                            <label> amountRequested:&emsp; </label>
                                                <input placeholder="amountRequested" name="amountRequested" className="form-control" value={this.state.amountRequested} onChange={this.changeamountRequestedHandler}/>

                                            <label> termMonths:&emsp; </label>
                                                <input type="number" placeholder="termMonths" name="termMonths" className="form-control" value={this.state.termMonths} onChange={this.changetermMonthsHandler}/>

                                            <label> submittedAt:&emsp; </label>
                                                <input type="time" placeholder="submittedAt" name="submittedAt" className="form-control" value={this.state.submittedAt} onChange={this.changesubmittedAtHandler}/>

                                            <label> Product:&emsp; </label>
                                                <select value={this.state.product} onChange={this.changeProductHandler}>
                      <option name="Product" className="form-control" >
                          PersonalLoan
                      </option>
                      <option name="Product" className="form-control" >
                          Mortgage
                      </option>
                      <option name="Product" className="form-control" >
                          InstallmentLoan
                      </option>
                      <option name="Product" className="form-control" >
                          CreditLine
                      </option>
                      <option name="Product" className="form-control" >
                          SME
                      </option>
                    </select>

                                            <label> Purpose:&emsp; </label>
                                                <select value={this.state.purpose} onChange={this.changePurposeHandler}>
                      <option name="Purpose" className="form-control" >
                          HomeImprovement
                      </option>
                      <option name="Purpose" className="form-control" >
                          Education
                      </option>
                      <option name="Purpose" className="form-control" >
                          DebtConsolidation
                      </option>
                      <option name="Purpose" className="form-control" >
                          Business
                      </option>
                      <option name="Purpose" className="form-control" >
                          Other
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Submitted
                      </option>
                      <option name="Status" className="form-control" >
                          Underwriting
                      </option>
                      <option name="Status" className="form-control" >
                          Approved
                      </option>
                      <option name="Status" className="form-control" >
                          Declined
                      </option>
                      <option name="Status" className="form-control" >
                          Withdrawn
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateLoanApplication}>Save</button>
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

export default CreateLoanApplicationComponent
