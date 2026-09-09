import React, { Component } from 'react'
import OpportunityLineItemService from '../services/OpportunityLineItemService'

class ViewOpportunityLineItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            opportunityLineItem: {}
        }
    }

    componentDidMount(){
        OpportunityLineItemService.getOpportunityLineItemById(this.state.id).then( res => {
            this.setState({opportunityLineItem: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View OpportunityLineItem Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.opportunityLineItem.quantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> unitPrice:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.opportunityLineItem.unitPrice }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> discountPercent:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.opportunityLineItem.discountPercent }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> totalPrice:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.opportunityLineItem.totalPrice }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewOpportunityLineItemComponent
