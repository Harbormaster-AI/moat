import React, { Component } from 'react'
import MerchantService from '../services/MerchantService'

class ViewMerchantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            merchant: {}
        }
    }

    componentDidMount(){
        MerchantService.getMerchantById(this.state.id).then( res => {
            this.setState({merchant: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Merchant Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.merchant.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> mcc:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.merchant.mcc }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> url:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.merchant.url }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> country:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.merchant.country }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> settlementCurrency:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.merchant.settlementCurrency }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewMerchantComponent
