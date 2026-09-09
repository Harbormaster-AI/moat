import React, { Component } from 'react'
import AccountService from '../services/AccountService';

class UpdateAccountComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                accountNumber: '',
                iban: '',
                bic: '',
                openedDate: '',
                currency: '',
                balance: '',
                availableBalance: '',
                accountType: '',
                status: ''
        }
        this.updateAccount = this.updateAccount.bind(this);

        this.changeaccountNumberHandler = this.changeaccountNumberHandler.bind(this);
        this.changeibanHandler = this.changeibanHandler.bind(this);
        this.changebicHandler = this.changebicHandler.bind(this);
        this.changeopenedDateHandler = this.changeopenedDateHandler.bind(this);
        this.changecurrencyHandler = this.changecurrencyHandler.bind(this);
        this.changebalanceHandler = this.changebalanceHandler.bind(this);
        this.changeavailableBalanceHandler = this.changeavailableBalanceHandler.bind(this);
        this.changeAccountTypeHandler = this.changeAccountTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        AccountService.getAccountById(this.state.id).then( (res) =>{
            let account = res.data;
            this.setState({
                accountNumber: account.accountNumber,
                iban: account.iban,
                bic: account.bic,
                openedDate: account.openedDate,
                currency: account.currency,
                balance: account.balance,
                availableBalance: account.availableBalance,
                accountType: account.accountType,
                status: account.status
            });
        });
    }

    updateAccount = (e) => {
        e.preventDefault();
        let account = {
            accountId: this.state.id,
            accountNumber: this.state.accountNumber,
            iban: this.state.iban,
            bic: this.state.bic,
            openedDate: this.state.openedDate,
            currency: this.state.currency,
            balance: this.state.balance,
            availableBalance: this.state.availableBalance,
            accountType: this.state.accountType,
            status: this.state.status
        };
        console.log('account => ' + JSON.stringify(account));
        console.log('id => ' + JSON.stringify(this.state.id));
        AccountService.updateAccount(account).then( res => {
            this.props.history.push('/accounts');
        });
    }

    changeaccountNumberHandler= (event) => {
        this.setState({accountNumber: event.target.value});
    }
    changeibanHandler= (event) => {
        this.setState({iban: event.target.value});
    }
    changebicHandler= (event) => {
        this.setState({bic: event.target.value});
    }
    changeopenedDateHandler= (event) => {
        this.setState({openedDate: event.target.value});
    }
    changecurrencyHandler= (event) => {
        this.setState({currency: event.target.value});
    }
    changebalanceHandler= (event) => {
        this.setState({balance: event.target.value});
    }
    changeavailableBalanceHandler= (event) => {
        this.setState({availableBalance: event.target.value});
    }
    changeAccountTypeHandler= (event) => {
        this.setState({accountType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/accounts');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Account</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> accountNumber: </label>
                                                <input placeholder="accountNumber" name="accountNumber" className="form-control" value={this.state.accountNumber} onChange={this.changeaccountNumberHandler}/>

                                            <label> iban: </label>
                                                <input placeholder="iban" name="iban" className="form-control" value={this.state.iban} onChange={this.changeibanHandler}/>

                                            <label> bic: </label>
                                                <input placeholder="bic" name="bic" className="form-control" value={this.state.bic} onChange={this.changebicHandler}/>

                                            <label> openedDate: </label>
                                                <input type="date" placeholder="openedDate" name="openedDate" className="form-control" value={this.state.openedDate} onChange={this.changeopenedDateHandler}/>

                                            <label> currency: </label>
                                                <input placeholder="currency" name="currency" className="form-control" value={this.state.currency} onChange={this.changecurrencyHandler}/>

                                            <label> balance: </label>
                                                <input placeholder="balance" name="balance" className="form-control" value={this.state.balance} onChange={this.changebalanceHandler}/>

                                            <label> availableBalance: </label>
                                                <input placeholder="availableBalance" name="availableBalance" className="form-control" value={this.state.availableBalance} onChange={this.changeavailableBalanceHandler}/>

                                            <label> AccountType: </label>
                                                <select value={this.state.accountType} onChange={this.changeAccountTypeHandler}>
                      <option name="AccountType" className="form-control" >
                          Checking
                      </option>
                      <option name="AccountType" className="form-control" >
                          Savings
                      </option>
                      <option name="AccountType" className="form-control" >
                          Current
                      </option>
                      <option name="AccountType" className="form-control" >
                          Brokerage
                      </option>
                      <option name="AccountType" className="form-control" >
                          Settlement
                      </option>
                      <option name="AccountType" className="form-control" >
                          Escrow
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
                                        <button className="btn btn-success" onClick={this.updateAccount}>Save</button>
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

export default UpdateAccountComponent
