import React, { Component } from 'react'
import TradeService from '../services/TradeService'

class ViewTradeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            trade: {}
        }
    }

    componentDidMount(){
        TradeService.getTradeById(this.state.id).then( res => {
            this.setState({trade: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Trade Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> executedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.trade.executedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.trade.quantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> price:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.trade.price }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> fees:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.trade.fees }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> settlementDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.trade.settlementDate }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewTradeComponent
