import React, { Component } from 'react'
import OrderLineService from '../services/OrderLineService'

class ListOrderLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                orderLines: []
        }
        this.addOrderLine = this.addOrderLine.bind(this);
        this.editOrderLine = this.editOrderLine.bind(this);
        this.deleteOrderLine = this.deleteOrderLine.bind(this);
    }

    deleteOrderLine(id){
        OrderLineService.deleteOrderLine(id).then( res => {
            this.setState({orderLines: this.state.orderLines.filter(orderLine => orderLine.orderLineId !== id)});
        });
    }
    viewOrderLine(id){
        this.props.history.push(`/view-orderLine/${id}`);
    }
    editOrderLine(id){
        this.props.history.push(`/add-orderLine/${id}`);
    }

    componentDidMount(){
        OrderLineService.getOrderLines().then((res) => {
            this.setState({ orderLines: res.data});
        });
    }

    addOrderLine(){
        this.props.history.push('/add-orderLine/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">OrderLine List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addOrderLine}> Add OrderLine</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Quantity </th>
                                    <th> UnitPrice </th>
                                    <th> TotalPrice </th>
                                    <th> TaxRate </th>
                                    <th> LineStatus </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.orderLines.map(
                                        orderLine => 
                                        <tr key = {orderLine.orderLineId}>
                                             <td> { orderLine.quantity } </td>
                                             <td> { orderLine.unitPrice } </td>
                                             <td> { orderLine.totalPrice } </td>
                                             <td> { orderLine.taxRate } </td>
                                             <td> { orderLine.lineStatus } </td>
                                             <td>
                                                 <button onClick={ () => this.editOrderLine(orderLine.orderLineId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteOrderLine(orderLine.orderLineId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewOrderLine(orderLine.orderLineId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListOrderLineComponent
