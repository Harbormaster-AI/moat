import React, { Component } from 'react'
import CartItemService from '../services/CartItemService'

class ViewCartItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            cartItem: {}
        }
    }

    componentDidMount(){
        CartItemService.getCartItemById(this.state.id).then( res => {
            this.setState({cartItem: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View CartItem Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cartItem.quantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> unitPrice:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cartItem.unitPrice }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> totalPrice:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cartItem.totalPrice }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCartItemComponent
