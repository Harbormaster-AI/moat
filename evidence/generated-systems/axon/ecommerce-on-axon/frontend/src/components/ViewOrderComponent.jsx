import React, { Component } from 'react'
import OrderService from '../services/OrderService'

class ViewOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            order: {}
        }
    }

    componentDidMount(){
        OrderService.getOrderById(this.state.id).then( res => {
            this.setState({order: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Order Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> orderNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.order.orderNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> placedDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.order.placedDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> subtotal:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.order.subtotal }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> discountTotal:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.order.discountTotal }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> shippingTotal:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.order.shippingTotal }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> taxTotal:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.order.taxTotal }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> grandTotal:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.order.grandTotal }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> shippingAddress:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.order.shippingAddress }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> billingAddress:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.order.billingAddress }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.order.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewOrderComponent
