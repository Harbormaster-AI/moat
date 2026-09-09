import React, { Component } from 'react'
import OrderItemService from '../services/OrderItemService';

class UpdateOrderItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                quantity: '',
                unitPrice: '',
                discountAmount: '',
                taxAmount: '',
                totalAmount: ''
        }
        this.updateOrderItem = this.updateOrderItem.bind(this);

        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changeunitPriceHandler = this.changeunitPriceHandler.bind(this);
        this.changediscountAmountHandler = this.changediscountAmountHandler.bind(this);
        this.changetaxAmountHandler = this.changetaxAmountHandler.bind(this);
        this.changetotalAmountHandler = this.changetotalAmountHandler.bind(this);
    }

    componentDidMount(){
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

    updateOrderItem = (e) => {
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
        console.log('id => ' + JSON.stringify(this.state.id));
        OrderItemService.updateOrderItem(orderItem).then( res => {
            this.props.history.push('/orderItems');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update OrderItem</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> quantity: </label>
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> unitPrice: </label>
                                                <input placeholder="unitPrice" name="unitPrice" className="form-control" value={this.state.unitPrice} onChange={this.changeunitPriceHandler}/>

                                            <label> discountAmount: </label>
                                                <input placeholder="discountAmount" name="discountAmount" className="form-control" value={this.state.discountAmount} onChange={this.changediscountAmountHandler}/>

                                            <label> taxAmount: </label>
                                                <input placeholder="taxAmount" name="taxAmount" className="form-control" value={this.state.taxAmount} onChange={this.changetaxAmountHandler}/>

                                            <label> totalAmount: </label>
                                                <input placeholder="totalAmount" name="totalAmount" className="form-control" value={this.state.totalAmount} onChange={this.changetotalAmountHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateOrderItem}>Save</button>
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

export default UpdateOrderItemComponent
