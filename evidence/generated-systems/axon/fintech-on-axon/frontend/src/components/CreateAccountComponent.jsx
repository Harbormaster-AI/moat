import React, { Component } from 'react'
import AccountService from '../services/AccountService';

class CreateAccountComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
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

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateAccount = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            account.accountId=''
            AccountService.createAccount(account).then(res =>{
                this.props.history.push('/accounts');
            });
        }else{
            AccountService.updateAccount(account).then( res => {
                this.props.history.push('/accounts');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Account</h3>
        }else{
            return <h3 className="text-center">Update Account</h3>
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
                                            <label> accountNumber:&emsp; </label>
                                                <input placeholder="accountNumber" name="accountNumber" className="form-control" value={this.state.accountNumber} onChange={this.changeaccountNumberHandler}/>

                                            <label> iban:&emsp; </label>
                                                <input placeholder="iban" name="iban" className="form-control" value={this.state.iban} onChange={this.changeibanHandler}/>

                                            <label> bic:&emsp; </label>
                                                <input placeholder="bic" name="bic" className="form-control" value={this.state.bic} onChange={this.changebicHandler}/>

                                            <label> openedDate:&emsp; </label>
                                                <input type="date" placeholder="openedDate" name="openedDate" className="form-control" value={this.state.openedDate} onChange={this.changeopenedDateHandler}/>

                                            <label> currency:&emsp; </label>
                                                <input placeholder="currency" name="currency" className="form-control" value={this.state.currency} onChange={this.changecurrencyHandler}/>

                                            <label> balance:&emsp; </label>
                                                <input placeholder="balance" name="balance" className="form-control" value={this.state.balance} onChange={this.changebalanceHandler}/>

                                            <label> availableBalance:&emsp; </label>
                                                <input placeholder="availableBalance" name="availableBalance" className="form-control" value={this.state.availableBalance} onChange={this.changeavailableBalanceHandler}/>

                                            <label> AccountType:&emsp; </label>
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

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAccount}>Save</button>
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

export default CreateAccountComponent
