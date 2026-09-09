import React, { Component } from 'react'
import ConsentService from '../services/ConsentService'

class ViewConsentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            consent: {}
        }
    }

    componentDidMount(){
        ConsentService.getConsentById(this.state.id).then( res => {
            this.setState({consent: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Consent Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> grantedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.consent.grantedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> expiresAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.consent.expiresAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> scope:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.consent.scope }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ConsentType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.consent.consentType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.consent.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewConsentComponent
