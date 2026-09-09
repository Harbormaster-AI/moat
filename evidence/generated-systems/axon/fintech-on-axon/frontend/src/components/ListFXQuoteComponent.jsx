import React, { Component } from 'react'
import FXQuoteService from '../services/FXQuoteService'

class ListFXQuoteComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                fXQuotes: []
        }
        this.addFXQuote = this.addFXQuote.bind(this);
        this.editFXQuote = this.editFXQuote.bind(this);
        this.deleteFXQuote = this.deleteFXQuote.bind(this);
    }

    deleteFXQuote(id){
        FXQuoteService.deleteFXQuote(id).then( res => {
            this.setState({fXQuotes: this.state.fXQuotes.filter(fXQuote => fXQuote.fXQuoteId !== id)});
        });
    }
    viewFXQuote(id){
        this.props.history.push(`/view-fXQuote/${id}`);
    }
    editFXQuote(id){
        this.props.history.push(`/add-fXQuote/${id}`);
    }

    componentDidMount(){
        FXQuoteService.getFXQuotes().then((res) => {
            this.setState({ fXQuotes: res.data});
        });
    }

    addFXQuote(){
        this.props.history.push('/add-fXQuote/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">FXQuote List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addFXQuote}> Add FXQuote</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> BaseCurrency </th>
                                    <th> QuoteCurrency </th>
                                    <th> Rate </th>
                                    <th> QuotedAt </th>
                                    <th> ExpiresAt </th>
                                    <th> PriceType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.fXQuotes.map(
                                        fXQuote => 
                                        <tr key = {fXQuote.fXQuoteId}>
                                             <td> { fXQuote.baseCurrency } </td>
                                             <td> { fXQuote.quoteCurrency } </td>
                                             <td> { fXQuote.rate } </td>
                                             <td> { fXQuote.quotedAt } </td>
                                             <td> { fXQuote.expiresAt } </td>
                                             <td> { fXQuote.priceType } </td>
                                             <td>
                                                 <button onClick={ () => this.editFXQuote(fXQuote.fXQuoteId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteFXQuote(fXQuote.fXQuoteId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewFXQuote(fXQuote.fXQuoteId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListFXQuoteComponent
