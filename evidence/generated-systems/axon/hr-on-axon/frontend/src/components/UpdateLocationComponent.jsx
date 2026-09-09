import React, { Component } from 'react'
import LocationService from '../services/LocationService';

class UpdateLocationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                address: '',
                timezone: ''
        }
        this.updateLocation = this.updateLocation.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
        this.changetimezoneHandler = this.changetimezoneHandler.bind(this);
    }

    componentDidMount(){
        LocationService.getLocationById(this.state.id).then( (res) =>{
            let location = res.data;
            this.setState({
                name: location.name,
                address: location.address,
                timezone: location.timezone
            });
        });
    }

    updateLocation = (e) => {
        e.preventDefault();
        let location = {
            locationId: this.state.id,
            name: this.state.name,
            address: this.state.address,
            timezone: this.state.timezone
        };
        console.log('location => ' + JSON.stringify(location));
        console.log('id => ' + JSON.stringify(this.state.id));
        LocationService.updateLocation(location).then( res => {
            this.props.history.push('/locations');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeaddressHandler= (event) => {
        this.setState({address: event.target.value});
    }
    changetimezoneHandler= (event) => {
        this.setState({timezone: event.target.value});
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
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> address: </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                            <label> timezone: </label>
                                                <input placeholder="timezone" name="timezone" className="form-control" value={this.state.timezone} onChange={this.changetimezoneHandler}/>

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
