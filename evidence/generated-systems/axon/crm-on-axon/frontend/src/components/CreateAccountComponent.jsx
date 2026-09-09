import React, { Component } from 'react'
import AccountService from '../services/AccountService';

class CreateAccountComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                accountNumber: '',
                industry: '',
                billingAddress: '',
                shippingAddress: '',
                website: '',
                phone: '',
                asActive: '',
                accountType: '',
                lifecycleStage: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeaccountNumberHandler = this.changeaccountNumberHandler.bind(this);
        this.changeindustryHandler = this.changeindustryHandler.bind(this);
        this.changebillingAddressHandler = this.changebillingAddressHandler.bind(this);
        this.changeshippingAddressHandler = this.changeshippingAddressHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
        this.changephoneHandler = this.changephoneHandler.bind(this);
        this.changeasActiveHandler = this.changeasActiveHandler.bind(this);
        this.changeAccountTypeHandler = this.changeAccountTypeHandler.bind(this);
        this.changeLifecycleStageHandler = this.changeLifecycleStageHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AccountService.getAccountById(this.state.id).then( (res) =>{
                let account = res.data;
                this.setState({
                    name: account.name,
                    accountNumber: account.accountNumber,
                    industry: account.industry,
                    billingAddress: account.billingAddress,
                    shippingAddress: account.shippingAddress,
                    website: account.website,
                    phone: account.phone,
                    asActive: account.asActive,
                    accountType: account.accountType,
                    lifecycleStage: account.lifecycleStage
                });
            });
        }        
    }
    saveOrUpdateAccount = (e) => {
        e.preventDefault();
        let account = {
                accountId: this.state.id,
                name: this.state.name,
                accountNumber: this.state.accountNumber,
                industry: this.state.industry,
                billingAddress: this.state.billingAddress,
                shippingAddress: this.state.shippingAddress,
                website: this.state.website,
                phone: this.state.phone,
                asActive: this.state.asActive,
                accountType: this.state.accountType,
                lifecycleStage: this.state.lifecycleStage
            };
        console.log('account => ' + JSON.stringify(account));

        // step 5
        if(this.state.id === '_add'){
            account.accountId=''
            AccountService.createAccount(account).then(res =>{
                this.props.history.push('/accounts');
            });
        }else{
            AccountService.updateAccount(account).then( res => {
                this.props.history.push('/accounts');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeaccountNumberHandler= (event) => {
        this.setState({accountNumber: event.target.value});
    }
    changeindustryHandler= (event) => {
        this.setState({industry: event.target.value});
    }
    changebillingAddressHandler= (event) => {
        this.setState({billingAddress: event.target.value});
    }
    changeshippingAddressHandler= (event) => {
        this.setState({shippingAddress: event.target.value});
    }
    changewebsiteHandler= (event) => {
        this.setState({website: event.target.value});
    }
    changephoneHandler= (event) => {
        this.setState({phone: event.target.value});
    }
    changeasActiveHandler= (event) => {
        this.setState({asActive: event.target.value});
    }
    changeAccountTypeHandler= (event) => {
        this.setState({accountType: event.target.value});
    }
    changeLifecycleStageHandler= (event) => {
        this.setState({lifecycleStage: event.target.value});
    }

    cancel(){
        this.props.history.push('/accounts');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Account</h3>
        }else{
            return <h3 className="text-center">Update Account</h3>
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

                                            <label> accountNumber:&emsp; </label>
                                                <input placeholder="accountNumber" name="accountNumber" className="form-control" value={this.state.accountNumber} onChange={this.changeaccountNumberHandler}/>

                                            <label> industry:&emsp; </label>
                                                <input placeholder="industry" name="industry" className="form-control" value={this.state.industry} onChange={this.changeindustryHandler}/>

                                            <label> billingAddress:&emsp; </label>
                                                <input placeholder="billingAddress" name="billingAddress" className="form-control" value={this.state.billingAddress} onChange={this.changebillingAddressHandler}/>

                                            <label> shippingAddress:&emsp; </label>
                                                <input placeholder="shippingAddress" name="shippingAddress" className="form-control" value={this.state.shippingAddress} onChange={this.changeshippingAddressHandler}/>

                                            <label> website:&emsp; </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                            <label> phone:&emsp; </label>
                                                <input placeholder="phone" name="phone" className="form-control" value={this.state.phone} onChange={this.changephoneHandler}/>

                                            <label> asActive:&emsp; </label>
                                                <input type="checkbox" placeholder="asActive" name="asActive" className="form-control" value={this.state.asActive} onChange={this.changeasActiveHandler}/>


                                            <label> AccountType:&emsp; </label>
                                                <select value={this.state.accountType} onChange={this.changeAccountTypeHandler}>
                      <option name="AccountType" className="form-control" >
                          Prospect
                      </option>
                      <option name="AccountType" className="form-control" >
                          Customer
                      </option>
                      <option name="AccountType" className="form-control" >
                          Partner
                      </option>
                      <option name="AccountType" className="form-control" >
                          Vendor
                      </option>
                      <option name="AccountType" className="form-control" >
                          Competitor
                      </option>
                    </select>

                                            <label> LifecycleStage:&emsp; </label>
                                                <select value={this.state.lifecycleStage} onChange={this.changeLifecycleStageHandler}>
                      <option name="LifecycleStage" className="form-control" >
                          Subscriber
                      </option>
                      <option name="LifecycleStage" className="form-control" >
                          Lead
                      </option>
                      <option name="LifecycleStage" className="form-control" >
                          MarketingQualified
                      </option>
                      <option name="LifecycleStage" className="form-control" >
                          SalesQualified
                      </option>
                      <option name="LifecycleStage" className="form-control" >
                          Customer
                      </option>
                      <option name="LifecycleStage" className="form-control" >
                          Evangelist
                      </option>
                      <option name="LifecycleStage" className="form-control" >
                          Churned
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAccount}>Save</button>
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

export default CreateAccountComponent
