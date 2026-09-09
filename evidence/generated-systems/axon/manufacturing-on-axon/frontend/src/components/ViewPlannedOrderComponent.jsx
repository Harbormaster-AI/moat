import React, { Component } from 'react'
import PlannedOrderService from '../services/PlannedOrderService'

class ViewPlannedOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            plannedOrder: {}
        }
    }

    componentDidMount(){
        PlannedOrderService.getPlannedOrderById(this.state.id).then( res => {
            this.setState({plannedOrder: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View PlannedOrder Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> plannedOrderNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.plannedOrder.plannedOrderNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.plannedOrder.quantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> dueDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.plannedOrder.dueDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> OrderType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.plannedOrder.orderType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.plannedOrder.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPlannedOrderComponent
