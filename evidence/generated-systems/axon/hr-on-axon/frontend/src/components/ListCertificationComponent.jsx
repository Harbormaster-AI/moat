import React, { Component } from 'react'
import CertificationService from '../services/CertificationService'

class ListCertificationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                certifications: []
        }
        this.addCertification = this.addCertification.bind(this);
        this.editCertification = this.editCertification.bind(this);
        this.deleteCertification = this.deleteCertification.bind(this);
    }

    deleteCertification(id){
        CertificationService.deleteCertification(id).then( res => {
            this.setState({certifications: this.state.certifications.filter(certification => certification.certificationId !== id)});
        });
    }
    viewCertification(id){
        this.props.history.push(`/view-certification/${id}`);
    }
    editCertification(id){
        this.props.history.push(`/add-certification/${id}`);
    }

    componentDidMount(){
        CertificationService.getCertifications().then((res) => {
            this.setState({ certifications: res.data});
        });
    }

    addCertification(){
        this.props.history.push('/add-certification/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Certification List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCertification}> Add Certification</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Issuer </th>
                                    <th> ValidFrom </th>
                                    <th> ValidTo </th>
                                    <th> CredentialId </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.certifications.map(
                                        certification => 
                                        <tr key = {certification.certificationId}>
                                             <td> { certification.name } </td>
                                             <td> { certification.issuer } </td>
                                             <td> { certification.validFrom } </td>
                                             <td> { certification.validTo } </td>
                                             <td> { certification.credentialId } </td>
                                             <td>
                                                 <button onClick={ () => this.editCertification(certification.certificationId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCertification(certification.certificationId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCertification(certification.certificationId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCertificationComponent
