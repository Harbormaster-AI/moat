import React, { Component } from 'react'
import WorkOrderService from '../services/WorkOrderService'

class ViewWorkOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            workOrder: {}
        }
    }

    componentDidMount(){
        WorkOrderService.getWorkOrderById(this.state.id).then( res => {
            this.setState({workOrder: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View WorkOrder Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> workOrderNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.workOrder.workOrderNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> plannedStart:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.workOrder.plannedStart }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> plannedEnd:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.workOrder.plannedEnd }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.workOrder.quantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> priority:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.workOrder.priority }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.workOrder.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewWorkOrderComponent
