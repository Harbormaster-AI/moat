import React, { Component } from 'react'
import InvoiceService from '../services/InvoiceService';

class CreateInvoiceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                invoiceNumber: '',
                issueDate: '',
                dueDate: '',
                total: '',
                currency: '',
                status: ''
        }
        this.changeinvoiceNumberHandler = this.changeinvoiceNumberHandler.bind(this);
        this.changeissueDateHandler = this.changeissueDateHandler.bind(this);
        this.changedueDateHandler = this.changedueDateHandler.bind(this);
        this.changetotalHandler = this.changetotalHandler.bind(this);
        this.changecurrencyHandler = this.changecurrencyHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            InvoiceService.getInvoiceById(this.state.id).then( (res) =>{
                let invoice = res.data;
                this.setState({
                    invoiceNumber: invoice.invoiceNumber,
                    issueDate: invoice.issueDate,
                    dueDate: invoice.dueDate,
                    total: invoice.total,
                    currency: invoice.currency,
                    status: invoice.status
                });
            });
        }        
    }
    saveOrUpdateInvoice = (e) => {
        e.preventDefault();
        let invoice = {
                invoiceId: this.state.id,
                invoiceNumber: this.state.invoiceNumber,
                issueDate: this.state.issueDate,
                dueDate: this.state.dueDate,
                total: this.state.total,
                currency: this.state.currency,
                status: this.state.status
            };
        console.log('invoice => ' + JSON.stringify(invoice));

        // step 5
        if(this.state.id === '_add'){
            invoice.invoiceId=''
            InvoiceService.createInvoice(invoice).then(res =>{
                this.props.history.push('/invoices');
            });
        }else{
            InvoiceService.updateInvoice(invoice).then( res => {
                this.props.history.push('/invoices');
            });
        }
    }
    
    changeinvoiceNumberHandler= (event) => {
        this.setState({invoiceNumber: event.target.value});
    }
    changeissueDateHandler= (event) => {
        this.setState({issueDate: event.target.value});
    }
    changedueDateHandler= (event) => {
        this.setState({dueDate: event.target.value});
    }
    changetotalHandler= (event) => {
        this.setState({total: event.target.value});
    }
    changecurrencyHandler= (event) => {
        this.setState({currency: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/invoices');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Invoice</h3>
        }else{
            return <h3 className="text-center">Update Invoice</h3>
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
                                            <label> invoiceNumber:&emsp; </label>
                                                <input placeholder="invoiceNumber" name="invoiceNumber" className="form-control" value={this.state.invoiceNumber} onChange={this.changeinvoiceNumberHandler}/>

                                            <label> issueDate:&emsp; </label>
                                                <input type="date" placeholder="issueDate" name="issueDate" className="form-control" value={this.state.issueDate} onChange={this.changeissueDateHandler}/>

                                            <label> dueDate:&emsp; </label>
                                                <input type="date" placeholder="dueDate" name="dueDate" className="form-control" value={this.state.dueDate} onChange={this.changedueDateHandler}/>

                                            <label> total:&emsp; </label>
                                                <input placeholder="total" name="total" className="form-control" value={this.state.total} onChange={this.changetotalHandler}/>

                                            <label> currency:&emsp; </label>
                                                <input placeholder="currency" name="currency" className="form-control" value={this.state.currency} onChange={this.changecurrencyHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateInvoice}>Save</button>
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

export default CreateInvoiceComponent
