import React, { Component } from 'react'
import BillingAccountService from '../services/BillingAccountService';

class CreateBillingAccountComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                accountNumber: '',
                balance: '',
                status: ''
        }
        this.changeaccountNumberHandler = this.changeaccountNumberHandler.bind(this);
        this.changebalanceHandler = this.changebalanceHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            BillingAccountService.getBillingAccountById(this.state.id).then( (res) =>{
                let billingAccount = res.data;
                this.setState({
                    accountNumber: billingAccount.accountNumber,
                    balance: billingAccount.balance,
                    status: billingAccount.status
                });
            });
        }        
    }
    saveOrUpdateBillingAccount = (e) => {
        e.preventDefault();
        let billingAccount = {
                billingAccountId: this.state.id,
                accountNumber: this.state.accountNumber,
                balance: this.state.balance,
                status: this.state.status
            };
        console.log('billingAccount => ' + JSON.stringify(billingAccount));

        // step 5
        if(this.state.id === '_add'){
            billingAccount.billingAccountId=''
            BillingAccountService.createBillingAccount(billingAccount).then(res =>{
                this.props.history.push('/billingAccounts');
            });
        }else{
            BillingAccountService.updateBillingAccount(billingAccount).then( res => {
                this.props.history.push('/billingAccounts');
            });
        }
    }
    
    changeaccountNumberHandler= (event) => {
        this.setState({accountNumber: event.target.value});
    }
    changebalanceHandler= (event) => {
        this.setState({balance: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/billingAccounts');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add BillingAccount</h3>
        }else{
            return <h3 className="text-center">Update BillingAccount</h3>
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
                                            <label> accountNumber:&emsp; </label>
                                                <input placeholder="accountNumber" name="accountNumber" className="form-control" value={this.state.accountNumber} onChange={this.changeaccountNumberHandler}/>

                                            <label> balance:&emsp; </label>
                                                <input placeholder="balance" name="balance" className="form-control" value={this.state.balance} onChange={this.changebalanceHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Current
                      </option>
                      <option name="Status" className="form-control" >
                          Delinquent
                      </option>
                      <option name="Status" className="form-control" >
                          Collections
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateBillingAccount}>Save</button>
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

export default CreateBillingAccountComponent
