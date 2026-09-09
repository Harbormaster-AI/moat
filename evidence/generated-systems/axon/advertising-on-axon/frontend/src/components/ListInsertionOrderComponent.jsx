import React, { Component } from 'react'
import InsertionOrderService from '../services/InsertionOrderService'

class ListInsertionOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                insertionOrders: []
        }
        this.addInsertionOrder = this.addInsertionOrder.bind(this);
        this.editInsertionOrder = this.editInsertionOrder.bind(this);
        this.deleteInsertionOrder = this.deleteInsertionOrder.bind(this);
    }

    deleteInsertionOrder(id){
        InsertionOrderService.deleteInsertionOrder(id).then( res => {
            this.setState({insertionOrders: this.state.insertionOrders.filter(insertionOrder => insertionOrder.insertionOrderId !== id)});
        });
    }
    viewInsertionOrder(id){
        this.props.history.push(`/view-insertionOrder/${id}`);
    }
    editInsertionOrder(id){
        this.props.history.push(`/add-insertionOrder/${id}`);
    }

    componentDidMount(){
        InsertionOrderService.getInsertionOrders().then((res) => {
            this.setState({ insertionOrders: res.data});
        });
    }

    addInsertionOrder(){
        this.props.history.push('/add-insertionOrder/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">InsertionOrder List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addInsertionOrder}> Add InsertionOrder</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> IoNumber </th>
                                    <th> AgreedBudget </th>
                                    <th> Flight </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.insertionOrders.map(
                                        insertionOrder => 
                                        <tr key = {insertionOrder.insertionOrderId}>
                                             <td> { insertionOrder.ioNumber } </td>
                                             <td> { insertionOrder.agreedBudget } </td>
                                             <td> { insertionOrder.flight } </td>
                                             <td> { insertionOrder.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editInsertionOrder(insertionOrder.insertionOrderId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteInsertionOrder(insertionOrder.insertionOrderId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewInsertionOrder(insertionOrder.insertionOrderId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListInsertionOrderComponent
