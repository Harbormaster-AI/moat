import React, { Component } from 'react'
import OutboundAllocationService from '../services/OutboundAllocationService'

class ViewOutboundAllocationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            outboundAllocation: {}
        }
    }

    componentDidMount(){
        OutboundAllocationService.getOutboundAllocationById(this.state.id).then( res => {
            this.setState({outboundAllocation: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View OutboundAllocation Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> allocationNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.outboundAllocation.allocationNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> allocatedQuantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.outboundAllocation.allocatedQuantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> allocationDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.outboundAllocation.allocationDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.outboundAllocation.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewOutboundAllocationComponent
