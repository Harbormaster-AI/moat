import React, { Component } from 'react'
import CartService from '../services/CartService';

class CreateCartComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                cartNumber: '',
                createdAt: '',
                currency: '',
                shippingAddress: '',
                billingAddress: '',
                status: ''
        }
        this.changecartNumberHandler = this.changecartNumberHandler.bind(this);
        this.changecreatedAtHandler = this.changecreatedAtHandler.bind(this);
        this.changecurrencyHandler = this.changecurrencyHandler.bind(this);
        this.changeshippingAddressHandler = this.changeshippingAddressHandler.bind(this);
        this.changebillingAddressHandler = this.changebillingAddressHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateCart = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            cart.cartId=''
            CartService.createCart(cart).then(res =>{
                this.props.history.push('/carts');
            });
        }else{
            CartService.updateCart(cart).then( res => {
                this.props.history.push('/carts');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Cart</h3>
        }else{
            return <h3 className="text-center">Update Cart</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> cartNumber:&emsp; </label>
                                                <input placeholder="cartNumber" name="cartNumber" className="form-control" value={this.state.cartNumber} onChange={this.changecartNumberHandler}/>

                                            <label> createdAt:&emsp; </label>
                                                <input type="date" placeholder="createdAt" name="createdAt" className="form-control" value={this.state.createdAt} onChange={this.changecreatedAtHandler}/>

                                            <label> currency:&emsp; </label>
                                                <input placeholder="currency" name="currency" className="form-control" value={this.state.currency} onChange={this.changecurrencyHandler}/>

                                            <label> shippingAddress:&emsp; </label>
                                                <input placeholder="shippingAddress" name="shippingAddress" className="form-control" value={this.state.shippingAddress} onChange={this.changeshippingAddressHandler}/>

                                            <label> billingAddress:&emsp; </label>
                                                <input placeholder="billingAddress" name="billingAddress" className="form-control" value={this.state.billingAddress} onChange={this.changebillingAddressHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCart}>Save</button>
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

export default CreateCartComponent
