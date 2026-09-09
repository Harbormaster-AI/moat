import React, { Component } from 'react'
import PaymentOrderService from '../services/PaymentOrderService'

class ListPaymentOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                paymentOrders: []
        }
        this.addPaymentOrder = this.addPaymentOrder.bind(this);
        this.editPaymentOrder = this.editPaymentOrder.bind(this);
        this.deletePaymentOrder = this.deletePaymentOrder.bind(this);
    }

    deletePaymentOrder(id){
        PaymentOrderService.deletePaymentOrder(id).then( res => {
            this.setState({paymentOrders: this.state.paymentOrders.filter(paymentOrder => paymentOrder.paymentOrderId !== id)});
        });
    }
    viewPaymentOrder(id){
        this.props.history.push(`/view-paymentOrder/${id}`);
    }
    editPaymentOrder(id){
        this.props.history.push(`/add-paymentOrder/${id}`);
    }

    componentDidMount(){
        PaymentOrderService.getPaymentOrders().then((res) => {
            this.setState({ paymentOrders: res.data});
        });
    }

    addPaymentOrder(){
        this.props.history.push('/add-paymentOrder/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">PaymentOrder List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPaymentOrder}> Add PaymentOrder</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> OrderReference </th>
                                    <th> RequestedExecutionDate </th>
                                    <th> Purpose </th>
                                    <th> PaymentMethod </th>
                                    <th> Status </th>
                                    <th> Priority </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.paymentOrders.map(
                                        paymentOrder => 
                                        <tr key = {paymentOrder.paymentOrderId}>
                                             <td> { paymentOrder.orderReference } </td>
                                             <td> { paymentOrder.requestedExecutionDate } </td>
                                             <td> { paymentOrder.purpose } </td>
                                             <td> { paymentOrder.paymentMethod } </td>
                                             <td> { paymentOrder.status } </td>
                                             <td> { paymentOrder.priority } </td>
                                             <td>
                                                 <button onClick={ () => this.editPaymentOrder(paymentOrder.paymentOrderId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePaymentOrder(paymentOrder.paymentOrderId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPaymentOrder(paymentOrder.paymentOrderId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPaymentOrderComponent
