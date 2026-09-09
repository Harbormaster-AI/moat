import React, { Component } from 'react'
import OrderService from '../services/OrderService';

class UpdateOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                orderNumber: '',
                orderDate: '',
                totalAmount: '',
                taxAmount: '',
                shippingAmount: '',
                status: ''
        }
        this.updateOrder = this.updateOrder.bind(this);

        this.changeorderNumberHandler = this.changeorderNumberHandler.bind(this);
        this.changeorderDateHandler = this.changeorderDateHandler.bind(this);
        this.changetotalAmountHandler = this.changetotalAmountHandler.bind(this);
        this.changetaxAmountHandler = this.changetaxAmountHandler.bind(this);
        this.changeshippingAmountHandler = this.changeshippingAmountHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        OrderService.getOrderById(this.state.id).then( (res) =>{
            let order = res.data;
            this.setState({
                orderNumber: order.orderNumber,
                orderDate: order.orderDate,
                totalAmount: order.totalAmount,
                taxAmount: order.taxAmount,
                shippingAmount: order.shippingAmount,
                status: order.status
            });
        });
    }

    updateOrder = (e) => {
        e.preventDefault();
        let order = {
            orderId: this.state.id,
            orderNumber: this.state.orderNumber,
            orderDate: this.state.orderDate,
            totalAmount: this.state.totalAmount,
            taxAmount: this.state.taxAmount,
            shippingAmount: this.state.shippingAmount,
            status: this.state.status
        };
        console.log('order => ' + JSON.stringify(order));
        console.log('id => ' + JSON.stringify(this.state.id));
        OrderService.updateOrder(order).then( res => {
            this.props.history.push('/orders');
        });
    }

    changeorderNumberHandler= (event) => {
        this.setState({orderNumber: event.target.value});
    }
    changeorderDateHandler= (event) => {
        this.setState({orderDate: event.target.value});
    }
    changetotalAmountHandler= (event) => {
        this.setState({totalAmount: event.target.value});
    }
    changetaxAmountHandler= (event) => {
        this.setState({taxAmount: event.target.value});
    }
    changeshippingAmountHandler= (event) => {
        this.setState({shippingAmount: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/orders');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Order</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> orderNumber: </label>
                                                <input placeholder="orderNumber" name="orderNumber" className="form-control" value={this.state.orderNumber} onChange={this.changeorderNumberHandler}/>

                                            <label> orderDate: </label>
                                                <input type="date" placeholder="orderDate" name="orderDate" className="form-control" value={this.state.orderDate} onChange={this.changeorderDateHandler}/>

                                            <label> totalAmount: </label>
                                                <input placeholder="totalAmount" name="totalAmount" className="form-control" value={this.state.totalAmount} onChange={this.changetotalAmountHandler}/>

                                            <label> taxAmount: </label>
                                                <input placeholder="taxAmount" name="taxAmount" className="form-control" value={this.state.taxAmount} onChange={this.changetaxAmountHandler}/>

                                            <label> shippingAmount: </label>
                                                <input placeholder="shippingAmount" name="shippingAmount" className="form-control" value={this.state.shippingAmount} onChange={this.changeshippingAmountHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Submitted
                      </option>
                      <option name="Status" className="form-control" >
                          PartiallyFulfilled
                      </option>
                      <option name="Status" className="form-control" >
                          Fulfilled
                      </option>
                      <option name="Status" className="form-control" >
                          Invoiced
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateOrder}>Save</button>
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

export default UpdateOrderComponent
