import React, { Component } from 'react'
import CartService from '../services/CartService'

class ListCartComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                carts: []
        }
        this.addCart = this.addCart.bind(this);
        this.editCart = this.editCart.bind(this);
        this.deleteCart = this.deleteCart.bind(this);
    }

    deleteCart(id){
        CartService.deleteCart(id).then( res => {
            this.setState({carts: this.state.carts.filter(cart => cart.cartId !== id)});
        });
    }
    viewCart(id){
        this.props.history.push(`/view-cart/${id}`);
    }
    editCart(id){
        this.props.history.push(`/add-cart/${id}`);
    }

    componentDidMount(){
        CartService.getCarts().then((res) => {
            this.setState({ carts: res.data});
        });
    }

    addCart(){
        this.props.history.push('/add-cart/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Cart List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCart}> Add Cart</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> CartNumber </th>
                                    <th> CreatedAt </th>
                                    <th> Currency </th>
                                    <th> ShippingAddress </th>
                                    <th> BillingAddress </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.carts.map(
                                        cart => 
                                        <tr key = {cart.cartId}>
                                             <td> { cart.cartNumber } </td>
                                             <td> { cart.createdAt } </td>
                                             <td> { cart.currency } </td>
                                             <td> { cart.shippingAddress } </td>
                                             <td> { cart.billingAddress } </td>
                                             <td> { cart.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editCart(cart.cartId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCart(cart.cartId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCart(cart.cartId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCartComponent
