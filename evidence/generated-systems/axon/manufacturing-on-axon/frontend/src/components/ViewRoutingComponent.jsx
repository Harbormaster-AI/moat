import React, { Component } from 'react'
import RoutingService from '../services/RoutingService'

class ViewRoutingComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            routing: {}
        }
    }

    componentDidMount(){
        RoutingService.getRoutingById(this.state.id).then( res => {
            this.setState({routing: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Routing Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> routingNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.routing.routingNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> revision:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.routing.revision }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> effectivityStart:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.routing.effectivityStart }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> effectivityEnd:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.routing.effectivityEnd }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> RoutingType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.routing.routingType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.routing.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewRoutingComponent
