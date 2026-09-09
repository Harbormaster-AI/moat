import React, { Component } from 'react'
import OrderLineService from '../services/OrderLineService';

class UpdateOrderLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                quantity: '',
                unitPrice: '',
                totalPrice: '',
                taxRate: '',
                lineStatus: ''
        }
        this.updateOrderLine = this.updateOrderLine.bind(this);

        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changeunitPriceHandler = this.changeunitPriceHandler.bind(this);
        this.changetotalPriceHandler = this.changetotalPriceHandler.bind(this);
        this.changetaxRateHandler = this.changetaxRateHandler.bind(this);
        this.changeLineStatusHandler = this.changeLineStatusHandler.bind(this);
    }

    componentDidMount(){
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

    updateOrderLine = (e) => {
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
        console.log('id => ' + JSON.stringify(this.state.id));
        OrderLineService.updateOrderLine(orderLine).then( res => {
            this.props.history.push('/orderLines');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update OrderLine</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> quantity: </label>
                                                <input type="number" placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> unitPrice: </label>
                                                <input placeholder="unitPrice" name="unitPrice" className="form-control" value={this.state.unitPrice} onChange={this.changeunitPriceHandler}/>

                                            <label> totalPrice: </label>
                                                <input placeholder="totalPrice" name="totalPrice" className="form-control" value={this.state.totalPrice} onChange={this.changetotalPriceHandler}/>

                                            <label> taxRate: </label>
                                                <input placeholder="taxRate" name="taxRate" className="form-control" value={this.state.taxRate} onChange={this.changetaxRateHandler}/>

                                            <label> LineStatus: </label>
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
                                        <button className="btn btn-success" onClick={this.updateOrderLine}>Save</button>
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

export default UpdateOrderLineComponent
