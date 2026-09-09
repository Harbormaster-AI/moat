import React, { Component } from 'react'
import WalletService from '../services/WalletService';

class CreateWalletComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                currency: '',
                balance: '',
                status: ''
        }
        this.changecurrencyHandler = this.changecurrencyHandler.bind(this);
        this.changebalanceHandler = this.changebalanceHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            WalletService.getWalletById(this.state.id).then( (res) =>{
                let wallet = res.data;
                this.setState({
                    currency: wallet.currency,
                    balance: wallet.balance,
                    status: wallet.status
                });
            });
        }        
    }
    saveOrUpdateWallet = (e) => {
        e.preventDefault();
        let wallet = {
                walletId: this.state.id,
                currency: this.state.currency,
                balance: this.state.balance,
                status: this.state.status
            };
        console.log('wallet => ' + JSON.stringify(wallet));

        // step 5
        if(this.state.id === '_add'){
            wallet.walletId=''
            WalletService.createWallet(wallet).then(res =>{
                this.props.history.push('/wallets');
            });
        }else{
            WalletService.updateWallet(wallet).then( res => {
                this.props.history.push('/wallets');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Wallet</h3>
        }else{
            return <h3 className="text-center">Update Wallet</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> currency:&emsp; </label>
                                                <input placeholder="currency" name="currency" className="form-control" value={this.state.currency} onChange={this.changecurrencyHandler}/>

                                            <label> balance:&emsp; </label>
                                                <input placeholder="balance" name="balance" className="form-control" value={this.state.balance} onChange={this.changebalanceHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateWallet}>Save</button>
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

export default CreateWalletComponent
