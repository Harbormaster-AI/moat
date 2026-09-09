import React, { Component } from 'react'
import InvoiceService from '../services/InvoiceService';

class UpdateInvoiceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                invoiceNumber: '',
                issuedDate: '',
                dueDate: '',
                total: '',
                status: ''
        }
        this.updateInvoice = this.updateInvoice.bind(this);

        this.changeinvoiceNumberHandler = this.changeinvoiceNumberHandler.bind(this);
        this.changeissuedDateHandler = this.changeissuedDateHandler.bind(this);
        this.changedueDateHandler = this.changedueDateHandler.bind(this);
        this.changetotalHandler = this.changetotalHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        InvoiceService.getInvoiceById(this.state.id).then( (res) =>{
            let invoice = res.data;
            this.setState({
                invoiceNumber: invoice.invoiceNumber,
                issuedDate: invoice.issuedDate,
                dueDate: invoice.dueDate,
                total: invoice.total,
                status: invoice.status
            });
        });
    }

    updateInvoice = (e) => {
        e.preventDefault();
        let invoice = {
            invoiceId: this.state.id,
            invoiceNumber: this.state.invoiceNumber,
            issuedDate: this.state.issuedDate,
            dueDate: this.state.dueDate,
            total: this.state.total,
            status: this.state.status
        };
        console.log('invoice => ' + JSON.stringify(invoice));
        console.log('id => ' + JSON.stringify(this.state.id));
        InvoiceService.updateInvoice(invoice).then( res => {
            this.props.history.push('/invoices');
        });
    }

    changeinvoiceNumberHandler= (event) => {
        this.setState({invoiceNumber: event.target.value});
    }
    changeissuedDateHandler= (event) => {
        this.setState({issuedDate: event.target.value});
    }
    changedueDateHandler= (event) => {
        this.setState({dueDate: event.target.value});
    }
    changetotalHandler= (event) => {
        this.setState({total: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/invoices');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Invoice</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> invoiceNumber: </label>
                                                <input placeholder="invoiceNumber" name="invoiceNumber" className="form-control" value={this.state.invoiceNumber} onChange={this.changeinvoiceNumberHandler}/>

                                            <label> issuedDate: </label>
                                                <input type="date" placeholder="issuedDate" name="issuedDate" className="form-control" value={this.state.issuedDate} onChange={this.changeissuedDateHandler}/>

                                            <label> dueDate: </label>
                                                <input type="date" placeholder="dueDate" name="dueDate" className="form-control" value={this.state.dueDate} onChange={this.changedueDateHandler}/>

                                            <label> total: </label>
                                                <input placeholder="total" name="total" className="form-control" value={this.state.total} onChange={this.changetotalHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Issued
                      </option>
                      <option name="Status" className="form-control" >
                          Paid
                      </option>
                      <option name="Status" className="form-control" >
                          Overdue
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateInvoice}>Save</button>
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

export default UpdateInvoiceComponent
