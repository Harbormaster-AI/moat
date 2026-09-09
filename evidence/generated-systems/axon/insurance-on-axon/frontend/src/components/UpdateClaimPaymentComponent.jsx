import React, { Component } from 'react'
import ClaimPaymentService from '../services/ClaimPaymentService';

class UpdateClaimPaymentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                paymentNumber: '',
                amount: '',
                paymentDate: '',
                payeeType: '',
                method: '',
                status: ''
        }
        this.updateClaimPayment = this.updateClaimPayment.bind(this);

        this.changepaymentNumberHandler = this.changepaymentNumberHandler.bind(this);
        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changepaymentDateHandler = this.changepaymentDateHandler.bind(this);
        this.changePayeeTypeHandler = this.changePayeeTypeHandler.bind(this);
        this.changeMethodHandler = this.changeMethodHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ClaimPaymentService.getClaimPaymentById(this.state.id).then( (res) =>{
            let claimPayment = res.data;
            this.setState({
                paymentNumber: claimPayment.paymentNumber,
                amount: claimPayment.amount,
                paymentDate: claimPayment.paymentDate,
                payeeType: claimPayment.payeeType,
                method: claimPayment.method,
                status: claimPayment.status
            });
        });
    }

    updateClaimPayment = (e) => {
        e.preventDefault();
        let claimPayment = {
            claimPaymentId: this.state.id,
            paymentNumber: this.state.paymentNumber,
            amount: this.state.amount,
            paymentDate: this.state.paymentDate,
            payeeType: this.state.payeeType,
            method: this.state.method,
            status: this.state.status
        };
        console.log('claimPayment => ' + JSON.stringify(claimPayment));
        console.log('id => ' + JSON.stringify(this.state.id));
        ClaimPaymentService.updateClaimPayment(claimPayment).then( res => {
            this.props.history.push('/claimPayments');
        });
    }

    changepaymentNumberHandler= (event) => {
        this.setState({paymentNumber: event.target.value});
    }
    changeamountHandler= (event) => {
        this.setState({amount: event.target.value});
    }
    changepaymentDateHandler= (event) => {
        this.setState({paymentDate: event.target.value});
    }
    changePayeeTypeHandler= (event) => {
        this.setState({payeeType: event.target.value});
    }
    changeMethodHandler= (event) => {
        this.setState({method: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/claimPayments');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ClaimPayment</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> paymentNumber: </label>
                                                <input placeholder="paymentNumber" name="paymentNumber" className="form-control" value={this.state.paymentNumber} onChange={this.changepaymentNumberHandler}/>

                                            <label> amount: </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> paymentDate: </label>
                                                <input type="date" placeholder="paymentDate" name="paymentDate" className="form-control" value={this.state.paymentDate} onChange={this.changepaymentDateHandler}/>

                                            <label> PayeeType: </label>
                                                <select value={this.state.payeeType} onChange={this.changePayeeTypeHandler}>
                      <option name="PayeeType" className="form-control" >
                          Claimant
                      </option>
                      <option name="PayeeType" className="form-control" >
                          Beneficiary
                      </option>
                      <option name="PayeeType" className="form-control" >
                          ServiceProvider
                      </option>
                      <option name="PayeeType" className="form-control" >
                          Lienholder
                      </option>
                      <option name="PayeeType" className="form-control" >
                          Attorney
                      </option>
                    </select>

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
                                        <button className="btn btn-success" onClick={this.updateClaimPayment}>Save</button>
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

export default UpdateClaimPaymentComponent
