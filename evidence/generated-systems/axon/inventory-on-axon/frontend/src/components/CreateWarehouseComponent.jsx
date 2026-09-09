import React, { Component } from 'react'
import WarehouseService from '../services/WarehouseService';

class CreateWarehouseComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                code: '',
                address: '',
                timeZone: '',
                allowsOverAllocation: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
        this.changetimeZoneHandler = this.changetimeZoneHandler.bind(this);
        this.changeallowsOverAllocationHandler = this.changeallowsOverAllocationHandler.bind(this);
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
                    code: warehouse.code,
                    address: warehouse.address,
                    timeZone: warehouse.timeZone,
                    allowsOverAllocation: warehouse.allowsOverAllocation
                });
            });
        }        
    }
    saveOrUpdateWarehouse = (e) => {
        e.preventDefault();
        let warehouse = {
                warehouseId: this.state.id,
                name: this.state.name,
                code: this.state.code,
                address: this.state.address,
                timeZone: this.state.timeZone,
                allowsOverAllocation: this.state.allowsOverAllocation
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
    changecodeHandler= (event) => {
        this.setState({code: event.target.value});
    }
    changeaddressHandler= (event) => {
        this.setState({address: event.target.value});
    }
    changetimeZoneHandler= (event) => {
        this.setState({timeZone: event.target.value});
    }
    changeallowsOverAllocationHandler= (event) => {
        this.setState({allowsOverAllocation: event.target.value});
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

                                            <label> code:&emsp; </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> address:&emsp; </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                            <label> timeZone:&emsp; </label>
                                                <input placeholder="timeZone" name="timeZone" className="form-control" value={this.state.timeZone} onChange={this.changetimeZoneHandler}/>

                                            <label> allowsOverAllocation:&emsp; </label>
                                                <input type="checkbox" placeholder="allowsOverAllocation" name="allowsOverAllocation" className="form-control" value={this.state.allowsOverAllocation} onChange={this.changeallowsOverAllocationHandler}/>


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
