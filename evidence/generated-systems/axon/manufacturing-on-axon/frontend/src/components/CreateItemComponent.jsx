import React, { Component } from 'react'
import ItemService from '../services/ItemService';

class CreateItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                itemNumber: '',
                name: '',
                standardCost: '',
                weight: '',
                asSerialControlled: '',
                itemType: '',
                procurementType: '',
                unitOfMeasure: '',
                lifecycleStatus: ''
        }
        this.changeitemNumberHandler = this.changeitemNumberHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changestandardCostHandler = this.changestandardCostHandler.bind(this);
        this.changeweightHandler = this.changeweightHandler.bind(this);
        this.changeasSerialControlledHandler = this.changeasSerialControlledHandler.bind(this);
        this.changeItemTypeHandler = this.changeItemTypeHandler.bind(this);
        this.changeProcurementTypeHandler = this.changeProcurementTypeHandler.bind(this);
        this.changeUnitOfMeasureHandler = this.changeUnitOfMeasureHandler.bind(this);
        this.changeLifecycleStatusHandler = this.changeLifecycleStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ItemService.getItemById(this.state.id).then( (res) =>{
                let item = res.data;
                this.setState({
                    itemNumber: item.itemNumber,
                    name: item.name,
                    standardCost: item.standardCost,
                    weight: item.weight,
                    asSerialControlled: item.asSerialControlled,
                    itemType: item.itemType,
                    procurementType: item.procurementType,
                    unitOfMeasure: item.unitOfMeasure,
                    lifecycleStatus: item.lifecycleStatus
                });
            });
        }        
    }
    saveOrUpdateItem = (e) => {
        e.preventDefault();
        let item = {
                itemId: this.state.id,
                itemNumber: this.state.itemNumber,
                name: this.state.name,
                standardCost: this.state.standardCost,
                weight: this.state.weight,
                asSerialControlled: this.state.asSerialControlled,
                itemType: this.state.itemType,
                procurementType: this.state.procurementType,
                unitOfMeasure: this.state.unitOfMeasure,
                lifecycleStatus: this.state.lifecycleStatus
            };
        console.log('item => ' + JSON.stringify(item));

        // step 5
        if(this.state.id === '_add'){
            item.itemId=''
            ItemService.createItem(item).then(res =>{
                this.props.history.push('/items');
            });
        }else{
            ItemService.updateItem(item).then( res => {
                this.props.history.push('/items');
            });
        }
    }
    
    changeitemNumberHandler= (event) => {
        this.setState({itemNumber: event.target.value});
    }
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changestandardCostHandler= (event) => {
        this.setState({standardCost: event.target.value});
    }
    changeweightHandler= (event) => {
        this.setState({weight: event.target.value});
    }
    changeasSerialControlledHandler= (event) => {
        this.setState({asSerialControlled: event.target.value});
    }
    changeItemTypeHandler= (event) => {
        this.setState({itemType: event.target.value});
    }
    changeProcurementTypeHandler= (event) => {
        this.setState({procurementType: event.target.value});
    }
    changeUnitOfMeasureHandler= (event) => {
        this.setState({unitOfMeasure: event.target.value});
    }
    changeLifecycleStatusHandler= (event) => {
        this.setState({lifecycleStatus: event.target.value});
    }

    cancel(){
        this.props.history.push('/items');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Item</h3>
        }else{
            return <h3 className="text-center">Update Item</h3>
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
                                            <label> itemNumber:&emsp; </label>
                                                <input placeholder="itemNumber" name="itemNumber" className="form-control" value={this.state.itemNumber} onChange={this.changeitemNumberHandler}/>

                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> standardCost:&emsp; </label>
                                                <input placeholder="standardCost" name="standardCost" className="form-control" value={this.state.standardCost} onChange={this.changestandardCostHandler}/>

                                            <label> weight:&emsp; </label>
                                                <input placeholder="weight" name="weight" className="form-control" value={this.state.weight} onChange={this.changeweightHandler}/>

                                            <label> asSerialControlled:&emsp; </label>
                                                <input type="checkbox" placeholder="asSerialControlled" name="asSerialControlled" className="form-control" value={this.state.asSerialControlled} onChange={this.changeasSerialControlledHandler}/>


                                            <label> ItemType:&emsp; </label>
                                                <select value={this.state.itemType} onChange={this.changeItemTypeHandler}>
                      <option name="ItemType" className="form-control" >
                          FinishedGood
                      </option>
                      <option name="ItemType" className="form-control" >
                          Subassembly
                      </option>
                      <option name="ItemType" className="form-control" >
                          Component
                      </option>
                      <option name="ItemType" className="form-control" >
                          RawMaterial
                      </option>
                      <option name="ItemType" className="form-control" >
                          Consumable
                      </option>
                      <option name="ItemType" className="form-control" >
                          Service
                      </option>
                    </select>

                                            <label> ProcurementType:&emsp; </label>
                                                <select value={this.state.procurementType} onChange={this.changeProcurementTypeHandler}>
                      <option name="ProcurementType" className="form-control" >
                          MakeToStock
                      </option>
                      <option name="ProcurementType" className="form-control" >
                          MakeToOrder
                      </option>
                      <option name="ProcurementType" className="form-control" >
                          Purchase
                      </option>
                      <option name="ProcurementType" className="form-control" >
                          Kanban
                      </option>
                      <option name="ProcurementType" className="form-control" >
                          Outsourced
                      </option>
                    </select>

                                            <label> UnitOfMeasure:&emsp; </label>
                                                <select value={this.state.unitOfMeasure} onChange={this.changeUnitOfMeasureHandler}>
                      <option name="UnitOfMeasure" className="form-control" >
                          Each
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Kilogram
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Gram
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Pound
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Liter
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Meter
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Centimeter
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Millimeter
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Hour
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Minute
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Box
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Pallet
                      </option>
                    </select>

                                            <label> LifecycleStatus:&emsp; </label>
                                                <select value={this.state.lifecycleStatus} onChange={this.changeLifecycleStatusHandler}>
                      <option name="LifecycleStatus" className="form-control" >
                          Active
                      </option>
                      <option name="LifecycleStatus" className="form-control" >
                          PendingApproval
                      </option>
                      <option name="LifecycleStatus" className="form-control" >
                          Discontinued
                      </option>
                      <option name="LifecycleStatus" className="form-control" >
                          Obsolete
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateItem}>Save</button>
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

export default CreateItemComponent
