import React, { Component } from 'react'
import OrderItemService from '../services/OrderItemService'

class ListOrderItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                orderItems: []
        }
        this.addOrderItem = this.addOrderItem.bind(this);
        this.editOrderItem = this.editOrderItem.bind(this);
        this.deleteOrderItem = this.deleteOrderItem.bind(this);
    }

    deleteOrderItem(id){
        OrderItemService.deleteOrderItem(id).then( res => {
            this.setState({orderItems: this.state.orderItems.filter(orderItem => orderItem.orderItemId !== id)});
        });
    }
    viewOrderItem(id){
        this.props.history.push(`/view-orderItem/${id}`);
    }
    editOrderItem(id){
        this.props.history.push(`/add-orderItem/${id}`);
    }

    componentDidMount(){
        OrderItemService.getOrderItems().then((res) => {
            this.setState({ orderItems: res.data});
        });
    }

    addOrderItem(){
        this.props.history.push('/add-orderItem/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">OrderItem List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addOrderItem}> Add OrderItem</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Quantity </th>
                                    <th> UnitPrice </th>
                                    <th> DiscountAmount </th>
                                    <th> TaxAmount </th>
                                    <th> TotalAmount </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.orderItems.map(
                                        orderItem => 
                                        <tr key = {orderItem.orderItemId}>
                                             <td> { orderItem.quantity } </td>
                                             <td> { orderItem.unitPrice } </td>
                                             <td> { orderItem.discountAmount } </td>
                                             <td> { orderItem.taxAmount } </td>
                                             <td> { orderItem.totalAmount } </td>
                                             <td>
                                                 <button onClick={ () => this.editOrderItem(orderItem.orderItemId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteOrderItem(orderItem.orderItemId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewOrderItem(orderItem.orderItemId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListOrderItemComponent
