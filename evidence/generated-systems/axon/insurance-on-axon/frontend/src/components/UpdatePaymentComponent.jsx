import React, { Component } from 'react'
import PaymentService from '../services/PaymentService';

class UpdatePaymentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                paymentReference: '',
                amount: '',
                paymentDate: '',
                method: '',
                status: ''
        }
        this.updatePayment = this.updatePayment.bind(this);

        this.changepaymentReferenceHandler = this.changepaymentReferenceHandler.bind(this);
        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changepaymentDateHandler = this.changepaymentDateHandler.bind(this);
        this.changeMethodHandler = this.changeMethodHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        PaymentService.getPaymentById(this.state.id).then( (res) =>{
            let payment = res.data;
            this.setState({
                paymentReference: payment.paymentReference,
                amount: payment.amount,
                paymentDate: payment.paymentDate,
                method: payment.method,
                status: payment.status
            });
        });
    }

    updatePayment = (e) => {
        e.preventDefault();
        let payment = {
            paymentId: this.state.id,
            paymentReference: this.state.paymentReference,
            amount: this.state.amount,
            paymentDate: this.state.paymentDate,
            method: this.state.method,
            status: this.state.status
        };
        console.log('payment => ' + JSON.stringify(payment));
        console.log('id => ' + JSON.stringify(this.state.id));
        PaymentService.updatePayment(payment).then( res => {
            this.props.history.push('/payments');
        });
    }

    changepaymentReferenceHandler= (event) => {
        this.setState({paymentReference: event.target.value});
    }
    changeamountHandler= (event) => {
        this.setState({amount: event.target.value});
    }
    changepaymentDateHandler= (event) => {
        this.setState({paymentDate: event.target.value});
    }
    changeMethodHandler= (event) => {
        this.setState({method: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
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
                                            <label> paymentReference: </label>
                                                <input placeholder="paymentReference" name="paymentReference" className="form-control" value={this.state.paymentReference} onChange={this.changepaymentReferenceHandler}/>

                                            <label> amount: </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> paymentDate: </label>
                                                <input type="date" placeholder="paymentDate" name="paymentDate" className="form-control" value={this.state.paymentDate} onChange={this.changepaymentDateHandler}/>

                                            <label> Method: </label>
                                                <select value={this.state.method} onChange={this.changeMethodHandler}>
                      <option name="Method" className="form-control" >
                          ACH
                      </option>
                      <option name="Method" className="form-control" >
                          CreditCard
                      </option>
                      <option name="Method" className="form-control" >
                          DebitCard
                      </option>
                      <option name="Method" className="form-control" >
                          Check
                      </option>
                      <option name="Method" className="form-control" >
                          Cash
                      </option>
                      <option name="Method" className="form-control" >
                          Wire
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Pending
                      </option>
                      <option name="Status" className="form-control" >
                          Settled
                      </option>
                      <option name="Status" className="form-control" >
                          Failed
                      </option>
                      <option name="Status" className="form-control" >
                          Refunded
                      </option>
                      <option name="Status" className="form-control" >
                          Reversed
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
