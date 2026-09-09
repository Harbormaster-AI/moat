import React, { Component } from 'react'
import AuditFindingService from '../services/AuditFindingService'

class ListAuditFindingComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                auditFindings: []
        }
        this.addAuditFinding = this.addAuditFinding.bind(this);
        this.editAuditFinding = this.editAuditFinding.bind(this);
        this.deleteAuditFinding = this.deleteAuditFinding.bind(this);
    }

    deleteAuditFinding(id){
        AuditFindingService.deleteAuditFinding(id).then( res => {
            this.setState({auditFindings: this.state.auditFindings.filter(auditFinding => auditFinding.auditFindingId !== id)});
        });
    }
    viewAuditFinding(id){
        this.props.history.push(`/view-auditFinding/${id}`);
    }
    editAuditFinding(id){
        this.props.history.push(`/add-auditFinding/${id}`);
    }

    componentDidMount(){
        AuditFindingService.getAuditFindings().then((res) => {
            this.setState({ auditFindings: res.data});
        });
    }

    addAuditFinding(){
        this.props.history.push('/add-auditFinding/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">AuditFinding List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAuditFinding}> Add AuditFinding</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Title </th>
                                    <th> Description </th>
                                    <th> DueDate </th>
                                    <th> Severity </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.auditFindings.map(
                                        auditFinding => 
                                        <tr key = {auditFinding.auditFindingId}>
                                             <td> { auditFinding.title } </td>
                                             <td> { auditFinding.description } </td>
                                             <td> { auditFinding.dueDate } </td>
                                             <td> { auditFinding.severity } </td>
                                             <td> { auditFinding.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editAuditFinding(auditFinding.auditFindingId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAuditFinding(auditFinding.auditFindingId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAuditFinding(auditFinding.auditFindingId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAuditFindingComponent
