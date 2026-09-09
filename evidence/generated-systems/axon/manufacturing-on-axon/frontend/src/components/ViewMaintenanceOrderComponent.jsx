import React, { Component } from 'react'
import MaintenanceOrderService from '../services/MaintenanceOrderService'

class ViewMaintenanceOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            maintenanceOrder: {}
        }
    }

    componentDidMount(){
        MaintenanceOrderService.getMaintenanceOrderById(this.state.id).then( res => {
            this.setState({maintenanceOrder: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View MaintenanceOrder Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> orderNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.maintenanceOrder.orderNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> priority:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.maintenanceOrder.priority }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> requestedDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.maintenanceOrder.requestedDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> completionDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.maintenanceOrder.completionDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.maintenanceOrder.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewMaintenanceOrderComponent
