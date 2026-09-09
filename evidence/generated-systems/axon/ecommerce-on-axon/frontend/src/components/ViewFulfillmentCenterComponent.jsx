import React, { Component } from 'react'
import FulfillmentCenterService from '../services/FulfillmentCenterService'

class ViewFulfillmentCenterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            fulfillmentCenter: {}
        }
    }

    componentDidMount(){
        FulfillmentCenterService.getFulfillmentCenterById(this.state.id).then( res => {
            this.setState({fulfillmentCenter: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View FulfillmentCenter Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.fulfillmentCenter.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> centerCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.fulfillmentCenter.centerCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> address:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.fulfillmentCenter.address }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> timezone:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.fulfillmentCenter.timezone }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> asActive:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.fulfillmentCenter.asActive }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewFulfillmentCenterComponent
