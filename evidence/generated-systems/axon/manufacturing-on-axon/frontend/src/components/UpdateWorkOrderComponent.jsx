import React, { Component } from 'react'
import WorkOrderService from '../services/WorkOrderService';

class UpdateWorkOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                workOrderNumber: '',
                plannedStart: '',
                plannedEnd: '',
                quantity: '',
                priority: '',
                status: ''
        }
        this.updateWorkOrder = this.updateWorkOrder.bind(this);

        this.changeworkOrderNumberHandler = this.changeworkOrderNumberHandler.bind(this);
        this.changeplannedStartHandler = this.changeplannedStartHandler.bind(this);
        this.changeplannedEndHandler = this.changeplannedEndHandler.bind(this);
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changepriorityHandler = this.changepriorityHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        WorkOrderService.getWorkOrderById(this.state.id).then( (res) =>{
            let workOrder = res.data;
            this.setState({
                workOrderNumber: workOrder.workOrderNumber,
                plannedStart: workOrder.plannedStart,
                plannedEnd: workOrder.plannedEnd,
                quantity: workOrder.quantity,
                priority: workOrder.priority,
                status: workOrder.status
            });
        });
    }

    updateWorkOrder = (e) => {
        e.preventDefault();
        let workOrder = {
            workOrderId: this.state.id,
            workOrderNumber: this.state.workOrderNumber,
            plannedStart: this.state.plannedStart,
            plannedEnd: this.state.plannedEnd,
            quantity: this.state.quantity,
            priority: this.state.priority,
            status: this.state.status
        };
        console.log('workOrder => ' + JSON.stringify(workOrder));
        console.log('id => ' + JSON.stringify(this.state.id));
        WorkOrderService.updateWorkOrder(workOrder).then( res => {
            this.props.history.push('/workOrders');
        });
    }

    changeworkOrderNumberHandler= (event) => {
        this.setState({workOrderNumber: event.target.value});
    }
    changeplannedStartHandler= (event) => {
        this.setState({plannedStart: event.target.value});
    }
    changeplannedEndHandler= (event) => {
        this.setState({plannedEnd: event.target.value});
    }
    changequantityHandler= (event) => {
        this.setState({quantity: event.target.value});
    }
    changepriorityHandler= (event) => {
        this.setState({priority: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/workOrders');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update WorkOrder</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> workOrderNumber: </label>
                                                <input placeholder="workOrderNumber" name="workOrderNumber" className="form-control" value={this.state.workOrderNumber} onChange={this.changeworkOrderNumberHandler}/>

                                            <label> plannedStart: </label>
                                                <input type="time" placeholder="plannedStart" name="plannedStart" className="form-control" value={this.state.plannedStart} onChange={this.changeplannedStartHandler}/>

                                            <label> plannedEnd: </label>
                                                <input type="time" placeholder="plannedEnd" name="plannedEnd" className="form-control" value={this.state.plannedEnd} onChange={this.changeplannedEndHandler}/>

                                            <label> quantity: </label>
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> priority: </label>
                                                <input type="number" placeholder="priority" name="priority" className="form-control" value={this.state.priority} onChange={this.changepriorityHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Planned
                      </option>
                      <option name="Status" className="form-control" >
                          Released
                      </option>
                      <option name="Status" className="form-control" >
                          InProcess
                      </option>
                      <option name="Status" className="form-control" >
                          Hold
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateWorkOrder}>Save</button>
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

export default UpdateWorkOrderComponent
