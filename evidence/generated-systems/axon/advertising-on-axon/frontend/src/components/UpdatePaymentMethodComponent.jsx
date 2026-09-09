import React, { Component } from 'react'
import PaymentMethodService from '../services/PaymentMethodService';

class UpdatePaymentMethodComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                last4: '',
                cardholderName: '',
                billingAddress: '',
                methodType: ''
        }
        this.updatePaymentMethod = this.updatePaymentMethod.bind(this);

        this.changelast4Handler = this.changelast4Handler.bind(this);
        this.changecardholderNameHandler = this.changecardholderNameHandler.bind(this);
        this.changebillingAddressHandler = this.changebillingAddressHandler.bind(this);
        this.changeMethodTypeHandler = this.changeMethodTypeHandler.bind(this);
    }

    componentDidMount(){
        PaymentMethodService.getPaymentMethodById(this.state.id).then( (res) =>{
            let paymentMethod = res.data;
            this.setState({
                last4: paymentMethod.last4,
                cardholderName: paymentMethod.cardholderName,
                billingAddress: paymentMethod.billingAddress,
                methodType: paymentMethod.methodType
            });
        });
    }

    updatePaymentMethod = (e) => {
        e.preventDefault();
        let paymentMethod = {
            paymentMethodId: this.state.id,
            last4: this.state.last4,
            cardholderName: this.state.cardholderName,
            billingAddress: this.state.billingAddress,
            methodType: this.state.methodType
        };
        console.log('paymentMethod => ' + JSON.stringify(paymentMethod));
        console.log('id => ' + JSON.stringify(this.state.id));
        PaymentMethodService.updatePaymentMethod(paymentMethod).then( res => {
            this.props.history.push('/paymentMethods');
        });
    }

    changelast4Handler= (event) => {
        this.setState({last4: event.target.value});
    }
    changecardholderNameHandler= (event) => {
        this.setState({cardholderName: event.target.value});
    }
    changebillingAddressHandler= (event) => {
        this.setState({billingAddress: event.target.value});
    }
    changeMethodTypeHandler= (event) => {
        this.setState({methodType: event.target.value});
    }

    cancel(){
        this.props.history.push('/paymentMethods');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update PaymentMethod</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> last4: </label>
                                                <input placeholder="last4" name="last4" className="form-control" value={this.state.last4} onChange={this.changelast4Handler}/>

                                            <label> cardholderName: </label>
                                                <input placeholder="cardholderName" name="cardholderName" className="form-control" value={this.state.cardholderName} onChange={this.changecardholderNameHandler}/>

                                            <label> billingAddress: </label>
                                                <input placeholder="billingAddress" name="billingAddress" className="form-control" value={this.state.billingAddress} onChange={this.changebillingAddressHandler}/>

                                            <label> MethodType: </label>
                                                <select value={this.state.methodType} onChange={this.changeMethodTypeHandler}>
                      <option name="MethodType" className="form-control" >
                          CreditCard
                      </option>
                      <option name="MethodType" className="form-control" >
                          Invoice
                      </option>
                      <option name="MethodType" className="form-control" >
                          Wire
                      </option>
                      <option name="MethodType" className="form-control" >
                          ACH
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePaymentMethod}>Save</button>
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

export default UpdatePaymentMethodComponent
