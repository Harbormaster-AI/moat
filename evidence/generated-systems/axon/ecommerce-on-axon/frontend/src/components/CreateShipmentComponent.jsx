import React, { Component } from 'react'
import ShipmentService from '../services/ShipmentService';

class CreateShipmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                shipmentNumber: '',
                shippedDate: '',
                deliveredDate: '',
                trackingNumber: '',
                shippingAddress: '',
                status: '',
                carrier: ''
        }
        this.changeshipmentNumberHandler = this.changeshipmentNumberHandler.bind(this);
        this.changeshippedDateHandler = this.changeshippedDateHandler.bind(this);
        this.changedeliveredDateHandler = this.changedeliveredDateHandler.bind(this);
        this.changetrackingNumberHandler = this.changetrackingNumberHandler.bind(this);
        this.changeshippingAddressHandler = this.changeshippingAddressHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changeCarrierHandler = this.changeCarrierHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateShipment = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            shipment.shipmentId=''
            ShipmentService.createShipment(shipment).then(res =>{
                this.props.history.push('/shipments');
            });
        }else{
            ShipmentService.updateShipment(shipment).then( res => {
                this.props.history.push('/shipments');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Shipment</h3>
        }else{
            return <h3 className="text-center">Update Shipment</h3>
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
                                            <label> shipmentNumber:&emsp; </label>
                                                <input placeholder="shipmentNumber" name="shipmentNumber" className="form-control" value={this.state.shipmentNumber} onChange={this.changeshipmentNumberHandler}/>

                                            <label> shippedDate:&emsp; </label>
                                                <input type="date" placeholder="shippedDate" name="shippedDate" className="form-control" value={this.state.shippedDate} onChange={this.changeshippedDateHandler}/>

                                            <label> deliveredDate:&emsp; </label>
                                                <input type="date" placeholder="deliveredDate" name="deliveredDate" className="form-control" value={this.state.deliveredDate} onChange={this.changedeliveredDateHandler}/>

                                            <label> trackingNumber:&emsp; </label>
                                                <input placeholder="trackingNumber" name="trackingNumber" className="form-control" value={this.state.trackingNumber} onChange={this.changetrackingNumberHandler}/>

                                            <label> shippingAddress:&emsp; </label>
                                                <input placeholder="shippingAddress" name="shippingAddress" className="form-control" value={this.state.shippingAddress} onChange={this.changeshippingAddressHandler}/>

                                            <label> Status:&emsp; </label>
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

                                            <label> Carrier:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateShipment}>Save</button>
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

export default CreateShipmentComponent
