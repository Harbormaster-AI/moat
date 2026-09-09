import React, { Component } from 'react'
import InventoryThresholdAlertService from '../services/InventoryThresholdAlertService'

class ListInventoryThresholdAlertComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                inventoryThresholdAlerts: []
        }
        this.addInventoryThresholdAlert = this.addInventoryThresholdAlert.bind(this);
        this.editInventoryThresholdAlert = this.editInventoryThresholdAlert.bind(this);
        this.deleteInventoryThresholdAlert = this.deleteInventoryThresholdAlert.bind(this);
    }

    deleteInventoryThresholdAlert(id){
        InventoryThresholdAlertService.deleteInventoryThresholdAlert(id).then( res => {
            this.setState({inventoryThresholdAlerts: this.state.inventoryThresholdAlerts.filter(inventoryThresholdAlert => inventoryThresholdAlert.inventoryThresholdAlertId !== id)});
        });
    }
    viewInventoryThresholdAlert(id){
        this.props.history.push(`/view-inventoryThresholdAlert/${id}`);
    }
    editInventoryThresholdAlert(id){
        this.props.history.push(`/add-inventoryThresholdAlert/${id}`);
    }

    componentDidMount(){
        InventoryThresholdAlertService.getInventoryThresholdAlerts().then((res) => {
            this.setState({ inventoryThresholdAlerts: res.data});
        });
    }

    addInventoryThresholdAlert(){
        this.props.history.push('/add-inventoryThresholdAlert/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">InventoryThresholdAlert List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addInventoryThresholdAlert}> Add InventoryThresholdAlert</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> AlertNumber </th>
                                    <th> DetectedAt </th>
                                    <th> Message </th>
                                    <th> AlertType </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.inventoryThresholdAlerts.map(
                                        inventoryThresholdAlert => 
                                        <tr key = {inventoryThresholdAlert.inventoryThresholdAlertId}>
                                             <td> { inventoryThresholdAlert.alertNumber } </td>
                                             <td> { inventoryThresholdAlert.detectedAt } </td>
                                             <td> { inventoryThresholdAlert.message } </td>
                                             <td> { inventoryThresholdAlert.alertType } </td>
                                             <td> { inventoryThresholdAlert.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editInventoryThresholdAlert(inventoryThresholdAlert.inventoryThresholdAlertId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteInventoryThresholdAlert(inventoryThresholdAlert.inventoryThresholdAlertId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewInventoryThresholdAlert(inventoryThresholdAlert.inventoryThresholdAlertId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListInventoryThresholdAlertComponent
