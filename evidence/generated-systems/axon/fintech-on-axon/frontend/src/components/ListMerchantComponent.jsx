import React, { Component } from 'react'
import MerchantService from '../services/MerchantService'

class ListMerchantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                merchants: []
        }
        this.addMerchant = this.addMerchant.bind(this);
        this.editMerchant = this.editMerchant.bind(this);
        this.deleteMerchant = this.deleteMerchant.bind(this);
    }

    deleteMerchant(id){
        MerchantService.deleteMerchant(id).then( res => {
            this.setState({merchants: this.state.merchants.filter(merchant => merchant.merchantId !== id)});
        });
    }
    viewMerchant(id){
        this.props.history.push(`/view-merchant/${id}`);
    }
    editMerchant(id){
        this.props.history.push(`/add-merchant/${id}`);
    }

    componentDidMount(){
        MerchantService.getMerchants().then((res) => {
            this.setState({ merchants: res.data});
        });
    }

    addMerchant(){
        this.props.history.push('/add-merchant/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Merchant List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addMerchant}> Add Merchant</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Mcc </th>
                                    <th> Url </th>
                                    <th> Country </th>
                                    <th> SettlementCurrency </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.merchants.map(
                                        merchant => 
                                        <tr key = {merchant.merchantId}>
                                             <td> { merchant.name } </td>
                                             <td> { merchant.mcc } </td>
                                             <td> { merchant.url } </td>
                                             <td> { merchant.country } </td>
                                             <td> { merchant.settlementCurrency } </td>
                                             <td>
                                                 <button onClick={ () => this.editMerchant(merchant.merchantId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteMerchant(merchant.merchantId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewMerchant(merchant.merchantId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListMerchantComponent
