import React, { Component } from 'react'
import SalesOrderService from '../services/SalesOrderService';

class UpdateSalesOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                orderNumber: '',
                orderDate: '',
                totalAmount: '',
                status: ''
        }
        this.updateSalesOrder = this.updateSalesOrder.bind(this);

        this.changeorderNumberHandler = this.changeorderNumberHandler.bind(this);
        this.changeorderDateHandler = this.changeorderDateHandler.bind(this);
        this.changetotalAmountHandler = this.changetotalAmountHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        SalesOrderService.getSalesOrderById(this.state.id).then( (res) =>{
            let salesOrder = res.data;
            this.setState({
                orderNumber: salesOrder.orderNumber,
                orderDate: salesOrder.orderDate,
                totalAmount: salesOrder.totalAmount,
                status: salesOrder.status
            });
        });
    }

    updateSalesOrder = (e) => {
        e.preventDefault();
        let salesOrder = {
            salesOrderId: this.state.id,
            orderNumber: this.state.orderNumber,
            orderDate: this.state.orderDate,
            totalAmount: this.state.totalAmount,
            status: this.state.status
        };
        console.log('salesOrder => ' + JSON.stringify(salesOrder));
        console.log('id => ' + JSON.stringify(this.state.id));
        SalesOrderService.updateSalesOrder(salesOrder).then( res => {
            this.props.history.push('/salesOrders');
        });
    }

    changeorderNumberHandler= (event) => {
        this.setState({orderNumber: event.target.value});
    }
    changeorderDateHandler= (event) => {
        this.setState({orderDate: event.target.value});
    }
    changetotalAmountHandler= (event) => {
        this.setState({totalAmount: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/salesOrders');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update SalesOrder</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> orderNumber: </label>
                                                <input placeholder="orderNumber" name="orderNumber" className="form-control" value={this.state.orderNumber} onChange={this.changeorderNumberHandler}/>

                                            <label> orderDate: </label>
                                                <input type="date" placeholder="orderDate" name="orderDate" className="form-control" value={this.state.orderDate} onChange={this.changeorderDateHandler}/>

                                            <label> totalAmount: </label>
                                                <input placeholder="totalAmount" name="totalAmount" className="form-control" value={this.state.totalAmount} onChange={this.changetotalAmountHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Confirmed
                      </option>
                      <option name="Status" className="form-control" >
                          Allocated
                      </option>
                      <option name="Status" className="form-control" >
                          InProduction
                      </option>
                      <option name="Status" className="form-control" >
                          Shipped
                      </option>
                      <option name="Status" className="form-control" >
                          Invoiced
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateSalesOrder}>Save</button>
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

export default UpdateSalesOrderComponent
