import React, { Component } from 'react'
import AdAccountService from '../services/AdAccountService'

class ListAdAccountComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                adAccounts: []
        }
        this.addAdAccount = this.addAdAccount.bind(this);
        this.editAdAccount = this.editAdAccount.bind(this);
        this.deleteAdAccount = this.deleteAdAccount.bind(this);
    }

    deleteAdAccount(id){
        AdAccountService.deleteAdAccount(id).then( res => {
            this.setState({adAccounts: this.state.adAccounts.filter(adAccount => adAccount.adAccountId !== id)});
        });
    }
    viewAdAccount(id){
        this.props.history.push(`/view-adAccount/${id}`);
    }
    editAdAccount(id){
        this.props.history.push(`/add-adAccount/${id}`);
    }

    componentDidMount(){
        AdAccountService.getAdAccounts().then((res) => {
            this.setState({ adAccounts: res.data});
        });
    }

    addAdAccount(){
        this.props.history.push('/add-adAccount/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">AdAccount List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAdAccount}> Add AdAccount</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> AccountCode </th>
                                    <th> DefaultCurrency </th>
                                    <th> DefaultTimezone </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.adAccounts.map(
                                        adAccount => 
                                        <tr key = {adAccount.adAccountId}>
                                             <td> { adAccount.name } </td>
                                             <td> { adAccount.accountCode } </td>
                                             <td> { adAccount.defaultCurrency } </td>
                                             <td> { adAccount.defaultTimezone } </td>
                                             <td>
                                                 <button onClick={ () => this.editAdAccount(adAccount.adAccountId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAdAccount(adAccount.adAccountId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAdAccount(adAccount.adAccountId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAdAccountComponent
