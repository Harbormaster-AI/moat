import React, { Component } from 'react'
import ComplianceRequirementService from '../services/ComplianceRequirementService'

class ListComplianceRequirementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                complianceRequirements: []
        }
        this.addComplianceRequirement = this.addComplianceRequirement.bind(this);
        this.editComplianceRequirement = this.editComplianceRequirement.bind(this);
        this.deleteComplianceRequirement = this.deleteComplianceRequirement.bind(this);
    }

    deleteComplianceRequirement(id){
        ComplianceRequirementService.deleteComplianceRequirement(id).then( res => {
            this.setState({complianceRequirements: this.state.complianceRequirements.filter(complianceRequirement => complianceRequirement.complianceRequirementId !== id)});
        });
    }
    viewComplianceRequirement(id){
        this.props.history.push(`/view-complianceRequirement/${id}`);
    }
    editComplianceRequirement(id){
        this.props.history.push(`/add-complianceRequirement/${id}`);
    }

    componentDidMount(){
        ComplianceRequirementService.getComplianceRequirements().then((res) => {
            this.setState({ complianceRequirements: res.data});
        });
    }

    addComplianceRequirement(){
        this.props.history.push('/add-complianceRequirement/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ComplianceRequirement List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addComplianceRequirement}> Add ComplianceRequirement</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Source </th>
                                    <th> Citation </th>
                                    <th> Applicability </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.complianceRequirements.map(
                                        complianceRequirement => 
                                        <tr key = {complianceRequirement.complianceRequirementId}>
                                             <td> { complianceRequirement.name } </td>
                                             <td> { complianceRequirement.source } </td>
                                             <td> { complianceRequirement.citation } </td>
                                             <td> { complianceRequirement.applicability } </td>
                                             <td> { complianceRequirement.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editComplianceRequirement(complianceRequirement.complianceRequirementId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteComplianceRequirement(complianceRequirement.complianceRequirementId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewComplianceRequirement(complianceRequirement.complianceRequirementId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListComplianceRequirementComponent
