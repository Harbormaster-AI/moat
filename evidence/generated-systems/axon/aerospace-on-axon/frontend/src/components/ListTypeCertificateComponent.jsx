import React, { Component } from 'react'
import TypeCertificateService from '../services/TypeCertificateService'

class ListTypeCertificateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                typeCertificates: []
        }
        this.addTypeCertificate = this.addTypeCertificate.bind(this);
        this.editTypeCertificate = this.editTypeCertificate.bind(this);
        this.deleteTypeCertificate = this.deleteTypeCertificate.bind(this);
    }

    deleteTypeCertificate(id){
        TypeCertificateService.deleteTypeCertificate(id).then( res => {
            this.setState({typeCertificates: this.state.typeCertificates.filter(typeCertificate => typeCertificate.typeCertificateId !== id)});
        });
    }
    viewTypeCertificate(id){
        this.props.history.push(`/view-typeCertificate/${id}`);
    }
    editTypeCertificate(id){
        this.props.history.push(`/add-typeCertificate/${id}`);
    }

    componentDidMount(){
        TypeCertificateService.getTypeCertificates().then((res) => {
            this.setState({ typeCertificates: res.data});
        });
    }

    addTypeCertificate(){
        this.props.history.push('/add-typeCertificate/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">TypeCertificate List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addTypeCertificate}> Add TypeCertificate</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> CertificateNumber </th>
                                    <th> Authority </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.typeCertificates.map(
                                        typeCertificate => 
                                        <tr key = {typeCertificate.typeCertificateId}>
                                             <td> { typeCertificate.certificateNumber } </td>
                                             <td> { typeCertificate.authority } </td>
                                             <td>
                                                 <button onClick={ () => this.editTypeCertificate(typeCertificate.typeCertificateId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteTypeCertificate(typeCertificate.typeCertificateId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewTypeCertificate(typeCertificate.typeCertificateId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListTypeCertificateComponent
