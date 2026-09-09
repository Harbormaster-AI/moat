import React, { Component } from 'react'
import DocumentService from '../services/DocumentService';

class UpdateDocumentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                fileName: '',
                uploadedDate: '',
                documentType: ''
        }
        this.updateDocument = this.updateDocument.bind(this);

        this.changefileNameHandler = this.changefileNameHandler.bind(this);
        this.changeuploadedDateHandler = this.changeuploadedDateHandler.bind(this);
        this.changeDocumentTypeHandler = this.changeDocumentTypeHandler.bind(this);
    }

    componentDidMount(){
        DocumentService.getDocumentById(this.state.id).then( (res) =>{
            let document = res.data;
            this.setState({
                fileName: document.fileName,
                uploadedDate: document.uploadedDate,
                documentType: document.documentType
            });
        });
    }

    updateDocument = (e) => {
        e.preventDefault();
        let document = {
            documentId: this.state.id,
            fileName: this.state.fileName,
            uploadedDate: this.state.uploadedDate,
            documentType: this.state.documentType
        };
        console.log('document => ' + JSON.stringify(document));
        console.log('id => ' + JSON.stringify(this.state.id));
        DocumentService.updateDocument(document).then( res => {
            this.props.history.push('/documents');
        });
    }

    changefileNameHandler= (event) => {
        this.setState({fileName: event.target.value});
    }
    changeuploadedDateHandler= (event) => {
        this.setState({uploadedDate: event.target.value});
    }
    changeDocumentTypeHandler= (event) => {
        this.setState({documentType: event.target.value});
    }

    cancel(){
        this.props.history.push('/documents');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Document</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> fileName: </label>
                                                <input placeholder="fileName" name="fileName" className="form-control" value={this.state.fileName} onChange={this.changefileNameHandler}/>

                                            <label> uploadedDate: </label>
                                                <input type="date" placeholder="uploadedDate" name="uploadedDate" className="form-control" value={this.state.uploadedDate} onChange={this.changeuploadedDateHandler}/>

                                            <label> DocumentType: </label>
                                                <select value={this.state.documentType} onChange={this.changeDocumentTypeHandler}>
                      <option name="DocumentType" className="form-control" >
                          ApplicationForm
                      </option>
                      <option name="DocumentType" className="form-control" >
                          PolicyDocument
                      </option>
                      <option name="DocumentType" className="form-control" >
                          Endorsement
                      </option>
                      <option name="DocumentType" className="form-control" >
                          Invoice
                      </option>
                      <option name="DocumentType" className="form-control" >
                          ClaimForm
                      </option>
                      <option name="DocumentType" className="form-control" >
                          PoliceReport
                      </option>
                      <option name="DocumentType" className="form-control" >
                          Estimate
                      </option>
                      <option name="DocumentType" className="form-control" >
                          Photo
                      </option>
                      <option name="DocumentType" className="form-control" >
                          MedicalRecord
                      </option>
                      <option name="DocumentType" className="form-control" >
                          Correspondence
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateDocument}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateDocumentComponent
