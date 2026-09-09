import React, { Component } from 'react'
import OrderService from '../services/OrderService'

class ListOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                orders: []
        }
        this.addOrder = this.addOrder.bind(this);
        this.editOrder = this.editOrder.bind(this);
        this.deleteOrder = this.deleteOrder.bind(this);
    }

    deleteOrder(id){
        OrderService.deleteOrder(id).then( res => {
            this.setState({orders: this.state.orders.filter(order => order.orderId !== id)});
        });
    }
    viewOrder(id){
        this.props.history.push(`/view-order/${id}`);
    }
    editOrder(id){
        this.props.history.push(`/add-order/${id}`);
    }

    componentDidMount(){
        OrderService.getOrders().then((res) => {
            this.setState({ orders: res.data});
        });
    }

    addOrder(){
        this.props.history.push('/add-order/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Order List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addOrder}> Add Order</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> OrderNumber </th>
                                    <th> PlacedDate </th>
                                    <th> Subtotal </th>
                                    <th> DiscountTotal </th>
                                    <th> ShippingTotal </th>
                                    <th> TaxTotal </th>
                                    <th> GrandTotal </th>
                                    <th> ShippingAddress </th>
                                    <th> BillingAddress </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.orders.map(
                                        order => 
                                        <tr key = {order.orderId}>
                                             <td> { order.orderNumber } </td>
                                             <td> { order.placedDate } </td>
                                             <td> { order.subtotal } </td>
                                             <td> { order.discountTotal } </td>
                                             <td> { order.shippingTotal } </td>
                                             <td> { order.taxTotal } </td>
                                             <td> { order.grandTotal } </td>
                                             <td> { order.shippingAddress } </td>
                                             <td> { order.billingAddress } </td>
                                             <td> { order.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editOrder(order.orderId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteOrder(order.orderId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewOrder(order.orderId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListOrderComponent
