import React, { Component } from 'react'
import InboundShipmentService from '../services/InboundShipmentService'

class ListInboundShipmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                inboundShipments: []
        }
        this.addInboundShipment = this.addInboundShipment.bind(this);
        this.editInboundShipment = this.editInboundShipment.bind(this);
        this.deleteInboundShipment = this.deleteInboundShipment.bind(this);
    }

    deleteInboundShipment(id){
        InboundShipmentService.deleteInboundShipment(id).then( res => {
            this.setState({inboundShipments: this.state.inboundShipments.filter(inboundShipment => inboundShipment.inboundShipmentId !== id)});
        });
    }
    viewInboundShipment(id){
        this.props.history.push(`/view-inboundShipment/${id}`);
    }
    editInboundShipment(id){
        this.props.history.push(`/add-inboundShipment/${id}`);
    }

    componentDidMount(){
        InboundShipmentService.getInboundShipments().then((res) => {
            this.setState({ inboundShipments: res.data});
        });
    }

    addInboundShipment(){
        this.props.history.push('/add-inboundShipment/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">InboundShipment List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addInboundShipment}> Add InboundShipment</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ShipmentNumber </th>
                                    <th> ExpectedArrivalDate </th>
                                    <th> ArrivalDate </th>
                                    <th> CarrierName </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.inboundShipments.map(
                                        inboundShipment => 
                                        <tr key = {inboundShipment.inboundShipmentId}>
                                             <td> { inboundShipment.shipmentNumber } </td>
                                             <td> { inboundShipment.expectedArrivalDate } </td>
                                             <td> { inboundShipment.arrivalDate } </td>
                                             <td> { inboundShipment.carrierName } </td>
                                             <td> { inboundShipment.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editInboundShipment(inboundShipment.inboundShipmentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteInboundShipment(inboundShipment.inboundShipmentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewInboundShipment(inboundShipment.inboundShipmentId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListInboundShipmentComponent
