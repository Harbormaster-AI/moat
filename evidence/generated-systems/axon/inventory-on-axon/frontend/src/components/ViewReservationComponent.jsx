import React, { Component } from 'react'
import ReservationService from '../services/ReservationService'

class ViewReservationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            reservation: {}
        }
    }

    componentDidMount(){
        ReservationService.getReservationById(this.state.id).then( res => {
            this.setState({reservation: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Reservation Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> referenceNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.reservation.referenceNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> reservedQuantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.reservation.reservedQuantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> promisedDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.reservation.promisedDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ReservationStatus:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.reservation.reservationStatus }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ReservationType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.reservation.reservationType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewReservationComponent
