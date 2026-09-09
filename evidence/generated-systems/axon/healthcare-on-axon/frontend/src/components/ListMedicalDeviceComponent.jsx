import React, { Component } from 'react'
import MedicalDeviceService from '../services/MedicalDeviceService'

class ListMedicalDeviceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                medicalDevices: []
        }
        this.addMedicalDevice = this.addMedicalDevice.bind(this);
        this.editMedicalDevice = this.editMedicalDevice.bind(this);
        this.deleteMedicalDevice = this.deleteMedicalDevice.bind(this);
    }

    deleteMedicalDevice(id){
        MedicalDeviceService.deleteMedicalDevice(id).then( res => {
            this.setState({medicalDevices: this.state.medicalDevices.filter(medicalDevice => medicalDevice.medicalDeviceId !== id)});
        });
    }
    viewMedicalDevice(id){
        this.props.history.push(`/view-medicalDevice/${id}`);
    }
    editMedicalDevice(id){
        this.props.history.push(`/add-medicalDevice/${id}`);
    }

    componentDidMount(){
        MedicalDeviceService.getMedicalDevices().then((res) => {
            this.setState({ medicalDevices: res.data});
        });
    }

    addMedicalDevice(){
        this.props.history.push('/add-medicalDevice/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">MedicalDevice List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addMedicalDevice}> Add MedicalDevice</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Udi </th>
                                    <th> Manufacturer </th>
                                    <th> DeviceType </th>
                                    <th> ConnectivityStatus </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.medicalDevices.map(
                                        medicalDevice => 
                                        <tr key = {medicalDevice.medicalDeviceId}>
                                             <td> { medicalDevice.udi } </td>
                                             <td> { medicalDevice.manufacturer } </td>
                                             <td> { medicalDevice.deviceType } </td>
                                             <td> { medicalDevice.connectivityStatus } </td>
                                             <td>
                                                 <button onClick={ () => this.editMedicalDevice(medicalDevice.medicalDeviceId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteMedicalDevice(medicalDevice.medicalDeviceId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewMedicalDevice(medicalDevice.medicalDeviceId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListMedicalDeviceComponent
