import React, { Component } from 'react'
import PriceBookEntryService from '../services/PriceBookEntryService'

class ViewPriceBookEntryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            priceBookEntry: {}
        }
    }

    componentDidMount(){
        PriceBookEntryService.getPriceBookEntryById(this.state.id).then( res => {
            this.setState({priceBookEntry: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View PriceBookEntry Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> unitPrice:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.priceBookEntry.unitPrice }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> effectiveDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.priceBookEntry.effectiveDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> expirationDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.priceBookEntry.expirationDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> asActive:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.priceBookEntry.asActive }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPriceBookEntryComponent
