import React, { Component } from 'react'
import CardTokenizationService from '../services/CardTokenizationService';

class CreateCardTokenizationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                tokenReference: '',
                createdAt: '',
                walletProvider: '',
                status: ''
        }
        this.changetokenReferenceHandler = this.changetokenReferenceHandler.bind(this);
        this.changecreatedAtHandler = this.changecreatedAtHandler.bind(this);
        this.changeWalletProviderHandler = this.changeWalletProviderHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            CardTokenizationService.getCardTokenizationById(this.state.id).then( (res) =>{
                let cardTokenization = res.data;
                this.setState({
                    tokenReference: cardTokenization.tokenReference,
                    createdAt: cardTokenization.createdAt,
                    walletProvider: cardTokenization.walletProvider,
                    status: cardTokenization.status
                });
            });
        }        
    }
    saveOrUpdateCardTokenization = (e) => {
        e.preventDefault();
        let cardTokenization = {
                cardTokenizationId: this.state.id,
                tokenReference: this.state.tokenReference,
                createdAt: this.state.createdAt,
                walletProvider: this.state.walletProvider,
                status: this.state.status
            };
        console.log('cardTokenization => ' + JSON.stringify(cardTokenization));

        // step 5
        if(this.state.id === '_add'){
            cardTokenization.cardTokenizationId=''
            CardTokenizationService.createCardTokenization(cardTokenization).then(res =>{
                this.props.history.push('/cardTokenizations');
            });
        }else{
            CardTokenizationService.updateCardTokenization(cardTokenization).then( res => {
                this.props.history.push('/cardTokenizations');
            });
        }
    }
    
    changetokenReferenceHandler= (event) => {
        this.setState({tokenReference: event.target.value});
    }
    changecreatedAtHandler= (event) => {
        this.setState({createdAt: event.target.value});
    }
    changeWalletProviderHandler= (event) => {
        this.setState({walletProvider: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/cardTokenizations');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add CardTokenization</h3>
        }else{
            return <h3 className="text-center">Update CardTokenization</h3>
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
                                            <label> tokenReference:&emsp; </label>
                                                <input placeholder="tokenReference" name="tokenReference" className="form-control" value={this.state.tokenReference} onChange={this.changetokenReferenceHandler}/>

                                            <label> createdAt:&emsp; </label>
                                                <input type="time" placeholder="createdAt" name="createdAt" className="form-control" value={this.state.createdAt} onChange={this.changecreatedAtHandler}/>

                                            <label> WalletProvider:&emsp; </label>
                                                <select value={this.state.walletProvider} onChange={this.changeWalletProviderHandler}>
                      <option name="WalletProvider" className="form-control" >
                          ApplePay
                      </option>
                      <option name="WalletProvider" className="form-control" >
                          GooglePay
                      </option>
                      <option name="WalletProvider" className="form-control" >
                          SamsungPay
                      </option>
                      <option name="WalletProvider" className="form-control" >
                          Other
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Suspended
                      </option>
                      <option name="Status" className="form-control" >
                          Deactivated
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCardTokenization}>Save</button>
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

export default CreateCardTokenizationComponent
