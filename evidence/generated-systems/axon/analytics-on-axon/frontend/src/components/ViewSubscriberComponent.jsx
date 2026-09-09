import React, { Component } from 'react'
import SubscriberService from '../services/SubscriberService'

class ViewSubscriberComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            subscriber: {}
        }
    }

    componentDidMount(){
        SubscriberService.getSubscriberById(this.state.id).then( res => {
            this.setState({subscriber: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Subscriber Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.subscriber.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> address:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.subscriber.address }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Channel:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.subscriber.channel }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewSubscriberComponent
