import React, { Component } from 'react'
import TradeService from '../services/TradeService'

class ListTradeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                trades: []
        }
        this.addTrade = this.addTrade.bind(this);
        this.editTrade = this.editTrade.bind(this);
        this.deleteTrade = this.deleteTrade.bind(this);
    }

    deleteTrade(id){
        TradeService.deleteTrade(id).then( res => {
            this.setState({trades: this.state.trades.filter(trade => trade.tradeId !== id)});
        });
    }
    viewTrade(id){
        this.props.history.push(`/view-trade/${id}`);
    }
    editTrade(id){
        this.props.history.push(`/add-trade/${id}`);
    }

    componentDidMount(){
        TradeService.getTrades().then((res) => {
            this.setState({ trades: res.data});
        });
    }

    addTrade(){
        this.props.history.push('/add-trade/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Trade List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addTrade}> Add Trade</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ExecutedAt </th>
                                    <th> Quantity </th>
                                    <th> Price </th>
                                    <th> Fees </th>
                                    <th> SettlementDate </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.trades.map(
                                        trade => 
                                        <tr key = {trade.tradeId}>
                                             <td> { trade.executedAt } </td>
                                             <td> { trade.quantity } </td>
                                             <td> { trade.price } </td>
                                             <td> { trade.fees } </td>
                                             <td> { trade.settlementDate } </td>
                                             <td>
                                                 <button onClick={ () => this.editTrade(trade.tradeId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteTrade(trade.tradeId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewTrade(trade.tradeId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListTradeComponent
