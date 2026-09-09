import React, { Component } from 'react'
import PayoutService from '../services/PayoutService'

class ViewPayoutComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            payout: {}
        }
    }

    componentDidMount(){
        PayoutService.getPayoutById(this.state.id).then( res => {
            this.setState({payout: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Payout Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> payoutReference:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payout.payoutReference }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> amount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payout.amount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> currency:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payout.currency }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> scheduledDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payout.scheduledDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> paidDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payout.paidDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payout.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPayoutComponent
