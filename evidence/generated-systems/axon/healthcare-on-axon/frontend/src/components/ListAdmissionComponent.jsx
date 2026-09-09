import React, { Component } from 'react'
import AdmissionService from '../services/AdmissionService'

class ListAdmissionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                admissions: []
        }
        this.addAdmission = this.addAdmission.bind(this);
        this.editAdmission = this.editAdmission.bind(this);
        this.deleteAdmission = this.deleteAdmission.bind(this);
    }

    deleteAdmission(id){
        AdmissionService.deleteAdmission(id).then( res => {
            this.setState({admissions: this.state.admissions.filter(admission => admission.admissionId !== id)});
        });
    }
    viewAdmission(id){
        this.props.history.push(`/view-admission/${id}`);
    }
    editAdmission(id){
        this.props.history.push(`/add-admission/${id}`);
    }

    componentDidMount(){
        AdmissionService.getAdmissions().then((res) => {
            this.setState({ admissions: res.data});
        });
    }

    addAdmission(){
        this.props.history.push('/add-admission/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Admission List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAdmission}> Add Admission</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> AdmitDateTime </th>
                                    <th> Bed </th>
                                    <th> AdmissionType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.admissions.map(
                                        admission => 
                                        <tr key = {admission.admissionId}>
                                             <td> { admission.admitDateTime } </td>
                                             <td> { admission.bed } </td>
                                             <td> { admission.admissionType } </td>
                                             <td>
                                                 <button onClick={ () => this.editAdmission(admission.admissionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAdmission(admission.admissionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAdmission(admission.admissionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAdmissionComponent
