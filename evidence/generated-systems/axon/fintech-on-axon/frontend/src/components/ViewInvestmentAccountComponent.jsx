import React, { Component } from 'react'
import InvestmentAccountService from '../services/InvestmentAccountService'

class ViewInvestmentAccountComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            investmentAccount: {}
        }
    }

    componentDidMount(){
        InvestmentAccountService.getInvestmentAccountById(this.state.id).then( res => {
            this.setState({investmentAccount: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View InvestmentAccount Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> accountNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.investmentAccount.accountNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> baseCurrency:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.investmentAccount.baseCurrency }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> balance:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.investmentAccount.balance }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> AccountType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.investmentAccount.accountType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.investmentAccount.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewInvestmentAccountComponent
