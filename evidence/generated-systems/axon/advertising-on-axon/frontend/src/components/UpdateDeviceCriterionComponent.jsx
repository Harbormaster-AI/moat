import React, { Component } from 'react'
import DeviceCriterionService from '../services/DeviceCriterionService';

class UpdateDeviceCriterionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                deviceType: '',
                platformType: '',
                operator: ''
        }
        this.updateDeviceCriterion = this.updateDeviceCriterion.bind(this);

        this.changeDeviceTypeHandler = this.changeDeviceTypeHandler.bind(this);
        this.changePlatformTypeHandler = this.changePlatformTypeHandler.bind(this);
        this.changeOperatorHandler = this.changeOperatorHandler.bind(this);
    }

    componentDidMount(){
        DeviceCriterionService.getDeviceCriterionById(this.state.id).then( (res) =>{
            let deviceCriterion = res.data;
            this.setState({
                deviceType: deviceCriterion.deviceType,
                platformType: deviceCriterion.platformType,
                operator: deviceCriterion.operator
            });
        });
    }

    updateDeviceCriterion = (e) => {
        e.preventDefault();
        let deviceCriterion = {
            deviceCriterionId: this.state.id,
            deviceType: this.state.deviceType,
            platformType: this.state.platformType,
            operator: this.state.operator
        };
        console.log('deviceCriterion => ' + JSON.stringify(deviceCriterion));
        console.log('id => ' + JSON.stringify(this.state.id));
        DeviceCriterionService.updateDeviceCriterion(deviceCriterion).then( res => {
            this.props.history.push('/deviceCriterions');
        });
    }

    changeDeviceTypeHandler= (event) => {
        this.setState({deviceType: event.target.value});
    }
    changePlatformTypeHandler= (event) => {
        this.setState({platformType: event.target.value});
    }
    changeOperatorHandler= (event) => {
        this.setState({operator: event.target.value});
    }

    cancel(){
        this.props.history.push('/deviceCriterions');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update DeviceCriterion</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> DeviceType: </label>
                                                <select value={this.state.deviceType} onChange={this.changeDeviceTypeHandler}>
                      <option name="DeviceType" className="form-control" >
                          Desktop
                      </option>
                      <option name="DeviceType" className="form-control" >
                          Mobile
                      </option>
                      <option name="DeviceType" className="form-control" >
                          Tablet
                      </option>
                      <option name="DeviceType" className="form-control" >
                          ConnectedTV
                      </option>
                    </select>

                                            <label> PlatformType: </label>
                                                <select value={this.state.platformType} onChange={this.changePlatformTypeHandler}>
                      <option name="PlatformType" className="form-control" >
                          Web
                      </option>
                      <option name="PlatformType" className="form-control" >
                          MobileApp
                      </option>
                      <option name="PlatformType" className="form-control" >
                          CTV
                      </option>
                    </select>

                                            <label> Operator: </label>
                                                <select value={this.state.operator} onChange={this.changeOperatorHandler}>
                      <option name="Operator" className="form-control" >
                          Include
                      </option>
                      <option name="Operator" className="form-control" >
                          Exclude
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateDeviceCriterion}>Save</button>
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

export default UpdateDeviceCriterionComponent
