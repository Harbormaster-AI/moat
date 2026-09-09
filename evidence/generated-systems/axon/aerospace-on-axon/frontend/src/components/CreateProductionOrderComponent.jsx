import React, { Component } from 'react'
import ProductionOrderService from '../services/ProductionOrderService';

class CreateProductionOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                orderNumber: '',
                status: ''
        }
        this.changeorderNumberHandler = this.changeorderNumberHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ProductionOrderService.getProductionOrderById(this.state.id).then( (res) =>{
                let productionOrder = res.data;
                this.setState({
                    orderNumber: productionOrder.orderNumber,
                    status: productionOrder.status
                });
            });
        }        
    }
    saveOrUpdateProductionOrder = (e) => {
        e.preventDefault();
        let productionOrder = {
                productionOrderId: this.state.id,
                orderNumber: this.state.orderNumber,
                status: this.state.status
            };
        console.log('productionOrder => ' + JSON.stringify(productionOrder));

        // step 5
        if(this.state.id === '_add'){
            productionOrder.productionOrderId=''
            ProductionOrderService.createProductionOrder(productionOrder).then(res =>{
                this.props.history.push('/productionOrders');
            });
        }else{
            ProductionOrderService.updateProductionOrder(productionOrder).then( res => {
                this.props.history.push('/productionOrders');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ProductionOrder</h3>
        }else{
            return <h3 className="text-center">Update ProductionOrder</h3>
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

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateProductionOrder}>Save</button>
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

export default CreateProductionOrderComponent
