import React, { Component } from 'react'
import BankAccountService from '../services/BankAccountService'

class ViewBankAccountComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            bankAccount: {}
        }
    }

    componentDidMount(){
        BankAccountService.getBankAccountById(this.state.id).then( res => {
            this.setState({bankAccount: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View BankAccount Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> accountHolder:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.bankAccount.accountHolder }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> bankName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.bankAccount.bankName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> iban:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.bankAccount.iban }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> bic:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.bankAccount.bic }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> accountNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.bankAccount.accountNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> routingNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.bankAccount.routingNumber }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewBankAccountComponent
