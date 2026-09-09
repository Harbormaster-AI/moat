import React, { Component } from 'react'
import InvoiceService from '../services/InvoiceService'

class ListInvoiceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                invoices: []
        }
        this.addInvoice = this.addInvoice.bind(this);
        this.editInvoice = this.editInvoice.bind(this);
        this.deleteInvoice = this.deleteInvoice.bind(this);
    }

    deleteInvoice(id){
        InvoiceService.deleteInvoice(id).then( res => {
            this.setState({invoices: this.state.invoices.filter(invoice => invoice.invoiceId !== id)});
        });
    }
    viewInvoice(id){
        this.props.history.push(`/view-invoice/${id}`);
    }
    editInvoice(id){
        this.props.history.push(`/add-invoice/${id}`);
    }

    componentDidMount(){
        InvoiceService.getInvoices().then((res) => {
            this.setState({ invoices: res.data});
        });
    }

    addInvoice(){
        this.props.history.push('/add-invoice/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Invoice List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addInvoice}> Add Invoice</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> InvoiceNumber </th>
                                    <th> IssueDate </th>
                                    <th> DueDate </th>
                                    <th> Total </th>
                                    <th> Currency </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.invoices.map(
                                        invoice => 
                                        <tr key = {invoice.invoiceId}>
                                             <td> { invoice.invoiceNumber } </td>
                                             <td> { invoice.issueDate } </td>
                                             <td> { invoice.dueDate } </td>
                                             <td> { invoice.total } </td>
                                             <td> { invoice.currency } </td>
                                             <td> { invoice.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editInvoice(invoice.invoiceId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteInvoice(invoice.invoiceId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewInvoice(invoice.invoiceId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListInvoiceComponent
