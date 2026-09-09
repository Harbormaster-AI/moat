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
                            <div className = "col" style={{textAlign:"right"}}><label> orderDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.order.orderDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> totalAmount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.order.totalAmount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> taxAmount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.order.taxAmount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> shippingAmount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.order.shippingAmount }</div>
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
