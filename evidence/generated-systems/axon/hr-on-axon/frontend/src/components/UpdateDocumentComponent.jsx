import React, { Component } from 'react'
import DocumentService from '../services/DocumentService';

class UpdateDocumentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                fileUrl: '',
                uploadedDate: '',
                documentType: ''
        }
        this.updateDocument = this.updateDocument.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changefileUrlHandler = this.changefileUrlHandler.bind(this);
        this.changeuploadedDateHandler = this.changeuploadedDateHandler.bind(this);
        this.changeDocumentTypeHandler = this.changeDocumentTypeHandler.bind(this);
    }

    componentDidMount(){
        DocumentService.getDocumentById(this.state.id).then( (res) =>{
            let document = res.data;
            this.setState({
                name: document.name,
                fileUrl: document.fileUrl,
                uploadedDate: document.uploadedDate,
                documentType: document.documentType
            });
        });
    }

    updateDocument = (e) => {
        e.preventDefault();
        let document = {
            documentId: this.state.id,
            name: this.state.name,
            fileUrl: this.state.fileUrl,
            uploadedDate: this.state.uploadedDate,
            documentType: this.state.documentType
        };
        console.log('document => ' + JSON.stringify(document));
        console.log('id => ' + JSON.stringify(this.state.id));
        DocumentService.updateDocument(document).then( res => {
            this.props.history.push('/documents');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changefileUrlHandler= (event) => {
        this.setState({fileUrl: event.target.value});
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
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> fileUrl: </label>
                                                <input placeholder="fileUrl" name="fileUrl" className="form-control" value={this.state.fileUrl} onChange={this.changefileUrlHandler}/>

                                            <label> uploadedDate: </label>
                                                <input type="date" placeholder="uploadedDate" name="uploadedDate" className="form-control" value={this.state.uploadedDate} onChange={this.changeuploadedDateHandler}/>

                                            <label> DocumentType: </label>
                                                <select value={this.state.documentType} onChange={this.changeDocumentTypeHandler}>
                      <option name="DocumentType" className="form-control" >
                          Resume
                      </option>
                      <option name="DocumentType" className="form-control" >
                          CoverLetter
                      </option>
                      <option name="DocumentType" className="form-control" >
                          ID
                      </option>
                      <option name="DocumentType" className="form-control" >
                          Certification
                      </option>
                      <option name="DocumentType" className="form-control" >
                          Contract
                      </option>
                      <option name="DocumentType" className="form-control" >
                          Policy
                      </option>
                      <option name="DocumentType" className="form-control" >
                          Other
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
