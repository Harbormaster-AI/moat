import React, { Component } from 'react'
import FXQuoteService from '../services/FXQuoteService';

class CreateFXQuoteComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                baseCurrency: '',
                quoteCurrency: '',
                rate: '',
                quotedAt: '',
                expiresAt: '',
                priceType: ''
        }
        this.changebaseCurrencyHandler = this.changebaseCurrencyHandler.bind(this);
        this.changequoteCurrencyHandler = this.changequoteCurrencyHandler.bind(this);
        this.changerateHandler = this.changerateHandler.bind(this);
        this.changequotedAtHandler = this.changequotedAtHandler.bind(this);
        this.changeexpiresAtHandler = this.changeexpiresAtHandler.bind(this);
        this.changePriceTypeHandler = this.changePriceTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateFXQuote = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            fXQuote.fXQuoteId=''
            FXQuoteService.createFXQuote(fXQuote).then(res =>{
                this.props.history.push('/fXQuotes');
            });
        }else{
            FXQuoteService.updateFXQuote(fXQuote).then( res => {
                this.props.history.push('/fXQuotes');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add FXQuote</h3>
        }else{
            return <h3 className="text-center">Update FXQuote</h3>
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
                                            <label> baseCurrency:&emsp; </label>
                                                <input placeholder="baseCurrency" name="baseCurrency" className="form-control" value={this.state.baseCurrency} onChange={this.changebaseCurrencyHandler}/>

                                            <label> quoteCurrency:&emsp; </label>
                                                <input placeholder="quoteCurrency" name="quoteCurrency" className="form-control" value={this.state.quoteCurrency} onChange={this.changequoteCurrencyHandler}/>

                                            <label> rate:&emsp; </label>
                                                <input placeholder="rate" name="rate" className="form-control" value={this.state.rate} onChange={this.changerateHandler}/>

                                            <label> quotedAt:&emsp; </label>
                                                <input type="time" placeholder="quotedAt" name="quotedAt" className="form-control" value={this.state.quotedAt} onChange={this.changequotedAtHandler}/>

                                            <label> expiresAt:&emsp; </label>
                                                <input type="time" placeholder="expiresAt" name="expiresAt" className="form-control" value={this.state.expiresAt} onChange={this.changeexpiresAtHandler}/>

                                            <label> PriceType:&emsp; </label>
                                                <select value={this.state.priceType} onChange={this.changePriceTypeHandler}>
                      <option name="PriceType" className="form-control" >
                          Indicative
                      </option>
                      <option name="PriceType" className="form-control" >
                          Firm
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateFXQuote}>Save</button>
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

export default CreateFXQuoteComponent
