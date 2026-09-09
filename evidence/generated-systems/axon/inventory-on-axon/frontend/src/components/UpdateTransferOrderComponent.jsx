import React, { Component } from 'react'
import TransferOrderService from '../services/TransferOrderService';

class UpdateTransferOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                orderNumber: '',
                requestedShipDate: '',
                requestedReceiveDate: '',
                shippedDate: '',
                receivedDate: '',
                status: ''
        }
        this.updateTransferOrder = this.updateTransferOrder.bind(this);

        this.changeorderNumberHandler = this.changeorderNumberHandler.bind(this);
        this.changerequestedShipDateHandler = this.changerequestedShipDateHandler.bind(this);
        this.changerequestedReceiveDateHandler = this.changerequestedReceiveDateHandler.bind(this);
        this.changeshippedDateHandler = this.changeshippedDateHandler.bind(this);
        this.changereceivedDateHandler = this.changereceivedDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        TransferOrderService.getTransferOrderById(this.state.id).then( (res) =>{
            let transferOrder = res.data;
            this.setState({
                orderNumber: transferOrder.orderNumber,
                requestedShipDate: transferOrder.requestedShipDate,
                requestedReceiveDate: transferOrder.requestedReceiveDate,
                shippedDate: transferOrder.shippedDate,
                receivedDate: transferOrder.receivedDate,
                status: transferOrder.status
            });
        });
    }

    updateTransferOrder = (e) => {
        e.preventDefault();
        let transferOrder = {
            transferOrderId: this.state.id,
            orderNumber: this.state.orderNumber,
            requestedShipDate: this.state.requestedShipDate,
            requestedReceiveDate: this.state.requestedReceiveDate,
            shippedDate: this.state.shippedDate,
            receivedDate: this.state.receivedDate,
            status: this.state.status
        };
        console.log('transferOrder => ' + JSON.stringify(transferOrder));
        console.log('id => ' + JSON.stringify(this.state.id));
        TransferOrderService.updateTransferOrder(transferOrder).then( res => {
            this.props.history.push('/transferOrders');
        });
    }

    changeorderNumberHandler= (event) => {
        this.setState({orderNumber: event.target.value});
    }
    changerequestedShipDateHandler= (event) => {
        this.setState({requestedShipDate: event.target.value});
    }
    changerequestedReceiveDateHandler= (event) => {
        this.setState({requestedReceiveDate: event.target.value});
    }
    changeshippedDateHandler= (event) => {
        this.setState({shippedDate: event.target.value});
    }
    changereceivedDateHandler= (event) => {
        this.setState({receivedDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/transferOrders');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update TransferOrder</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> orderNumber: </label>
                                                <input placeholder="orderNumber" name="orderNumber" className="form-control" value={this.state.orderNumber} onChange={this.changeorderNumberHandler}/>

                                            <label> requestedShipDate: </label>
                                                <input type="date" placeholder="requestedShipDate" name="requestedShipDate" className="form-control" value={this.state.requestedShipDate} onChange={this.changerequestedShipDateHandler}/>

                                            <label> requestedReceiveDate: </label>
                                                <input type="date" placeholder="requestedReceiveDate" name="requestedReceiveDate" className="form-control" value={this.state.requestedReceiveDate} onChange={this.changerequestedReceiveDateHandler}/>

                                            <label> shippedDate: </label>
                                                <input type="date" placeholder="shippedDate" name="shippedDate" className="form-control" value={this.state.shippedDate} onChange={this.changeshippedDateHandler}/>

                                            <label> receivedDate: </label>
                                                <input type="date" placeholder="receivedDate" name="receivedDate" className="form-control" value={this.state.receivedDate} onChange={this.changereceivedDateHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Released
                      </option>
                      <option name="Status" className="form-control" >
                          InTransit
                      </option>
                      <option name="Status" className="form-control" >
                          Received
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateTransferOrder}>Save</button>
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

export default UpdateTransferOrderComponent
