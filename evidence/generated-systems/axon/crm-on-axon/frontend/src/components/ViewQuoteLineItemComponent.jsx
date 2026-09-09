import React, { Component } from 'react'
import QuoteLineItemService from '../services/QuoteLineItemService'

class ViewQuoteLineItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            quoteLineItem: {}
        }
    }

    componentDidMount(){
        QuoteLineItemService.getQuoteLineItemById(this.state.id).then( res => {
            this.setState({quoteLineItem: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View QuoteLineItem Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.quoteLineItem.quantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> unitPrice:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.quoteLineItem.unitPrice }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> discountAmount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.quoteLineItem.discountAmount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> taxAmount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.quoteLineItem.taxAmount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> totalAmount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.quoteLineItem.totalAmount }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewQuoteLineItemComponent
