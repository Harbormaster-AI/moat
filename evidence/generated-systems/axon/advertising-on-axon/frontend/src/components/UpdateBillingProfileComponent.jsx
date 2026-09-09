import React, { Component } from 'react'
import BillingProfileService from '../services/BillingProfileService';

class UpdateBillingProfileComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                billingName: '',
                taxId: '',
                billingAddress: '',
                paymentTerms: ''
        }
        this.updateBillingProfile = this.updateBillingProfile.bind(this);

        this.changebillingNameHandler = this.changebillingNameHandler.bind(this);
        this.changetaxIdHandler = this.changetaxIdHandler.bind(this);
        this.changebillingAddressHandler = this.changebillingAddressHandler.bind(this);
        this.changePaymentTermsHandler = this.changePaymentTermsHandler.bind(this);
    }

    componentDidMount(){
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

    updateBillingProfile = (e) => {
        e.preventDefault();
        let billingProfile = {
            billingProfileId: this.state.id,
            billingName: this.state.billingName,
            taxId: this.state.taxId,
            billingAddress: this.state.billingAddress,
            paymentTerms: this.state.paymentTerms
        };
        console.log('billingProfile => ' + JSON.stringify(billingProfile));
        console.log('id => ' + JSON.stringify(this.state.id));
        BillingProfileService.updateBillingProfile(billingProfile).then( res => {
            this.props.history.push('/billingProfiles');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update BillingProfile</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> billingName: </label>
                                                <input placeholder="billingName" name="billingName" className="form-control" value={this.state.billingName} onChange={this.changebillingNameHandler}/>

                                            <label> taxId: </label>
                                                <input placeholder="taxId" name="taxId" className="form-control" value={this.state.taxId} onChange={this.changetaxIdHandler}/>

                                            <label> billingAddress: </label>
                                                <input placeholder="billingAddress" name="billingAddress" className="form-control" value={this.state.billingAddress} onChange={this.changebillingAddressHandler}/>

                                            <label> PaymentTerms: </label>
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
                                        <button className="btn btn-success" onClick={this.updateBillingProfile}>Save</button>
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

export default UpdateBillingProfileComponent
