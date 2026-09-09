import React, { Component } from 'react'
import SalesOrderLineService from '../services/SalesOrderLineService'

class ListSalesOrderLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                salesOrderLines: []
        }
        this.addSalesOrderLine = this.addSalesOrderLine.bind(this);
        this.editSalesOrderLine = this.editSalesOrderLine.bind(this);
        this.deleteSalesOrderLine = this.deleteSalesOrderLine.bind(this);
    }

    deleteSalesOrderLine(id){
        SalesOrderLineService.deleteSalesOrderLine(id).then( res => {
            this.setState({salesOrderLines: this.state.salesOrderLines.filter(salesOrderLine => salesOrderLine.salesOrderLineId !== id)});
        });
    }
    viewSalesOrderLine(id){
        this.props.history.push(`/view-salesOrderLine/${id}`);
    }
    editSalesOrderLine(id){
        this.props.history.push(`/add-salesOrderLine/${id}`);
    }

    componentDidMount(){
        SalesOrderLineService.getSalesOrderLines().then((res) => {
            this.setState({ salesOrderLines: res.data});
        });
    }

    addSalesOrderLine(){
        this.props.history.push('/add-salesOrderLine/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">SalesOrderLine List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addSalesOrderLine}> Add SalesOrderLine</button>
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
                                    this.state.salesOrderLines.map(
                                        salesOrderLine => 
                                        <tr key = {salesOrderLine.salesOrderLineId}>
                                             <td> { salesOrderLine.lineNumber } </td>
                                             <td> { salesOrderLine.quantity } </td>
                                             <td> { salesOrderLine.unitPrice } </td>
                                             <td> { salesOrderLine.dueDate } </td>
                                             <td>
                                                 <button onClick={ () => this.editSalesOrderLine(salesOrderLine.salesOrderLineId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteSalesOrderLine(salesOrderLine.salesOrderLineId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewSalesOrderLine(salesOrderLine.salesOrderLineId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListSalesOrderLineComponent
