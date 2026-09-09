import React, { Component } from 'react'
import ConnectedAircraftService from '../services/ConnectedAircraftService';

class UpdateConnectedAircraftComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                communicationsProvider: '',
                connectivityStatus: ''
        }
        this.updateConnectedAircraft = this.updateConnectedAircraft.bind(this);

        this.changecommunicationsProviderHandler = this.changecommunicationsProviderHandler.bind(this);
        this.changeConnectivityStatusHandler = this.changeConnectivityStatusHandler.bind(this);
    }

    componentDidMount(){
        ConnectedAircraftService.getConnectedAircraftById(this.state.id).then( (res) =>{
            let connectedAircraft = res.data;
            this.setState({
                communicationsProvider: connectedAircraft.communicationsProvider,
                connectivityStatus: connectedAircraft.connectivityStatus
            });
        });
    }

    updateConnectedAircraft = (e) => {
        e.preventDefault();
        let connectedAircraft = {
            connectedAircraftId: this.state.id,
            communicationsProvider: this.state.communicationsProvider,
            connectivityStatus: this.state.connectivityStatus
        };
        console.log('connectedAircraft => ' + JSON.stringify(connectedAircraft));
        console.log('id => ' + JSON.stringify(this.state.id));
        ConnectedAircraftService.updateConnectedAircraft(connectedAircraft).then( res => {
            this.props.history.push('/connectedAircrafts');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ConnectedAircraft</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> communicationsProvider: </label>
                                                <input placeholder="communicationsProvider" name="communicationsProvider" className="form-control" value={this.state.communicationsProvider} onChange={this.changecommunicationsProviderHandler}/>

                                            <label> ConnectivityStatus: </label>
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
                                        <button className="btn btn-success" onClick={this.updateConnectedAircraft}>Save</button>
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

export default UpdateConnectedAircraftComponent
