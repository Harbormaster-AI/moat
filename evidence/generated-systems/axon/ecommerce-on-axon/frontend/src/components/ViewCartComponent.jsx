import React, { Component } from 'react'
import CartService from '../services/CartService'

class ViewCartComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            cart: {}
        }
    }

    componentDidMount(){
        CartService.getCartById(this.state.id).then( res => {
            this.setState({cart: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Cart Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> cartNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cart.cartNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> createdAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cart.createdAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> currency:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cart.currency }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> shippingAddress:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cart.shippingAddress }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> billingAddress:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cart.billingAddress }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cart.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCartComponent
