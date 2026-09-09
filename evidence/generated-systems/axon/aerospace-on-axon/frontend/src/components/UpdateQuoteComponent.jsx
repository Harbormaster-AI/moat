import React, { Component } from 'react'
import QuoteService from '../services/QuoteService';

class UpdateQuoteComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                quoteNumber: '',
                totalAmount: ''
        }
        this.updateQuote = this.updateQuote.bind(this);

        this.changequoteNumberHandler = this.changequoteNumberHandler.bind(this);
        this.changetotalAmountHandler = this.changetotalAmountHandler.bind(this);
    }

    componentDidMount(){
        QuoteService.getQuoteById(this.state.id).then( (res) =>{
            let quote = res.data;
            this.setState({
                quoteNumber: quote.quoteNumber,
                totalAmount: quote.totalAmount
            });
        });
    }

    updateQuote = (e) => {
        e.preventDefault();
        let quote = {
            quoteId: this.state.id,
            quoteNumber: this.state.quoteNumber,
            totalAmount: this.state.totalAmount
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
    changetotalAmountHandler= (event) => {
        this.setState({totalAmount: event.target.value});
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

                                            <label> totalAmount: </label>
                                                <input placeholder="totalAmount" name="totalAmount" className="form-control" value={this.state.totalAmount} onChange={this.changetotalAmountHandler}/>

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
