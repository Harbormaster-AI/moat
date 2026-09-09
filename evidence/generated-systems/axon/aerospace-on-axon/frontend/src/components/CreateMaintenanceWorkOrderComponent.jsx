import React, { Component } from 'react'
import MaintenanceWorkOrderService from '../services/MaintenanceWorkOrderService';

class CreateMaintenanceWorkOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                workOrderNumber: '',
                status: ''
        }
        this.changeworkOrderNumberHandler = this.changeworkOrderNumberHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            MaintenanceWorkOrderService.getMaintenanceWorkOrderById(this.state.id).then( (res) =>{
                let maintenanceWorkOrder = res.data;
                this.setState({
                    workOrderNumber: maintenanceWorkOrder.workOrderNumber,
                    status: maintenanceWorkOrder.status
                });
            });
        }        
    }
    saveOrUpdateMaintenanceWorkOrder = (e) => {
        e.preventDefault();
        let maintenanceWorkOrder = {
                maintenanceWorkOrderId: this.state.id,
                workOrderNumber: this.state.workOrderNumber,
                status: this.state.status
            };
        console.log('maintenanceWorkOrder => ' + JSON.stringify(maintenanceWorkOrder));

        // step 5
        if(this.state.id === '_add'){
            maintenanceWorkOrder.maintenanceWorkOrderId=''
            MaintenanceWorkOrderService.createMaintenanceWorkOrder(maintenanceWorkOrder).then(res =>{
                this.props.history.push('/maintenanceWorkOrders');
            });
        }else{
            MaintenanceWorkOrderService.updateMaintenanceWorkOrder(maintenanceWorkOrder).then( res => {
                this.props.history.push('/maintenanceWorkOrders');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add MaintenanceWorkOrder</h3>
        }else{
            return <h3 className="text-center">Update MaintenanceWorkOrder</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> workOrderNumber:&emsp; </label>
                                                <input placeholder="workOrderNumber" name="workOrderNumber" className="form-control" value={this.state.workOrderNumber} onChange={this.changeworkOrderNumberHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateMaintenanceWorkOrder}>Save</button>
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

export default CreateMaintenanceWorkOrderComponent
