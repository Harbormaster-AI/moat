import React, { Component } from 'react'
import ClaimPaymentService from '../services/ClaimPaymentService';

class CreateClaimPaymentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                paymentNumber: '',
                amount: '',
                paymentDate: '',
                payeeType: '',
                method: '',
                status: ''
        }
        this.changepaymentNumberHandler = this.changepaymentNumberHandler.bind(this);
        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changepaymentDateHandler = this.changepaymentDateHandler.bind(this);
        this.changePayeeTypeHandler = this.changePayeeTypeHandler.bind(this);
        this.changeMethodHandler = this.changeMethodHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateClaimPayment = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            claimPayment.claimPaymentId=''
            ClaimPaymentService.createClaimPayment(claimPayment).then(res =>{
                this.props.history.push('/claimPayments');
            });
        }else{
            ClaimPaymentService.updateClaimPayment(claimPayment).then( res => {
                this.props.history.push('/claimPayments');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ClaimPayment</h3>
        }else{
            return <h3 className="text-center">Update ClaimPayment</h3>
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
                                            <label> paymentNumber:&emsp; </label>
                                                <input placeholder="paymentNumber" name="paymentNumber" className="form-control" value={this.state.paymentNumber} onChange={this.changepaymentNumberHandler}/>

                                            <label> amount:&emsp; </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> paymentDate:&emsp; </label>
                                                <input type="date" placeholder="paymentDate" name="paymentDate" className="form-control" value={this.state.paymentDate} onChange={this.changepaymentDateHandler}/>

                                            <label> PayeeType:&emsp; </label>
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

                                            <label> Method:&emsp; </label>
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

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateClaimPayment}>Save</button>
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

export default CreateClaimPaymentComponent
