import React, { Component } from 'react'
import PurchaseOrderService from '../services/PurchaseOrderService';

class CreatePurchaseOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                poNumber: '',
                orderDate: '',
                totalAmount: '',
                status: ''
        }
        this.changepoNumberHandler = this.changepoNumberHandler.bind(this);
        this.changeorderDateHandler = this.changeorderDateHandler.bind(this);
        this.changetotalAmountHandler = this.changetotalAmountHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PurchaseOrderService.getPurchaseOrderById(this.state.id).then( (res) =>{
                let purchaseOrder = res.data;
                this.setState({
                    poNumber: purchaseOrder.poNumber,
                    orderDate: purchaseOrder.orderDate,
                    totalAmount: purchaseOrder.totalAmount,
                    status: purchaseOrder.status
                });
            });
        }        
    }
    saveOrUpdatePurchaseOrder = (e) => {
        e.preventDefault();
        let purchaseOrder = {
                purchaseOrderId: this.state.id,
                poNumber: this.state.poNumber,
                orderDate: this.state.orderDate,
                totalAmount: this.state.totalAmount,
                status: this.state.status
            };
        console.log('purchaseOrder => ' + JSON.stringify(purchaseOrder));

        // step 5
        if(this.state.id === '_add'){
            purchaseOrder.purchaseOrderId=''
            PurchaseOrderService.createPurchaseOrder(purchaseOrder).then(res =>{
                this.props.history.push('/purchaseOrders');
            });
        }else{
            PurchaseOrderService.updatePurchaseOrder(purchaseOrder).then( res => {
                this.props.history.push('/purchaseOrders');
            });
        }
    }
    
    changepoNumberHandler= (event) => {
        this.setState({poNumber: event.target.value});
    }
    changeorderDateHandler= (event) => {
        this.setState({orderDate: event.target.value});
    }
    changetotalAmountHandler= (event) => {
        this.setState({totalAmount: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/purchaseOrders');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add PurchaseOrder</h3>
        }else{
            return <h3 className="text-center">Update PurchaseOrder</h3>
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
                                            <label> poNumber:&emsp; </label>
                                                <input placeholder="poNumber" name="poNumber" className="form-control" value={this.state.poNumber} onChange={this.changepoNumberHandler}/>

                                            <label> orderDate:&emsp; </label>
                                                <input type="date" placeholder="orderDate" name="orderDate" className="form-control" value={this.state.orderDate} onChange={this.changeorderDateHandler}/>

                                            <label> totalAmount:&emsp; </label>
                                                <input placeholder="totalAmount" name="totalAmount" className="form-control" value={this.state.totalAmount} onChange={this.changetotalAmountHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Submitted
                      </option>
                      <option name="Status" className="form-control" >
                          Acknowledged
                      </option>
                      <option name="Status" className="form-control" >
                          PartiallyReceived
                      </option>
                      <option name="Status" className="form-control" >
                          Received
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePurchaseOrder}>Save</button>
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

export default CreatePurchaseOrderComponent
