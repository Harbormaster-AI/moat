import React, { Component } from 'react'
import ConnectedAircraftService from '../services/ConnectedAircraftService'

class ViewConnectedAircraftComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            connectedAircraft: {}
        }
    }

    componentDidMount(){
        ConnectedAircraftService.getConnectedAircraftById(this.state.id).then( res => {
            this.setState({connectedAircraft: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ConnectedAircraft Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> communicationsProvider:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.connectedAircraft.communicationsProvider }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ConnectivityStatus:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.connectedAircraft.connectivityStatus }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewConnectedAircraftComponent
