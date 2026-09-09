import React, { Component } from 'react'
import TaxRuleService from '../services/TaxRuleService'

class ViewTaxRuleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            taxRule: {}
        }
    }

    componentDidMount(){
        TaxRuleService.getTaxRuleById(this.state.id).then( res => {
            this.setState({taxRule: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View TaxRule Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.taxRule.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> country:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.taxRule.country }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> region:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.taxRule.region }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> rate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.taxRule.rate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> taxInclusive:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.taxRule.taxInclusive }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> TaxClass:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.taxRule.taxClass }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewTaxRuleComponent
