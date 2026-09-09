import React, { Component } from 'react'
import BankAccountService from '../services/BankAccountService'

class ListBankAccountComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                bankAccounts: []
        }
        this.addBankAccount = this.addBankAccount.bind(this);
        this.editBankAccount = this.editBankAccount.bind(this);
        this.deleteBankAccount = this.deleteBankAccount.bind(this);
    }

    deleteBankAccount(id){
        BankAccountService.deleteBankAccount(id).then( res => {
            this.setState({bankAccounts: this.state.bankAccounts.filter(bankAccount => bankAccount.bankAccountId !== id)});
        });
    }
    viewBankAccount(id){
        this.props.history.push(`/view-bankAccount/${id}`);
    }
    editBankAccount(id){
        this.props.history.push(`/add-bankAccount/${id}`);
    }

    componentDidMount(){
        BankAccountService.getBankAccounts().then((res) => {
            this.setState({ bankAccounts: res.data});
        });
    }

    addBankAccount(){
        this.props.history.push('/add-bankAccount/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">BankAccount List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addBankAccount}> Add BankAccount</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> AccountHolder </th>
                                    <th> BankName </th>
                                    <th> Iban </th>
                                    <th> Bic </th>
                                    <th> AccountNumber </th>
                                    <th> RoutingNumber </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.bankAccounts.map(
                                        bankAccount => 
                                        <tr key = {bankAccount.bankAccountId}>
                                             <td> { bankAccount.accountHolder } </td>
                                             <td> { bankAccount.bankName } </td>
                                             <td> { bankAccount.iban } </td>
                                             <td> { bankAccount.bic } </td>
                                             <td> { bankAccount.accountNumber } </td>
                                             <td> { bankAccount.routingNumber } </td>
                                             <td>
                                                 <button onClick={ () => this.editBankAccount(bankAccount.bankAccountId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteBankAccount(bankAccount.bankAccountId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewBankAccount(bankAccount.bankAccountId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListBankAccountComponent
