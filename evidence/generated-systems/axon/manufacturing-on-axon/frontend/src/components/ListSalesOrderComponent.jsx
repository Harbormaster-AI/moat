import React, { Component } from 'react'
import SalesOrderService from '../services/SalesOrderService'

class ListSalesOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                salesOrders: []
        }
        this.addSalesOrder = this.addSalesOrder.bind(this);
        this.editSalesOrder = this.editSalesOrder.bind(this);
        this.deleteSalesOrder = this.deleteSalesOrder.bind(this);
    }

    deleteSalesOrder(id){
        SalesOrderService.deleteSalesOrder(id).then( res => {
            this.setState({salesOrders: this.state.salesOrders.filter(salesOrder => salesOrder.salesOrderId !== id)});
        });
    }
    viewSalesOrder(id){
        this.props.history.push(`/view-salesOrder/${id}`);
    }
    editSalesOrder(id){
        this.props.history.push(`/add-salesOrder/${id}`);
    }

    componentDidMount(){
        SalesOrderService.getSalesOrders().then((res) => {
            this.setState({ salesOrders: res.data});
        });
    }

    addSalesOrder(){
        this.props.history.push('/add-salesOrder/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">SalesOrder List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addSalesOrder}> Add SalesOrder</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> OrderNumber </th>
                                    <th> OrderDate </th>
                                    <th> TotalAmount </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.salesOrders.map(
                                        salesOrder => 
                                        <tr key = {salesOrder.salesOrderId}>
                                             <td> { salesOrder.orderNumber } </td>
                                             <td> { salesOrder.orderDate } </td>
                                             <td> { salesOrder.totalAmount } </td>
                                             <td> { salesOrder.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editSalesOrder(salesOrder.salesOrderId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteSalesOrder(salesOrder.salesOrderId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewSalesOrder(salesOrder.salesOrderId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListSalesOrderComponent
