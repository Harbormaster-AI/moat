import React, { Component } from 'react'
import PaymentOrderService from '../services/PaymentOrderService';

class CreatePaymentOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                orderReference: '',
                requestedExecutionDate: '',
                purpose: '',
                paymentMethod: '',
                status: '',
                priority: ''
        }
        this.changeorderReferenceHandler = this.changeorderReferenceHandler.bind(this);
        this.changerequestedExecutionDateHandler = this.changerequestedExecutionDateHandler.bind(this);
        this.changepurposeHandler = this.changepurposeHandler.bind(this);
        this.changePaymentMethodHandler = this.changePaymentMethodHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changePriorityHandler = this.changePriorityHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PaymentOrderService.getPaymentOrderById(this.state.id).then( (res) =>{
                let paymentOrder = res.data;
                this.setState({
                    orderReference: paymentOrder.orderReference,
                    requestedExecutionDate: paymentOrder.requestedExecutionDate,
                    purpose: paymentOrder.purpose,
                    paymentMethod: paymentOrder.paymentMethod,
                    status: paymentOrder.status,
                    priority: paymentOrder.priority
                });
            });
        }        
    }
    saveOrUpdatePaymentOrder = (e) => {
        e.preventDefault();
        let paymentOrder = {
                paymentOrderId: this.state.id,
                orderReference: this.state.orderReference,
                requestedExecutionDate: this.state.requestedExecutionDate,
                purpose: this.state.purpose,
                paymentMethod: this.state.paymentMethod,
                status: this.state.status,
                priority: this.state.priority
            };
        console.log('paymentOrder => ' + JSON.stringify(paymentOrder));

        // step 5
        if(this.state.id === '_add'){
            paymentOrder.paymentOrderId=''
            PaymentOrderService.createPaymentOrder(paymentOrder).then(res =>{
                this.props.history.push('/paymentOrders');
            });
        }else{
            PaymentOrderService.updatePaymentOrder(paymentOrder).then( res => {
                this.props.history.push('/paymentOrders');
            });
        }
    }
    
    changeorderReferenceHandler= (event) => {
        this.setState({orderReference: event.target.value});
    }
    changerequestedExecutionDateHandler= (event) => {
        this.setState({requestedExecutionDate: event.target.value});
    }
    changepurposeHandler= (event) => {
        this.setState({purpose: event.target.value});
    }
    changePaymentMethodHandler= (event) => {
        this.setState({paymentMethod: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changePriorityHandler= (event) => {
        this.setState({priority: event.target.value});
    }

    cancel(){
        this.props.history.push('/paymentOrders');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add PaymentOrder</h3>
        }else{
            return <h3 className="text-center">Update PaymentOrder</h3>
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
                                            <label> orderReference:&emsp; </label>
                                                <input placeholder="orderReference" name="orderReference" className="form-control" value={this.state.orderReference} onChange={this.changeorderReferenceHandler}/>

                                            <label> requestedExecutionDate:&emsp; </label>
                                                <input type="date" placeholder="requestedExecutionDate" name="requestedExecutionDate" className="form-control" value={this.state.requestedExecutionDate} onChange={this.changerequestedExecutionDateHandler}/>

                                            <label> purpose:&emsp; </label>
                                                <input placeholder="purpose" name="purpose" className="form-control" value={this.state.purpose} onChange={this.changepurposeHandler}/>

                                            <label> PaymentMethod:&emsp; </label>
                                                <select value={this.state.paymentMethod} onChange={this.changePaymentMethodHandler}>
                      <option name="PaymentMethod" className="form-control" >
                          Card
                      </option>
                      <option name="PaymentMethod" className="form-control" >
                          BankTransfer
                      </option>
                      <option name="PaymentMethod" className="form-control" >
                          DirectDebit
                      </option>
                      <option name="PaymentMethod" className="form-control" >
                          Wallet
                      </option>
                      <option name="PaymentMethod" className="form-control" >
                          Cash
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Submitted
                      </option>
                      <option name="Status" className="form-control" >
                          Processing
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                      <option name="Status" className="form-control" >
                          Failed
                      </option>
                    </select>

                                            <label> Priority:&emsp; </label>
                                                <select value={this.state.priority} onChange={this.changePriorityHandler}>
                      <option name="Priority" className="form-control" >
                          Normal
                      </option>
                      <option name="Priority" className="form-control" >
                          Urgent
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePaymentOrder}>Save</button>
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

export default CreatePaymentOrderComponent
