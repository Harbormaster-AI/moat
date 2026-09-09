import React, { Component } from 'react'
import DiagnosisService from '../services/DiagnosisService'

class ListDiagnosisComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                diagnosiss: []
        }
        this.addDiagnosis = this.addDiagnosis.bind(this);
        this.editDiagnosis = this.editDiagnosis.bind(this);
        this.deleteDiagnosis = this.deleteDiagnosis.bind(this);
    }

    deleteDiagnosis(id){
        DiagnosisService.deleteDiagnosis(id).then( res => {
            this.setState({diagnosiss: this.state.diagnosiss.filter(diagnosis => diagnosis.diagnosisId !== id)});
        });
    }
    viewDiagnosis(id){
        this.props.history.push(`/view-diagnosis/${id}`);
    }
    editDiagnosis(id){
        this.props.history.push(`/add-diagnosis/${id}`);
    }

    componentDidMount(){
        DiagnosisService.getDiagnosiss().then((res) => {
            this.setState({ diagnosiss: res.data});
        });
    }

    addDiagnosis(){
        this.props.history.push('/add-diagnosis/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Diagnosis List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDiagnosis}> Add Diagnosis</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Code </th>
                                    <th> Description </th>
                                    <th> OnsetDate </th>
                                    <th> Certainty </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.diagnosiss.map(
                                        diagnosis => 
                                        <tr key = {diagnosis.diagnosisId}>
                                             <td> { diagnosis.code } </td>
                                             <td> { diagnosis.description } </td>
                                             <td> { diagnosis.onsetDate } </td>
                                             <td> { diagnosis.certainty } </td>
                                             <td>
                                                 <button onClick={ () => this.editDiagnosis(diagnosis.diagnosisId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDiagnosis(diagnosis.diagnosisId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDiagnosis(diagnosis.diagnosisId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListDiagnosisComponent
