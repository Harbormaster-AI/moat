import React, { Component } from 'react'
import ReservationService from '../services/ReservationService';

class CreateReservationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                referenceNumber: '',
                reservedQuantity: '',
                promisedDate: '',
                reservationStatus: '',
                reservationType: ''
        }
        this.changereferenceNumberHandler = this.changereferenceNumberHandler.bind(this);
        this.changereservedQuantityHandler = this.changereservedQuantityHandler.bind(this);
        this.changepromisedDateHandler = this.changepromisedDateHandler.bind(this);
        this.changeReservationStatusHandler = this.changeReservationStatusHandler.bind(this);
        this.changeReservationTypeHandler = this.changeReservationTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateReservation = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            reservation.reservationId=''
            ReservationService.createReservation(reservation).then(res =>{
                this.props.history.push('/reservations');
            });
        }else{
            ReservationService.updateReservation(reservation).then( res => {
                this.props.history.push('/reservations');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Reservation</h3>
        }else{
            return <h3 className="text-center">Update Reservation</h3>
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
                                            <label> referenceNumber:&emsp; </label>
                                                <input placeholder="referenceNumber" name="referenceNumber" className="form-control" value={this.state.referenceNumber} onChange={this.changereferenceNumberHandler}/>

                                            <label> reservedQuantity:&emsp; </label>
                                                <input placeholder="reservedQuantity" name="reservedQuantity" className="form-control" value={this.state.reservedQuantity} onChange={this.changereservedQuantityHandler}/>

                                            <label> promisedDate:&emsp; </label>
                                                <input type="date" placeholder="promisedDate" name="promisedDate" className="form-control" value={this.state.promisedDate} onChange={this.changepromisedDateHandler}/>

                                            <label> ReservationStatus:&emsp; </label>
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

                                            <label> ReservationType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateReservation}>Save</button>
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

export default CreateReservationComponent
