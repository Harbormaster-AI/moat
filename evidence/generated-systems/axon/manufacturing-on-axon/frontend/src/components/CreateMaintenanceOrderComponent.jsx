import React, { Component } from 'react'
import MaintenanceOrderService from '../services/MaintenanceOrderService';

class CreateMaintenanceOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                orderNumber: '',
                priority: '',
                requestedDate: '',
                completionDate: '',
                status: ''
        }
        this.changeorderNumberHandler = this.changeorderNumberHandler.bind(this);
        this.changepriorityHandler = this.changepriorityHandler.bind(this);
        this.changerequestedDateHandler = this.changerequestedDateHandler.bind(this);
        this.changecompletionDateHandler = this.changecompletionDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateMaintenanceOrder = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            maintenanceOrder.maintenanceOrderId=''
            MaintenanceOrderService.createMaintenanceOrder(maintenanceOrder).then(res =>{
                this.props.history.push('/maintenanceOrders');
            });
        }else{
            MaintenanceOrderService.updateMaintenanceOrder(maintenanceOrder).then( res => {
                this.props.history.push('/maintenanceOrders');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add MaintenanceOrder</h3>
        }else{
            return <h3 className="text-center">Update MaintenanceOrder</h3>
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
                                            <label> orderNumber:&emsp; </label>
                                                <input placeholder="orderNumber" name="orderNumber" className="form-control" value={this.state.orderNumber} onChange={this.changeorderNumberHandler}/>

                                            <label> priority:&emsp; </label>
                                                <input type="number" placeholder="priority" name="priority" className="form-control" value={this.state.priority} onChange={this.changepriorityHandler}/>

                                            <label> requestedDate:&emsp; </label>
                                                <input type="date" placeholder="requestedDate" name="requestedDate" className="form-control" value={this.state.requestedDate} onChange={this.changerequestedDateHandler}/>

                                            <label> completionDate:&emsp; </label>
                                                <input type="date" placeholder="completionDate" name="completionDate" className="form-control" value={this.state.completionDate} onChange={this.changecompletionDateHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateMaintenanceOrder}>Save</button>
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

export default CreateMaintenanceOrderComponent
