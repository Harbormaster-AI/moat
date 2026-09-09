import React, { Component } from 'react'
import QuoteService from '../services/QuoteService'

class ViewQuoteComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            quote: {}
        }
    }

    componentDidMount(){
        QuoteService.getQuoteById(this.state.id).then( res => {
            this.setState({quote: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Quote Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quoteNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.quote.quoteNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> totalPremium:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.quote.totalPremium }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ratingDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.quote.ratingDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> asBound:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.quote.asBound }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewQuoteComponent
