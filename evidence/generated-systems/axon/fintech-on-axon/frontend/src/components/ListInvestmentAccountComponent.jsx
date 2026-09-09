import React, { Component } from 'react'
import InvestmentAccountService from '../services/InvestmentAccountService'

class ListInvestmentAccountComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                investmentAccounts: []
        }
        this.addInvestmentAccount = this.addInvestmentAccount.bind(this);
        this.editInvestmentAccount = this.editInvestmentAccount.bind(this);
        this.deleteInvestmentAccount = this.deleteInvestmentAccount.bind(this);
    }

    deleteInvestmentAccount(id){
        InvestmentAccountService.deleteInvestmentAccount(id).then( res => {
            this.setState({investmentAccounts: this.state.investmentAccounts.filter(investmentAccount => investmentAccount.investmentAccountId !== id)});
        });
    }
    viewInvestmentAccount(id){
        this.props.history.push(`/view-investmentAccount/${id}`);
    }
    editInvestmentAccount(id){
        this.props.history.push(`/add-investmentAccount/${id}`);
    }

    componentDidMount(){
        InvestmentAccountService.getInvestmentAccounts().then((res) => {
            this.setState({ investmentAccounts: res.data});
        });
    }

    addInvestmentAccount(){
        this.props.history.push('/add-investmentAccount/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">InvestmentAccount List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addInvestmentAccount}> Add InvestmentAccount</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> AccountNumber </th>
                                    <th> BaseCurrency </th>
                                    <th> Balance </th>
                                    <th> AccountType </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.investmentAccounts.map(
                                        investmentAccount => 
                                        <tr key = {investmentAccount.investmentAccountId}>
                                             <td> { investmentAccount.accountNumber } </td>
                                             <td> { investmentAccount.baseCurrency } </td>
                                             <td> { investmentAccount.balance } </td>
                                             <td> { investmentAccount.accountType } </td>
                                             <td> { investmentAccount.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editInvestmentAccount(investmentAccount.investmentAccountId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteInvestmentAccount(investmentAccount.investmentAccountId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewInvestmentAccount(investmentAccount.investmentAccountId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListInvestmentAccountComponent
