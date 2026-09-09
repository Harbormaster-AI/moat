import React, { Component } from 'react'
import TradeOrderService from '../services/TradeOrderService'

class ViewTradeOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            tradeOrder: {}
        }
    }

    componentDidMount(){
        TradeOrderService.getTradeOrderById(this.state.id).then( res => {
            this.setState({tradeOrder: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View TradeOrder Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> orderId:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.tradeOrder.orderId }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.tradeOrder.quantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> limitPrice:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.tradeOrder.limitPrice }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> placedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.tradeOrder.placedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Side:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.tradeOrder.side }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Type:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.tradeOrder.type }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.tradeOrder.status }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> TimeInForce:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.tradeOrder.timeInForce }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewTradeOrderComponent
