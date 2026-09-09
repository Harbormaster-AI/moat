import React, { Component } from 'react'
import BillingProfileService from '../services/BillingProfileService';

class CreateBillingProfileComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                billingName: '',
                taxId: '',
                billingAddress: '',
                paymentTerms: ''
        }
        this.changebillingNameHandler = this.changebillingNameHandler.bind(this);
        this.changetaxIdHandler = this.changetaxIdHandler.bind(this);
        this.changebillingAddressHandler = this.changebillingAddressHandler.bind(this);
        this.changePaymentTermsHandler = this.changePaymentTermsHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            BillingProfileService.getBillingProfileById(this.state.id).then( (res) =>{
                let billingProfile = res.data;
                this.setState({
                    billingName: billingProfile.billingName,
                    taxId: billingProfile.taxId,
                    billingAddress: billingProfile.billingAddress,
                    paymentTerms: billingProfile.paymentTerms
                });
            });
        }        
    }
    saveOrUpdateBillingProfile = (e) => {
        e.preventDefault();
        let billingProfile = {
                billingProfileId: this.state.id,
                billingName: this.state.billingName,
                taxId: this.state.taxId,
                billingAddress: this.state.billingAddress,
                paymentTerms: this.state.paymentTerms
            };
        console.log('billingProfile => ' + JSON.stringify(billingProfile));

        // step 5
        if(this.state.id === '_add'){
            billingProfile.billingProfileId=''
            BillingProfileService.createBillingProfile(billingProfile).then(res =>{
                this.props.history.push('/billingProfiles');
            });
        }else{
            BillingProfileService.updateBillingProfile(billingProfile).then( res => {
                this.props.history.push('/billingProfiles');
            });
        }
    }
    
    changebillingNameHandler= (event) => {
        this.setState({billingName: event.target.value});
    }
    changetaxIdHandler= (event) => {
        this.setState({taxId: event.target.value});
    }
    changebillingAddressHandler= (event) => {
        this.setState({billingAddress: event.target.value});
    }
    changePaymentTermsHandler= (event) => {
        this.setState({paymentTerms: event.target.value});
    }

    cancel(){
        this.props.history.push('/billingProfiles');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add BillingProfile</h3>
        }else{
            return <h3 className="text-center">Update BillingProfile</h3>
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
                                            <label> billingName:&emsp; </label>
                                                <input placeholder="billingName" name="billingName" className="form-control" value={this.state.billingName} onChange={this.changebillingNameHandler}/>

                                            <label> taxId:&emsp; </label>
                                                <input placeholder="taxId" name="taxId" className="form-control" value={this.state.taxId} onChange={this.changetaxIdHandler}/>

                                            <label> billingAddress:&emsp; </label>
                                                <input placeholder="billingAddress" name="billingAddress" className="form-control" value={this.state.billingAddress} onChange={this.changebillingAddressHandler}/>

                                            <label> PaymentTerms:&emsp; </label>
                                                <select value={this.state.paymentTerms} onChange={this.changePaymentTermsHandler}>
                      <option name="PaymentTerms" className="form-control" >
                          Prepaid
                      </option>
                      <option name="PaymentTerms" className="form-control" >
                          NetFifteen
                      </option>
                      <option name="PaymentTerms" className="form-control" >
                          NetThirty
                      </option>
                      <option name="PaymentTerms" className="form-control" >
                          NetSixty
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateBillingProfile}>Save</button>
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

export default CreateBillingProfileComponent
