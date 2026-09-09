import React, { Component } from 'react'
import CandidateService from '../services/CandidateService'

class ListCandidateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                candidates: []
        }
        this.addCandidate = this.addCandidate.bind(this);
        this.editCandidate = this.editCandidate.bind(this);
        this.deleteCandidate = this.deleteCandidate.bind(this);
    }

    deleteCandidate(id){
        CandidateService.deleteCandidate(id).then( res => {
            this.setState({candidates: this.state.candidates.filter(candidate => candidate.candidateId !== id)});
        });
    }
    viewCandidate(id){
        this.props.history.push(`/view-candidate/${id}`);
    }
    editCandidate(id){
        this.props.history.push(`/add-candidate/${id}`);
    }

    componentDidMount(){
        CandidateService.getCandidates().then((res) => {
            this.setState({ candidates: res.data});
        });
    }

    addCandidate(){
        this.props.history.push('/add-candidate/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Candidate List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCandidate}> Add Candidate</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Email </th>
                                    <th> Phone </th>
                                    <th> Source </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.candidates.map(
                                        candidate => 
                                        <tr key = {candidate.candidateId}>
                                             <td> { candidate.name } </td>
                                             <td> { candidate.email } </td>
                                             <td> { candidate.phone } </td>
                                             <td> { candidate.source } </td>
                                             <td>
                                                 <button onClick={ () => this.editCandidate(candidate.candidateId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCandidate(candidate.candidateId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCandidate(candidate.candidateId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCandidateComponent
