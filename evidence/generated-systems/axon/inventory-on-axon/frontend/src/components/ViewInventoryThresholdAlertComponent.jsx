import React, { Component } from 'react'
import InventoryThresholdAlertService from '../services/InventoryThresholdAlertService'

class ViewInventoryThresholdAlertComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            inventoryThresholdAlert: {}
        }
    }

    componentDidMount(){
        InventoryThresholdAlertService.getInventoryThresholdAlertById(this.state.id).then( res => {
            this.setState({inventoryThresholdAlert: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View InventoryThresholdAlert Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> alertNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inventoryThresholdAlert.alertNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> detectedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inventoryThresholdAlert.detectedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> message:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inventoryThresholdAlert.message }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> AlertType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inventoryThresholdAlert.alertType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inventoryThresholdAlert.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewInventoryThresholdAlertComponent
