import React, { Component } from 'react'
import InventoryItemService from '../services/InventoryItemService';

class UpdateInventoryItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                quantityOnHand: '',
                quantityReserved: '',
                safetyStock: '',
                status: ''
        }
        this.updateInventoryItem = this.updateInventoryItem.bind(this);

        this.changequantityOnHandHandler = this.changequantityOnHandHandler.bind(this);
        this.changequantityReservedHandler = this.changequantityReservedHandler.bind(this);
        this.changesafetyStockHandler = this.changesafetyStockHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        InventoryItemService.getInventoryItemById(this.state.id).then( (res) =>{
            let inventoryItem = res.data;
            this.setState({
                quantityOnHand: inventoryItem.quantityOnHand,
                quantityReserved: inventoryItem.quantityReserved,
                safetyStock: inventoryItem.safetyStock,
                status: inventoryItem.status
            });
        });
    }

    updateInventoryItem = (e) => {
        e.preventDefault();
        let inventoryItem = {
            inventoryItemId: this.state.id,
            quantityOnHand: this.state.quantityOnHand,
            quantityReserved: this.state.quantityReserved,
            safetyStock: this.state.safetyStock,
            status: this.state.status
        };
        console.log('inventoryItem => ' + JSON.stringify(inventoryItem));
        console.log('id => ' + JSON.stringify(this.state.id));
        InventoryItemService.updateInventoryItem(inventoryItem).then( res => {
            this.props.history.push('/inventoryItems');
        });
    }

    changequantityOnHandHandler= (event) => {
        this.setState({quantityOnHand: event.target.value});
    }
    changequantityReservedHandler= (event) => {
        this.setState({quantityReserved: event.target.value});
    }
    changesafetyStockHandler= (event) => {
        this.setState({safetyStock: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/inventoryItems');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update InventoryItem</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> quantityOnHand: </label>
                                                <input type="number" placeholder="quantityOnHand" name="quantityOnHand" className="form-control" value={this.state.quantityOnHand} onChange={this.changequantityOnHandHandler}/>

                                            <label> quantityReserved: </label>
                                                <input type="number" placeholder="quantityReserved" name="quantityReserved" className="form-control" value={this.state.quantityReserved} onChange={this.changequantityReservedHandler}/>

                                            <label> safetyStock: </label>
                                                <input type="number" placeholder="safetyStock" name="safetyStock" className="form-control" value={this.state.safetyStock} onChange={this.changesafetyStockHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          InStock
                      </option>
                      <option name="Status" className="form-control" >
                          LowStock
                      </option>
                      <option name="Status" className="form-control" >
                          OutOfStock
                      </option>
                      <option name="Status" className="form-control" >
                          Backorder
                      </option>
                      <option name="Status" className="form-control" >
                          Preorder
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateInventoryItem}>Save</button>
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

export default UpdateInventoryItemComponent
