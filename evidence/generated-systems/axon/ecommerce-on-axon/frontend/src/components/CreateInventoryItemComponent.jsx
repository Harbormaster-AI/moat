import React, { Component } from 'react'
import InventoryItemService from '../services/InventoryItemService';

class CreateInventoryItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                quantityOnHand: '',
                quantityReserved: '',
                safetyStock: '',
                status: ''
        }
        this.changequantityOnHandHandler = this.changequantityOnHandHandler.bind(this);
        this.changequantityReservedHandler = this.changequantityReservedHandler.bind(this);
        this.changesafetyStockHandler = this.changesafetyStockHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateInventoryItem = (e) => {
        e.preventDefault();
        let inventoryItem = {
                inventoryItemId: this.state.id,
                quantityOnHand: this.state.quantityOnHand,
                quantityReserved: this.state.quantityReserved,
                safetyStock: this.state.safetyStock,
                status: this.state.status
            };
        console.log('inventoryItem => ' + JSON.stringify(inventoryItem));

        // step 5
        if(this.state.id === '_add'){
            inventoryItem.inventoryItemId=''
            InventoryItemService.createInventoryItem(inventoryItem).then(res =>{
                this.props.history.push('/inventoryItems');
            });
        }else{
            InventoryItemService.updateInventoryItem(inventoryItem).then( res => {
                this.props.history.push('/inventoryItems');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add InventoryItem</h3>
        }else{
            return <h3 className="text-center">Update InventoryItem</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> quantityOnHand:&emsp; </label>
                                                <input type="number" placeholder="quantityOnHand" name="quantityOnHand" className="form-control" value={this.state.quantityOnHand} onChange={this.changequantityOnHandHandler}/>

                                            <label> quantityReserved:&emsp; </label>
                                                <input type="number" placeholder="quantityReserved" name="quantityReserved" className="form-control" value={this.state.quantityReserved} onChange={this.changequantityReservedHandler}/>

                                            <label> safetyStock:&emsp; </label>
                                                <input type="number" placeholder="safetyStock" name="safetyStock" className="form-control" value={this.state.safetyStock} onChange={this.changesafetyStockHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateInventoryItem}>Save</button>
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

export default CreateInventoryItemComponent
