import React, { Component } from 'react'
import OrderLineService from '../services/OrderLineService';

class CreateOrderLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                quantity: '',
                unitPrice: '',
                totalPrice: '',
                taxRate: '',
                lineStatus: ''
        }
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changeunitPriceHandler = this.changeunitPriceHandler.bind(this);
        this.changetotalPriceHandler = this.changetotalPriceHandler.bind(this);
        this.changetaxRateHandler = this.changetaxRateHandler.bind(this);
        this.changeLineStatusHandler = this.changeLineStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            OrderLineService.getOrderLineById(this.state.id).then( (res) =>{
                let orderLine = res.data;
                this.setState({
                    quantity: orderLine.quantity,
                    unitPrice: orderLine.unitPrice,
                    totalPrice: orderLine.totalPrice,
                    taxRate: orderLine.taxRate,
                    lineStatus: orderLine.lineStatus
                });
            });
        }        
    }
    saveOrUpdateOrderLine = (e) => {
        e.preventDefault();
        let orderLine = {
                orderLineId: this.state.id,
                quantity: this.state.quantity,
                unitPrice: this.state.unitPrice,
                totalPrice: this.state.totalPrice,
                taxRate: this.state.taxRate,
                lineStatus: this.state.lineStatus
            };
        console.log('orderLine => ' + JSON.stringify(orderLine));

        // step 5
        if(this.state.id === '_add'){
            orderLine.orderLineId=''
            OrderLineService.createOrderLine(orderLine).then(res =>{
                this.props.history.push('/orderLines');
            });
        }else{
            OrderLineService.updateOrderLine(orderLine).then( res => {
                this.props.history.push('/orderLines');
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
    changetaxRateHandler= (event) => {
        this.setState({taxRate: event.target.value});
    }
    changeLineStatusHandler= (event) => {
        this.setState({lineStatus: event.target.value});
    }

    cancel(){
        this.props.history.push('/orderLines');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add OrderLine</h3>
        }else{
            return <h3 className="text-center">Update OrderLine</h3>
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

                                            <label> taxRate:&emsp; </label>
                                                <input placeholder="taxRate" name="taxRate" className="form-control" value={this.state.taxRate} onChange={this.changetaxRateHandler}/>

                                            <label> LineStatus:&emsp; </label>
                                                <select value={this.state.lineStatus} onChange={this.changeLineStatusHandler}>
                      <option name="LineStatus" className="form-control" >
                          Pending
                      </option>
                      <option name="LineStatus" className="form-control" >
                          Fulfilled
                      </option>
                      <option name="LineStatus" className="form-control" >
                          Cancelled
                      </option>
                      <option name="LineStatus" className="form-control" >
                          Backordered
                      </option>
                      <option name="LineStatus" className="form-control" >
                          Returned
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateOrderLine}>Save</button>
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

export default CreateOrderLineComponent
