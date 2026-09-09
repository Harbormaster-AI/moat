import React, { Component } from 'react'
import MerchantService from '../services/MerchantService';

class CreateMerchantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                legalName: '',
                website: '',
                defaultCurrency: '',
                defaultLocale: '',
                supportEmail: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changelegalNameHandler = this.changelegalNameHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
        this.changedefaultCurrencyHandler = this.changedefaultCurrencyHandler.bind(this);
        this.changedefaultLocaleHandler = this.changedefaultLocaleHandler.bind(this);
        this.changesupportEmailHandler = this.changesupportEmailHandler.bind(this);
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
                    legalName: merchant.legalName,
                    website: merchant.website,
                    defaultCurrency: merchant.defaultCurrency,
                    defaultLocale: merchant.defaultLocale,
                    supportEmail: merchant.supportEmail
                });
            });
        }        
    }
    saveOrUpdateMerchant = (e) => {
        e.preventDefault();
        let merchant = {
                merchantId: this.state.id,
                name: this.state.name,
                legalName: this.state.legalName,
                website: this.state.website,
                defaultCurrency: this.state.defaultCurrency,
                defaultLocale: this.state.defaultLocale,
                supportEmail: this.state.supportEmail
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
    changelegalNameHandler= (event) => {
        this.setState({legalName: event.target.value});
    }
    changewebsiteHandler= (event) => {
        this.setState({website: event.target.value});
    }
    changedefaultCurrencyHandler= (event) => {
        this.setState({defaultCurrency: event.target.value});
    }
    changedefaultLocaleHandler= (event) => {
        this.setState({defaultLocale: event.target.value});
    }
    changesupportEmailHandler= (event) => {
        this.setState({supportEmail: event.target.value});
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

                                            <label> legalName:&emsp; </label>
                                                <input placeholder="legalName" name="legalName" className="form-control" value={this.state.legalName} onChange={this.changelegalNameHandler}/>

                                            <label> website:&emsp; </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                            <label> defaultCurrency:&emsp; </label>
                                                <input placeholder="defaultCurrency" name="defaultCurrency" className="form-control" value={this.state.defaultCurrency} onChange={this.changedefaultCurrencyHandler}/>

                                            <label> defaultLocale:&emsp; </label>
                                                <input placeholder="defaultLocale" name="defaultLocale" className="form-control" value={this.state.defaultLocale} onChange={this.changedefaultLocaleHandler}/>

                                            <label> supportEmail:&emsp; </label>
                                                <input placeholder="supportEmail" name="supportEmail" className="form-control" value={this.state.supportEmail} onChange={this.changesupportEmailHandler}/>

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
