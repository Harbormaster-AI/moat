import React, { Component } from 'react'
import OrganizationService from '../services/OrganizationService'

class ListOrganizationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                organizations: []
        }
        this.addOrganization = this.addOrganization.bind(this);
        this.editOrganization = this.editOrganization.bind(this);
        this.deleteOrganization = this.deleteOrganization.bind(this);
    }

    deleteOrganization(id){
        OrganizationService.deleteOrganization(id).then( res => {
            this.setState({organizations: this.state.organizations.filter(organization => organization.organizationId !== id)});
        });
    }
    viewOrganization(id){
        this.props.history.push(`/view-organization/${id}`);
    }
    editOrganization(id){
        this.props.history.push(`/add-organization/${id}`);
    }

    componentDidMount(){
        OrganizationService.getOrganizations().then((res) => {
            this.setState({ organizations: res.data});
        });
    }

    addOrganization(){
        this.props.history.push('/add-organization/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Organization List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addOrganization}> Add Organization</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> LegalName </th>
                                    <th> Jurisdiction </th>
                                    <th> IndustrySector </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.organizations.map(
                                        organization => 
                                        <tr key = {organization.organizationId}>
                                             <td> { organization.name } </td>
                                             <td> { organization.legalName } </td>
                                             <td> { organization.jurisdiction } </td>
                                             <td> { organization.industrySector } </td>
                                             <td>
                                                 <button onClick={ () => this.editOrganization(organization.organizationId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteOrganization(organization.organizationId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewOrganization(organization.organizationId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListOrganizationComponent
