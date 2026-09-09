import React, { Component } from 'react'
import InvoiceService from '../services/InvoiceService'

class ViewInvoiceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            invoice: {}
        }
    }

    componentDidMount(){
        InvoiceService.getInvoiceById(this.state.id).then( res => {
            this.setState({invoice: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Invoice Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> invoiceNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.invoice.invoiceNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> issueDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.invoice.issueDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> dueDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.invoice.dueDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> total:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.invoice.total }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> currency:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.invoice.currency }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.invoice.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewInvoiceComponent
