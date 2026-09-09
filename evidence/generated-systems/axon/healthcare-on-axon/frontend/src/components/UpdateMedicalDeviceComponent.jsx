import React, { Component } from 'react'
import MedicalDeviceService from '../services/MedicalDeviceService';

class UpdateMedicalDeviceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                udi: '',
                manufacturer: '',
                deviceType: '',
                connectivityStatus: ''
        }
        this.updateMedicalDevice = this.updateMedicalDevice.bind(this);

        this.changeudiHandler = this.changeudiHandler.bind(this);
        this.changemanufacturerHandler = this.changemanufacturerHandler.bind(this);
        this.changeDeviceTypeHandler = this.changeDeviceTypeHandler.bind(this);
        this.changeConnectivityStatusHandler = this.changeConnectivityStatusHandler.bind(this);
    }

    componentDidMount(){
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

    updateMedicalDevice = (e) => {
        e.preventDefault();
        let medicalDevice = {
            medicalDeviceId: this.state.id,
            udi: this.state.udi,
            manufacturer: this.state.manufacturer,
            deviceType: this.state.deviceType,
            connectivityStatus: this.state.connectivityStatus
        };
        console.log('medicalDevice => ' + JSON.stringify(medicalDevice));
        console.log('id => ' + JSON.stringify(this.state.id));
        MedicalDeviceService.updateMedicalDevice(medicalDevice).then( res => {
            this.props.history.push('/medicalDevices');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update MedicalDevice</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> udi: </label>
                                                <input placeholder="udi" name="udi" className="form-control" value={this.state.udi} onChange={this.changeudiHandler}/>

                                            <label> manufacturer: </label>
                                                <input placeholder="manufacturer" name="manufacturer" className="form-control" value={this.state.manufacturer} onChange={this.changemanufacturerHandler}/>

                                            <label> DeviceType: </label>
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

                                            <label> ConnectivityStatus: </label>
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
                                        <button className="btn btn-success" onClick={this.updateMedicalDevice}>Save</button>
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

export default UpdateMedicalDeviceComponent
