import React, { Component } from 'react'
import FXQuoteService from '../services/FXQuoteService'

class ViewFXQuoteComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            fXQuote: {}
        }
    }

    componentDidMount(){
        FXQuoteService.getFXQuoteById(this.state.id).then( res => {
            this.setState({fXQuote: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View FXQuote Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> baseCurrency:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.fXQuote.baseCurrency }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quoteCurrency:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.fXQuote.quoteCurrency }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> rate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.fXQuote.rate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quotedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.fXQuote.quotedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> expiresAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.fXQuote.expiresAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> PriceType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.fXQuote.priceType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewFXQuoteComponent
