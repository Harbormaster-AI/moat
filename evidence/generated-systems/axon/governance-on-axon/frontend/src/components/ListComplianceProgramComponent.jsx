import React, { Component } from 'react'
import ComplianceProgramService from '../services/ComplianceProgramService'

class ListComplianceProgramComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                compliancePrograms: []
        }
        this.addComplianceProgram = this.addComplianceProgram.bind(this);
        this.editComplianceProgram = this.editComplianceProgram.bind(this);
        this.deleteComplianceProgram = this.deleteComplianceProgram.bind(this);
    }

    deleteComplianceProgram(id){
        ComplianceProgramService.deleteComplianceProgram(id).then( res => {
            this.setState({compliancePrograms: this.state.compliancePrograms.filter(complianceProgram => complianceProgram.complianceProgramId !== id)});
        });
    }
    viewComplianceProgram(id){
        this.props.history.push(`/view-complianceProgram/${id}`);
    }
    editComplianceProgram(id){
        this.props.history.push(`/add-complianceProgram/${id}`);
    }

    componentDidMount(){
        ComplianceProgramService.getCompliancePrograms().then((res) => {
            this.setState({ compliancePrograms: res.data});
        });
    }

    addComplianceProgram(){
        this.props.history.push('/add-complianceProgram/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ComplianceProgram List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addComplianceProgram}> Add ComplianceProgram</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Framework </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.compliancePrograms.map(
                                        complianceProgram => 
                                        <tr key = {complianceProgram.complianceProgramId}>
                                             <td> { complianceProgram.name } </td>
                                             <td> { complianceProgram.framework } </td>
                                             <td> { complianceProgram.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editComplianceProgram(complianceProgram.complianceProgramId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteComplianceProgram(complianceProgram.complianceProgramId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewComplianceProgram(complianceProgram.complianceProgramId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListComplianceProgramComponent
