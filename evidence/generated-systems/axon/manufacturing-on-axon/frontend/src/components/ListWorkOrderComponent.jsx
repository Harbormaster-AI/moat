import React, { Component } from 'react'
import WorkOrderService from '../services/WorkOrderService'

class ListWorkOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                workOrders: []
        }
        this.addWorkOrder = this.addWorkOrder.bind(this);
        this.editWorkOrder = this.editWorkOrder.bind(this);
        this.deleteWorkOrder = this.deleteWorkOrder.bind(this);
    }

    deleteWorkOrder(id){
        WorkOrderService.deleteWorkOrder(id).then( res => {
            this.setState({workOrders: this.state.workOrders.filter(workOrder => workOrder.workOrderId !== id)});
        });
    }
    viewWorkOrder(id){
        this.props.history.push(`/view-workOrder/${id}`);
    }
    editWorkOrder(id){
        this.props.history.push(`/add-workOrder/${id}`);
    }

    componentDidMount(){
        WorkOrderService.getWorkOrders().then((res) => {
            this.setState({ workOrders: res.data});
        });
    }

    addWorkOrder(){
        this.props.history.push('/add-workOrder/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">WorkOrder List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addWorkOrder}> Add WorkOrder</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> WorkOrderNumber </th>
                                    <th> PlannedStart </th>
                                    <th> PlannedEnd </th>
                                    <th> Quantity </th>
                                    <th> Priority </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.workOrders.map(
                                        workOrder => 
                                        <tr key = {workOrder.workOrderId}>
                                             <td> { workOrder.workOrderNumber } </td>
                                             <td> { workOrder.plannedStart } </td>
                                             <td> { workOrder.plannedEnd } </td>
                                             <td> { workOrder.quantity } </td>
                                             <td> { workOrder.priority } </td>
                                             <td> { workOrder.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editWorkOrder(workOrder.workOrderId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteWorkOrder(workOrder.workOrderId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewWorkOrder(workOrder.workOrderId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListWorkOrderComponent
