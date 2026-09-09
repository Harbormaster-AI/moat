import React, { Component } from 'react'
import FXQuoteService from '../services/FXQuoteService';

class UpdateFXQuoteComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                baseCurrency: '',
                quoteCurrency: '',
                rate: '',
                quotedAt: '',
                expiresAt: '',
                priceType: ''
        }
        this.updateFXQuote = this.updateFXQuote.bind(this);

        this.changebaseCurrencyHandler = this.changebaseCurrencyHandler.bind(this);
        this.changequoteCurrencyHandler = this.changequoteCurrencyHandler.bind(this);
        this.changerateHandler = this.changerateHandler.bind(this);
        this.changequotedAtHandler = this.changequotedAtHandler.bind(this);
        this.changeexpiresAtHandler = this.changeexpiresAtHandler.bind(this);
        this.changePriceTypeHandler = this.changePriceTypeHandler.bind(this);
    }

    componentDidMount(){
        FXQuoteService.getFXQuoteById(this.state.id).then( (res) =>{
            let fXQuote = res.data;
            this.setState({
                baseCurrency: fXQuote.baseCurrency,
                quoteCurrency: fXQuote.quoteCurrency,
                rate: fXQuote.rate,
                quotedAt: fXQuote.quotedAt,
                expiresAt: fXQuote.expiresAt,
                priceType: fXQuote.priceType
            });
        });
    }

    updateFXQuote = (e) => {
        e.preventDefault();
        let fXQuote = {
            fXQuoteId: this.state.id,
            baseCurrency: this.state.baseCurrency,
            quoteCurrency: this.state.quoteCurrency,
            rate: this.state.rate,
            quotedAt: this.state.quotedAt,
            expiresAt: this.state.expiresAt,
            priceType: this.state.priceType
        };
        console.log('fXQuote => ' + JSON.stringify(fXQuote));
        console.log('id => ' + JSON.stringify(this.state.id));
        FXQuoteService.updateFXQuote(fXQuote).then( res => {
            this.props.history.push('/fXQuotes');
        });
    }

    changebaseCurrencyHandler= (event) => {
        this.setState({baseCurrency: event.target.value});
    }
    changequoteCurrencyHandler= (event) => {
        this.setState({quoteCurrency: event.target.value});
    }
    changerateHandler= (event) => {
        this.setState({rate: event.target.value});
    }
    changequotedAtHandler= (event) => {
        this.setState({quotedAt: event.target.value});
    }
    changeexpiresAtHandler= (event) => {
        this.setState({expiresAt: event.target.value});
    }
    changePriceTypeHandler= (event) => {
        this.setState({priceType: event.target.value});
    }

    cancel(){
        this.props.history.push('/fXQuotes');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update FXQuote</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> baseCurrency: </label>
                                                <input placeholder="baseCurrency" name="baseCurrency" className="form-control" value={this.state.baseCurrency} onChange={this.changebaseCurrencyHandler}/>

                                            <label> quoteCurrency: </label>
                                                <input placeholder="quoteCurrency" name="quoteCurrency" className="form-control" value={this.state.quoteCurrency} onChange={this.changequoteCurrencyHandler}/>

                                            <label> rate: </label>
                                                <input placeholder="rate" name="rate" className="form-control" value={this.state.rate} onChange={this.changerateHandler}/>

                                            <label> quotedAt: </label>
                                                <input type="time" placeholder="quotedAt" name="quotedAt" className="form-control" value={this.state.quotedAt} onChange={this.changequotedAtHandler}/>

                                            <label> expiresAt: </label>
                                                <input type="time" placeholder="expiresAt" name="expiresAt" className="form-control" value={this.state.expiresAt} onChange={this.changeexpiresAtHandler}/>

                                            <label> PriceType: </label>
                                                <select value={this.state.priceType} onChange={this.changePriceTypeHandler}>
                      <option name="PriceType" className="form-control" >
                          Indicative
                      </option>
                      <option name="PriceType" className="form-control" >
                          Firm
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateFXQuote}>Save</button>
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

export default UpdateFXQuoteComponent
