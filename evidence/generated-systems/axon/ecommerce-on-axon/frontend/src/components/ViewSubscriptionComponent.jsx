import React, { Component } from 'react'
import SubscriptionService from '../services/SubscriptionService'

class ViewSubscriptionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            subscription: {}
        }
    }

    componentDidMount(){
        SubscriptionService.getSubscriptionById(this.state.id).then( res => {
            this.setState({subscription: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Subscription Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> subscriptionNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.subscription.subscriptionNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> nextBillingDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.subscription.nextBillingDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> startDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.subscription.startDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> endDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.subscription.endDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.subscription.status }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Interval:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.subscription.interval }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewSubscriptionComponent
