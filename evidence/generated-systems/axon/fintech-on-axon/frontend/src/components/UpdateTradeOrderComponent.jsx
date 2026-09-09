import React, { Component } from 'react'
import TradeOrderService from '../services/TradeOrderService';

class UpdateTradeOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                orderId: '',
                quantity: '',
                limitPrice: '',
                placedAt: '',
                side: '',
                type: '',
                status: '',
                timeInForce: ''
        }
        this.updateTradeOrder = this.updateTradeOrder.bind(this);

        this.changeorderIdHandler = this.changeorderIdHandler.bind(this);
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changelimitPriceHandler = this.changelimitPriceHandler.bind(this);
        this.changeplacedAtHandler = this.changeplacedAtHandler.bind(this);
        this.changeSideHandler = this.changeSideHandler.bind(this);
        this.changeTypeHandler = this.changeTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changeTimeInForceHandler = this.changeTimeInForceHandler.bind(this);
    }

    componentDidMount(){
        TradeOrderService.getTradeOrderById(this.state.id).then( (res) =>{
            let tradeOrder = res.data;
            this.setState({
                orderId: tradeOrder.orderId,
                quantity: tradeOrder.quantity,
                limitPrice: tradeOrder.limitPrice,
                placedAt: tradeOrder.placedAt,
                side: tradeOrder.side,
                type: tradeOrder.type,
                status: tradeOrder.status,
                timeInForce: tradeOrder.timeInForce
            });
        });
    }

    updateTradeOrder = (e) => {
        e.preventDefault();
        let tradeOrder = {
            tradeOrderId: this.state.id,
            orderId: this.state.orderId,
            quantity: this.state.quantity,
            limitPrice: this.state.limitPrice,
            placedAt: this.state.placedAt,
            side: this.state.side,
            type: this.state.type,
            status: this.state.status,
            timeInForce: this.state.timeInForce
        };
        console.log('tradeOrder => ' + JSON.stringify(tradeOrder));
        console.log('id => ' + JSON.stringify(this.state.id));
        TradeOrderService.updateTradeOrder(tradeOrder).then( res => {
            this.props.history.push('/tradeOrders');
        });
    }

    changeorderIdHandler= (event) => {
        this.setState({orderId: event.target.value});
    }
    changequantityHandler= (event) => {
        this.setState({quantity: event.target.value});
    }
    changelimitPriceHandler= (event) => {
        this.setState({limitPrice: event.target.value});
    }
    changeplacedAtHandler= (event) => {
        this.setState({placedAt: event.target.value});
    }
    changeSideHandler= (event) => {
        this.setState({side: event.target.value});
    }
    changeTypeHandler= (event) => {
        this.setState({type: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changeTimeInForceHandler= (event) => {
        this.setState({timeInForce: event.target.value});
    }

    cancel(){
        this.props.history.push('/tradeOrders');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update TradeOrder</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> orderId: </label>
                                                <input placeholder="orderId" name="orderId" className="form-control" value={this.state.orderId} onChange={this.changeorderIdHandler}/>

                                            <label> quantity: </label>
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> limitPrice: </label>
                                                <input placeholder="limitPrice" name="limitPrice" className="form-control" value={this.state.limitPrice} onChange={this.changelimitPriceHandler}/>

                                            <label> placedAt: </label>
                                                <input type="time" placeholder="placedAt" name="placedAt" className="form-control" value={this.state.placedAt} onChange={this.changeplacedAtHandler}/>

                                            <label> Side: </label>
                                                <select value={this.state.side} onChange={this.changeSideHandler}>
                      <option name="Side" className="form-control" >
                          Buy
                      </option>
                      <option name="Side" className="form-control" >
                          Sell
                      </option>
                    </select>

                                            <label> Type: </label>
                                                <select value={this.state.type} onChange={this.changeTypeHandler}>
                      <option name="Type" className="form-control" >
                          Market
                      </option>
                      <option name="Type" className="form-control" >
                          Limit
                      </option>
                      <option name="Type" className="form-control" >
                          Stop
                      </option>
                      <option name="Type" className="form-control" >
                          StopLimit
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          New
                      </option>
                      <option name="Status" className="form-control" >
                          PartiallyFilled
                      </option>
                      <option name="Status" className="form-control" >
                          Filled
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                      <option name="Status" className="form-control" >
                          Rejected
                      </option>
                      <option name="Status" className="form-control" >
                          Expired
                      </option>
                    </select>

                                            <label> TimeInForce: </label>
                                                <select value={this.state.timeInForce} onChange={this.changeTimeInForceHandler}>
                      <option name="TimeInForce" className="form-control" >
                          Day
                      </option>
                      <option name="TimeInForce" className="form-control" >
                          GTC
                      </option>
                      <option name="TimeInForce" className="form-control" >
                          IOC
                      </option>
                      <option name="TimeInForce" className="form-control" >
                          FOK
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateTradeOrder}>Save</button>
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

export default UpdateTradeOrderComponent
