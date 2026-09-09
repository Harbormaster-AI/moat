import React, { Component } from 'react'
import AdAccountService from '../services/AdAccountService';

class CreateAdAccountComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                accountCode: '',
                defaultCurrency: '',
                defaultTimezone: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeaccountCodeHandler = this.changeaccountCodeHandler.bind(this);
        this.changedefaultCurrencyHandler = this.changedefaultCurrencyHandler.bind(this);
        this.changedefaultTimezoneHandler = this.changedefaultTimezoneHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AdAccountService.getAdAccountById(this.state.id).then( (res) =>{
                let adAccount = res.data;
                this.setState({
                    name: adAccount.name,
                    accountCode: adAccount.accountCode,
                    defaultCurrency: adAccount.defaultCurrency,
                    defaultTimezone: adAccount.defaultTimezone
                });
            });
        }        
    }
    saveOrUpdateAdAccount = (e) => {
        e.preventDefault();
        let adAccount = {
                adAccountId: this.state.id,
                name: this.state.name,
                accountCode: this.state.accountCode,
                defaultCurrency: this.state.defaultCurrency,
                defaultTimezone: this.state.defaultTimezone
            };
        console.log('adAccount => ' + JSON.stringify(adAccount));

        // step 5
        if(this.state.id === '_add'){
            adAccount.adAccountId=''
            AdAccountService.createAdAccount(adAccount).then(res =>{
                this.props.history.push('/adAccounts');
            });
        }else{
            AdAccountService.updateAdAccount(adAccount).then( res => {
                this.props.history.push('/adAccounts');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeaccountCodeHandler= (event) => {
        this.setState({accountCode: event.target.value});
    }
    changedefaultCurrencyHandler= (event) => {
        this.setState({defaultCurrency: event.target.value});
    }
    changedefaultTimezoneHandler= (event) => {
        this.setState({defaultTimezone: event.target.value});
    }

    cancel(){
        this.props.history.push('/adAccounts');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add AdAccount</h3>
        }else{
            return <h3 className="text-center">Update AdAccount</h3>
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

                                            <label> accountCode:&emsp; </label>
                                                <input placeholder="accountCode" name="accountCode" className="form-control" value={this.state.accountCode} onChange={this.changeaccountCodeHandler}/>

                                            <label> defaultCurrency:&emsp; </label>
                                                <input placeholder="defaultCurrency" name="defaultCurrency" className="form-control" value={this.state.defaultCurrency} onChange={this.changedefaultCurrencyHandler}/>

                                            <label> defaultTimezone:&emsp; </label>
                                                <input placeholder="defaultTimezone" name="defaultTimezone" className="form-control" value={this.state.defaultTimezone} onChange={this.changedefaultTimezoneHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAdAccount}>Save</button>
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

export default CreateAdAccountComponent
