import React, { Component } from 'react'
import PurchaseOrderService from '../services/PurchaseOrderService'

class ListPurchaseOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                purchaseOrders: []
        }
        this.addPurchaseOrder = this.addPurchaseOrder.bind(this);
        this.editPurchaseOrder = this.editPurchaseOrder.bind(this);
        this.deletePurchaseOrder = this.deletePurchaseOrder.bind(this);
    }

    deletePurchaseOrder(id){
        PurchaseOrderService.deletePurchaseOrder(id).then( res => {
            this.setState({purchaseOrders: this.state.purchaseOrders.filter(purchaseOrder => purchaseOrder.purchaseOrderId !== id)});
        });
    }
    viewPurchaseOrder(id){
        this.props.history.push(`/view-purchaseOrder/${id}`);
    }
    editPurchaseOrder(id){
        this.props.history.push(`/add-purchaseOrder/${id}`);
    }

    componentDidMount(){
        PurchaseOrderService.getPurchaseOrders().then((res) => {
            this.setState({ purchaseOrders: res.data});
        });
    }

    addPurchaseOrder(){
        this.props.history.push('/add-purchaseOrder/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">PurchaseOrder List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPurchaseOrder}> Add PurchaseOrder</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> PoNumber </th>
                                    <th> OrderDate </th>
                                    <th> TotalAmount </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.purchaseOrders.map(
                                        purchaseOrder => 
                                        <tr key = {purchaseOrder.purchaseOrderId}>
                                             <td> { purchaseOrder.poNumber } </td>
                                             <td> { purchaseOrder.orderDate } </td>
                                             <td> { purchaseOrder.totalAmount } </td>
                                             <td> { purchaseOrder.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editPurchaseOrder(purchaseOrder.purchaseOrderId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePurchaseOrder(purchaseOrder.purchaseOrderId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPurchaseOrder(purchaseOrder.purchaseOrderId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPurchaseOrderComponent
