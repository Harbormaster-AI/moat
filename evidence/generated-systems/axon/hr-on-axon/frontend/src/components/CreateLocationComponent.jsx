import React, { Component } from 'react'
import LocationService from '../services/LocationService';

class CreateLocationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                address: '',
                timezone: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
        this.changetimezoneHandler = this.changetimezoneHandler.bind(this);
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
                    name: location.name,
                    address: location.address,
                    timezone: location.timezone
                });
            });
        }        
    }
    saveOrUpdateLocation = (e) => {
        e.preventDefault();
        let location = {
                locationId: this.state.id,
                name: this.state.name,
                address: this.state.address,
                timezone: this.state.timezone
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> address:&emsp; </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                            <label> timezone:&emsp; </label>
                                                <input placeholder="timezone" name="timezone" className="form-control" value={this.state.timezone} onChange={this.changetimezoneHandler}/>

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
