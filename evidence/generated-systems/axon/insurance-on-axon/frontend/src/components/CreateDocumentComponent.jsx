import React, { Component } from 'react'
import DocumentService from '../services/DocumentService';

class CreateDocumentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                fileName: '',
                uploadedDate: '',
                documentType: ''
        }
        this.changefileNameHandler = this.changefileNameHandler.bind(this);
        this.changeuploadedDateHandler = this.changeuploadedDateHandler.bind(this);
        this.changeDocumentTypeHandler = this.changeDocumentTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            DocumentService.getDocumentById(this.state.id).then( (res) =>{
                let document = res.data;
                this.setState({
                    fileName: document.fileName,
                    uploadedDate: document.uploadedDate,
                    documentType: document.documentType
                });
            });
        }        
    }
    saveOrUpdateDocument = (e) => {
        e.preventDefault();
        let document = {
                documentId: this.state.id,
                fileName: this.state.fileName,
                uploadedDate: this.state.uploadedDate,
                documentType: this.state.documentType
            };
        console.log('document => ' + JSON.stringify(document));

        // step 5
        if(this.state.id === '_add'){
            document.documentId=''
            DocumentService.createDocument(document).then(res =>{
                this.props.history.push('/documents');
            });
        }else{
            DocumentService.updateDocument(document).then( res => {
                this.props.history.push('/documents');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Document</h3>
        }else{
            return <h3 className="text-center">Update Document</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> fileName:&emsp; </label>
                                                <input placeholder="fileName" name="fileName" className="form-control" value={this.state.fileName} onChange={this.changefileNameHandler}/>

                                            <label> uploadedDate:&emsp; </label>
                                                <input type="date" placeholder="uploadedDate" name="uploadedDate" className="form-control" value={this.state.uploadedDate} onChange={this.changeuploadedDateHandler}/>

                                            <label> DocumentType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateDocument}>Save</button>
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

export default CreateDocumentComponent
