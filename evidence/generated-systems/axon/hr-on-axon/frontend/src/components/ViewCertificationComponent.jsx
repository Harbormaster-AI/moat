import React, { Component } from 'react'
import CertificationService from '../services/CertificationService'

class ViewCertificationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            certification: {}
        }
    }

    componentDidMount(){
        CertificationService.getCertificationById(this.state.id).then( res => {
            this.setState({certification: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Certification Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.certification.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> issuer:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.certification.issuer }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> validFrom:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.certification.validFrom }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> validTo:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.certification.validTo }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> credentialId:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.certification.credentialId }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCertificationComponent
