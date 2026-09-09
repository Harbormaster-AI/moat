import React, { Component } from 'react'
import WarehouseService from '../services/WarehouseService';

class UpdateWarehouseComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: ''
        }
        this.updateWarehouse = this.updateWarehouse.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
    }

    componentDidMount(){
        WarehouseService.getWarehouseById(this.state.id).then( (res) =>{
            let warehouse = res.data;
            this.setState({
                name: warehouse.name
            });
        });
    }

    updateWarehouse = (e) => {
        e.preventDefault();
        let warehouse = {
            warehouseId: this.state.id,
            name: this.state.name
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
