import React, { Component } from 'react'
import EvidenceService from '../services/EvidenceService'

class ListEvidenceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                evidences: []
        }
        this.addEvidence = this.addEvidence.bind(this);
        this.editEvidence = this.editEvidence.bind(this);
        this.deleteEvidence = this.deleteEvidence.bind(this);
    }

    deleteEvidence(id){
        EvidenceService.deleteEvidence(id).then( res => {
            this.setState({evidences: this.state.evidences.filter(evidence => evidence.evidenceId !== id)});
        });
    }
    viewEvidence(id){
        this.props.history.push(`/view-evidence/${id}`);
    }
    editEvidence(id){
        this.props.history.push(`/add-evidence/${id}`);
    }

    componentDidMount(){
        EvidenceService.getEvidences().then((res) => {
            this.setState({ evidences: res.data});
        });
    }

    addEvidence(){
        this.props.history.push('/add-evidence/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Evidence List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addEvidence}> Add Evidence</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Title </th>
                                    <th> LocationUrl </th>
                                    <th> ReceivedDate </th>
                                    <th> EvidenceType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.evidences.map(
                                        evidence => 
                                        <tr key = {evidence.evidenceId}>
                                             <td> { evidence.title } </td>
                                             <td> { evidence.locationUrl } </td>
                                             <td> { evidence.receivedDate } </td>
                                             <td> { evidence.evidenceType } </td>
                                             <td>
                                                 <button onClick={ () => this.editEvidence(evidence.evidenceId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteEvidence(evidence.evidenceId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewEvidence(evidence.evidenceId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListEvidenceComponent
