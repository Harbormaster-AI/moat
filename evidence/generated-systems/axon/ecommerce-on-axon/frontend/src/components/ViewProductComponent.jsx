import React, { Component } from 'react'
import ProductService from '../services/ProductService'

class ViewProductComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            product: {}
        }
    }

    componentDidMount(){
        ProductService.getProductById(this.state.id).then( res => {
            this.setState({product: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Product Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.product.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> slug:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.product.slug }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> asActive:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.product.asActive }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ProductType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.product.productType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> DefaultTaxClass:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.product.defaultTaxClass }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewProductComponent
