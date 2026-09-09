import React, { Component } from 'react'
import ProductionOrderService from '../services/ProductionOrderService'

class ListProductionOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                productionOrders: []
        }
        this.addProductionOrder = this.addProductionOrder.bind(this);
        this.editProductionOrder = this.editProductionOrder.bind(this);
        this.deleteProductionOrder = this.deleteProductionOrder.bind(this);
    }

    deleteProductionOrder(id){
        ProductionOrderService.deleteProductionOrder(id).then( res => {
            this.setState({productionOrders: this.state.productionOrders.filter(productionOrder => productionOrder.productionOrderId !== id)});
        });
    }
    viewProductionOrder(id){
        this.props.history.push(`/view-productionOrder/${id}`);
    }
    editProductionOrder(id){
        this.props.history.push(`/add-productionOrder/${id}`);
    }

    componentDidMount(){
        ProductionOrderService.getProductionOrders().then((res) => {
            this.setState({ productionOrders: res.data});
        });
    }

    addProductionOrder(){
        this.props.history.push('/add-productionOrder/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ProductionOrder List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addProductionOrder}> Add ProductionOrder</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> OrderNumber </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.productionOrders.map(
                                        productionOrder => 
                                        <tr key = {productionOrder.productionOrderId}>
                                             <td> { productionOrder.orderNumber } </td>
                                             <td> { productionOrder.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editProductionOrder(productionOrder.productionOrderId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteProductionOrder(productionOrder.productionOrderId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewProductionOrder(productionOrder.productionOrderId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListProductionOrderComponent
