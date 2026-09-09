import React, { Component } from 'react'
import AuditProgramService from '../services/AuditProgramService'

class ListAuditProgramComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                auditPrograms: []
        }
        this.addAuditProgram = this.addAuditProgram.bind(this);
        this.editAuditProgram = this.editAuditProgram.bind(this);
        this.deleteAuditProgram = this.deleteAuditProgram.bind(this);
    }

    deleteAuditProgram(id){
        AuditProgramService.deleteAuditProgram(id).then( res => {
            this.setState({auditPrograms: this.state.auditPrograms.filter(auditProgram => auditProgram.auditProgramId !== id)});
        });
    }
    viewAuditProgram(id){
        this.props.history.push(`/view-auditProgram/${id}`);
    }
    editAuditProgram(id){
        this.props.history.push(`/add-auditProgram/${id}`);
    }

    componentDidMount(){
        AuditProgramService.getAuditPrograms().then((res) => {
            this.setState({ auditPrograms: res.data});
        });
    }

    addAuditProgram(){
        this.props.history.push('/add-auditProgram/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">AuditProgram List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAuditProgram}> Add AuditProgram</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Scope </th>
                                    <th> Cycle </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.auditPrograms.map(
                                        auditProgram => 
                                        <tr key = {auditProgram.auditProgramId}>
                                             <td> { auditProgram.name } </td>
                                             <td> { auditProgram.scope } </td>
                                             <td> { auditProgram.cycle } </td>
                                             <td> { auditProgram.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editAuditProgram(auditProgram.auditProgramId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAuditProgram(auditProgram.auditProgramId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAuditProgram(auditProgram.auditProgramId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAuditProgramComponent
