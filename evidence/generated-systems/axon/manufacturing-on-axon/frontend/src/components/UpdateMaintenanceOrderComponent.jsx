import React, { Component } from 'react'
import MaintenanceOrderService from '../services/MaintenanceOrderService';

class UpdateMaintenanceOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                orderNumber: '',
                priority: '',
                requestedDate: '',
                completionDate: '',
                status: ''
        }
        this.updateMaintenanceOrder = this.updateMaintenanceOrder.bind(this);

        this.changeorderNumberHandler = this.changeorderNumberHandler.bind(this);
        this.changepriorityHandler = this.changepriorityHandler.bind(this);
        this.changerequestedDateHandler = this.changerequestedDateHandler.bind(this);
        this.changecompletionDateHandler = this.changecompletionDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        MaintenanceOrderService.getMaintenanceOrderById(this.state.id).then( (res) =>{
            let maintenanceOrder = res.data;
            this.setState({
                orderNumber: maintenanceOrder.orderNumber,
                priority: maintenanceOrder.priority,
                requestedDate: maintenanceOrder.requestedDate,
                completionDate: maintenanceOrder.completionDate,
                status: maintenanceOrder.status
            });
        });
    }

    updateMaintenanceOrder = (e) => {
        e.preventDefault();
        let maintenanceOrder = {
            maintenanceOrderId: this.state.id,
            orderNumber: this.state.orderNumber,
            priority: this.state.priority,
            requestedDate: this.state.requestedDate,
            completionDate: this.state.completionDate,
            status: this.state.status
        };
        console.log('maintenanceOrder => ' + JSON.stringify(maintenanceOrder));
        console.log('id => ' + JSON.stringify(this.state.id));
        MaintenanceOrderService.updateMaintenanceOrder(maintenanceOrder).then( res => {
            this.props.history.push('/maintenanceOrders');
        });
    }

    changeorderNumberHandler= (event) => {
        this.setState({orderNumber: event.target.value});
    }
    changepriorityHandler= (event) => {
        this.setState({priority: event.target.value});
    }
    changerequestedDateHandler= (event) => {
        this.setState({requestedDate: event.target.value});
    }
    changecompletionDateHandler= (event) => {
        this.setState({completionDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/maintenanceOrders');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update MaintenanceOrder</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> orderNumber: </label>
                                                <input placeholder="orderNumber" name="orderNumber" className="form-control" value={this.state.orderNumber} onChange={this.changeorderNumberHandler}/>

                                            <label> priority: </label>
                                                <input type="number" placeholder="priority" name="priority" className="form-control" value={this.state.priority} onChange={this.changepriorityHandler}/>

                                            <label> requestedDate: </label>
                                                <input type="date" placeholder="requestedDate" name="requestedDate" className="form-control" value={this.state.requestedDate} onChange={this.changerequestedDateHandler}/>

                                            <label> completionDate: </label>
                                                <input type="date" placeholder="completionDate" name="completionDate" className="form-control" value={this.state.completionDate} onChange={this.changecompletionDateHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Created
                      </option>
                      <option name="Status" className="form-control" >
                          Approved
                      </option>
                      <option name="Status" className="form-control" >
                          Scheduled
                      </option>
                      <option name="Status" className="form-control" >
                          InProgress
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateMaintenanceOrder}>Save</button>
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

export default UpdateMaintenanceOrderComponent
