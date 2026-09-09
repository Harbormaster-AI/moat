import React, { Component } from 'react'
import TransferOrderService from '../services/TransferOrderService'

class ListTransferOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                transferOrders: []
        }
        this.addTransferOrder = this.addTransferOrder.bind(this);
        this.editTransferOrder = this.editTransferOrder.bind(this);
        this.deleteTransferOrder = this.deleteTransferOrder.bind(this);
    }

    deleteTransferOrder(id){
        TransferOrderService.deleteTransferOrder(id).then( res => {
            this.setState({transferOrders: this.state.transferOrders.filter(transferOrder => transferOrder.transferOrderId !== id)});
        });
    }
    viewTransferOrder(id){
        this.props.history.push(`/view-transferOrder/${id}`);
    }
    editTransferOrder(id){
        this.props.history.push(`/add-transferOrder/${id}`);
    }

    componentDidMount(){
        TransferOrderService.getTransferOrders().then((res) => {
            this.setState({ transferOrders: res.data});
        });
    }

    addTransferOrder(){
        this.props.history.push('/add-transferOrder/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">TransferOrder List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addTransferOrder}> Add TransferOrder</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> OrderNumber </th>
                                    <th> RequestedShipDate </th>
                                    <th> RequestedReceiveDate </th>
                                    <th> ShippedDate </th>
                                    <th> ReceivedDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.transferOrders.map(
                                        transferOrder => 
                                        <tr key = {transferOrder.transferOrderId}>
                                             <td> { transferOrder.orderNumber } </td>
                                             <td> { transferOrder.requestedShipDate } </td>
                                             <td> { transferOrder.requestedReceiveDate } </td>
                                             <td> { transferOrder.shippedDate } </td>
                                             <td> { transferOrder.receivedDate } </td>
                                             <td> { transferOrder.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editTransferOrder(transferOrder.transferOrderId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteTransferOrder(transferOrder.transferOrderId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewTransferOrder(transferOrder.transferOrderId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListTransferOrderComponent
