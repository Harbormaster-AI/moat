import React, { Component } from 'react'
import PaymentService from '../services/PaymentService';

class CreatePaymentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                paymentNumber: '',
                amount: '',
                paymentDate: '',
                method: ''
        }
        this.changepaymentNumberHandler = this.changepaymentNumberHandler.bind(this);
        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changepaymentDateHandler = this.changepaymentDateHandler.bind(this);
        this.changeMethodHandler = this.changeMethodHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PaymentService.getPaymentById(this.state.id).then( (res) =>{
                let payment = res.data;
                this.setState({
                    paymentNumber: payment.paymentNumber,
                    amount: payment.amount,
                    paymentDate: payment.paymentDate,
                    method: payment.method
                });
            });
        }        
    }
    saveOrUpdatePayment = (e) => {
        e.preventDefault();
        let payment = {
                paymentId: this.state.id,
                paymentNumber: this.state.paymentNumber,
                amount: this.state.amount,
                paymentDate: this.state.paymentDate,
                method: this.state.method
            };
        console.log('payment => ' + JSON.stringify(payment));

        // step 5
        if(this.state.id === '_add'){
            payment.paymentId=''
            PaymentService.createPayment(payment).then(res =>{
                this.props.history.push('/payments');
            });
        }else{
            PaymentService.updatePayment(payment).then( res => {
                this.props.history.push('/payments');
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
    changeMethodHandler= (event) => {
        this.setState({method: event.target.value});
    }

    cancel(){
        this.props.history.push('/payments');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Payment</h3>
        }else{
            return <h3 className="text-center">Update Payment</h3>
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

                                            <label> Method:&emsp; </label>
                                                <select value={this.state.method} onChange={this.changeMethodHandler}>
                      <option name="Method" className="form-control" >
                          ACH
                      </option>
                      <option name="Method" className="form-control" >
                          Check
                      </option>
                      <option name="Method" className="form-control" >
                          CreditCard
                      </option>
                      <option name="Method" className="form-control" >
                          EFT
                      </option>
                      <option name="Method" className="form-control" >
                          Cash
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePayment}>Save</button>
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

export default CreatePaymentComponent
