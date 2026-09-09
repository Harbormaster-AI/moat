import React, { Component } from 'react'
import PaymentService from '../services/PaymentService';

class UpdatePaymentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                paymentNumber: '',
                amount: '',
                transactionId: '',
                authorizedAt: '',
                capturedAt: '',
                status: '',
                paymentMethod: ''
        }
        this.updatePayment = this.updatePayment.bind(this);

        this.changepaymentNumberHandler = this.changepaymentNumberHandler.bind(this);
        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changetransactionIdHandler = this.changetransactionIdHandler.bind(this);
        this.changeauthorizedAtHandler = this.changeauthorizedAtHandler.bind(this);
        this.changecapturedAtHandler = this.changecapturedAtHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changePaymentMethodHandler = this.changePaymentMethodHandler.bind(this);
    }

    componentDidMount(){
        PaymentService.getPaymentById(this.state.id).then( (res) =>{
            let payment = res.data;
            this.setState({
                paymentNumber: payment.paymentNumber,
                amount: payment.amount,
                transactionId: payment.transactionId,
                authorizedAt: payment.authorizedAt,
                capturedAt: payment.capturedAt,
                status: payment.status,
                paymentMethod: payment.paymentMethod
            });
        });
    }

    updatePayment = (e) => {
        e.preventDefault();
        let payment = {
            paymentId: this.state.id,
            paymentNumber: this.state.paymentNumber,
            amount: this.state.amount,
            transactionId: this.state.transactionId,
            authorizedAt: this.state.authorizedAt,
            capturedAt: this.state.capturedAt,
            status: this.state.status,
            paymentMethod: this.state.paymentMethod
        };
        console.log('payment => ' + JSON.stringify(payment));
        console.log('id => ' + JSON.stringify(this.state.id));
        PaymentService.updatePayment(payment).then( res => {
            this.props.history.push('/payments');
        });
    }

    changepaymentNumberHandler= (event) => {
        this.setState({paymentNumber: event.target.value});
    }
    changeamountHandler= (event) => {
        this.setState({amount: event.target.value});
    }
    changetransactionIdHandler= (event) => {
        this.setState({transactionId: event.target.value});
    }
    changeauthorizedAtHandler= (event) => {
        this.setState({authorizedAt: event.target.value});
    }
    changecapturedAtHandler= (event) => {
        this.setState({capturedAt: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changePaymentMethodHandler= (event) => {
        this.setState({paymentMethod: event.target.value});
    }

    cancel(){
        this.props.history.push('/payments');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Payment</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> paymentNumber: </label>
                                                <input placeholder="paymentNumber" name="paymentNumber" className="form-control" value={this.state.paymentNumber} onChange={this.changepaymentNumberHandler}/>

                                            <label> amount: </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> transactionId: </label>
                                                <input placeholder="transactionId" name="transactionId" className="form-control" value={this.state.transactionId} onChange={this.changetransactionIdHandler}/>

                                            <label> authorizedAt: </label>
                                                <input type="date" placeholder="authorizedAt" name="authorizedAt" className="form-control" value={this.state.authorizedAt} onChange={this.changeauthorizedAtHandler}/>

                                            <label> capturedAt: </label>
                                                <input type="date" placeholder="capturedAt" name="capturedAt" className="form-control" value={this.state.capturedAt} onChange={this.changecapturedAtHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Authorized
                      </option>
                      <option name="Status" className="form-control" >
                          Captured
                      </option>
                      <option name="Status" className="form-control" >
                          PartiallyCaptured
                      </option>
                      <option name="Status" className="form-control" >
                          Declined
                      </option>
                      <option name="Status" className="form-control" >
                          Refunded
                      </option>
                      <option name="Status" className="form-control" >
                          PartiallyRefunded
                      </option>
                      <option name="Status" className="form-control" >
                          Voided
                      </option>
                      <option name="Status" className="form-control" >
                          Pending
                      </option>
                    </select>

                                            <label> PaymentMethod: </label>
                                                <select value={this.state.paymentMethod} onChange={this.changePaymentMethodHandler}>
                      <option name="PaymentMethod" className="form-control" >
                          CreditCard
                      </option>
                      <option name="PaymentMethod" className="form-control" >
                          DebitCard
                      </option>
                      <option name="PaymentMethod" className="form-control" >
                          PayPal
                      </option>
                      <option name="PaymentMethod" className="form-control" >
                          BankTransfer
                      </option>
                      <option name="PaymentMethod" className="form-control" >
                          CashOnDelivery
                      </option>
                      <option name="PaymentMethod" className="form-control" >
                          GiftCard
                      </option>
                      <option name="PaymentMethod" className="form-control" >
                          ApplePay
                      </option>
                      <option name="PaymentMethod" className="form-control" >
                          GooglePay
                      </option>
                      <option name="PaymentMethod" className="form-control" >
                          BuyNowPayLater
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePayment}>Save</button>
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

export default UpdatePaymentComponent
