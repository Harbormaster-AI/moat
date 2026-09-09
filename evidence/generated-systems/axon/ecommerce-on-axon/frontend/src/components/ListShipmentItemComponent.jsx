import React, { Component } from 'react'
import ShipmentItemService from '../services/ShipmentItemService'

class ListShipmentItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                shipmentItems: []
        }
        this.addShipmentItem = this.addShipmentItem.bind(this);
        this.editShipmentItem = this.editShipmentItem.bind(this);
        this.deleteShipmentItem = this.deleteShipmentItem.bind(this);
    }

    deleteShipmentItem(id){
        ShipmentItemService.deleteShipmentItem(id).then( res => {
            this.setState({shipmentItems: this.state.shipmentItems.filter(shipmentItem => shipmentItem.shipmentItemId !== id)});
        });
    }
    viewShipmentItem(id){
        this.props.history.push(`/view-shipmentItem/${id}`);
    }
    editShipmentItem(id){
        this.props.history.push(`/add-shipmentItem/${id}`);
    }

    componentDidMount(){
        ShipmentItemService.getShipmentItems().then((res) => {
            this.setState({ shipmentItems: res.data});
        });
    }

    addShipmentItem(){
        this.props.history.push('/add-shipmentItem/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ShipmentItem List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addShipmentItem}> Add ShipmentItem</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Quantity </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.shipmentItems.map(
                                        shipmentItem => 
                                        <tr key = {shipmentItem.shipmentItemId}>
                                             <td> { shipmentItem.quantity } </td>
                                             <td>
                                                 <button onClick={ () => this.editShipmentItem(shipmentItem.shipmentItemId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteShipmentItem(shipmentItem.shipmentItemId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewShipmentItem(shipmentItem.shipmentItemId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListShipmentItemComponent
