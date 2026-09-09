import React, { Component } from 'react'
import CartItemService from '../services/CartItemService';

class CreateCartItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                quantity: '',
                unitPrice: '',
                totalPrice: ''
        }
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changeunitPriceHandler = this.changeunitPriceHandler.bind(this);
        this.changetotalPriceHandler = this.changetotalPriceHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            CartItemService.getCartItemById(this.state.id).then( (res) =>{
                let cartItem = res.data;
                this.setState({
                    quantity: cartItem.quantity,
                    unitPrice: cartItem.unitPrice,
                    totalPrice: cartItem.totalPrice
                });
            });
        }        
    }
    saveOrUpdateCartItem = (e) => {
        e.preventDefault();
        let cartItem = {
                cartItemId: this.state.id,
                quantity: this.state.quantity,
                unitPrice: this.state.unitPrice,
                totalPrice: this.state.totalPrice
            };
        console.log('cartItem => ' + JSON.stringify(cartItem));

        // step 5
        if(this.state.id === '_add'){
            cartItem.cartItemId=''
            CartItemService.createCartItem(cartItem).then(res =>{
                this.props.history.push('/cartItems');
            });
        }else{
            CartItemService.updateCartItem(cartItem).then( res => {
                this.props.history.push('/cartItems');
            });
        }
    }
    
    changequantityHandler= (event) => {
        this.setState({quantity: event.target.value});
    }
    changeunitPriceHandler= (event) => {
        this.setState({unitPrice: event.target.value});
    }
    changetotalPriceHandler= (event) => {
        this.setState({totalPrice: event.target.value});
    }

    cancel(){
        this.props.history.push('/cartItems');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add CartItem</h3>
        }else{
            return <h3 className="text-center">Update CartItem</h3>
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
                                            <label> quantity:&emsp; </label>
                                                <input type="number" placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> unitPrice:&emsp; </label>
                                                <input placeholder="unitPrice" name="unitPrice" className="form-control" value={this.state.unitPrice} onChange={this.changeunitPriceHandler}/>

                                            <label> totalPrice:&emsp; </label>
                                                <input placeholder="totalPrice" name="totalPrice" className="form-control" value={this.state.totalPrice} onChange={this.changetotalPriceHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCartItem}>Save</button>
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

export default CreateCartItemComponent
