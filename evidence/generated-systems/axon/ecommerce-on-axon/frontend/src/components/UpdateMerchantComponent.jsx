import React, { Component } from 'react'
import MerchantService from '../services/MerchantService';

class UpdateMerchantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                legalName: '',
                website: '',
                defaultCurrency: '',
                defaultLocale: '',
                supportEmail: ''
        }
        this.updateMerchant = this.updateMerchant.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changelegalNameHandler = this.changelegalNameHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
        this.changedefaultCurrencyHandler = this.changedefaultCurrencyHandler.bind(this);
        this.changedefaultLocaleHandler = this.changedefaultLocaleHandler.bind(this);
        this.changesupportEmailHandler = this.changesupportEmailHandler.bind(this);
    }

    componentDidMount(){
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

    updateMerchant = (e) => {
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
        console.log('id => ' + JSON.stringify(this.state.id));
        MerchantService.updateMerchant(merchant).then( res => {
            this.props.history.push('/merchants');
        });
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

                                            <label> legalName: </label>
                                                <input placeholder="legalName" name="legalName" className="form-control" value={this.state.legalName} onChange={this.changelegalNameHandler}/>

                                            <label> website: </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                            <label> defaultCurrency: </label>
                                                <input placeholder="defaultCurrency" name="defaultCurrency" className="form-control" value={this.state.defaultCurrency} onChange={this.changedefaultCurrencyHandler}/>

                                            <label> defaultLocale: </label>
                                                <input placeholder="defaultLocale" name="defaultLocale" className="form-control" value={this.state.defaultLocale} onChange={this.changedefaultLocaleHandler}/>

                                            <label> supportEmail: </label>
                                                <input placeholder="supportEmail" name="supportEmail" className="form-control" value={this.state.supportEmail} onChange={this.changesupportEmailHandler}/>

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
