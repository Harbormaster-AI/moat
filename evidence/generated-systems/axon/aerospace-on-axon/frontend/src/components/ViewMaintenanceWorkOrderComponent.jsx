import React, { Component } from 'react'
import MaintenanceWorkOrderService from '../services/MaintenanceWorkOrderService'

class ViewMaintenanceWorkOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            maintenanceWorkOrder: {}
        }
    }

    componentDidMount(){
        MaintenanceWorkOrderService.getMaintenanceWorkOrderById(this.state.id).then( res => {
            this.setState({maintenanceWorkOrder: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View MaintenanceWorkOrder Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> workOrderNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.maintenanceWorkOrder.workOrderNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.maintenanceWorkOrder.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewMaintenanceWorkOrderComponent
