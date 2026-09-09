import React, { Component } from 'react'
import TypeCertificateService from '../services/TypeCertificateService'

class ViewTypeCertificateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            typeCertificate: {}
        }
    }

    componentDidMount(){
        TypeCertificateService.getTypeCertificateById(this.state.id).then( res => {
            this.setState({typeCertificate: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View TypeCertificate Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> certificateNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.typeCertificate.certificateNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> authority:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.typeCertificate.authority }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewTypeCertificateComponent
