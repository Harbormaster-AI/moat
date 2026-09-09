import React, { Component } from 'react'
import InventoryThresholdAlertService from '../services/InventoryThresholdAlertService';

class CreateInventoryThresholdAlertComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                alertNumber: '',
                detectedAt: '',
                message: '',
                alertType: '',
                status: ''
        }
        this.changealertNumberHandler = this.changealertNumberHandler.bind(this);
        this.changedetectedAtHandler = this.changedetectedAtHandler.bind(this);
        this.changemessageHandler = this.changemessageHandler.bind(this);
        this.changeAlertTypeHandler = this.changeAlertTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            InventoryThresholdAlertService.getInventoryThresholdAlertById(this.state.id).then( (res) =>{
                let inventoryThresholdAlert = res.data;
                this.setState({
                    alertNumber: inventoryThresholdAlert.alertNumber,
                    detectedAt: inventoryThresholdAlert.detectedAt,
                    message: inventoryThresholdAlert.message,
                    alertType: inventoryThresholdAlert.alertType,
                    status: inventoryThresholdAlert.status
                });
            });
        }        
    }
    saveOrUpdateInventoryThresholdAlert = (e) => {
        e.preventDefault();
        let inventoryThresholdAlert = {
                inventoryThresholdAlertId: this.state.id,
                alertNumber: this.state.alertNumber,
                detectedAt: this.state.detectedAt,
                message: this.state.message,
                alertType: this.state.alertType,
                status: this.state.status
            };
        console.log('inventoryThresholdAlert => ' + JSON.stringify(inventoryThresholdAlert));

        // step 5
        if(this.state.id === '_add'){
            inventoryThresholdAlert.inventoryThresholdAlertId=''
            InventoryThresholdAlertService.createInventoryThresholdAlert(inventoryThresholdAlert).then(res =>{
                this.props.history.push('/inventoryThresholdAlerts');
            });
        }else{
            InventoryThresholdAlertService.updateInventoryThresholdAlert(inventoryThresholdAlert).then( res => {
                this.props.history.push('/inventoryThresholdAlerts');
            });
        }
    }
    
    changealertNumberHandler= (event) => {
        this.setState({alertNumber: event.target.value});
    }
    changedetectedAtHandler= (event) => {
        this.setState({detectedAt: event.target.value});
    }
    changemessageHandler= (event) => {
        this.setState({message: event.target.value});
    }
    changeAlertTypeHandler= (event) => {
        this.setState({alertType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/inventoryThresholdAlerts');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add InventoryThresholdAlert</h3>
        }else{
            return <h3 className="text-center">Update InventoryThresholdAlert</h3>
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
                                            <label> alertNumber:&emsp; </label>
                                                <input placeholder="alertNumber" name="alertNumber" className="form-control" value={this.state.alertNumber} onChange={this.changealertNumberHandler}/>

                                            <label> detectedAt:&emsp; </label>
                                                <input type="date" placeholder="detectedAt" name="detectedAt" className="form-control" value={this.state.detectedAt} onChange={this.changedetectedAtHandler}/>

                                            <label> message:&emsp; </label>
                                                <input placeholder="message" name="message" className="form-control" value={this.state.message} onChange={this.changemessageHandler}/>

                                            <label> AlertType:&emsp; </label>
                                                <select value={this.state.alertType} onChange={this.changeAlertTypeHandler}>
                      <option name="AlertType" className="form-control" >
                          BelowMin
                      </option>
                      <option name="AlertType" className="form-control" >
                          AboveMax
                      </option>
                      <option name="AlertType" className="form-control" >
                          StockoutRisk
                      </option>
                      <option name="AlertType" className="form-control" >
                          ExcessStock
                      </option>
                      <option name="AlertType" className="form-control" >
                          ExpiryRisk
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          New
                      </option>
                      <option name="Status" className="form-control" >
                          Acknowledged
                      </option>
                      <option name="Status" className="form-control" >
                          Resolved
                      </option>
                      <option name="Status" className="form-control" >
                          Dismissed
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateInventoryThresholdAlert}>Save</button>
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

export default CreateInventoryThresholdAlertComponent
