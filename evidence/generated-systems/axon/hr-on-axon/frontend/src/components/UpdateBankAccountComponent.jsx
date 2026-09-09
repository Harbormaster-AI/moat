import React, { Component } from 'react'
import BankAccountService from '../services/BankAccountService';

class UpdateBankAccountComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                accountHolder: '',
                bankName: '',
                iban: '',
                bic: '',
                accountNumber: '',
                routingNumber: ''
        }
        this.updateBankAccount = this.updateBankAccount.bind(this);

        this.changeaccountHolderHandler = this.changeaccountHolderHandler.bind(this);
        this.changebankNameHandler = this.changebankNameHandler.bind(this);
        this.changeibanHandler = this.changeibanHandler.bind(this);
        this.changebicHandler = this.changebicHandler.bind(this);
        this.changeaccountNumberHandler = this.changeaccountNumberHandler.bind(this);
        this.changeroutingNumberHandler = this.changeroutingNumberHandler.bind(this);
    }

    componentDidMount(){
        BankAccountService.getBankAccountById(this.state.id).then( (res) =>{
            let bankAccount = res.data;
            this.setState({
                accountHolder: bankAccount.accountHolder,
                bankName: bankAccount.bankName,
                iban: bankAccount.iban,
                bic: bankAccount.bic,
                accountNumber: bankAccount.accountNumber,
                routingNumber: bankAccount.routingNumber
            });
        });
    }

    updateBankAccount = (e) => {
        e.preventDefault();
        let bankAccount = {
            bankAccountId: this.state.id,
            accountHolder: this.state.accountHolder,
            bankName: this.state.bankName,
            iban: this.state.iban,
            bic: this.state.bic,
            accountNumber: this.state.accountNumber,
            routingNumber: this.state.routingNumber
        };
        console.log('bankAccount => ' + JSON.stringify(bankAccount));
        console.log('id => ' + JSON.stringify(this.state.id));
        BankAccountService.updateBankAccount(bankAccount).then( res => {
            this.props.history.push('/bankAccounts');
        });
    }

    changeaccountHolderHandler= (event) => {
        this.setState({accountHolder: event.target.value});
    }
    changebankNameHandler= (event) => {
        this.setState({bankName: event.target.value});
    }
    changeibanHandler= (event) => {
        this.setState({iban: event.target.value});
    }
    changebicHandler= (event) => {
        this.setState({bic: event.target.value});
    }
    changeaccountNumberHandler= (event) => {
        this.setState({accountNumber: event.target.value});
    }
    changeroutingNumberHandler= (event) => {
        this.setState({routingNumber: event.target.value});
    }

    cancel(){
        this.props.history.push('/bankAccounts');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update BankAccount</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> accountHolder: </label>
                                                <input placeholder="accountHolder" name="accountHolder" className="form-control" value={this.state.accountHolder} onChange={this.changeaccountHolderHandler}/>

                                            <label> bankName: </label>
                                                <input placeholder="bankName" name="bankName" className="form-control" value={this.state.bankName} onChange={this.changebankNameHandler}/>

                                            <label> iban: </label>
                                                <input placeholder="iban" name="iban" className="form-control" value={this.state.iban} onChange={this.changeibanHandler}/>

                                            <label> bic: </label>
                                                <input placeholder="bic" name="bic" className="form-control" value={this.state.bic} onChange={this.changebicHandler}/>

                                            <label> accountNumber: </label>
                                                <input placeholder="accountNumber" name="accountNumber" className="form-control" value={this.state.accountNumber} onChange={this.changeaccountNumberHandler}/>

                                            <label> routingNumber: </label>
                                                <input placeholder="routingNumber" name="routingNumber" className="form-control" value={this.state.routingNumber} onChange={this.changeroutingNumberHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateBankAccount}>Save</button>
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

export default UpdateBankAccountComponent
