import React, { Component } from 'react'
import LocationService from '../services/LocationService';

class UpdateLocationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                locationCode: '',
                description: '',
                locationType: ''
        }
        this.updateLocation = this.updateLocation.bind(this);

        this.changelocationCodeHandler = this.changelocationCodeHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changeLocationTypeHandler = this.changeLocationTypeHandler.bind(this);
    }

    componentDidMount(){
        LocationService.getLocationById(this.state.id).then( (res) =>{
            let location = res.data;
            this.setState({
                locationCode: location.locationCode,
                description: location.description,
                locationType: location.locationType
            });
        });
    }

    updateLocation = (e) => {
        e.preventDefault();
        let location = {
            locationId: this.state.id,
            locationCode: this.state.locationCode,
            description: this.state.description,
            locationType: this.state.locationType
        };
        console.log('location => ' + JSON.stringify(location));
        console.log('id => ' + JSON.stringify(this.state.id));
        LocationService.updateLocation(location).then( res => {
            this.props.history.push('/locations');
        });
    }

    changelocationCodeHandler= (event) => {
        this.setState({locationCode: event.target.value});
    }
    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }
    changeLocationTypeHandler= (event) => {
        this.setState({locationType: event.target.value});
    }

    cancel(){
        this.props.history.push('/locations');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Location</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> locationCode: </label>
                                                <input placeholder="locationCode" name="locationCode" className="form-control" value={this.state.locationCode} onChange={this.changelocationCodeHandler}/>

                                            <label> description: </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> LocationType: </label>
                                                <select value={this.state.locationType} onChange={this.changeLocationTypeHandler}>
                      <option name="LocationType" className="form-control" >
                          Bin
                      </option>
                      <option name="LocationType" className="form-control" >
                          Dock
                      </option>
                      <option name="LocationType" className="form-control" >
                          Staging
                      </option>
                      <option name="LocationType" className="form-control" >
                          QAHold
                      </option>
                      <option name="LocationType" className="form-control" >
                          Scrap
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateLocation}>Save</button>
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

export default UpdateLocationComponent
