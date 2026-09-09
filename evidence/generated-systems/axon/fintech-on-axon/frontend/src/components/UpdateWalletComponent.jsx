import React, { Component } from 'react'
import WalletService from '../services/WalletService';

class UpdateWalletComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                currency: '',
                balance: '',
                status: ''
        }
        this.updateWallet = this.updateWallet.bind(this);

        this.changecurrencyHandler = this.changecurrencyHandler.bind(this);
        this.changebalanceHandler = this.changebalanceHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        WalletService.getWalletById(this.state.id).then( (res) =>{
            let wallet = res.data;
            this.setState({
                currency: wallet.currency,
                balance: wallet.balance,
                status: wallet.status
            });
        });
    }

    updateWallet = (e) => {
        e.preventDefault();
        let wallet = {
            walletId: this.state.id,
            currency: this.state.currency,
            balance: this.state.balance,
            status: this.state.status
        };
        console.log('wallet => ' + JSON.stringify(wallet));
        console.log('id => ' + JSON.stringify(this.state.id));
        WalletService.updateWallet(wallet).then( res => {
            this.props.history.push('/wallets');
        });
    }

    changecurrencyHandler= (event) => {
        this.setState({currency: event.target.value});
    }
    changebalanceHandler= (event) => {
        this.setState({balance: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/wallets');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Wallet</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> currency: </label>
                                                <input placeholder="currency" name="currency" className="form-control" value={this.state.currency} onChange={this.changecurrencyHandler}/>

                                            <label> balance: </label>
                                                <input placeholder="balance" name="balance" className="form-control" value={this.state.balance} onChange={this.changebalanceHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Suspended
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateWallet}>Save</button>
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

export default UpdateWalletComponent
