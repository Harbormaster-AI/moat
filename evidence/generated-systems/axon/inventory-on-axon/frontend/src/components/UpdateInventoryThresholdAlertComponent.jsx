import React, { Component } from 'react'
import InventoryThresholdAlertService from '../services/InventoryThresholdAlertService';

class UpdateInventoryThresholdAlertComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                alertNumber: '',
                detectedAt: '',
                message: '',
                alertType: '',
                status: ''
        }
        this.updateInventoryThresholdAlert = this.updateInventoryThresholdAlert.bind(this);

        this.changealertNumberHandler = this.changealertNumberHandler.bind(this);
        this.changedetectedAtHandler = this.changedetectedAtHandler.bind(this);
        this.changemessageHandler = this.changemessageHandler.bind(this);
        this.changeAlertTypeHandler = this.changeAlertTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
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

    updateInventoryThresholdAlert = (e) => {
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
        console.log('id => ' + JSON.stringify(this.state.id));
        InventoryThresholdAlertService.updateInventoryThresholdAlert(inventoryThresholdAlert).then( res => {
            this.props.history.push('/inventoryThresholdAlerts');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update InventoryThresholdAlert</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> alertNumber: </label>
                                                <input placeholder="alertNumber" name="alertNumber" className="form-control" value={this.state.alertNumber} onChange={this.changealertNumberHandler}/>

                                            <label> detectedAt: </label>
                                                <input type="date" placeholder="detectedAt" name="detectedAt" className="form-control" value={this.state.detectedAt} onChange={this.changedetectedAtHandler}/>

                                            <label> message: </label>
                                                <input placeholder="message" name="message" className="form-control" value={this.state.message} onChange={this.changemessageHandler}/>

                                            <label> AlertType: </label>
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

                                            <label> Status: </label>
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
                                        <button className="btn btn-success" onClick={this.updateInventoryThresholdAlert}>Save</button>
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

export default UpdateInventoryThresholdAlertComponent
