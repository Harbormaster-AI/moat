import React, { Component } from 'react'
import AttestationService from '../services/AttestationService'

class ListAttestationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                attestations: []
        }
        this.addAttestation = this.addAttestation.bind(this);
        this.editAttestation = this.editAttestation.bind(this);
        this.deleteAttestation = this.deleteAttestation.bind(this);
    }

    deleteAttestation(id){
        AttestationService.deleteAttestation(id).then( res => {
            this.setState({attestations: this.state.attestations.filter(attestation => attestation.attestationId !== id)});
        });
    }
    viewAttestation(id){
        this.props.history.push(`/view-attestation/${id}`);
    }
    editAttestation(id){
        this.props.history.push(`/add-attestation/${id}`);
    }

    componentDidMount(){
        AttestationService.getAttestations().then((res) => {
            this.setState({ attestations: res.data});
        });
    }

    addAttestation(){
        this.props.history.push('/add-attestation/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Attestation List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAttestation}> Add Attestation</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Statement </th>
                                    <th> Attestor </th>
                                    <th> DateSigned </th>
                                    <th> Result </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.attestations.map(
                                        attestation => 
                                        <tr key = {attestation.attestationId}>
                                             <td> { attestation.statement } </td>
                                             <td> { attestation.attestor } </td>
                                             <td> { attestation.dateSigned } </td>
                                             <td> { attestation.result } </td>
                                             <td>
                                                 <button onClick={ () => this.editAttestation(attestation.attestationId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAttestation(attestation.attestationId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAttestation(attestation.attestationId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAttestationComponent
