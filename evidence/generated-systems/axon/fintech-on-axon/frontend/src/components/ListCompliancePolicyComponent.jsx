import React, { Component } from 'react'
import CompliancePolicyService from '../services/CompliancePolicyService'

class ListCompliancePolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                compliancePolicys: []
        }
        this.addCompliancePolicy = this.addCompliancePolicy.bind(this);
        this.editCompliancePolicy = this.editCompliancePolicy.bind(this);
        this.deleteCompliancePolicy = this.deleteCompliancePolicy.bind(this);
    }

    deleteCompliancePolicy(id){
        CompliancePolicyService.deleteCompliancePolicy(id).then( res => {
            this.setState({compliancePolicys: this.state.compliancePolicys.filter(compliancePolicy => compliancePolicy.compliancePolicyId !== id)});
        });
    }
    viewCompliancePolicy(id){
        this.props.history.push(`/view-compliancePolicy/${id}`);
    }
    editCompliancePolicy(id){
        this.props.history.push(`/add-compliancePolicy/${id}`);
    }

    componentDidMount(){
        CompliancePolicyService.getCompliancePolicys().then((res) => {
            this.setState({ compliancePolicys: res.data});
        });
    }

    addCompliancePolicy(){
        this.props.history.push('/add-compliancePolicy/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CompliancePolicy List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCompliancePolicy}> Add CompliancePolicy</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> PolicyCode </th>
                                    <th> Description </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.compliancePolicys.map(
                                        compliancePolicy => 
                                        <tr key = {compliancePolicy.compliancePolicyId}>
                                             <td> { compliancePolicy.name } </td>
                                             <td> { compliancePolicy.policyCode } </td>
                                             <td> { compliancePolicy.description } </td>
                                             <td> { compliancePolicy.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editCompliancePolicy(compliancePolicy.compliancePolicyId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCompliancePolicy(compliancePolicy.compliancePolicyId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCompliancePolicy(compliancePolicy.compliancePolicyId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCompliancePolicyComponent
