import React, { Component } from 'react'
import EnterpriseService from '../services/EnterpriseService'

class ListEnterpriseComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                enterprises: []
        }
        this.addEnterprise = this.addEnterprise.bind(this);
        this.editEnterprise = this.editEnterprise.bind(this);
        this.deleteEnterprise = this.deleteEnterprise.bind(this);
    }

    deleteEnterprise(id){
        EnterpriseService.deleteEnterprise(id).then( res => {
            this.setState({enterprises: this.state.enterprises.filter(enterprise => enterprise.enterpriseId !== id)});
        });
    }
    viewEnterprise(id){
        this.props.history.push(`/view-enterprise/${id}`);
    }
    editEnterprise(id){
        this.props.history.push(`/add-enterprise/${id}`);
    }

    componentDidMount(){
        EnterpriseService.getEnterprises().then((res) => {
            this.setState({ enterprises: res.data});
        });
    }

    addEnterprise(){
        this.props.history.push('/add-enterprise/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Enterprise List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addEnterprise}> Add Enterprise</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> LegalName </th>
                                    <th> RegistrationCountry </th>
                                    <th> Website </th>
                                    <th> TaxId </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.enterprises.map(
                                        enterprise => 
                                        <tr key = {enterprise.enterpriseId}>
                                             <td> { enterprise.name } </td>
                                             <td> { enterprise.legalName } </td>
                                             <td> { enterprise.registrationCountry } </td>
                                             <td> { enterprise.website } </td>
                                             <td> { enterprise.taxId } </td>
                                             <td>
                                                 <button onClick={ () => this.editEnterprise(enterprise.enterpriseId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteEnterprise(enterprise.enterpriseId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewEnterprise(enterprise.enterpriseId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListEnterpriseComponent
