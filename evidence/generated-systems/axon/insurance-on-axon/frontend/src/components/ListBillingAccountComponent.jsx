import React, { Component } from 'react'
import BillingAccountService from '../services/BillingAccountService'

class ListBillingAccountComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                billingAccounts: []
        }
        this.addBillingAccount = this.addBillingAccount.bind(this);
        this.editBillingAccount = this.editBillingAccount.bind(this);
        this.deleteBillingAccount = this.deleteBillingAccount.bind(this);
    }

    deleteBillingAccount(id){
        BillingAccountService.deleteBillingAccount(id).then( res => {
            this.setState({billingAccounts: this.state.billingAccounts.filter(billingAccount => billingAccount.billingAccountId !== id)});
        });
    }
    viewBillingAccount(id){
        this.props.history.push(`/view-billingAccount/${id}`);
    }
    editBillingAccount(id){
        this.props.history.push(`/add-billingAccount/${id}`);
    }

    componentDidMount(){
        BillingAccountService.getBillingAccounts().then((res) => {
            this.setState({ billingAccounts: res.data});
        });
    }

    addBillingAccount(){
        this.props.history.push('/add-billingAccount/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">BillingAccount List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addBillingAccount}> Add BillingAccount</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> AccountNumber </th>
                                    <th> Balance </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.billingAccounts.map(
                                        billingAccount => 
                                        <tr key = {billingAccount.billingAccountId}>
                                             <td> { billingAccount.accountNumber } </td>
                                             <td> { billingAccount.balance } </td>
                                             <td> { billingAccount.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editBillingAccount(billingAccount.billingAccountId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteBillingAccount(billingAccount.billingAccountId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewBillingAccount(billingAccount.billingAccountId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListBillingAccountComponent
