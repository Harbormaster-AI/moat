import React, { Component } from 'react'
import WarehouseService from '../services/WarehouseService';

class CreateWarehouseComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                warehouseCode: '',
                address: '',
                warehouseType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changewarehouseCodeHandler = this.changewarehouseCodeHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
        this.changeWarehouseTypeHandler = this.changeWarehouseTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            WarehouseService.getWarehouseById(this.state.id).then( (res) =>{
                let warehouse = res.data;
                this.setState({
                    name: warehouse.name,
                    warehouseCode: warehouse.warehouseCode,
                    address: warehouse.address,
                    warehouseType: warehouse.warehouseType
                });
            });
        }        
    }
    saveOrUpdateWarehouse = (e) => {
        e.preventDefault();
        let warehouse = {
                warehouseId: this.state.id,
                name: this.state.name,
                warehouseCode: this.state.warehouseCode,
                address: this.state.address,
                warehouseType: this.state.warehouseType
            };
        console.log('warehouse => ' + JSON.stringify(warehouse));

        // step 5
        if(this.state.id === '_add'){
            warehouse.warehouseId=''
            WarehouseService.createWarehouse(warehouse).then(res =>{
                this.props.history.push('/warehouses');
            });
        }else{
            WarehouseService.updateWarehouse(warehouse).then( res => {
                this.props.history.push('/warehouses');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changewarehouseCodeHandler= (event) => {
        this.setState({warehouseCode: event.target.value});
    }
    changeaddressHandler= (event) => {
        this.setState({address: event.target.value});
    }
    changeWarehouseTypeHandler= (event) => {
        this.setState({warehouseType: event.target.value});
    }

    cancel(){
        this.props.history.push('/warehouses');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Warehouse</h3>
        }else{
            return <h3 className="text-center">Update Warehouse</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> warehouseCode:&emsp; </label>
                                                <input placeholder="warehouseCode" name="warehouseCode" className="form-control" value={this.state.warehouseCode} onChange={this.changewarehouseCodeHandler}/>

                                            <label> address:&emsp; </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                            <label> WarehouseType:&emsp; </label>
                                                <select value={this.state.warehouseType} onChange={this.changeWarehouseTypeHandler}>
                      <option name="WarehouseType" className="form-control" >
                          RawMaterial
                      </option>
                      <option name="WarehouseType" className="form-control" >
                          WIP
                      </option>
                      <option name="WarehouseType" className="form-control" >
                          FinishedGoods
                      </option>
                      <option name="WarehouseType" className="form-control" >
                          Distribution
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateWarehouse}>Save</button>
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

export default CreateWarehouseComponent
