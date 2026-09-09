import React, { Component } from 'react'
import ProductVariantService from '../services/ProductVariantService'

class ViewProductVariantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            productVariant: {}
        }
    }

    componentDidMount(){
        ProductVariantService.getProductVariantById(this.state.id).then( res => {
            this.setState({productVariant: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ProductVariant Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> sku:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.productVariant.sku }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> barcode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.productVariant.barcode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> title:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.productVariant.title }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> weight:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.productVariant.weight }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> requiresShipping:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.productVariant.requiresShipping }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> WeightUnit:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.productVariant.weightUnit }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewProductVariantComponent
