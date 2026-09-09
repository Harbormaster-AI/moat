import React, { Component } from 'react'
import AircraftOrderService from '../services/AircraftOrderService';

class UpdateAircraftOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                orderNumber: '',
                totalAmount: '',
                status: ''
        }
        this.updateAircraftOrder = this.updateAircraftOrder.bind(this);

        this.changeorderNumberHandler = this.changeorderNumberHandler.bind(this);
        this.changetotalAmountHandler = this.changetotalAmountHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        AircraftOrderService.getAircraftOrderById(this.state.id).then( (res) =>{
            let aircraftOrder = res.data;
            this.setState({
                orderNumber: aircraftOrder.orderNumber,
                totalAmount: aircraftOrder.totalAmount,
                status: aircraftOrder.status
            });
        });
    }

    updateAircraftOrder = (e) => {
        e.preventDefault();
        let aircraftOrder = {
            aircraftOrderId: this.state.id,
            orderNumber: this.state.orderNumber,
            totalAmount: this.state.totalAmount,
            status: this.state.status
        };
        console.log('aircraftOrder => ' + JSON.stringify(aircraftOrder));
        console.log('id => ' + JSON.stringify(this.state.id));
        AircraftOrderService.updateAircraftOrder(aircraftOrder).then( res => {
            this.props.history.push('/aircraftOrders');
        });
    }

    changeorderNumberHandler= (event) => {
        this.setState({orderNumber: event.target.value});
    }
    changetotalAmountHandler= (event) => {
        this.setState({totalAmount: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/aircraftOrders');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update AircraftOrder</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> orderNumber: </label>
                                                <input placeholder="orderNumber" name="orderNumber" className="form-control" value={this.state.orderNumber} onChange={this.changeorderNumberHandler}/>

                                            <label> totalAmount: </label>
                                                <input placeholder="totalAmount" name="totalAmount" className="form-control" value={this.state.totalAmount} onChange={this.changetotalAmountHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Committed
                      </option>
                      <option name="Status" className="form-control" >
                          InProduction
                      </option>
                      <option name="Status" className="form-control" >
                          Delivered
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateAircraftOrder}>Save</button>
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

export default UpdateAircraftOrderComponent
