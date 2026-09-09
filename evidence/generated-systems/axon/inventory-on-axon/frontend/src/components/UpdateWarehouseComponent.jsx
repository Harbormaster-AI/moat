import React, { Component } from 'react'
import WarehouseService from '../services/WarehouseService';

class UpdateWarehouseComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                code: '',
                address: '',
                timeZone: '',
                allowsOverAllocation: ''
        }
        this.updateWarehouse = this.updateWarehouse.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
        this.changetimeZoneHandler = this.changetimeZoneHandler.bind(this);
        this.changeallowsOverAllocationHandler = this.changeallowsOverAllocationHandler.bind(this);
    }

    componentDidMount(){
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

    updateWarehouse = (e) => {
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
        console.log('id => ' + JSON.stringify(this.state.id));
        WarehouseService.updateWarehouse(warehouse).then( res => {
            this.props.history.push('/warehouses');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Warehouse</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> code: </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> address: </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                            <label> timeZone: </label>
                                                <input placeholder="timeZone" name="timeZone" className="form-control" value={this.state.timeZone} onChange={this.changetimeZoneHandler}/>

                                            <label> allowsOverAllocation: </label>
                                                <input type="checkbox" placeholder="allowsOverAllocation" name="allowsOverAllocation" className="form-control" value={this.state.allowsOverAllocation} onChange={this.changeallowsOverAllocationHandler}/>


                                        </div>
                                        <button className="btn btn-success" onClick={this.updateWarehouse}>Save</button>
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

export default UpdateWarehouseComponent
