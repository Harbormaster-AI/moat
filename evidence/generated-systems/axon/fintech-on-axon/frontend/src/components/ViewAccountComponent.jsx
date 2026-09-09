import React, { Component } from 'react'
import AccountService from '../services/AccountService'

class ViewAccountComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            account: {}
        }
    }

    componentDidMount(){
        AccountService.getAccountById(this.state.id).then( res => {
            this.setState({account: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Account Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> accountNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.account.accountNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> iban:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.account.iban }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> bic:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.account.bic }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> openedDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.account.openedDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> currency:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.account.currency }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> balance:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.account.balance }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> availableBalance:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.account.availableBalance }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> AccountType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.account.accountType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.account.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAccountComponent
