import React, { Component } from 'react'
import LocationService from '../services/LocationService';

class CreateLocationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                locationCode: '',
                description: '',
                locationType: ''
        }
        this.changelocationCodeHandler = this.changelocationCodeHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changeLocationTypeHandler = this.changeLocationTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            LocationService.getLocationById(this.state.id).then( (res) =>{
                let location = res.data;
                this.setState({
                    locationCode: location.locationCode,
                    description: location.description,
                    locationType: location.locationType
                });
            });
        }        
    }
    saveOrUpdateLocation = (e) => {
        e.preventDefault();
        let location = {
                locationId: this.state.id,
                locationCode: this.state.locationCode,
                description: this.state.description,
                locationType: this.state.locationType
            };
        console.log('location => ' + JSON.stringify(location));

        // step 5
        if(this.state.id === '_add'){
            location.locationId=''
            LocationService.createLocation(location).then(res =>{
                this.props.history.push('/locations');
            });
        }else{
            LocationService.updateLocation(location).then( res => {
                this.props.history.push('/locations');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Location</h3>
        }else{
            return <h3 className="text-center">Update Location</h3>
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
                                            <label> locationCode:&emsp; </label>
                                                <input placeholder="locationCode" name="locationCode" className="form-control" value={this.state.locationCode} onChange={this.changelocationCodeHandler}/>

                                            <label> description:&emsp; </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> LocationType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateLocation}>Save</button>
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

export default CreateLocationComponent
