import React, { Component } from 'react'
import ReservationService from '../services/ReservationService';

class UpdateReservationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                referenceNumber: '',
                reservedQuantity: '',
                promisedDate: '',
                reservationStatus: '',
                reservationType: ''
        }
        this.updateReservation = this.updateReservation.bind(this);

        this.changereferenceNumberHandler = this.changereferenceNumberHandler.bind(this);
        this.changereservedQuantityHandler = this.changereservedQuantityHandler.bind(this);
        this.changepromisedDateHandler = this.changepromisedDateHandler.bind(this);
        this.changeReservationStatusHandler = this.changeReservationStatusHandler.bind(this);
        this.changeReservationTypeHandler = this.changeReservationTypeHandler.bind(this);
    }

    componentDidMount(){
        ReservationService.getReservationById(this.state.id).then( (res) =>{
            let reservation = res.data;
            this.setState({
                referenceNumber: reservation.referenceNumber,
                reservedQuantity: reservation.reservedQuantity,
                promisedDate: reservation.promisedDate,
                reservationStatus: reservation.reservationStatus,
                reservationType: reservation.reservationType
            });
        });
    }

    updateReservation = (e) => {
        e.preventDefault();
        let reservation = {
            reservationId: this.state.id,
            referenceNumber: this.state.referenceNumber,
            reservedQuantity: this.state.reservedQuantity,
            promisedDate: this.state.promisedDate,
            reservationStatus: this.state.reservationStatus,
            reservationType: this.state.reservationType
        };
        console.log('reservation => ' + JSON.stringify(reservation));
        console.log('id => ' + JSON.stringify(this.state.id));
        ReservationService.updateReservation(reservation).then( res => {
            this.props.history.push('/reservations');
        });
    }

    changereferenceNumberHandler= (event) => {
        this.setState({referenceNumber: event.target.value});
    }
    changereservedQuantityHandler= (event) => {
        this.setState({reservedQuantity: event.target.value});
    }
    changepromisedDateHandler= (event) => {
        this.setState({promisedDate: event.target.value});
    }
    changeReservationStatusHandler= (event) => {
        this.setState({reservationStatus: event.target.value});
    }
    changeReservationTypeHandler= (event) => {
        this.setState({reservationType: event.target.value});
    }

    cancel(){
        this.props.history.push('/reservations');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Reservation</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> referenceNumber: </label>
                                                <input placeholder="referenceNumber" name="referenceNumber" className="form-control" value={this.state.referenceNumber} onChange={this.changereferenceNumberHandler}/>

                                            <label> reservedQuantity: </label>
                                                <input placeholder="reservedQuantity" name="reservedQuantity" className="form-control" value={this.state.reservedQuantity} onChange={this.changereservedQuantityHandler}/>

                                            <label> promisedDate: </label>
                                                <input type="date" placeholder="promisedDate" name="promisedDate" className="form-control" value={this.state.promisedDate} onChange={this.changepromisedDateHandler}/>

                                            <label> ReservationStatus: </label>
                                                <select value={this.state.reservationStatus} onChange={this.changeReservationStatusHandler}>
                      <option name="ReservationStatus" className="form-control" >
                          Draft
                      </option>
                      <option name="ReservationStatus" className="form-control" >
                          Confirmed
                      </option>
                      <option name="ReservationStatus" className="form-control" >
                          Released
                      </option>
                      <option name="ReservationStatus" className="form-control" >
                          Fulfilled
                      </option>
                      <option name="ReservationStatus" className="form-control" >
                          Cancelled
                      </option>
                      <option name="ReservationStatus" className="form-control" >
                          Expired
                      </option>
                    </select>

                                            <label> ReservationType: </label>
                                                <select value={this.state.reservationType} onChange={this.changeReservationTypeHandler}>
                      <option name="ReservationType" className="form-control" >
                          SalesOrder
                      </option>
                      <option name="ReservationType" className="form-control" >
                          WorkOrder
                      </option>
                      <option name="ReservationType" className="form-control" >
                          TransferOrder
                      </option>
                      <option name="ReservationType" className="form-control" >
                          ServiceOrder
                      </option>
                      <option name="ReservationType" className="form-control" >
                          Other
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateReservation}>Save</button>
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

export default UpdateReservationComponent
