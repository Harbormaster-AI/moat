import React, { Component } from 'react'
import TradeOrderService from '../services/TradeOrderService'

class ListTradeOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                tradeOrders: []
        }
        this.addTradeOrder = this.addTradeOrder.bind(this);
        this.editTradeOrder = this.editTradeOrder.bind(this);
        this.deleteTradeOrder = this.deleteTradeOrder.bind(this);
    }

    deleteTradeOrder(id){
        TradeOrderService.deleteTradeOrder(id).then( res => {
            this.setState({tradeOrders: this.state.tradeOrders.filter(tradeOrder => tradeOrder.tradeOrderId !== id)});
        });
    }
    viewTradeOrder(id){
        this.props.history.push(`/view-tradeOrder/${id}`);
    }
    editTradeOrder(id){
        this.props.history.push(`/add-tradeOrder/${id}`);
    }

    componentDidMount(){
        TradeOrderService.getTradeOrders().then((res) => {
            this.setState({ tradeOrders: res.data});
        });
    }

    addTradeOrder(){
        this.props.history.push('/add-tradeOrder/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">TradeOrder List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addTradeOrder}> Add TradeOrder</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> OrderId </th>
                                    <th> Quantity </th>
                                    <th> LimitPrice </th>
                                    <th> PlacedAt </th>
                                    <th> Side </th>
                                    <th> Type </th>
                                    <th> Status </th>
                                    <th> TimeInForce </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.tradeOrders.map(
                                        tradeOrder => 
                                        <tr key = {tradeOrder.tradeOrderId}>
                                             <td> { tradeOrder.orderId } </td>
                                             <td> { tradeOrder.quantity } </td>
                                             <td> { tradeOrder.limitPrice } </td>
                                             <td> { tradeOrder.placedAt } </td>
                                             <td> { tradeOrder.side } </td>
                                             <td> { tradeOrder.type } </td>
                                             <td> { tradeOrder.status } </td>
                                             <td> { tradeOrder.timeInForce } </td>
                                             <td>
                                                 <button onClick={ () => this.editTradeOrder(tradeOrder.tradeOrderId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteTradeOrder(tradeOrder.tradeOrderId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewTradeOrder(tradeOrder.tradeOrderId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListTradeOrderComponent
