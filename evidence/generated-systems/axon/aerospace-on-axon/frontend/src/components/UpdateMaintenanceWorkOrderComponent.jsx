import React, { Component } from 'react'
import MaintenanceWorkOrderService from '../services/MaintenanceWorkOrderService';

class UpdateMaintenanceWorkOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                workOrderNumber: '',
                status: ''
        }
        this.updateMaintenanceWorkOrder = this.updateMaintenanceWorkOrder.bind(this);

        this.changeworkOrderNumberHandler = this.changeworkOrderNumberHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        MaintenanceWorkOrderService.getMaintenanceWorkOrderById(this.state.id).then( (res) =>{
            let maintenanceWorkOrder = res.data;
            this.setState({
                workOrderNumber: maintenanceWorkOrder.workOrderNumber,
                status: maintenanceWorkOrder.status
            });
        });
    }

    updateMaintenanceWorkOrder = (e) => {
        e.preventDefault();
        let maintenanceWorkOrder = {
            maintenanceWorkOrderId: this.state.id,
            workOrderNumber: this.state.workOrderNumber,
            status: this.state.status
        };
        console.log('maintenanceWorkOrder => ' + JSON.stringify(maintenanceWorkOrder));
        console.log('id => ' + JSON.stringify(this.state.id));
        MaintenanceWorkOrderService.updateMaintenanceWorkOrder(maintenanceWorkOrder).then( res => {
            this.props.history.push('/maintenanceWorkOrders');
        });
    }

    changeworkOrderNumberHandler= (event) => {
        this.setState({workOrderNumber: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/maintenanceWorkOrders');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update MaintenanceWorkOrder</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> workOrderNumber: </label>
                                                <input placeholder="workOrderNumber" name="workOrderNumber" className="form-control" value={this.state.workOrderNumber} onChange={this.changeworkOrderNumberHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Open
                      </option>
                      <option name="Status" className="form-control" >
                          InProgress
                      </option>
                      <option name="Status" className="form-control" >
                          AwaitingParts
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                      <option name="Status" className="form-control" >
                          Deferred
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateMaintenanceWorkOrder}>Save</button>
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

export default UpdateMaintenanceWorkOrderComponent
