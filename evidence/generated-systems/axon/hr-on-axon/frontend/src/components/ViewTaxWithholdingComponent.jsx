import React, { Component } from 'react'
import TaxWithholdingService from '../services/TaxWithholdingService'

class ViewTaxWithholdingComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            taxWithholding: {}
        }
    }

    componentDidMount(){
        TaxWithholdingService.getTaxWithholdingById(this.state.id).then( res => {
            this.setState({taxWithholding: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View TaxWithholding Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> taxId:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.taxWithholding.taxId }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> allowances:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.taxWithholding.allowances }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> additionalAmount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.taxWithholding.additionalAmount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> FilingStatus:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.taxWithholding.filingStatus }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewTaxWithholdingComponent
