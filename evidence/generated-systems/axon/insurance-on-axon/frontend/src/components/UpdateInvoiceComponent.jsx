import React, { Component } from 'react'
import InvoiceService from '../services/InvoiceService';

class UpdateInvoiceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                invoiceNumber: '',
                dueDate: '',
                totalDue: '',
                status: ''
        }
        this.updateInvoice = this.updateInvoice.bind(this);

        this.changeinvoiceNumberHandler = this.changeinvoiceNumberHandler.bind(this);
        this.changedueDateHandler = this.changedueDateHandler.bind(this);
        this.changetotalDueHandler = this.changetotalDueHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        InvoiceService.getInvoiceById(this.state.id).then( (res) =>{
            let invoice = res.data;
            this.setState({
                invoiceNumber: invoice.invoiceNumber,
                dueDate: invoice.dueDate,
                totalDue: invoice.totalDue,
                status: invoice.status
            });
        });
    }

    updateInvoice = (e) => {
        e.preventDefault();
        let invoice = {
            invoiceId: this.state.id,
            invoiceNumber: this.state.invoiceNumber,
            dueDate: this.state.dueDate,
            totalDue: this.state.totalDue,
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
    changedueDateHandler= (event) => {
        this.setState({dueDate: event.target.value});
    }
    changetotalDueHandler= (event) => {
        this.setState({totalDue: event.target.value});
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

                                            <label> dueDate: </label>
                                                <input type="date" placeholder="dueDate" name="dueDate" className="form-control" value={this.state.dueDate} onChange={this.changedueDateHandler}/>

                                            <label> totalDue: </label>
                                                <input placeholder="totalDue" name="totalDue" className="form-control" value={this.state.totalDue} onChange={this.changetotalDueHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Open
                      </option>
                      <option name="Status" className="form-control" >
                          Paid
                      </option>
                      <option name="Status" className="form-control" >
                          PartiallyPaid
                      </option>
                      <option name="Status" className="form-control" >
                          Void
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
