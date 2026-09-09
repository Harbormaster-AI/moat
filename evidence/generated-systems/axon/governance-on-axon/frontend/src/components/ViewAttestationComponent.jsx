import React, { Component } from 'react'
import AttestationService from '../services/AttestationService'

class ViewAttestationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            attestation: {}
        }
    }

    componentDidMount(){
        AttestationService.getAttestationById(this.state.id).then( res => {
            this.setState({attestation: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Attestation Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> statement:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.attestation.statement }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> attestor:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.attestation.attestor }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> dateSigned:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.attestation.dateSigned }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Result:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.attestation.result }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAttestationComponent
