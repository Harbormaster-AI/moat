import React, { Component } from 'react'
import LoanService from '../services/LoanService';

class UpdateLoanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                loanNumber: '',
                principal: '',
                interestRate: '',
                originationDate: '',
                maturityDate: '',
                rateType: '',
                status: ''
        }
        this.updateLoan = this.updateLoan.bind(this);

        this.changeloanNumberHandler = this.changeloanNumberHandler.bind(this);
        this.changeprincipalHandler = this.changeprincipalHandler.bind(this);
        this.changeinterestRateHandler = this.changeinterestRateHandler.bind(this);
        this.changeoriginationDateHandler = this.changeoriginationDateHandler.bind(this);
        this.changematurityDateHandler = this.changematurityDateHandler.bind(this);
        this.changeRateTypeHandler = this.changeRateTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        LoanService.getLoanById(this.state.id).then( (res) =>{
            let loan = res.data;
            this.setState({
                loanNumber: loan.loanNumber,
                principal: loan.principal,
                interestRate: loan.interestRate,
                originationDate: loan.originationDate,
                maturityDate: loan.maturityDate,
                rateType: loan.rateType,
                status: loan.status
            });
        });
    }

    updateLoan = (e) => {
        e.preventDefault();
        let loan = {
            loanId: this.state.id,
            loanNumber: this.state.loanNumber,
            principal: this.state.principal,
            interestRate: this.state.interestRate,
            originationDate: this.state.originationDate,
            maturityDate: this.state.maturityDate,
            rateType: this.state.rateType,
            status: this.state.status
        };
        console.log('loan => ' + JSON.stringify(loan));
        console.log('id => ' + JSON.stringify(this.state.id));
        LoanService.updateLoan(loan).then( res => {
            this.props.history.push('/loans');
        });
    }

    changeloanNumberHandler= (event) => {
        this.setState({loanNumber: event.target.value});
    }
    changeprincipalHandler= (event) => {
        this.setState({principal: event.target.value});
    }
    changeinterestRateHandler= (event) => {
        this.setState({interestRate: event.target.value});
    }
    changeoriginationDateHandler= (event) => {
        this.setState({originationDate: event.target.value});
    }
    changematurityDateHandler= (event) => {
        this.setState({maturityDate: event.target.value});
    }
    changeRateTypeHandler= (event) => {
        this.setState({rateType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/loans');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Loan</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> loanNumber: </label>
                                                <input placeholder="loanNumber" name="loanNumber" className="form-control" value={this.state.loanNumber} onChange={this.changeloanNumberHandler}/>

                                            <label> principal: </label>
                                                <input placeholder="principal" name="principal" className="form-control" value={this.state.principal} onChange={this.changeprincipalHandler}/>

                                            <label> interestRate: </label>
                                                <input placeholder="interestRate" name="interestRate" className="form-control" value={this.state.interestRate} onChange={this.changeinterestRateHandler}/>

                                            <label> originationDate: </label>
                                                <input type="date" placeholder="originationDate" name="originationDate" className="form-control" value={this.state.originationDate} onChange={this.changeoriginationDateHandler}/>

                                            <label> maturityDate: </label>
                                                <input type="date" placeholder="maturityDate" name="maturityDate" className="form-control" value={this.state.maturityDate} onChange={this.changematurityDateHandler}/>

                                            <label> RateType: </label>
                                                <select value={this.state.rateType} onChange={this.changeRateTypeHandler}>
                      <option name="RateType" className="form-control" >
                          Fixed
                      </option>
                      <option name="RateType" className="form-control" >
                          Variable
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Delinquent
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                      <option name="Status" className="form-control" >
                          ChargedOff
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateLoan}>Save</button>
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

export default UpdateLoanComponent
