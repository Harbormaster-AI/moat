import React, { Component } from 'react'
import ShipmentService from '../services/ShipmentService';

class UpdateShipmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                shipmentNumber: '',
                shippedDate: '',
                deliveredDate: '',
                trackingNumber: '',
                shippingAddress: '',
                status: '',
                carrier: ''
        }
        this.updateShipment = this.updateShipment.bind(this);

        this.changeshipmentNumberHandler = this.changeshipmentNumberHandler.bind(this);
        this.changeshippedDateHandler = this.changeshippedDateHandler.bind(this);
        this.changedeliveredDateHandler = this.changedeliveredDateHandler.bind(this);
        this.changetrackingNumberHandler = this.changetrackingNumberHandler.bind(this);
        this.changeshippingAddressHandler = this.changeshippingAddressHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changeCarrierHandler = this.changeCarrierHandler.bind(this);
    }

    componentDidMount(){
        ShipmentService.getShipmentById(this.state.id).then( (res) =>{
            let shipment = res.data;
            this.setState({
                shipmentNumber: shipment.shipmentNumber,
                shippedDate: shipment.shippedDate,
                deliveredDate: shipment.deliveredDate,
                trackingNumber: shipment.trackingNumber,
                shippingAddress: shipment.shippingAddress,
                status: shipment.status,
                carrier: shipment.carrier
            });
        });
    }

    updateShipment = (e) => {
        e.preventDefault();
        let shipment = {
            shipmentId: this.state.id,
            shipmentNumber: this.state.shipmentNumber,
            shippedDate: this.state.shippedDate,
            deliveredDate: this.state.deliveredDate,
            trackingNumber: this.state.trackingNumber,
            shippingAddress: this.state.shippingAddress,
            status: this.state.status,
            carrier: this.state.carrier
        };
        console.log('shipment => ' + JSON.stringify(shipment));
        console.log('id => ' + JSON.stringify(this.state.id));
        ShipmentService.updateShipment(shipment).then( res => {
            this.props.history.push('/shipments');
        });
    }

    changeshipmentNumberHandler= (event) => {
        this.setState({shipmentNumber: event.target.value});
    }
    changeshippedDateHandler= (event) => {
        this.setState({shippedDate: event.target.value});
    }
    changedeliveredDateHandler= (event) => {
        this.setState({deliveredDate: event.target.value});
    }
    changetrackingNumberHandler= (event) => {
        this.setState({trackingNumber: event.target.value});
    }
    changeshippingAddressHandler= (event) => {
        this.setState({shippingAddress: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changeCarrierHandler= (event) => {
        this.setState({carrier: event.target.value});
    }

    cancel(){
        this.props.history.push('/shipments');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Shipment</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> shipmentNumber: </label>
                                                <input placeholder="shipmentNumber" name="shipmentNumber" className="form-control" value={this.state.shipmentNumber} onChange={this.changeshipmentNumberHandler}/>

                                            <label> shippedDate: </label>
                                                <input type="date" placeholder="shippedDate" name="shippedDate" className="form-control" value={this.state.shippedDate} onChange={this.changeshippedDateHandler}/>

                                            <label> deliveredDate: </label>
                                                <input type="date" placeholder="deliveredDate" name="deliveredDate" className="form-control" value={this.state.deliveredDate} onChange={this.changedeliveredDateHandler}/>

                                            <label> trackingNumber: </label>
                                                <input placeholder="trackingNumber" name="trackingNumber" className="form-control" value={this.state.trackingNumber} onChange={this.changetrackingNumberHandler}/>

                                            <label> shippingAddress: </label>
                                                <input placeholder="shippingAddress" name="shippingAddress" className="form-control" value={this.state.shippingAddress} onChange={this.changeshippingAddressHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Pending
                      </option>
                      <option name="Status" className="form-control" >
                          Packed
                      </option>
                      <option name="Status" className="form-control" >
                          Shipped
                      </option>
                      <option name="Status" className="form-control" >
                          InTransit
                      </option>
                      <option name="Status" className="form-control" >
                          Delivered
                      </option>
                      <option name="Status" className="form-control" >
                          Delayed
                      </option>
                      <option name="Status" className="form-control" >
                          Returned
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                            <label> Carrier: </label>
                                                <select value={this.state.carrier} onChange={this.changeCarrierHandler}>
                      <option name="Carrier" className="form-control" >
                          UPS
                      </option>
                      <option name="Carrier" className="form-control" >
                          FedEx
                      </option>
                      <option name="Carrier" className="form-control" >
                          USPS
                      </option>
                      <option name="Carrier" className="form-control" >
                          DHL
                      </option>
                      <option name="Carrier" className="form-control" >
                          RoyalMail
                      </option>
                      <option name="Carrier" className="form-control" >
                          CanadaPost
                      </option>
                      <option name="Carrier" className="form-control" >
                          LocalCourier
                      </option>
                      <option name="Carrier" className="form-control" >
                          Other
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateShipment}>Save</button>
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

export default UpdateShipmentComponent
