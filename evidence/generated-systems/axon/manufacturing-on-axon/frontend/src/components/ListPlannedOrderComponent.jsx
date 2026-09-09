import React, { Component } from 'react'
import PlannedOrderService from '../services/PlannedOrderService'

class ListPlannedOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                plannedOrders: []
        }
        this.addPlannedOrder = this.addPlannedOrder.bind(this);
        this.editPlannedOrder = this.editPlannedOrder.bind(this);
        this.deletePlannedOrder = this.deletePlannedOrder.bind(this);
    }

    deletePlannedOrder(id){
        PlannedOrderService.deletePlannedOrder(id).then( res => {
            this.setState({plannedOrders: this.state.plannedOrders.filter(plannedOrder => plannedOrder.plannedOrderId !== id)});
        });
    }
    viewPlannedOrder(id){
        this.props.history.push(`/view-plannedOrder/${id}`);
    }
    editPlannedOrder(id){
        this.props.history.push(`/add-plannedOrder/${id}`);
    }

    componentDidMount(){
        PlannedOrderService.getPlannedOrders().then((res) => {
            this.setState({ plannedOrders: res.data});
        });
    }

    addPlannedOrder(){
        this.props.history.push('/add-plannedOrder/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">PlannedOrder List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPlannedOrder}> Add PlannedOrder</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> PlannedOrderNumber </th>
                                    <th> Quantity </th>
                                    <th> DueDate </th>
                                    <th> OrderType </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.plannedOrders.map(
                                        plannedOrder => 
                                        <tr key = {plannedOrder.plannedOrderId}>
                                             <td> { plannedOrder.plannedOrderNumber } </td>
                                             <td> { plannedOrder.quantity } </td>
                                             <td> { plannedOrder.dueDate } </td>
                                             <td> { plannedOrder.orderType } </td>
                                             <td> { plannedOrder.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editPlannedOrder(plannedOrder.plannedOrderId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePlannedOrder(plannedOrder.plannedOrderId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPlannedOrder(plannedOrder.plannedOrderId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPlannedOrderComponent
