import React, { Component } from 'react'
import TradeService from '../services/TradeService';

class UpdateTradeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                executedAt: '',
                quantity: '',
                price: '',
                fees: '',
                settlementDate: ''
        }
        this.updateTrade = this.updateTrade.bind(this);

        this.changeexecutedAtHandler = this.changeexecutedAtHandler.bind(this);
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changepriceHandler = this.changepriceHandler.bind(this);
        this.changefeesHandler = this.changefeesHandler.bind(this);
        this.changesettlementDateHandler = this.changesettlementDateHandler.bind(this);
    }

    componentDidMount(){
        TradeService.getTradeById(this.state.id).then( (res) =>{
            let trade = res.data;
            this.setState({
                executedAt: trade.executedAt,
                quantity: trade.quantity,
                price: trade.price,
                fees: trade.fees,
                settlementDate: trade.settlementDate
            });
        });
    }

    updateTrade = (e) => {
        e.preventDefault();
        let trade = {
            tradeId: this.state.id,
            executedAt: this.state.executedAt,
            quantity: this.state.quantity,
            price: this.state.price,
            fees: this.state.fees,
            settlementDate: this.state.settlementDate
        };
        console.log('trade => ' + JSON.stringify(trade));
        console.log('id => ' + JSON.stringify(this.state.id));
        TradeService.updateTrade(trade).then( res => {
            this.props.history.push('/trades');
        });
    }

    changeexecutedAtHandler= (event) => {
        this.setState({executedAt: event.target.value});
    }
    changequantityHandler= (event) => {
        this.setState({quantity: event.target.value});
    }
    changepriceHandler= (event) => {
        this.setState({price: event.target.value});
    }
    changefeesHandler= (event) => {
        this.setState({fees: event.target.value});
    }
    changesettlementDateHandler= (event) => {
        this.setState({settlementDate: event.target.value});
    }

    cancel(){
        this.props.history.push('/trades');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Trade</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> executedAt: </label>
                                                <input type="time" placeholder="executedAt" name="executedAt" className="form-control" value={this.state.executedAt} onChange={this.changeexecutedAtHandler}/>

                                            <label> quantity: </label>
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> price: </label>
                                                <input placeholder="price" name="price" className="form-control" value={this.state.price} onChange={this.changepriceHandler}/>

                                            <label> fees: </label>
                                                <input placeholder="fees" name="fees" className="form-control" value={this.state.fees} onChange={this.changefeesHandler}/>

                                            <label> settlementDate: </label>
                                                <input type="date" placeholder="settlementDate" name="settlementDate" className="form-control" value={this.state.settlementDate} onChange={this.changesettlementDateHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateTrade}>Save</button>
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

export default UpdateTradeComponent
