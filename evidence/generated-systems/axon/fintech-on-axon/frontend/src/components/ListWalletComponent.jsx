import React, { Component } from 'react'
import WalletService from '../services/WalletService'

class ListWalletComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                wallets: []
        }
        this.addWallet = this.addWallet.bind(this);
        this.editWallet = this.editWallet.bind(this);
        this.deleteWallet = this.deleteWallet.bind(this);
    }

    deleteWallet(id){
        WalletService.deleteWallet(id).then( res => {
            this.setState({wallets: this.state.wallets.filter(wallet => wallet.walletId !== id)});
        });
    }
    viewWallet(id){
        this.props.history.push(`/view-wallet/${id}`);
    }
    editWallet(id){
        this.props.history.push(`/add-wallet/${id}`);
    }

    componentDidMount(){
        WalletService.getWallets().then((res) => {
            this.setState({ wallets: res.data});
        });
    }

    addWallet(){
        this.props.history.push('/add-wallet/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Wallet List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addWallet}> Add Wallet</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Currency </th>
                                    <th> Balance </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.wallets.map(
                                        wallet => 
                                        <tr key = {wallet.walletId}>
                                             <td> { wallet.currency } </td>
                                             <td> { wallet.balance } </td>
                                             <td> { wallet.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editWallet(wallet.walletId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteWallet(wallet.walletId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewWallet(wallet.walletId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListWalletComponent
