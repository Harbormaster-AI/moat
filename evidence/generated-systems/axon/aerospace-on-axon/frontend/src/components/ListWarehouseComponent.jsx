import React, { Component } from 'react'
import WarehouseService from '../services/WarehouseService'

class ListWarehouseComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                warehouses: []
        }
        this.addWarehouse = this.addWarehouse.bind(this);
        this.editWarehouse = this.editWarehouse.bind(this);
        this.deleteWarehouse = this.deleteWarehouse.bind(this);
    }

    deleteWarehouse(id){
        WarehouseService.deleteWarehouse(id).then( res => {
            this.setState({warehouses: this.state.warehouses.filter(warehouse => warehouse.warehouseId !== id)});
        });
    }
    viewWarehouse(id){
        this.props.history.push(`/view-warehouse/${id}`);
    }
    editWarehouse(id){
        this.props.history.push(`/add-warehouse/${id}`);
    }

    componentDidMount(){
        WarehouseService.getWarehouses().then((res) => {
            this.setState({ warehouses: res.data});
        });
    }

    addWarehouse(){
        this.props.history.push('/add-warehouse/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Warehouse List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addWarehouse}> Add Warehouse</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.warehouses.map(
                                        warehouse => 
                                        <tr key = {warehouse.warehouseId}>
                                             <td> { warehouse.name } </td>
                                             <td>
                                                 <button onClick={ () => this.editWarehouse(warehouse.warehouseId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteWarehouse(warehouse.warehouseId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewWarehouse(warehouse.warehouseId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListWarehouseComponent
