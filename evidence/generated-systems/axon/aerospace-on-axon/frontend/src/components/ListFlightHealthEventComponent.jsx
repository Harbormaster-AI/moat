import React, { Component } from 'react'
import FlightHealthEventService from '../services/FlightHealthEventService'

class ListFlightHealthEventComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                flightHealthEvents: []
        }
        this.addFlightHealthEvent = this.addFlightHealthEvent.bind(this);
        this.editFlightHealthEvent = this.editFlightHealthEvent.bind(this);
        this.deleteFlightHealthEvent = this.deleteFlightHealthEvent.bind(this);
    }

    deleteFlightHealthEvent(id){
        FlightHealthEventService.deleteFlightHealthEvent(id).then( res => {
            this.setState({flightHealthEvents: this.state.flightHealthEvents.filter(flightHealthEvent => flightHealthEvent.flightHealthEventId !== id)});
        });
    }
    viewFlightHealthEvent(id){
        this.props.history.push(`/view-flightHealthEvent/${id}`);
    }
    editFlightHealthEvent(id){
        this.props.history.push(`/add-flightHealthEvent/${id}`);
    }

    componentDidMount(){
        FlightHealthEventService.getFlightHealthEvents().then((res) => {
            this.setState({ flightHealthEvents: res.data});
        });
    }

    addFlightHealthEvent(){
        this.props.history.push('/add-flightHealthEvent/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">FlightHealthEvent List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addFlightHealthEvent}> Add FlightHealthEvent</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> EventCode </th>
                                    <th> Severity </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.flightHealthEvents.map(
                                        flightHealthEvent => 
                                        <tr key = {flightHealthEvent.flightHealthEventId}>
                                             <td> { flightHealthEvent.eventCode } </td>
                                             <td> { flightHealthEvent.severity } </td>
                                             <td>
                                                 <button onClick={ () => this.editFlightHealthEvent(flightHealthEvent.flightHealthEventId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteFlightHealthEvent(flightHealthEvent.flightHealthEventId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewFlightHealthEvent(flightHealthEvent.flightHealthEventId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListFlightHealthEventComponent
