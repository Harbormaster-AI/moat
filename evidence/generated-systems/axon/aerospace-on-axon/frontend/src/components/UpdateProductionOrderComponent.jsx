import React, { Component } from 'react'
import ProductionOrderService from '../services/ProductionOrderService';

class UpdateProductionOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                orderNumber: '',
                status: ''
        }
        this.updateProductionOrder = this.updateProductionOrder.bind(this);

        this.changeorderNumberHandler = this.changeorderNumberHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ProductionOrderService.getProductionOrderById(this.state.id).then( (res) =>{
            let productionOrder = res.data;
            this.setState({
                orderNumber: productionOrder.orderNumber,
                status: productionOrder.status
            });
        });
    }

    updateProductionOrder = (e) => {
        e.preventDefault();
        let productionOrder = {
            productionOrderId: this.state.id,
            orderNumber: this.state.orderNumber,
            status: this.state.status
        };
        console.log('productionOrder => ' + JSON.stringify(productionOrder));
        console.log('id => ' + JSON.stringify(this.state.id));
        ProductionOrderService.updateProductionOrder(productionOrder).then( res => {
            this.props.history.push('/productionOrders');
        });
    }

    changeorderNumberHandler= (event) => {
        this.setState({orderNumber: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/productionOrders');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ProductionOrder</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> orderNumber: </label>
                                                <input placeholder="orderNumber" name="orderNumber" className="form-control" value={this.state.orderNumber} onChange={this.changeorderNumberHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Planned
                      </option>
                      <option name="Status" className="form-control" >
                          Released
                      </option>
                      <option name="Status" className="form-control" >
                          InAssembly
                      </option>
                      <option name="Status" className="form-control" >
                          FlightTest
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateProductionOrder}>Save</button>
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

export default UpdateProductionOrderComponent
