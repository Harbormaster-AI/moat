import React, { Component } from 'react'
import PurchaseOrderLineService from '../services/PurchaseOrderLineService'

class ListPurchaseOrderLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                purchaseOrderLines: []
        }
        this.addPurchaseOrderLine = this.addPurchaseOrderLine.bind(this);
        this.editPurchaseOrderLine = this.editPurchaseOrderLine.bind(this);
        this.deletePurchaseOrderLine = this.deletePurchaseOrderLine.bind(this);
    }

    deletePurchaseOrderLine(id){
        PurchaseOrderLineService.deletePurchaseOrderLine(id).then( res => {
            this.setState({purchaseOrderLines: this.state.purchaseOrderLines.filter(purchaseOrderLine => purchaseOrderLine.purchaseOrderLineId !== id)});
        });
    }
    viewPurchaseOrderLine(id){
        this.props.history.push(`/view-purchaseOrderLine/${id}`);
    }
    editPurchaseOrderLine(id){
        this.props.history.push(`/add-purchaseOrderLine/${id}`);
    }

    componentDidMount(){
        PurchaseOrderLineService.getPurchaseOrderLines().then((res) => {
            this.setState({ purchaseOrderLines: res.data});
        });
    }

    addPurchaseOrderLine(){
        this.props.history.push('/add-purchaseOrderLine/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">PurchaseOrderLine List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPurchaseOrderLine}> Add PurchaseOrderLine</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> LineNumber </th>
                                    <th> Quantity </th>
                                    <th> UnitPrice </th>
                                    <th> DueDate </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.purchaseOrderLines.map(
                                        purchaseOrderLine => 
                                        <tr key = {purchaseOrderLine.purchaseOrderLineId}>
                                             <td> { purchaseOrderLine.lineNumber } </td>
                                             <td> { purchaseOrderLine.quantity } </td>
                                             <td> { purchaseOrderLine.unitPrice } </td>
                                             <td> { purchaseOrderLine.dueDate } </td>
                                             <td>
                                                 <button onClick={ () => this.editPurchaseOrderLine(purchaseOrderLine.purchaseOrderLineId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePurchaseOrderLine(purchaseOrderLine.purchaseOrderLineId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPurchaseOrderLine(purchaseOrderLine.purchaseOrderLineId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPurchaseOrderLineComponent
