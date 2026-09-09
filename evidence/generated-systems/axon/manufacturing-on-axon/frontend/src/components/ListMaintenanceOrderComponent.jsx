import React, { Component } from 'react'
import MaintenanceOrderService from '../services/MaintenanceOrderService'

class ListMaintenanceOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                maintenanceOrders: []
        }
        this.addMaintenanceOrder = this.addMaintenanceOrder.bind(this);
        this.editMaintenanceOrder = this.editMaintenanceOrder.bind(this);
        this.deleteMaintenanceOrder = this.deleteMaintenanceOrder.bind(this);
    }

    deleteMaintenanceOrder(id){
        MaintenanceOrderService.deleteMaintenanceOrder(id).then( res => {
            this.setState({maintenanceOrders: this.state.maintenanceOrders.filter(maintenanceOrder => maintenanceOrder.maintenanceOrderId !== id)});
        });
    }
    viewMaintenanceOrder(id){
        this.props.history.push(`/view-maintenanceOrder/${id}`);
    }
    editMaintenanceOrder(id){
        this.props.history.push(`/add-maintenanceOrder/${id}`);
    }

    componentDidMount(){
        MaintenanceOrderService.getMaintenanceOrders().then((res) => {
            this.setState({ maintenanceOrders: res.data});
        });
    }

    addMaintenanceOrder(){
        this.props.history.push('/add-maintenanceOrder/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">MaintenanceOrder List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addMaintenanceOrder}> Add MaintenanceOrder</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> OrderNumber </th>
                                    <th> Priority </th>
                                    <th> RequestedDate </th>
                                    <th> CompletionDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.maintenanceOrders.map(
                                        maintenanceOrder => 
                                        <tr key = {maintenanceOrder.maintenanceOrderId}>
                                             <td> { maintenanceOrder.orderNumber } </td>
                                             <td> { maintenanceOrder.priority } </td>
                                             <td> { maintenanceOrder.requestedDate } </td>
                                             <td> { maintenanceOrder.completionDate } </td>
                                             <td> { maintenanceOrder.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editMaintenanceOrder(maintenanceOrder.maintenanceOrderId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteMaintenanceOrder(maintenanceOrder.maintenanceOrderId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewMaintenanceOrder(maintenanceOrder.maintenanceOrderId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListMaintenanceOrderComponent
