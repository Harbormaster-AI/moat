import React, { Component } from 'react'
import VerifiedAddressService from '../services/VerifiedAddressService'

class ViewVerifiedAddressComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            verifiedAddress: {}
        }
    }

    componentDidMount(){
        VerifiedAddressService.getVerifiedAddressById(this.state.id).then( res => {
            this.setState({verifiedAddress: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View VerifiedAddress Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> address:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.verifiedAddress.address }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> verifiedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.verifiedAddress.verifiedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> VerificationStatus:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.verifiedAddress.verificationStatus }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewVerifiedAddressComponent
