import React, { Component } from 'react'
import MaintenanceWorkOrderService from '../services/MaintenanceWorkOrderService'

class ListMaintenanceWorkOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                maintenanceWorkOrders: []
        }
        this.addMaintenanceWorkOrder = this.addMaintenanceWorkOrder.bind(this);
        this.editMaintenanceWorkOrder = this.editMaintenanceWorkOrder.bind(this);
        this.deleteMaintenanceWorkOrder = this.deleteMaintenanceWorkOrder.bind(this);
    }

    deleteMaintenanceWorkOrder(id){
        MaintenanceWorkOrderService.deleteMaintenanceWorkOrder(id).then( res => {
            this.setState({maintenanceWorkOrders: this.state.maintenanceWorkOrders.filter(maintenanceWorkOrder => maintenanceWorkOrder.maintenanceWorkOrderId !== id)});
        });
    }
    viewMaintenanceWorkOrder(id){
        this.props.history.push(`/view-maintenanceWorkOrder/${id}`);
    }
    editMaintenanceWorkOrder(id){
        this.props.history.push(`/add-maintenanceWorkOrder/${id}`);
    }

    componentDidMount(){
        MaintenanceWorkOrderService.getMaintenanceWorkOrders().then((res) => {
            this.setState({ maintenanceWorkOrders: res.data});
        });
    }

    addMaintenanceWorkOrder(){
        this.props.history.push('/add-maintenanceWorkOrder/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">MaintenanceWorkOrder List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addMaintenanceWorkOrder}> Add MaintenanceWorkOrder</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> WorkOrderNumber </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.maintenanceWorkOrders.map(
                                        maintenanceWorkOrder => 
                                        <tr key = {maintenanceWorkOrder.maintenanceWorkOrderId}>
                                             <td> { maintenanceWorkOrder.workOrderNumber } </td>
                                             <td> { maintenanceWorkOrder.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editMaintenanceWorkOrder(maintenanceWorkOrder.maintenanceWorkOrderId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteMaintenanceWorkOrder(maintenanceWorkOrder.maintenanceWorkOrderId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewMaintenanceWorkOrder(maintenanceWorkOrder.maintenanceWorkOrderId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListMaintenanceWorkOrderComponent
