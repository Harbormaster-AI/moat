import React, { Component } from 'react'
import DocumentService from '../services/DocumentService'

class ViewDocumentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            document: {}
        }
    }

    componentDidMount(){
        DocumentService.getDocumentById(this.state.id).then( res => {
            this.setState({document: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Document Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> fileName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.document.fileName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> uploadedDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.document.uploadedDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> DocumentType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.document.documentType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewDocumentComponent
