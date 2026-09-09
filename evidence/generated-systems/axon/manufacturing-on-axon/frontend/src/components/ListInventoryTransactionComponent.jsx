import React, { Component } from 'react'
import InventoryTransactionService from '../services/InventoryTransactionService'

class ListInventoryTransactionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                inventoryTransactions: []
        }
        this.addInventoryTransaction = this.addInventoryTransaction.bind(this);
        this.editInventoryTransaction = this.editInventoryTransaction.bind(this);
        this.deleteInventoryTransaction = this.deleteInventoryTransaction.bind(this);
    }

    deleteInventoryTransaction(id){
        InventoryTransactionService.deleteInventoryTransaction(id).then( res => {
            this.setState({inventoryTransactions: this.state.inventoryTransactions.filter(inventoryTransaction => inventoryTransaction.inventoryTransactionId !== id)});
        });
    }
    viewInventoryTransaction(id){
        this.props.history.push(`/view-inventoryTransaction/${id}`);
    }
    editInventoryTransaction(id){
        this.props.history.push(`/add-inventoryTransaction/${id}`);
    }

    componentDidMount(){
        InventoryTransactionService.getInventoryTransactions().then((res) => {
            this.setState({ inventoryTransactions: res.data});
        });
    }

    addInventoryTransaction(){
        this.props.history.push('/add-inventoryTransaction/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">InventoryTransaction List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addInventoryTransaction}> Add InventoryTransaction</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> TransactionNumber </th>
                                    <th> Quantity </th>
                                    <th> TransactionDateTime </th>
                                    <th> ReferenceDocument </th>
                                    <th> TransactionType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.inventoryTransactions.map(
                                        inventoryTransaction => 
                                        <tr key = {inventoryTransaction.inventoryTransactionId}>
                                             <td> { inventoryTransaction.transactionNumber } </td>
                                             <td> { inventoryTransaction.quantity } </td>
                                             <td> { inventoryTransaction.transactionDateTime } </td>
                                             <td> { inventoryTransaction.referenceDocument } </td>
                                             <td> { inventoryTransaction.transactionType } </td>
                                             <td>
                                                 <button onClick={ () => this.editInventoryTransaction(inventoryTransaction.inventoryTransactionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteInventoryTransaction(inventoryTransaction.inventoryTransactionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewInventoryTransaction(inventoryTransaction.inventoryTransactionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListInventoryTransactionComponent
