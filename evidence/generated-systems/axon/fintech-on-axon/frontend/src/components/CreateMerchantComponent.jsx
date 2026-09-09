import React, { Component } from 'react'
import MerchantService from '../services/MerchantService';

class CreateMerchantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                mcc: '',
                url: '',
                country: '',
                settlementCurrency: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changemccHandler = this.changemccHandler.bind(this);
        this.changeurlHandler = this.changeurlHandler.bind(this);
        this.changecountryHandler = this.changecountryHandler.bind(this);
        this.changesettlementCurrencyHandler = this.changesettlementCurrencyHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateMerchant = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            merchant.merchantId=''
            MerchantService.createMerchant(merchant).then(res =>{
                this.props.history.push('/merchants');
            });
        }else{
            MerchantService.updateMerchant(merchant).then( res => {
                this.props.history.push('/merchants');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Merchant</h3>
        }else{
            return <h3 className="text-center">Update Merchant</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> mcc:&emsp; </label>
                                                <input placeholder="mcc" name="mcc" className="form-control" value={this.state.mcc} onChange={this.changemccHandler}/>

                                            <label> url:&emsp; </label>
                                                <input placeholder="url" name="url" className="form-control" value={this.state.url} onChange={this.changeurlHandler}/>

                                            <label> country:&emsp; </label>
                                                <input placeholder="country" name="country" className="form-control" value={this.state.country} onChange={this.changecountryHandler}/>

                                            <label> settlementCurrency:&emsp; </label>
                                                <input placeholder="settlementCurrency" name="settlementCurrency" className="form-control" value={this.state.settlementCurrency} onChange={this.changesettlementCurrencyHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateMerchant}>Save</button>
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

export default CreateMerchantComponent
