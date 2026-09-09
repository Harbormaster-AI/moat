import React, { Component } from 'react'
import InventoryItemService from '../services/InventoryItemService';

class CreateInventoryItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                quantityOnHand: '',
                quantityAvailable: '',
                quantityReserved: '',
                unitCost: '',
                lastUpdated: '',
                stockStatus: ''
        }
        this.changequantityOnHandHandler = this.changequantityOnHandHandler.bind(this);
        this.changequantityAvailableHandler = this.changequantityAvailableHandler.bind(this);
        this.changequantityReservedHandler = this.changequantityReservedHandler.bind(this);
        this.changeunitCostHandler = this.changeunitCostHandler.bind(this);
        this.changelastUpdatedHandler = this.changelastUpdatedHandler.bind(this);
        this.changeStockStatusHandler = this.changeStockStatusHandler.bind(this);
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
                    quantityAvailable: inventoryItem.quantityAvailable,
                    quantityReserved: inventoryItem.quantityReserved,
                    unitCost: inventoryItem.unitCost,
                    lastUpdated: inventoryItem.lastUpdated,
                    stockStatus: inventoryItem.stockStatus
                });
            });
        }        
    }
    saveOrUpdateInventoryItem = (e) => {
        e.preventDefault();
        let inventoryItem = {
                inventoryItemId: this.state.id,
                quantityOnHand: this.state.quantityOnHand,
                quantityAvailable: this.state.quantityAvailable,
                quantityReserved: this.state.quantityReserved,
                unitCost: this.state.unitCost,
                lastUpdated: this.state.lastUpdated,
                stockStatus: this.state.stockStatus
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
    changequantityAvailableHandler= (event) => {
        this.setState({quantityAvailable: event.target.value});
    }
    changequantityReservedHandler= (event) => {
        this.setState({quantityReserved: event.target.value});
    }
    changeunitCostHandler= (event) => {
        this.setState({unitCost: event.target.value});
    }
    changelastUpdatedHandler= (event) => {
        this.setState({lastUpdated: event.target.value});
    }
    changeStockStatusHandler= (event) => {
        this.setState({stockStatus: event.target.value});
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
                                                <input placeholder="quantityOnHand" name="quantityOnHand" className="form-control" value={this.state.quantityOnHand} onChange={this.changequantityOnHandHandler}/>

                                            <label> quantityAvailable:&emsp; </label>
                                                <input placeholder="quantityAvailable" name="quantityAvailable" className="form-control" value={this.state.quantityAvailable} onChange={this.changequantityAvailableHandler}/>

                                            <label> quantityReserved:&emsp; </label>
                                                <input placeholder="quantityReserved" name="quantityReserved" className="form-control" value={this.state.quantityReserved} onChange={this.changequantityReservedHandler}/>

                                            <label> unitCost:&emsp; </label>
                                                <input placeholder="unitCost" name="unitCost" className="form-control" value={this.state.unitCost} onChange={this.changeunitCostHandler}/>

                                            <label> lastUpdated:&emsp; </label>
                                                <input type="date" placeholder="lastUpdated" name="lastUpdated" className="form-control" value={this.state.lastUpdated} onChange={this.changelastUpdatedHandler}/>

                                            <label> StockStatus:&emsp; </label>
                                                <select value={this.state.stockStatus} onChange={this.changeStockStatusHandler}>
                      <option name="StockStatus" className="form-control" >
                          Available
                      </option>
                      <option name="StockStatus" className="form-control" >
                          Reserved
                      </option>
                      <option name="StockStatus" className="form-control" >
                          Damaged
                      </option>
                      <option name="StockStatus" className="form-control" >
                          Hold
                      </option>
                      <option name="StockStatus" className="form-control" >
                          Quarantined
                      </option>
                      <option name="StockStatus" className="form-control" >
                          InTransit
                      </option>
                      <option name="StockStatus" className="form-control" >
                          PendingInspection
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
