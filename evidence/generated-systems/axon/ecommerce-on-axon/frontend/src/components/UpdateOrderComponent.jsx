import React, { Component } from 'react'
import OrderService from '../services/OrderService';

class UpdateOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                orderNumber: '',
                placedDate: '',
                subtotal: '',
                discountTotal: '',
                shippingTotal: '',
                taxTotal: '',
                grandTotal: '',
                shippingAddress: '',
                billingAddress: '',
                status: ''
        }
        this.updateOrder = this.updateOrder.bind(this);

        this.changeorderNumberHandler = this.changeorderNumberHandler.bind(this);
        this.changeplacedDateHandler = this.changeplacedDateHandler.bind(this);
        this.changesubtotalHandler = this.changesubtotalHandler.bind(this);
        this.changediscountTotalHandler = this.changediscountTotalHandler.bind(this);
        this.changeshippingTotalHandler = this.changeshippingTotalHandler.bind(this);
        this.changetaxTotalHandler = this.changetaxTotalHandler.bind(this);
        this.changegrandTotalHandler = this.changegrandTotalHandler.bind(this);
        this.changeshippingAddressHandler = this.changeshippingAddressHandler.bind(this);
        this.changebillingAddressHandler = this.changebillingAddressHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        OrderService.getOrderById(this.state.id).then( (res) =>{
            let order = res.data;
            this.setState({
                orderNumber: order.orderNumber,
                placedDate: order.placedDate,
                subtotal: order.subtotal,
                discountTotal: order.discountTotal,
                shippingTotal: order.shippingTotal,
                taxTotal: order.taxTotal,
                grandTotal: order.grandTotal,
                shippingAddress: order.shippingAddress,
                billingAddress: order.billingAddress,
                status: order.status
            });
        });
    }

    updateOrder = (e) => {
        e.preventDefault();
        let order = {
            orderId: this.state.id,
            orderNumber: this.state.orderNumber,
            placedDate: this.state.placedDate,
            subtotal: this.state.subtotal,
            discountTotal: this.state.discountTotal,
            shippingTotal: this.state.shippingTotal,
            taxTotal: this.state.taxTotal,
            grandTotal: this.state.grandTotal,
            shippingAddress: this.state.shippingAddress,
            billingAddress: this.state.billingAddress,
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
    changeplacedDateHandler= (event) => {
        this.setState({placedDate: event.target.value});
    }
    changesubtotalHandler= (event) => {
        this.setState({subtotal: event.target.value});
    }
    changediscountTotalHandler= (event) => {
        this.setState({discountTotal: event.target.value});
    }
    changeshippingTotalHandler= (event) => {
        this.setState({shippingTotal: event.target.value});
    }
    changetaxTotalHandler= (event) => {
        this.setState({taxTotal: event.target.value});
    }
    changegrandTotalHandler= (event) => {
        this.setState({grandTotal: event.target.value});
    }
    changeshippingAddressHandler= (event) => {
        this.setState({shippingAddress: event.target.value});
    }
    changebillingAddressHandler= (event) => {
        this.setState({billingAddress: event.target.value});
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

                                            <label> placedDate: </label>
                                                <input type="date" placeholder="placedDate" name="placedDate" className="form-control" value={this.state.placedDate} onChange={this.changeplacedDateHandler}/>

                                            <label> subtotal: </label>
                                                <input placeholder="subtotal" name="subtotal" className="form-control" value={this.state.subtotal} onChange={this.changesubtotalHandler}/>

                                            <label> discountTotal: </label>
                                                <input placeholder="discountTotal" name="discountTotal" className="form-control" value={this.state.discountTotal} onChange={this.changediscountTotalHandler}/>

                                            <label> shippingTotal: </label>
                                                <input placeholder="shippingTotal" name="shippingTotal" className="form-control" value={this.state.shippingTotal} onChange={this.changeshippingTotalHandler}/>

                                            <label> taxTotal: </label>
                                                <input placeholder="taxTotal" name="taxTotal" className="form-control" value={this.state.taxTotal} onChange={this.changetaxTotalHandler}/>

                                            <label> grandTotal: </label>
                                                <input placeholder="grandTotal" name="grandTotal" className="form-control" value={this.state.grandTotal} onChange={this.changegrandTotalHandler}/>

                                            <label> shippingAddress: </label>
                                                <input placeholder="shippingAddress" name="shippingAddress" className="form-control" value={this.state.shippingAddress} onChange={this.changeshippingAddressHandler}/>

                                            <label> billingAddress: </label>
                                                <input placeholder="billingAddress" name="billingAddress" className="form-control" value={this.state.billingAddress} onChange={this.changebillingAddressHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Pending
                      </option>
                      <option name="Status" className="form-control" >
                          Confirmed
                      </option>
                      <option name="Status" className="form-control" >
                          Paid
                      </option>
                      <option name="Status" className="form-control" >
                          PartiallyShipped
                      </option>
                      <option name="Status" className="form-control" >
                          Shipped
                      </option>
                      <option name="Status" className="form-control" >
                          Delivered
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                      <option name="Status" className="form-control" >
                          Refunded
                      </option>
                      <option name="Status" className="form-control" >
                          PartiallyRefunded
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
