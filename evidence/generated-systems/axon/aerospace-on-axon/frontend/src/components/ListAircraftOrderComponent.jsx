import React, { Component } from 'react'
import AircraftOrderService from '../services/AircraftOrderService'

class ListAircraftOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                aircraftOrders: []
        }
        this.addAircraftOrder = this.addAircraftOrder.bind(this);
        this.editAircraftOrder = this.editAircraftOrder.bind(this);
        this.deleteAircraftOrder = this.deleteAircraftOrder.bind(this);
    }

    deleteAircraftOrder(id){
        AircraftOrderService.deleteAircraftOrder(id).then( res => {
            this.setState({aircraftOrders: this.state.aircraftOrders.filter(aircraftOrder => aircraftOrder.aircraftOrderId !== id)});
        });
    }
    viewAircraftOrder(id){
        this.props.history.push(`/view-aircraftOrder/${id}`);
    }
    editAircraftOrder(id){
        this.props.history.push(`/add-aircraftOrder/${id}`);
    }

    componentDidMount(){
        AircraftOrderService.getAircraftOrders().then((res) => {
            this.setState({ aircraftOrders: res.data});
        });
    }

    addAircraftOrder(){
        this.props.history.push('/add-aircraftOrder/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">AircraftOrder List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAircraftOrder}> Add AircraftOrder</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> OrderNumber </th>
                                    <th> TotalAmount </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.aircraftOrders.map(
                                        aircraftOrder => 
                                        <tr key = {aircraftOrder.aircraftOrderId}>
                                             <td> { aircraftOrder.orderNumber } </td>
                                             <td> { aircraftOrder.totalAmount } </td>
                                             <td> { aircraftOrder.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editAircraftOrder(aircraftOrder.aircraftOrderId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAircraftOrder(aircraftOrder.aircraftOrderId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAircraftOrder(aircraftOrder.aircraftOrderId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAircraftOrderComponent
