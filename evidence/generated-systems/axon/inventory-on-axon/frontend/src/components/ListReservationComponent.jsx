import React, { Component } from 'react'
import ReservationService from '../services/ReservationService'

class ListReservationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                reservations: []
        }
        this.addReservation = this.addReservation.bind(this);
        this.editReservation = this.editReservation.bind(this);
        this.deleteReservation = this.deleteReservation.bind(this);
    }

    deleteReservation(id){
        ReservationService.deleteReservation(id).then( res => {
            this.setState({reservations: this.state.reservations.filter(reservation => reservation.reservationId !== id)});
        });
    }
    viewReservation(id){
        this.props.history.push(`/view-reservation/${id}`);
    }
    editReservation(id){
        this.props.history.push(`/add-reservation/${id}`);
    }

    componentDidMount(){
        ReservationService.getReservations().then((res) => {
            this.setState({ reservations: res.data});
        });
    }

    addReservation(){
        this.props.history.push('/add-reservation/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Reservation List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addReservation}> Add Reservation</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ReferenceNumber </th>
                                    <th> ReservedQuantity </th>
                                    <th> PromisedDate </th>
                                    <th> ReservationStatus </th>
                                    <th> ReservationType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.reservations.map(
                                        reservation => 
                                        <tr key = {reservation.reservationId}>
                                             <td> { reservation.referenceNumber } </td>
                                             <td> { reservation.reservedQuantity } </td>
                                             <td> { reservation.promisedDate } </td>
                                             <td> { reservation.reservationStatus } </td>
                                             <td> { reservation.reservationType } </td>
                                             <td>
                                                 <button onClick={ () => this.editReservation(reservation.reservationId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteReservation(reservation.reservationId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewReservation(reservation.reservationId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListReservationComponent
