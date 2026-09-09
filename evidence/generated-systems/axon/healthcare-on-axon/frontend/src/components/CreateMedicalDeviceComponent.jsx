import React, { Component } from 'react'
import MedicalDeviceService from '../services/MedicalDeviceService';

class CreateMedicalDeviceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                udi: '',
                manufacturer: '',
                deviceType: '',
                connectivityStatus: ''
        }
        this.changeudiHandler = this.changeudiHandler.bind(this);
        this.changemanufacturerHandler = this.changemanufacturerHandler.bind(this);
        this.changeDeviceTypeHandler = this.changeDeviceTypeHandler.bind(this);
        this.changeConnectivityStatusHandler = this.changeConnectivityStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            MedicalDeviceService.getMedicalDeviceById(this.state.id).then( (res) =>{
                let medicalDevice = res.data;
                this.setState({
                    udi: medicalDevice.udi,
                    manufacturer: medicalDevice.manufacturer,
                    deviceType: medicalDevice.deviceType,
                    connectivityStatus: medicalDevice.connectivityStatus
                });
            });
        }        
    }
    saveOrUpdateMedicalDevice = (e) => {
        e.preventDefault();
        let medicalDevice = {
                medicalDeviceId: this.state.id,
                udi: this.state.udi,
                manufacturer: this.state.manufacturer,
                deviceType: this.state.deviceType,
                connectivityStatus: this.state.connectivityStatus
            };
        console.log('medicalDevice => ' + JSON.stringify(medicalDevice));

        // step 5
        if(this.state.id === '_add'){
            medicalDevice.medicalDeviceId=''
            MedicalDeviceService.createMedicalDevice(medicalDevice).then(res =>{
                this.props.history.push('/medicalDevices');
            });
        }else{
            MedicalDeviceService.updateMedicalDevice(medicalDevice).then( res => {
                this.props.history.push('/medicalDevices');
            });
        }
    }
    
    changeudiHandler= (event) => {
        this.setState({udi: event.target.value});
    }
    changemanufacturerHandler= (event) => {
        this.setState({manufacturer: event.target.value});
    }
    changeDeviceTypeHandler= (event) => {
        this.setState({deviceType: event.target.value});
    }
    changeConnectivityStatusHandler= (event) => {
        this.setState({connectivityStatus: event.target.value});
    }

    cancel(){
        this.props.history.push('/medicalDevices');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add MedicalDevice</h3>
        }else{
            return <h3 className="text-center">Update MedicalDevice</h3>
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
                                            <label> udi:&emsp; </label>
                                                <input placeholder="udi" name="udi" className="form-control" value={this.state.udi} onChange={this.changeudiHandler}/>

                                            <label> manufacturer:&emsp; </label>
                                                <input placeholder="manufacturer" name="manufacturer" className="form-control" value={this.state.manufacturer} onChange={this.changemanufacturerHandler}/>

                                            <label> DeviceType:&emsp; </label>
                                                <select value={this.state.deviceType} onChange={this.changeDeviceTypeHandler}>
                      <option name="DeviceType" className="form-control" >
                          Pacemaker
                      </option>
                      <option name="DeviceType" className="form-control" >
                          InsulinPump
                      </option>
                      <option name="DeviceType" className="form-control" >
                          BloodPressureMonitor
                      </option>
                      <option name="DeviceType" className="form-control" >
                          GlucoseMeter
                      </option>
                      <option name="DeviceType" className="form-control" >
                          PulseOximeter
                      </option>
                      <option name="DeviceType" className="form-control" >
                          Ventilator
                      </option>
                      <option name="DeviceType" className="form-control" >
                          InfusionPump
                      </option>
                      <option name="DeviceType" className="form-control" >
                          WearableTracker
                      </option>
                    </select>

                                            <label> ConnectivityStatus:&emsp; </label>
                                                <select value={this.state.connectivityStatus} onChange={this.changeConnectivityStatusHandler}>
                      <option name="ConnectivityStatus" className="form-control" >
                          Connected
                      </option>
                      <option name="ConnectivityStatus" className="form-control" >
                          Disconnected
                      </option>
                      <option name="ConnectivityStatus" className="form-control" >
                          Standby
                      </option>
                      <option name="ConnectivityStatus" className="form-control" >
                          Fault
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateMedicalDevice}>Save</button>
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

export default CreateMedicalDeviceComponent
