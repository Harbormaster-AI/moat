import React, { Component } from 'react'
import CartService from '../services/CartService';

class UpdateCartComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                cartNumber: '',
                createdAt: '',
                currency: '',
                shippingAddress: '',
                billingAddress: '',
                status: ''
        }
        this.updateCart = this.updateCart.bind(this);

        this.changecartNumberHandler = this.changecartNumberHandler.bind(this);
        this.changecreatedAtHandler = this.changecreatedAtHandler.bind(this);
        this.changecurrencyHandler = this.changecurrencyHandler.bind(this);
        this.changeshippingAddressHandler = this.changeshippingAddressHandler.bind(this);
        this.changebillingAddressHandler = this.changebillingAddressHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        CartService.getCartById(this.state.id).then( (res) =>{
            let cart = res.data;
            this.setState({
                cartNumber: cart.cartNumber,
                createdAt: cart.createdAt,
                currency: cart.currency,
                shippingAddress: cart.shippingAddress,
                billingAddress: cart.billingAddress,
                status: cart.status
            });
        });
    }

    updateCart = (e) => {
        e.preventDefault();
        let cart = {
            cartId: this.state.id,
            cartNumber: this.state.cartNumber,
            createdAt: this.state.createdAt,
            currency: this.state.currency,
            shippingAddress: this.state.shippingAddress,
            billingAddress: this.state.billingAddress,
            status: this.state.status
        };
        console.log('cart => ' + JSON.stringify(cart));
        console.log('id => ' + JSON.stringify(this.state.id));
        CartService.updateCart(cart).then( res => {
            this.props.history.push('/carts');
        });
    }

    changecartNumberHandler= (event) => {
        this.setState({cartNumber: event.target.value});
    }
    changecreatedAtHandler= (event) => {
        this.setState({createdAt: event.target.value});
    }
    changecurrencyHandler= (event) => {
        this.setState({currency: event.target.value});
    }
    changeshippingAddressHandler= (event) => {
        this.setState({shippingAddress: event.target.value});
    }
    changebillingAddressHandler= (event) => {
        this.setState({billingAddress: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/carts');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Cart</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> cartNumber: </label>
                                                <input placeholder="cartNumber" name="cartNumber" className="form-control" value={this.state.cartNumber} onChange={this.changecartNumberHandler}/>

                                            <label> createdAt: </label>
                                                <input type="date" placeholder="createdAt" name="createdAt" className="form-control" value={this.state.createdAt} onChange={this.changecreatedAtHandler}/>

                                            <label> currency: </label>
                                                <input placeholder="currency" name="currency" className="form-control" value={this.state.currency} onChange={this.changecurrencyHandler}/>

                                            <label> shippingAddress: </label>
                                                <input placeholder="shippingAddress" name="shippingAddress" className="form-control" value={this.state.shippingAddress} onChange={this.changeshippingAddressHandler}/>

                                            <label> billingAddress: </label>
                                                <input placeholder="billingAddress" name="billingAddress" className="form-control" value={this.state.billingAddress} onChange={this.changebillingAddressHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Merged
                      </option>
                      <option name="Status" className="form-control" >
                          Ordered
                      </option>
                      <option name="Status" className="form-control" >
                          Abandoned
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCart}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateCartComponent
