import React, { Component } from 'react'
import InventoryItemService from '../services/InventoryItemService'

class ListInventoryItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                inventoryItems: []
        }
        this.addInventoryItem = this.addInventoryItem.bind(this);
        this.editInventoryItem = this.editInventoryItem.bind(this);
        this.deleteInventoryItem = this.deleteInventoryItem.bind(this);
    }

    deleteInventoryItem(id){
        InventoryItemService.deleteInventoryItem(id).then( res => {
            this.setState({inventoryItems: this.state.inventoryItems.filter(inventoryItem => inventoryItem.inventoryItemId !== id)});
        });
    }
    viewInventoryItem(id){
        this.props.history.push(`/view-inventoryItem/${id}`);
    }
    editInventoryItem(id){
        this.props.history.push(`/add-inventoryItem/${id}`);
    }

    componentDidMount(){
        InventoryItemService.getInventoryItems().then((res) => {
            this.setState({ inventoryItems: res.data});
        });
    }

    addInventoryItem(){
        this.props.history.push('/add-inventoryItem/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">InventoryItem List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addInventoryItem}> Add InventoryItem</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> QuantityOnHand </th>
                                    <th> QuantityReserved </th>
                                    <th> LotNumber </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.inventoryItems.map(
                                        inventoryItem => 
                                        <tr key = {inventoryItem.inventoryItemId}>
                                             <td> { inventoryItem.quantityOnHand } </td>
                                             <td> { inventoryItem.quantityReserved } </td>
                                             <td> { inventoryItem.lotNumber } </td>
                                             <td>
                                                 <button onClick={ () => this.editInventoryItem(inventoryItem.inventoryItemId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteInventoryItem(inventoryItem.inventoryItemId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewInventoryItem(inventoryItem.inventoryItemId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListInventoryItemComponent
