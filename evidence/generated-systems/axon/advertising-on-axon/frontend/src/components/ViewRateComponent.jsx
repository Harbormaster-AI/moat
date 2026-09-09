import React, { Component } from 'react'
import RateService from '../services/RateService'

class ViewRateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            rate: {}
        }
    }

    componentDidMount(){
        RateService.getRateById(this.state.id).then( res => {
            this.setState({rate: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Rate Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> unitPrice:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.rate.unitPrice }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> AdFormat:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.rate.adFormat }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> PricingModel:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.rate.pricingModel }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewRateComponent
