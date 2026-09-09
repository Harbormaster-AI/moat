import React, { Component } from 'react'
import MedicalDeviceService from '../services/MedicalDeviceService'

class ViewMedicalDeviceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            medicalDevice: {}
        }
    }

    componentDidMount(){
        MedicalDeviceService.getMedicalDeviceById(this.state.id).then( res => {
            this.setState({medicalDevice: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View MedicalDevice Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> udi:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.medicalDevice.udi }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> manufacturer:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.medicalDevice.manufacturer }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> DeviceType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.medicalDevice.deviceType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ConnectivityStatus:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.medicalDevice.connectivityStatus }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewMedicalDeviceComponent
