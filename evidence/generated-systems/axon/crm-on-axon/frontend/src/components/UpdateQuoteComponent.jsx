import React, { Component } from 'react'
import QuoteService from '../services/QuoteService';

class UpdateQuoteComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                quoteNumber: '',
                validityStart: '',
                validityEnd: '',
                totalAmount: '',
                discountPercent: '',
                taxAmount: '',
                shippingAmount: '',
                status: ''
        }
        this.updateQuote = this.updateQuote.bind(this);

        this.changequoteNumberHandler = this.changequoteNumberHandler.bind(this);
        this.changevalidityStartHandler = this.changevalidityStartHandler.bind(this);
        this.changevalidityEndHandler = this.changevalidityEndHandler.bind(this);
        this.changetotalAmountHandler = this.changetotalAmountHandler.bind(this);
        this.changediscountPercentHandler = this.changediscountPercentHandler.bind(this);
        this.changetaxAmountHandler = this.changetaxAmountHandler.bind(this);
        this.changeshippingAmountHandler = this.changeshippingAmountHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        QuoteService.getQuoteById(this.state.id).then( (res) =>{
            let quote = res.data;
            this.setState({
                quoteNumber: quote.quoteNumber,
                validityStart: quote.validityStart,
                validityEnd: quote.validityEnd,
                totalAmount: quote.totalAmount,
                discountPercent: quote.discountPercent,
                taxAmount: quote.taxAmount,
                shippingAmount: quote.shippingAmount,
                status: quote.status
            });
        });
    }

    updateQuote = (e) => {
        e.preventDefault();
        let quote = {
            quoteId: this.state.id,
            quoteNumber: this.state.quoteNumber,
            validityStart: this.state.validityStart,
            validityEnd: this.state.validityEnd,
            totalAmount: this.state.totalAmount,
            discountPercent: this.state.discountPercent,
            taxAmount: this.state.taxAmount,
            shippingAmount: this.state.shippingAmount,
            status: this.state.status
        };
        console.log('quote => ' + JSON.stringify(quote));
        console.log('id => ' + JSON.stringify(this.state.id));
        QuoteService.updateQuote(quote).then( res => {
            this.props.history.push('/quotes');
        });
    }

    changequoteNumberHandler= (event) => {
        this.setState({quoteNumber: event.target.value});
    }
    changevalidityStartHandler= (event) => {
        this.setState({validityStart: event.target.value});
    }
    changevalidityEndHandler= (event) => {
        this.setState({validityEnd: event.target.value});
    }
    changetotalAmountHandler= (event) => {
        this.setState({totalAmount: event.target.value});
    }
    changediscountPercentHandler= (event) => {
        this.setState({discountPercent: event.target.value});
    }
    changetaxAmountHandler= (event) => {
        this.setState({taxAmount: event.target.value});
    }
    changeshippingAmountHandler= (event) => {
        this.setState({shippingAmount: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/quotes');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Quote</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> quoteNumber: </label>
                                                <input placeholder="quoteNumber" name="quoteNumber" className="form-control" value={this.state.quoteNumber} onChange={this.changequoteNumberHandler}/>

                                            <label> validityStart: </label>
                                                <input type="date" placeholder="validityStart" name="validityStart" className="form-control" value={this.state.validityStart} onChange={this.changevalidityStartHandler}/>

                                            <label> validityEnd: </label>
                                                <input type="date" placeholder="validityEnd" name="validityEnd" className="form-control" value={this.state.validityEnd} onChange={this.changevalidityEndHandler}/>

                                            <label> totalAmount: </label>
                                                <input placeholder="totalAmount" name="totalAmount" className="form-control" value={this.state.totalAmount} onChange={this.changetotalAmountHandler}/>

                                            <label> discountPercent: </label>
                                                <input placeholder="discountPercent" name="discountPercent" className="form-control" value={this.state.discountPercent} onChange={this.changediscountPercentHandler}/>

                                            <label> taxAmount: </label>
                                                <input placeholder="taxAmount" name="taxAmount" className="form-control" value={this.state.taxAmount} onChange={this.changetaxAmountHandler}/>

                                            <label> shippingAmount: </label>
                                                <input placeholder="shippingAmount" name="shippingAmount" className="form-control" value={this.state.shippingAmount} onChange={this.changeshippingAmountHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Presented
                      </option>
                      <option name="Status" className="form-control" >
                          Approved
                      </option>
                      <option name="Status" className="form-control" >
                          Rejected
                      </option>
                      <option name="Status" className="form-control" >
                          Accepted
                      </option>
                      <option name="Status" className="form-control" >
                          Expired
                      </option>
                      <option name="Status" className="form-control" >
                          Withdrawn
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateQuote}>Save</button>
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

export default UpdateQuoteComponent
