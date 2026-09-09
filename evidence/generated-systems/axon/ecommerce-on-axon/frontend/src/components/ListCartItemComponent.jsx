import React, { Component } from 'react'
import CartItemService from '../services/CartItemService'

class ListCartItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                cartItems: []
        }
        this.addCartItem = this.addCartItem.bind(this);
        this.editCartItem = this.editCartItem.bind(this);
        this.deleteCartItem = this.deleteCartItem.bind(this);
    }

    deleteCartItem(id){
        CartItemService.deleteCartItem(id).then( res => {
            this.setState({cartItems: this.state.cartItems.filter(cartItem => cartItem.cartItemId !== id)});
        });
    }
    viewCartItem(id){
        this.props.history.push(`/view-cartItem/${id}`);
    }
    editCartItem(id){
        this.props.history.push(`/add-cartItem/${id}`);
    }

    componentDidMount(){
        CartItemService.getCartItems().then((res) => {
            this.setState({ cartItems: res.data});
        });
    }

    addCartItem(){
        this.props.history.push('/add-cartItem/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CartItem List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCartItem}> Add CartItem</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Quantity </th>
                                    <th> UnitPrice </th>
                                    <th> TotalPrice </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.cartItems.map(
                                        cartItem => 
                                        <tr key = {cartItem.cartItemId}>
                                             <td> { cartItem.quantity } </td>
                                             <td> { cartItem.unitPrice } </td>
                                             <td> { cartItem.totalPrice } </td>
                                             <td>
                                                 <button onClick={ () => this.editCartItem(cartItem.cartItemId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCartItem(cartItem.cartItemId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCartItem(cartItem.cartItemId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListCartItemComponent
