import React, { Component } from 'react'
import InventoryItemService from '../services/InventoryItemService';

class UpdateInventoryItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                quantityOnHand: '',
                quantityAvailable: '',
                quantityReserved: '',
                unitCost: '',
                lastUpdated: '',
                stockStatus: ''
        }
        this.updateInventoryItem = this.updateInventoryItem.bind(this);

        this.changequantityOnHandHandler = this.changequantityOnHandHandler.bind(this);
        this.changequantityAvailableHandler = this.changequantityAvailableHandler.bind(this);
        this.changequantityReservedHandler = this.changequantityReservedHandler.bind(this);
        this.changeunitCostHandler = this.changeunitCostHandler.bind(this);
        this.changelastUpdatedHandler = this.changelastUpdatedHandler.bind(this);
        this.changeStockStatusHandler = this.changeStockStatusHandler.bind(this);
    }

    componentDidMount(){
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

    updateInventoryItem = (e) => {
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
        console.log('id => ' + JSON.stringify(this.state.id));
        InventoryItemService.updateInventoryItem(inventoryItem).then( res => {
            this.props.history.push('/inventoryItems');
        });
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
                                                <input placeholder="quantityOnHand" name="quantityOnHand" className="form-control" value={this.state.quantityOnHand} onChange={this.changequantityOnHandHandler}/>

                                            <label> quantityAvailable: </label>
                                                <input placeholder="quantityAvailable" name="quantityAvailable" className="form-control" value={this.state.quantityAvailable} onChange={this.changequantityAvailableHandler}/>

                                            <label> quantityReserved: </label>
                                                <input placeholder="quantityReserved" name="quantityReserved" className="form-control" value={this.state.quantityReserved} onChange={this.changequantityReservedHandler}/>

                                            <label> unitCost: </label>
                                                <input placeholder="unitCost" name="unitCost" className="form-control" value={this.state.unitCost} onChange={this.changeunitCostHandler}/>

                                            <label> lastUpdated: </label>
                                                <input type="date" placeholder="lastUpdated" name="lastUpdated" className="form-control" value={this.state.lastUpdated} onChange={this.changelastUpdatedHandler}/>

                                            <label> StockStatus: </label>
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
