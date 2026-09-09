import React, { Component } from 'react'
import KYCDocumentService from '../services/KYCDocumentService'

class ViewKYCDocumentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            kYCDocument: {}
        }
    }

    componentDidMount(){
        KYCDocumentService.getKYCDocumentById(this.state.id).then( res => {
            this.setState({kYCDocument: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View KYCDocument Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> reference:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.kYCDocument.reference }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> issuedCountry:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.kYCDocument.issuedCountry }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> expirationDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.kYCDocument.expirationDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> DocumentType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.kYCDocument.documentType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.kYCDocument.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewKYCDocumentComponent
