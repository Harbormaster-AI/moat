import React, { Component } from 'react'
import KYCDocumentService from '../services/KYCDocumentService'

class ListKYCDocumentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                kYCDocuments: []
        }
        this.addKYCDocument = this.addKYCDocument.bind(this);
        this.editKYCDocument = this.editKYCDocument.bind(this);
        this.deleteKYCDocument = this.deleteKYCDocument.bind(this);
    }

    deleteKYCDocument(id){
        KYCDocumentService.deleteKYCDocument(id).then( res => {
            this.setState({kYCDocuments: this.state.kYCDocuments.filter(kYCDocument => kYCDocument.kYCDocumentId !== id)});
        });
    }
    viewKYCDocument(id){
        this.props.history.push(`/view-kYCDocument/${id}`);
    }
    editKYCDocument(id){
        this.props.history.push(`/add-kYCDocument/${id}`);
    }

    componentDidMount(){
        KYCDocumentService.getKYCDocuments().then((res) => {
            this.setState({ kYCDocuments: res.data});
        });
    }

    addKYCDocument(){
        this.props.history.push('/add-kYCDocument/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">KYCDocument List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addKYCDocument}> Add KYCDocument</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Reference </th>
                                    <th> IssuedCountry </th>
                                    <th> ExpirationDate </th>
                                    <th> DocumentType </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.kYCDocuments.map(
                                        kYCDocument => 
                                        <tr key = {kYCDocument.kYCDocumentId}>
                                             <td> { kYCDocument.reference } </td>
                                             <td> { kYCDocument.issuedCountry } </td>
                                             <td> { kYCDocument.expirationDate } </td>
                                             <td> { kYCDocument.documentType } </td>
                                             <td> { kYCDocument.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editKYCDocument(kYCDocument.kYCDocumentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteKYCDocument(kYCDocument.kYCDocumentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewKYCDocument(kYCDocument.kYCDocumentId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListKYCDocumentComponent
