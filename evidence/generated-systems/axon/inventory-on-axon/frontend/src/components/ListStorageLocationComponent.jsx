import React, { Component } from 'react'
import StorageLocationService from '../services/StorageLocationService'

class ListStorageLocationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                storageLocations: []
        }
        this.addStorageLocation = this.addStorageLocation.bind(this);
        this.editStorageLocation = this.editStorageLocation.bind(this);
        this.deleteStorageLocation = this.deleteStorageLocation.bind(this);
    }

    deleteStorageLocation(id){
        StorageLocationService.deleteStorageLocation(id).then( res => {
            this.setState({storageLocations: this.state.storageLocations.filter(storageLocation => storageLocation.storageLocationId !== id)});
        });
    }
    viewStorageLocation(id){
        this.props.history.push(`/view-storageLocation/${id}`);
    }
    editStorageLocation(id){
        this.props.history.push(`/add-storageLocation/${id}`);
    }

    componentDidMount(){
        StorageLocationService.getStorageLocations().then((res) => {
            this.setState({ storageLocations: res.data});
        });
    }

    addStorageLocation(){
        this.props.history.push('/add-storageLocation/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">StorageLocation List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addStorageLocation}> Add StorageLocation</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Code </th>
                                    <th> TemperatureControlled </th>
                                    <th> Capacity </th>
                                    <th> CapacityUnit </th>
                                    <th> LocationType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.storageLocations.map(
                                        storageLocation => 
                                        <tr key = {storageLocation.storageLocationId}>
                                             <td> { storageLocation.code } </td>
                                             <td> { storageLocation.temperatureControlled } </td>
                                             <td> { storageLocation.capacity } </td>
                                             <td> { storageLocation.capacityUnit } </td>
                                             <td> { storageLocation.locationType } </td>
                                             <td>
                                                 <button onClick={ () => this.editStorageLocation(storageLocation.storageLocationId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteStorageLocation(storageLocation.storageLocationId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewStorageLocation(storageLocation.storageLocationId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListStorageLocationComponent
