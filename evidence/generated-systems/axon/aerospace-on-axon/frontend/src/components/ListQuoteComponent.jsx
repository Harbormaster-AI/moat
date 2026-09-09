import React, { Component } from 'react'
import QuoteService from '../services/QuoteService'

class ListQuoteComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                quotes: []
        }
        this.addQuote = this.addQuote.bind(this);
        this.editQuote = this.editQuote.bind(this);
        this.deleteQuote = this.deleteQuote.bind(this);
    }

    deleteQuote(id){
        QuoteService.deleteQuote(id).then( res => {
            this.setState({quotes: this.state.quotes.filter(quote => quote.quoteId !== id)});
        });
    }
    viewQuote(id){
        this.props.history.push(`/view-quote/${id}`);
    }
    editQuote(id){
        this.props.history.push(`/add-quote/${id}`);
    }

    componentDidMount(){
        QuoteService.getQuotes().then((res) => {
            this.setState({ quotes: res.data});
        });
    }

    addQuote(){
        this.props.history.push('/add-quote/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Quote List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addQuote}> Add Quote</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> QuoteNumber </th>
                                    <th> TotalAmount </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.quotes.map(
                                        quote => 
                                        <tr key = {quote.quoteId}>
                                             <td> { quote.quoteNumber } </td>
                                             <td> { quote.totalAmount } </td>
                                             <td>
                                                 <button onClick={ () => this.editQuote(quote.quoteId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteQuote(quote.quoteId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewQuote(quote.quoteId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListQuoteComponent
