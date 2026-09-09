import React, { Component } from 'react'
import InboundShipmentService from '../services/InboundShipmentService';

class UpdateInboundShipmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                shipmentNumber: '',
                expectedArrivalDate: '',
                arrivalDate: '',
                carrierName: '',
                status: ''
        }
        this.updateInboundShipment = this.updateInboundShipment.bind(this);

        this.changeshipmentNumberHandler = this.changeshipmentNumberHandler.bind(this);
        this.changeexpectedArrivalDateHandler = this.changeexpectedArrivalDateHandler.bind(this);
        this.changearrivalDateHandler = this.changearrivalDateHandler.bind(this);
        this.changecarrierNameHandler = this.changecarrierNameHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        InboundShipmentService.getInboundShipmentById(this.state.id).then( (res) =>{
            let inboundShipment = res.data;
            this.setState({
                shipmentNumber: inboundShipment.shipmentNumber,
                expectedArrivalDate: inboundShipment.expectedArrivalDate,
                arrivalDate: inboundShipment.arrivalDate,
                carrierName: inboundShipment.carrierName,
                status: inboundShipment.status
            });
        });
    }

    updateInboundShipment = (e) => {
        e.preventDefault();
        let inboundShipment = {
            inboundShipmentId: this.state.id,
            shipmentNumber: this.state.shipmentNumber,
            expectedArrivalDate: this.state.expectedArrivalDate,
            arrivalDate: this.state.arrivalDate,
            carrierName: this.state.carrierName,
            status: this.state.status
        };
        console.log('inboundShipment => ' + JSON.stringify(inboundShipment));
        console.log('id => ' + JSON.stringify(this.state.id));
        InboundShipmentService.updateInboundShipment(inboundShipment).then( res => {
            this.props.history.push('/inboundShipments');
        });
    }

    changeshipmentNumberHandler= (event) => {
        this.setState({shipmentNumber: event.target.value});
    }
    changeexpectedArrivalDateHandler= (event) => {
        this.setState({expectedArrivalDate: event.target.value});
    }
    changearrivalDateHandler= (event) => {
        this.setState({arrivalDate: event.target.value});
    }
    changecarrierNameHandler= (event) => {
        this.setState({carrierName: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/inboundShipments');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update InboundShipment</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> shipmentNumber: </label>
                                                <input placeholder="shipmentNumber" name="shipmentNumber" className="form-control" value={this.state.shipmentNumber} onChange={this.changeshipmentNumberHandler}/>

                                            <label> expectedArrivalDate: </label>
                                                <input type="date" placeholder="expectedArrivalDate" name="expectedArrivalDate" className="form-control" value={this.state.expectedArrivalDate} onChange={this.changeexpectedArrivalDateHandler}/>

                                            <label> arrivalDate: </label>
                                                <input type="date" placeholder="arrivalDate" name="arrivalDate" className="form-control" value={this.state.arrivalDate} onChange={this.changearrivalDateHandler}/>

                                            <label> carrierName: </label>
                                                <input placeholder="carrierName" name="carrierName" className="form-control" value={this.state.carrierName} onChange={this.changecarrierNameHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Planned
                      </option>
                      <option name="Status" className="form-control" >
                          Arrived
                      </option>
                      <option name="Status" className="form-control" >
                          Received
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateInboundShipment}>Save</button>
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

export default UpdateInboundShipmentComponent
