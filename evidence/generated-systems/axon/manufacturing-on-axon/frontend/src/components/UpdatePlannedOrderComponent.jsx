import React, { Component } from 'react'
import PlannedOrderService from '../services/PlannedOrderService';

class UpdatePlannedOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                plannedOrderNumber: '',
                quantity: '',
                dueDate: '',
                orderType: '',
                status: ''
        }
        this.updatePlannedOrder = this.updatePlannedOrder.bind(this);

        this.changeplannedOrderNumberHandler = this.changeplannedOrderNumberHandler.bind(this);
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changedueDateHandler = this.changedueDateHandler.bind(this);
        this.changeOrderTypeHandler = this.changeOrderTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        PlannedOrderService.getPlannedOrderById(this.state.id).then( (res) =>{
            let plannedOrder = res.data;
            this.setState({
                plannedOrderNumber: plannedOrder.plannedOrderNumber,
                quantity: plannedOrder.quantity,
                dueDate: plannedOrder.dueDate,
                orderType: plannedOrder.orderType,
                status: plannedOrder.status
            });
        });
    }

    updatePlannedOrder = (e) => {
        e.preventDefault();
        let plannedOrder = {
            plannedOrderId: this.state.id,
            plannedOrderNumber: this.state.plannedOrderNumber,
            quantity: this.state.quantity,
            dueDate: this.state.dueDate,
            orderType: this.state.orderType,
            status: this.state.status
        };
        console.log('plannedOrder => ' + JSON.stringify(plannedOrder));
        console.log('id => ' + JSON.stringify(this.state.id));
        PlannedOrderService.updatePlannedOrder(plannedOrder).then( res => {
            this.props.history.push('/plannedOrders');
        });
    }

    changeplannedOrderNumberHandler= (event) => {
        this.setState({plannedOrderNumber: event.target.value});
    }
    changequantityHandler= (event) => {
        this.setState({quantity: event.target.value});
    }
    changedueDateHandler= (event) => {
        this.setState({dueDate: event.target.value});
    }
    changeOrderTypeHandler= (event) => {
        this.setState({orderType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/plannedOrders');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update PlannedOrder</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> plannedOrderNumber: </label>
                                                <input placeholder="plannedOrderNumber" name="plannedOrderNumber" className="form-control" value={this.state.plannedOrderNumber} onChange={this.changeplannedOrderNumberHandler}/>

                                            <label> quantity: </label>
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> dueDate: </label>
                                                <input type="date" placeholder="dueDate" name="dueDate" className="form-control" value={this.state.dueDate} onChange={this.changedueDateHandler}/>

                                            <label> OrderType: </label>
                                                <select value={this.state.orderType} onChange={this.changeOrderTypeHandler}>
                      <option name="OrderType" className="form-control" >
                          WorkOrder
                      </option>
                      <option name="OrderType" className="form-control" >
                          PurchaseRequisition
                      </option>
                      <option name="OrderType" className="form-control" >
                          TransferOrder
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Planned
                      </option>
                      <option name="Status" className="form-control" >
                          Firmed
                      </option>
                      <option name="Status" className="form-control" >
                          Released
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePlannedOrder}>Save</button>
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

export default UpdatePlannedOrderComponent
