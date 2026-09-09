import React, { Component } from 'react'
import OrderItemService from '../services/OrderItemService';

class CreateOrderItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                quantity: '',
                unitPrice: '',
                discountAmount: '',
                taxAmount: '',
                totalAmount: ''
        }
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changeunitPriceHandler = this.changeunitPriceHandler.bind(this);
        this.changediscountAmountHandler = this.changediscountAmountHandler.bind(this);
        this.changetaxAmountHandler = this.changetaxAmountHandler.bind(this);
        this.changetotalAmountHandler = this.changetotalAmountHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            OrderItemService.getOrderItemById(this.state.id).then( (res) =>{
                let orderItem = res.data;
                this.setState({
                    quantity: orderItem.quantity,
                    unitPrice: orderItem.unitPrice,
                    discountAmount: orderItem.discountAmount,
                    taxAmount: orderItem.taxAmount,
                    totalAmount: orderItem.totalAmount
                });
            });
        }        
    }
    saveOrUpdateOrderItem = (e) => {
        e.preventDefault();
        let orderItem = {
                orderItemId: this.state.id,
                quantity: this.state.quantity,
                unitPrice: this.state.unitPrice,
                discountAmount: this.state.discountAmount,
                taxAmount: this.state.taxAmount,
                totalAmount: this.state.totalAmount
            };
        console.log('orderItem => ' + JSON.stringify(orderItem));

        // step 5
        if(this.state.id === '_add'){
            orderItem.orderItemId=''
            OrderItemService.createOrderItem(orderItem).then(res =>{
                this.props.history.push('/orderItems');
            });
        }else{
            OrderItemService.updateOrderItem(orderItem).then( res => {
                this.props.history.push('/orderItems');
            });
        }
    }
    
    changequantityHandler= (event) => {
        this.setState({quantity: event.target.value});
    }
    changeunitPriceHandler= (event) => {
        this.setState({unitPrice: event.target.value});
    }
    changediscountAmountHandler= (event) => {
        this.setState({discountAmount: event.target.value});
    }
    changetaxAmountHandler= (event) => {
        this.setState({taxAmount: event.target.value});
    }
    changetotalAmountHandler= (event) => {
        this.setState({totalAmount: event.target.value});
    }

    cancel(){
        this.props.history.push('/orderItems');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add OrderItem</h3>
        }else{
            return <h3 className="text-center">Update OrderItem</h3>
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
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> unitPrice:&emsp; </label>
                                                <input placeholder="unitPrice" name="unitPrice" className="form-control" value={this.state.unitPrice} onChange={this.changeunitPriceHandler}/>

                                            <label> discountAmount:&emsp; </label>
                                                <input placeholder="discountAmount" name="discountAmount" className="form-control" value={this.state.discountAmount} onChange={this.changediscountAmountHandler}/>

                                            <label> taxAmount:&emsp; </label>
                                                <input placeholder="taxAmount" name="taxAmount" className="form-control" value={this.state.taxAmount} onChange={this.changetaxAmountHandler}/>

                                            <label> totalAmount:&emsp; </label>
                                                <input placeholder="totalAmount" name="totalAmount" className="form-control" value={this.state.totalAmount} onChange={this.changetotalAmountHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateOrderItem}>Save</button>
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

export default CreateOrderItemComponent
