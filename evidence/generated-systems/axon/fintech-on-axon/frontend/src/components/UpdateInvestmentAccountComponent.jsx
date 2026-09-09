import React, { Component } from 'react'
import InvestmentAccountService from '../services/InvestmentAccountService';

class UpdateInvestmentAccountComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                accountNumber: '',
                baseCurrency: '',
                balance: '',
                accountType: '',
                status: ''
        }
        this.updateInvestmentAccount = this.updateInvestmentAccount.bind(this);

        this.changeaccountNumberHandler = this.changeaccountNumberHandler.bind(this);
        this.changebaseCurrencyHandler = this.changebaseCurrencyHandler.bind(this);
        this.changebalanceHandler = this.changebalanceHandler.bind(this);
        this.changeAccountTypeHandler = this.changeAccountTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        InvestmentAccountService.getInvestmentAccountById(this.state.id).then( (res) =>{
            let investmentAccount = res.data;
            this.setState({
                accountNumber: investmentAccount.accountNumber,
                baseCurrency: investmentAccount.baseCurrency,
                balance: investmentAccount.balance,
                accountType: investmentAccount.accountType,
                status: investmentAccount.status
            });
        });
    }

    updateInvestmentAccount = (e) => {
        e.preventDefault();
        let investmentAccount = {
            investmentAccountId: this.state.id,
            accountNumber: this.state.accountNumber,
            baseCurrency: this.state.baseCurrency,
            balance: this.state.balance,
            accountType: this.state.accountType,
            status: this.state.status
        };
        console.log('investmentAccount => ' + JSON.stringify(investmentAccount));
        console.log('id => ' + JSON.stringify(this.state.id));
        InvestmentAccountService.updateInvestmentAccount(investmentAccount).then( res => {
            this.props.history.push('/investmentAccounts');
        });
    }

    changeaccountNumberHandler= (event) => {
        this.setState({accountNumber: event.target.value});
    }
    changebaseCurrencyHandler= (event) => {
        this.setState({baseCurrency: event.target.value});
    }
    changebalanceHandler= (event) => {
        this.setState({balance: event.target.value});
    }
    changeAccountTypeHandler= (event) => {
        this.setState({accountType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/investmentAccounts');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update InvestmentAccount</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> accountNumber: </label>
                                                <input placeholder="accountNumber" name="accountNumber" className="form-control" value={this.state.accountNumber} onChange={this.changeaccountNumberHandler}/>

                                            <label> baseCurrency: </label>
                                                <input placeholder="baseCurrency" name="baseCurrency" className="form-control" value={this.state.baseCurrency} onChange={this.changebaseCurrencyHandler}/>

                                            <label> balance: </label>
                                                <input placeholder="balance" name="balance" className="form-control" value={this.state.balance} onChange={this.changebalanceHandler}/>

                                            <label> AccountType: </label>
                                                <select value={this.state.accountType} onChange={this.changeAccountTypeHandler}>
                      <option name="AccountType" className="form-control" >
                          Brokerage
                      </option>
                      <option name="AccountType" className="form-control" >
                          Retirement
                      </option>
                      <option name="AccountType" className="form-control" >
                          Custody
                      </option>
                      <option name="AccountType" className="form-control" >
                          Margin
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Pending
                      </option>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Frozen
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateInvestmentAccount}>Save</button>
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

export default UpdateInvestmentAccountComponent
