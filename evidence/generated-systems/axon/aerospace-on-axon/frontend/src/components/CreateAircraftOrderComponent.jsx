import React, { Component } from 'react'
import AircraftOrderService from '../services/AircraftOrderService';

class CreateAircraftOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                orderNumber: '',
                totalAmount: '',
                status: ''
        }
        this.changeorderNumberHandler = this.changeorderNumberHandler.bind(this);
        this.changetotalAmountHandler = this.changetotalAmountHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AircraftOrderService.getAircraftOrderById(this.state.id).then( (res) =>{
                let aircraftOrder = res.data;
                this.setState({
                    orderNumber: aircraftOrder.orderNumber,
                    totalAmount: aircraftOrder.totalAmount,
                    status: aircraftOrder.status
                });
            });
        }        
    }
    saveOrUpdateAircraftOrder = (e) => {
        e.preventDefault();
        let aircraftOrder = {
                aircraftOrderId: this.state.id,
                orderNumber: this.state.orderNumber,
                totalAmount: this.state.totalAmount,
                status: this.state.status
            };
        console.log('aircraftOrder => ' + JSON.stringify(aircraftOrder));

        // step 5
        if(this.state.id === '_add'){
            aircraftOrder.aircraftOrderId=''
            AircraftOrderService.createAircraftOrder(aircraftOrder).then(res =>{
                this.props.history.push('/aircraftOrders');
            });
        }else{
            AircraftOrderService.updateAircraftOrder(aircraftOrder).then( res => {
                this.props.history.push('/aircraftOrders');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add AircraftOrder</h3>
        }else{
            return <h3 className="text-center">Update AircraftOrder</h3>
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

                                            <label> totalAmount:&emsp; </label>
                                                <input placeholder="totalAmount" name="totalAmount" className="form-control" value={this.state.totalAmount} onChange={this.changetotalAmountHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAircraftOrder}>Save</button>
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

export default CreateAircraftOrderComponent
