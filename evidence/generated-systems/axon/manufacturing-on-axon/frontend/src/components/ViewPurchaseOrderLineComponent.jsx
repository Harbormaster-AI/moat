import React, { Component } from 'react'
import PurchaseOrderLineService from '../services/PurchaseOrderLineService'

class ViewPurchaseOrderLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            purchaseOrderLine: {}
        }
    }

    componentDidMount(){
        PurchaseOrderLineService.getPurchaseOrderLineById(this.state.id).then( res => {
            this.setState({purchaseOrderLine: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View PurchaseOrderLine Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> lineNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.purchaseOrderLine.lineNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.purchaseOrderLine.quantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> unitPrice:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.purchaseOrderLine.unitPrice }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> dueDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.purchaseOrderLine.dueDate }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPurchaseOrderLineComponent
