import React, { Component } from 'react'
import ShipmentService from '../services/ShipmentService'

class ListShipmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                shipments: []
        }
        this.addShipment = this.addShipment.bind(this);
        this.editShipment = this.editShipment.bind(this);
        this.deleteShipment = this.deleteShipment.bind(this);
    }

    deleteShipment(id){
        ShipmentService.deleteShipment(id).then( res => {
            this.setState({shipments: this.state.shipments.filter(shipment => shipment.shipmentId !== id)});
        });
    }
    viewShipment(id){
        this.props.history.push(`/view-shipment/${id}`);
    }
    editShipment(id){
        this.props.history.push(`/add-shipment/${id}`);
    }

    componentDidMount(){
        ShipmentService.getShipments().then((res) => {
            this.setState({ shipments: res.data});
        });
    }

    addShipment(){
        this.props.history.push('/add-shipment/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Shipment List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addShipment}> Add Shipment</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ShipmentNumber </th>
                                    <th> ShippedDate </th>
                                    <th> DeliveredDate </th>
                                    <th> TrackingNumber </th>
                                    <th> ShippingAddress </th>
                                    <th> Status </th>
                                    <th> Carrier </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.shipments.map(
                                        shipment => 
                                        <tr key = {shipment.shipmentId}>
                                             <td> { shipment.shipmentNumber } </td>
                                             <td> { shipment.shippedDate } </td>
                                             <td> { shipment.deliveredDate } </td>
                                             <td> { shipment.trackingNumber } </td>
                                             <td> { shipment.shippingAddress } </td>
                                             <td> { shipment.status } </td>
                                             <td> { shipment.carrier } </td>
                                             <td>
                                                 <button onClick={ () => this.editShipment(shipment.shipmentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteShipment(shipment.shipmentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewShipment(shipment.shipmentId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListShipmentComponent
