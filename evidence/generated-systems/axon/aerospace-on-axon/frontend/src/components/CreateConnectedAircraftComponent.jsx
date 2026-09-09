import React, { Component } from 'react'
import ConnectedAircraftService from '../services/ConnectedAircraftService';

class CreateConnectedAircraftComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                communicationsProvider: '',
                connectivityStatus: ''
        }
        this.changecommunicationsProviderHandler = this.changecommunicationsProviderHandler.bind(this);
        this.changeConnectivityStatusHandler = this.changeConnectivityStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ConnectedAircraftService.getConnectedAircraftById(this.state.id).then( (res) =>{
                let connectedAircraft = res.data;
                this.setState({
                    communicationsProvider: connectedAircraft.communicationsProvider,
                    connectivityStatus: connectedAircraft.connectivityStatus
                });
            });
        }        
    }
    saveOrUpdateConnectedAircraft = (e) => {
        e.preventDefault();
        let connectedAircraft = {
                connectedAircraftId: this.state.id,
                communicationsProvider: this.state.communicationsProvider,
                connectivityStatus: this.state.connectivityStatus
            };
        console.log('connectedAircraft => ' + JSON.stringify(connectedAircraft));

        // step 5
        if(this.state.id === '_add'){
            connectedAircraft.connectedAircraftId=''
            ConnectedAircraftService.createConnectedAircraft(connectedAircraft).then(res =>{
                this.props.history.push('/connectedAircrafts');
            });
        }else{
            ConnectedAircraftService.updateConnectedAircraft(connectedAircraft).then( res => {
                this.props.history.push('/connectedAircrafts');
            });
        }
    }
    
    changecommunicationsProviderHandler= (event) => {
        this.setState({communicationsProvider: event.target.value});
    }
    changeConnectivityStatusHandler= (event) => {
        this.setState({connectivityStatus: event.target.value});
    }

    cancel(){
        this.props.history.push('/connectedAircrafts');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ConnectedAircraft</h3>
        }else{
            return <h3 className="text-center">Update ConnectedAircraft</h3>
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
                                            <label> communicationsProvider:&emsp; </label>
                                                <input placeholder="communicationsProvider" name="communicationsProvider" className="form-control" value={this.state.communicationsProvider} onChange={this.changecommunicationsProviderHandler}/>

                                            <label> ConnectivityStatus:&emsp; </label>
                                                <select value={this.state.connectivityStatus} onChange={this.changeConnectivityStatusHandler}>
                      <option name="ConnectivityStatus" className="form-control" >
                          Offline
                      </option>
                      <option name="ConnectivityStatus" className="form-control" >
                          Online
                      </option>
                      <option name="ConnectivityStatus" className="form-control" >
                          Degraded
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateConnectedAircraft}>Save</button>
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

export default CreateConnectedAircraftComponent
