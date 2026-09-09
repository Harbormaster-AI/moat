import React, { Component } from 'react'
import PaymentProcessorService from '../services/PaymentProcessorService'

class ViewPaymentProcessorComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            paymentProcessor: {}
        }
    }

    componentDidMount(){
        PaymentProcessorService.getPaymentProcessorById(this.state.id).then( res => {
            this.setState({paymentProcessor: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View PaymentProcessor Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.paymentProcessor.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> processorCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.paymentProcessor.processorCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> networkSupport:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.paymentProcessor.networkSupport }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPaymentProcessorComponent
