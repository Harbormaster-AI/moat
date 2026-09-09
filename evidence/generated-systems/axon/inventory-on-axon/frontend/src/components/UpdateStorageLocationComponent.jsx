import React, { Component } from 'react'
import StorageLocationService from '../services/StorageLocationService';

class UpdateStorageLocationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                code: '',
                temperatureControlled: '',
                capacity: '',
                capacityUnit: '',
                locationType: ''
        }
        this.updateStorageLocation = this.updateStorageLocation.bind(this);

        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changetemperatureControlledHandler = this.changetemperatureControlledHandler.bind(this);
        this.changecapacityHandler = this.changecapacityHandler.bind(this);
        this.changecapacityUnitHandler = this.changecapacityUnitHandler.bind(this);
        this.changeLocationTypeHandler = this.changeLocationTypeHandler.bind(this);
    }

    componentDidMount(){
        StorageLocationService.getStorageLocationById(this.state.id).then( (res) =>{
            let storageLocation = res.data;
            this.setState({
                code: storageLocation.code,
                temperatureControlled: storageLocation.temperatureControlled,
                capacity: storageLocation.capacity,
                capacityUnit: storageLocation.capacityUnit,
                locationType: storageLocation.locationType
            });
        });
    }

    updateStorageLocation = (e) => {
        e.preventDefault();
        let storageLocation = {
            storageLocationId: this.state.id,
            code: this.state.code,
            temperatureControlled: this.state.temperatureControlled,
            capacity: this.state.capacity,
            capacityUnit: this.state.capacityUnit,
            locationType: this.state.locationType
        };
        console.log('storageLocation => ' + JSON.stringify(storageLocation));
        console.log('id => ' + JSON.stringify(this.state.id));
        StorageLocationService.updateStorageLocation(storageLocation).then( res => {
            this.props.history.push('/storageLocations');
        });
    }

    changecodeHandler= (event) => {
        this.setState({code: event.target.value});
    }
    changetemperatureControlledHandler= (event) => {
        this.setState({temperatureControlled: event.target.value});
    }
    changecapacityHandler= (event) => {
        this.setState({capacity: event.target.value});
    }
    changecapacityUnitHandler= (event) => {
        this.setState({capacityUnit: event.target.value});
    }
    changeLocationTypeHandler= (event) => {
        this.setState({locationType: event.target.value});
    }

    cancel(){
        this.props.history.push('/storageLocations');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update StorageLocation</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> code: </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> temperatureControlled: </label>
                                                <input type="checkbox" placeholder="temperatureControlled" name="temperatureControlled" className="form-control" value={this.state.temperatureControlled} onChange={this.changetemperatureControlledHandler}/>


                                            <label> capacity: </label>
                                                <input placeholder="capacity" name="capacity" className="form-control" value={this.state.capacity} onChange={this.changecapacityHandler}/>

                                            <label> capacityUnit: </label>
                                                <input placeholder="capacityUnit" name="capacityUnit" className="form-control" value={this.state.capacityUnit} onChange={this.changecapacityUnitHandler}/>

                                            <label> LocationType: </label>
                                                <select value={this.state.locationType} onChange={this.changeLocationTypeHandler}>
                      <option name="LocationType" className="form-control" >
                          Bin
                      </option>
                      <option name="LocationType" className="form-control" >
                          Bulk
                      </option>
                      <option name="LocationType" className="form-control" >
                          Staging
                      </option>
                      <option name="LocationType" className="form-control" >
                          Dock
                      </option>
                      <option name="LocationType" className="form-control" >
                          Picking
                      </option>
                      <option name="LocationType" className="form-control" >
                          Packing
                      </option>
                      <option name="LocationType" className="form-control" >
                          Quality
                      </option>
                      <option name="LocationType" className="form-control" >
                          Return
                      </option>
                      <option name="LocationType" className="form-control" >
                          ColdStorage
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateStorageLocation}>Save</button>
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

export default UpdateStorageLocationComponent
