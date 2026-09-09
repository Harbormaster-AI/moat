import React, { Component } from 'react'
import PaymentProviderService from '../services/PaymentProviderService'

class ViewPaymentProviderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            paymentProvider: {}
        }
    }

    componentDidMount(){
        PaymentProviderService.getPaymentProviderById(this.state.id).then( res => {
            this.setState({paymentProvider: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View PaymentProvider Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.paymentProvider.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> enabled:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.paymentProvider.enabled }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> merchantAccountId:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.paymentProvider.merchantAccountId }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ProviderType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.paymentProvider.providerType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPaymentProviderComponent
