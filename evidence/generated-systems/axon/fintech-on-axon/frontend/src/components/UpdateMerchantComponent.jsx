import React, { Component } from 'react'
import MerchantService from '../services/MerchantService';

class UpdateMerchantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                mcc: '',
                url: '',
                country: '',
                settlementCurrency: ''
        }
        this.updateMerchant = this.updateMerchant.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changemccHandler = this.changemccHandler.bind(this);
        this.changeurlHandler = this.changeurlHandler.bind(this);
        this.changecountryHandler = this.changecountryHandler.bind(this);
        this.changesettlementCurrencyHandler = this.changesettlementCurrencyHandler.bind(this);
    }

    componentDidMount(){
        MerchantService.getMerchantById(this.state.id).then( (res) =>{
            let merchant = res.data;
            this.setState({
                name: merchant.name,
                mcc: merchant.mcc,
                url: merchant.url,
                country: merchant.country,
                settlementCurrency: merchant.settlementCurrency
            });
        });
    }

    updateMerchant = (e) => {
        e.preventDefault();
        let merchant = {
            merchantId: this.state.id,
            name: this.state.name,
            mcc: this.state.mcc,
            url: this.state.url,
            country: this.state.country,
            settlementCurrency: this.state.settlementCurrency
        };
        console.log('merchant => ' + JSON.stringify(merchant));
        console.log('id => ' + JSON.stringify(this.state.id));
        MerchantService.updateMerchant(merchant).then( res => {
            this.props.history.push('/merchants');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changemccHandler= (event) => {
        this.setState({mcc: event.target.value});
    }
    changeurlHandler= (event) => {
        this.setState({url: event.target.value});
    }
    changecountryHandler= (event) => {
        this.setState({country: event.target.value});
    }
    changesettlementCurrencyHandler= (event) => {
        this.setState({settlementCurrency: event.target.value});
    }

    cancel(){
        this.props.history.push('/merchants');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Merchant</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> mcc: </label>
                                                <input placeholder="mcc" name="mcc" className="form-control" value={this.state.mcc} onChange={this.changemccHandler}/>

                                            <label> url: </label>
                                                <input placeholder="url" name="url" className="form-control" value={this.state.url} onChange={this.changeurlHandler}/>

                                            <label> country: </label>
                                                <input placeholder="country" name="country" className="form-control" value={this.state.country} onChange={this.changecountryHandler}/>

                                            <label> settlementCurrency: </label>
                                                <input placeholder="settlementCurrency" name="settlementCurrency" className="form-control" value={this.state.settlementCurrency} onChange={this.changesettlementCurrencyHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateMerchant}>Save</button>
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

export default UpdateMerchantComponent
