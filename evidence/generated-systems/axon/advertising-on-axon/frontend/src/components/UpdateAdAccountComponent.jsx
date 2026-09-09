import React, { Component } from 'react'
import AdAccountService from '../services/AdAccountService';

class UpdateAdAccountComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                accountCode: '',
                defaultCurrency: '',
                defaultTimezone: ''
        }
        this.updateAdAccount = this.updateAdAccount.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeaccountCodeHandler = this.changeaccountCodeHandler.bind(this);
        this.changedefaultCurrencyHandler = this.changedefaultCurrencyHandler.bind(this);
        this.changedefaultTimezoneHandler = this.changedefaultTimezoneHandler.bind(this);
    }

    componentDidMount(){
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

    updateAdAccount = (e) => {
        e.preventDefault();
        let adAccount = {
            adAccountId: this.state.id,
            name: this.state.name,
            accountCode: this.state.accountCode,
            defaultCurrency: this.state.defaultCurrency,
            defaultTimezone: this.state.defaultTimezone
        };
        console.log('adAccount => ' + JSON.stringify(adAccount));
        console.log('id => ' + JSON.stringify(this.state.id));
        AdAccountService.updateAdAccount(adAccount).then( res => {
            this.props.history.push('/adAccounts');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update AdAccount</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> accountCode: </label>
                                                <input placeholder="accountCode" name="accountCode" className="form-control" value={this.state.accountCode} onChange={this.changeaccountCodeHandler}/>

                                            <label> defaultCurrency: </label>
                                                <input placeholder="defaultCurrency" name="defaultCurrency" className="form-control" value={this.state.defaultCurrency} onChange={this.changedefaultCurrencyHandler}/>

                                            <label> defaultTimezone: </label>
                                                <input placeholder="defaultTimezone" name="defaultTimezone" className="form-control" value={this.state.defaultTimezone} onChange={this.changedefaultTimezoneHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateAdAccount}>Save</button>
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

export default UpdateAdAccountComponent
