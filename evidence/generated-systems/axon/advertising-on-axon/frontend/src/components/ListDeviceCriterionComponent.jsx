import React, { Component } from 'react'
import DeviceCriterionService from '../services/DeviceCriterionService'

class ListDeviceCriterionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                deviceCriterions: []
        }
        this.addDeviceCriterion = this.addDeviceCriterion.bind(this);
        this.editDeviceCriterion = this.editDeviceCriterion.bind(this);
        this.deleteDeviceCriterion = this.deleteDeviceCriterion.bind(this);
    }

    deleteDeviceCriterion(id){
        DeviceCriterionService.deleteDeviceCriterion(id).then( res => {
            this.setState({deviceCriterions: this.state.deviceCriterions.filter(deviceCriterion => deviceCriterion.deviceCriterionId !== id)});
        });
    }
    viewDeviceCriterion(id){
        this.props.history.push(`/view-deviceCriterion/${id}`);
    }
    editDeviceCriterion(id){
        this.props.history.push(`/add-deviceCriterion/${id}`);
    }

    componentDidMount(){
        DeviceCriterionService.getDeviceCriterions().then((res) => {
            this.setState({ deviceCriterions: res.data});
        });
    }

    addDeviceCriterion(){
        this.props.history.push('/add-deviceCriterion/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">DeviceCriterion List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDeviceCriterion}> Add DeviceCriterion</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> DeviceType </th>
                                    <th> PlatformType </th>
                                    <th> Operator </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.deviceCriterions.map(
                                        deviceCriterion => 
                                        <tr key = {deviceCriterion.deviceCriterionId}>
                                             <td> { deviceCriterion.deviceType } </td>
                                             <td> { deviceCriterion.platformType } </td>
                                             <td> { deviceCriterion.operator } </td>
                                             <td>
                                                 <button onClick={ () => this.editDeviceCriterion(deviceCriterion.deviceCriterionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDeviceCriterion(deviceCriterion.deviceCriterionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDeviceCriterion(deviceCriterion.deviceCriterionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListDeviceCriterionComponent
