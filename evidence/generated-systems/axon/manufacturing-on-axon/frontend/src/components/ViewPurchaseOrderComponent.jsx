import React, { Component } from 'react'
import PurchaseOrderService from '../services/PurchaseOrderService'

class ViewPurchaseOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            purchaseOrder: {}
        }
    }

    componentDidMount(){
        PurchaseOrderService.getPurchaseOrderById(this.state.id).then( res => {
            this.setState({purchaseOrder: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View PurchaseOrder Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> poNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.purchaseOrder.poNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> orderDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.purchaseOrder.orderDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> totalAmount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.purchaseOrder.totalAmount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.purchaseOrder.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPurchaseOrderComponent
