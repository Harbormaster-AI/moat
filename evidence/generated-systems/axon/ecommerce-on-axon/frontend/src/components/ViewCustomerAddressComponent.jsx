import React, { Component } from 'react'
import CustomerAddressService from '../services/CustomerAddressService'

class ViewCustomerAddressComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            customerAddress: {}
        }
    }

    componentDidMount(){
        CustomerAddressService.getCustomerAddressById(this.state.id).then( res => {
            this.setState({customerAddress: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View CustomerAddress Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> label:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.customerAddress.label }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> address:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.customerAddress.address }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> asDefaultShipping:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.customerAddress.asDefaultShipping }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> asDefaultBilling:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.customerAddress.asDefaultBilling }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCustomerAddressComponent
