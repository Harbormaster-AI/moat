import React, { Component } from 'react'
import QuoteService from '../services/QuoteService';

class UpdateQuoteComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                quoteNumber: '',
                totalPremium: '',
                ratingDate: '',
                asBound: ''
        }
        this.updateQuote = this.updateQuote.bind(this);

        this.changequoteNumberHandler = this.changequoteNumberHandler.bind(this);
        this.changetotalPremiumHandler = this.changetotalPremiumHandler.bind(this);
        this.changeratingDateHandler = this.changeratingDateHandler.bind(this);
        this.changeasBoundHandler = this.changeasBoundHandler.bind(this);
    }

    componentDidMount(){
        QuoteService.getQuoteById(this.state.id).then( (res) =>{
            let quote = res.data;
            this.setState({
                quoteNumber: quote.quoteNumber,
                totalPremium: quote.totalPremium,
                ratingDate: quote.ratingDate,
                asBound: quote.asBound
            });
        });
    }

    updateQuote = (e) => {
        e.preventDefault();
        let quote = {
            quoteId: this.state.id,
            quoteNumber: this.state.quoteNumber,
            totalPremium: this.state.totalPremium,
            ratingDate: this.state.ratingDate,
            asBound: this.state.asBound
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
    changetotalPremiumHandler= (event) => {
        this.setState({totalPremium: event.target.value});
    }
    changeratingDateHandler= (event) => {
        this.setState({ratingDate: event.target.value});
    }
    changeasBoundHandler= (event) => {
        this.setState({asBound: event.target.value});
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

                                            <label> totalPremium: </label>
                                                <input placeholder="totalPremium" name="totalPremium" className="form-control" value={this.state.totalPremium} onChange={this.changetotalPremiumHandler}/>

                                            <label> ratingDate: </label>
                                                <input type="date" placeholder="ratingDate" name="ratingDate" className="form-control" value={this.state.ratingDate} onChange={this.changeratingDateHandler}/>

                                            <label> asBound: </label>
                                                <input type="checkbox" placeholder="asBound" name="asBound" className="form-control" value={this.state.asBound} onChange={this.changeasBoundHandler}/>


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
