import React, { Component } from 'react'
import ConnectedAircraftService from '../services/ConnectedAircraftService'

class ListConnectedAircraftComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                connectedAircrafts: []
        }
        this.addConnectedAircraft = this.addConnectedAircraft.bind(this);
        this.editConnectedAircraft = this.editConnectedAircraft.bind(this);
        this.deleteConnectedAircraft = this.deleteConnectedAircraft.bind(this);
    }

    deleteConnectedAircraft(id){
        ConnectedAircraftService.deleteConnectedAircraft(id).then( res => {
            this.setState({connectedAircrafts: this.state.connectedAircrafts.filter(connectedAircraft => connectedAircraft.connectedAircraftId !== id)});
        });
    }
    viewConnectedAircraft(id){
        this.props.history.push(`/view-connectedAircraft/${id}`);
    }
    editConnectedAircraft(id){
        this.props.history.push(`/add-connectedAircraft/${id}`);
    }

    componentDidMount(){
        ConnectedAircraftService.getConnectedAircrafts().then((res) => {
            this.setState({ connectedAircrafts: res.data});
        });
    }

    addConnectedAircraft(){
        this.props.history.push('/add-connectedAircraft/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ConnectedAircraft List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addConnectedAircraft}> Add ConnectedAircraft</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> CommunicationsProvider </th>
                                    <th> ConnectivityStatus </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.connectedAircrafts.map(
                                        connectedAircraft => 
                                        <tr key = {connectedAircraft.connectedAircraftId}>
                                             <td> { connectedAircraft.communicationsProvider } </td>
                                             <td> { connectedAircraft.connectivityStatus } </td>
                                             <td>
                                                 <button onClick={ () => this.editConnectedAircraft(connectedAircraft.connectedAircraftId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteConnectedAircraft(connectedAircraft.connectedAircraftId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewConnectedAircraft(connectedAircraft.connectedAircraftId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListConnectedAircraftComponent
