import React, { Component } from 'react'
import OrderItemService from '../services/OrderItemService'

class ViewOrderItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            orderItem: {}
        }
    }

    componentDidMount(){
        OrderItemService.getOrderItemById(this.state.id).then( res => {
            this.setState({orderItem: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View OrderItem Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.orderItem.quantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> unitPrice:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.orderItem.unitPrice }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> discountAmount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.orderItem.discountAmount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> taxAmount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.orderItem.taxAmount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> totalAmount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.orderItem.totalAmount }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewOrderItemComponent
