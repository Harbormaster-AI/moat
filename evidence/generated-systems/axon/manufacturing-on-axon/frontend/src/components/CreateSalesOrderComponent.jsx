import React, { Component } from 'react'
import SalesOrderService from '../services/SalesOrderService';

class CreateSalesOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                orderNumber: '',
                orderDate: '',
                totalAmount: '',
                status: ''
        }
        this.changeorderNumberHandler = this.changeorderNumberHandler.bind(this);
        this.changeorderDateHandler = this.changeorderDateHandler.bind(this);
        this.changetotalAmountHandler = this.changetotalAmountHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateSalesOrder = (e) => {
        e.preventDefault();
        let salesOrder = {
                salesOrderId: this.state.id,
                orderNumber: this.state.orderNumber,
                orderDate: this.state.orderDate,
                totalAmount: this.state.totalAmount,
                status: this.state.status
            };
        console.log('salesOrder => ' + JSON.stringify(salesOrder));

        // step 5
        if(this.state.id === '_add'){
            salesOrder.salesOrderId=''
            SalesOrderService.createSalesOrder(salesOrder).then(res =>{
                this.props.history.push('/salesOrders');
            });
        }else{
            SalesOrderService.updateSalesOrder(salesOrder).then( res => {
                this.props.history.push('/salesOrders');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add SalesOrder</h3>
        }else{
            return <h3 className="text-center">Update SalesOrder</h3>
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
                                            <label> orderNumber:&emsp; </label>
                                                <input placeholder="orderNumber" name="orderNumber" className="form-control" value={this.state.orderNumber} onChange={this.changeorderNumberHandler}/>

                                            <label> orderDate:&emsp; </label>
                                                <input type="date" placeholder="orderDate" name="orderDate" className="form-control" value={this.state.orderDate} onChange={this.changeorderDateHandler}/>

                                            <label> totalAmount:&emsp; </label>
                                                <input placeholder="totalAmount" name="totalAmount" className="form-control" value={this.state.totalAmount} onChange={this.changetotalAmountHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateSalesOrder}>Save</button>
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

export default CreateSalesOrderComponent
