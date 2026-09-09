import React, { Component } from 'react'
import AircraftService from '../services/AircraftService'

class ListAircraftComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                aircrafts: []
        }
        this.addAircraft = this.addAircraft.bind(this);
        this.editAircraft = this.editAircraft.bind(this);
        this.deleteAircraft = this.deleteAircraft.bind(this);
    }

    deleteAircraft(id){
        AircraftService.deleteAircraft(id).then( res => {
            this.setState({aircrafts: this.state.aircrafts.filter(aircraft => aircraft.aircraftId !== id)});
        });
    }
    viewAircraft(id){
        this.props.history.push(`/view-aircraft/${id}`);
    }
    editAircraft(id){
        this.props.history.push(`/add-aircraft/${id}`);
    }

    componentDidMount(){
        AircraftService.getAircrafts().then((res) => {
            this.setState({ aircrafts: res.data});
        });
    }

    addAircraft(){
        this.props.history.push('/add-aircraft/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Aircraft List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAircraft}> Add Aircraft</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Msn </th>
                                    <th> DeliveryDate </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.aircrafts.map(
                                        aircraft => 
                                        <tr key = {aircraft.aircraftId}>
                                             <td> { aircraft.msn } </td>
                                             <td> { aircraft.deliveryDate } </td>
                                             <td>
                                                 <button onClick={ () => this.editAircraft(aircraft.aircraftId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAircraft(aircraft.aircraftId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAircraft(aircraft.aircraftId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAircraftComponent
