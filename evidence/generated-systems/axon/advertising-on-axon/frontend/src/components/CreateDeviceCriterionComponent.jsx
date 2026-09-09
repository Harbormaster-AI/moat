import React, { Component } from 'react'
import DeviceCriterionService from '../services/DeviceCriterionService';

class CreateDeviceCriterionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                deviceType: '',
                platformType: '',
                operator: ''
        }
        this.changeDeviceTypeHandler = this.changeDeviceTypeHandler.bind(this);
        this.changePlatformTypeHandler = this.changePlatformTypeHandler.bind(this);
        this.changeOperatorHandler = this.changeOperatorHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            DeviceCriterionService.getDeviceCriterionById(this.state.id).then( (res) =>{
                let deviceCriterion = res.data;
                this.setState({
                    deviceType: deviceCriterion.deviceType,
                    platformType: deviceCriterion.platformType,
                    operator: deviceCriterion.operator
                });
            });
        }        
    }
    saveOrUpdateDeviceCriterion = (e) => {
        e.preventDefault();
        let deviceCriterion = {
                deviceCriterionId: this.state.id,
                deviceType: this.state.deviceType,
                platformType: this.state.platformType,
                operator: this.state.operator
            };
        console.log('deviceCriterion => ' + JSON.stringify(deviceCriterion));

        // step 5
        if(this.state.id === '_add'){
            deviceCriterion.deviceCriterionId=''
            DeviceCriterionService.createDeviceCriterion(deviceCriterion).then(res =>{
                this.props.history.push('/deviceCriterions');
            });
        }else{
            DeviceCriterionService.updateDeviceCriterion(deviceCriterion).then( res => {
                this.props.history.push('/deviceCriterions');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add DeviceCriterion</h3>
        }else{
            return <h3 className="text-center">Update DeviceCriterion</h3>
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
                                            <label> DeviceType:&emsp; </label>
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

                                            <label> PlatformType:&emsp; </label>
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

                                            <label> Operator:&emsp; </label>
                                                <select value={this.state.operator} onChange={this.changeOperatorHandler}>
                      <option name="Operator" className="form-control" >
                          Include
                      </option>
                      <option name="Operator" className="form-control" >
                          Exclude
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateDeviceCriterion}>Save</button>
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

export default CreateDeviceCriterionComponent
