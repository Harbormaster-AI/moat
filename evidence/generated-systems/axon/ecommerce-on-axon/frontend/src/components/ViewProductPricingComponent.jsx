import React, { Component } from 'react'
import ProductPricingService from '../services/ProductPricingService'

class ViewProductPricingComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            productPricing: {}
        }
    }

    componentDidMount(){
        ProductPricingService.getProductPricingById(this.state.id).then( res => {
            this.setState({productPricing: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ProductPricing Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> listPrice:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.productPricing.listPrice }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> salePrice:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.productPricing.salePrice }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> validFrom:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.productPricing.validFrom }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> validTo:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.productPricing.validTo }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewProductPricingComponent
