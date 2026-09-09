import React, { Component } from 'react'
import DocumentService from '../services/DocumentService'

class ListDocumentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                documents: []
        }
        this.addDocument = this.addDocument.bind(this);
        this.editDocument = this.editDocument.bind(this);
        this.deleteDocument = this.deleteDocument.bind(this);
    }

    deleteDocument(id){
        DocumentService.deleteDocument(id).then( res => {
            this.setState({documents: this.state.documents.filter(document => document.documentId !== id)});
        });
    }
    viewDocument(id){
        this.props.history.push(`/view-document/${id}`);
    }
    editDocument(id){
        this.props.history.push(`/add-document/${id}`);
    }

    componentDidMount(){
        DocumentService.getDocuments().then((res) => {
            this.setState({ documents: res.data});
        });
    }

    addDocument(){
        this.props.history.push('/add-document/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Document List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDocument}> Add Document</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> FileUrl </th>
                                    <th> UploadedDate </th>
                                    <th> DocumentType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.documents.map(
                                        document => 
                                        <tr key = {document.documentId}>
                                             <td> { document.name } </td>
                                             <td> { document.fileUrl } </td>
                                             <td> { document.uploadedDate } </td>
                                             <td> { document.documentType } </td>
                                             <td>
                                                 <button onClick={ () => this.editDocument(document.documentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDocument(document.documentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDocument(document.documentId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListDocumentComponent
