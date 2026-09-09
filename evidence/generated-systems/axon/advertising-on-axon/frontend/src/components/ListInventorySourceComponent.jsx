import React, { Component } from 'react'
import InventorySourceService from '../services/InventorySourceService'

class ListInventorySourceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                inventorySources: []
        }
        this.addInventorySource = this.addInventorySource.bind(this);
        this.editInventorySource = this.editInventorySource.bind(this);
        this.deleteInventorySource = this.deleteInventorySource.bind(this);
    }

    deleteInventorySource(id){
        InventorySourceService.deleteInventorySource(id).then( res => {
            this.setState({inventorySources: this.state.inventorySources.filter(inventorySource => inventorySource.inventorySourceId !== id)});
        });
    }
    viewInventorySource(id){
        this.props.history.push(`/view-inventorySource/${id}`);
    }
    editInventorySource(id){
        this.props.history.push(`/add-inventorySource/${id}`);
    }

    componentDidMount(){
        InventorySourceService.getInventorySources().then((res) => {
            this.setState({ inventorySources: res.data});
        });
    }

    addInventorySource(){
        this.props.history.push('/add-inventorySource/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">InventorySource List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addInventorySource}> Add InventorySource</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Domain </th>
                                    <th> Channel </th>
                                    <th> PrimaryFormat </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.inventorySources.map(
                                        inventorySource => 
                                        <tr key = {inventorySource.inventorySourceId}>
                                             <td> { inventorySource.name } </td>
                                             <td> { inventorySource.domain } </td>
                                             <td> { inventorySource.channel } </td>
                                             <td> { inventorySource.primaryFormat } </td>
                                             <td>
                                                 <button onClick={ () => this.editInventorySource(inventorySource.inventorySourceId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteInventorySource(inventorySource.inventorySourceId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewInventorySource(inventorySource.inventorySourceId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListInventorySourceComponent
