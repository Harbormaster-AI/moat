import React, { Component } from 'react'
import SalesOrderService from '../services/SalesOrderService'

class ViewSalesOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            salesOrder: {}
        }
    }

    componentDidMount(){
        SalesOrderService.getSalesOrderById(this.state.id).then( res => {
            this.setState({salesOrder: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View SalesOrder Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> orderNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.salesOrder.orderNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> orderDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.salesOrder.orderDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> totalAmount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.salesOrder.totalAmount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.salesOrder.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewSalesOrderComponent
