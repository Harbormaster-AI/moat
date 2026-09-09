import React, { Component } from 'react'
import FinancialInstitutionService from '../services/FinancialInstitutionService'

class ListFinancialInstitutionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                financialInstitutions: []
        }
        this.addFinancialInstitution = this.addFinancialInstitution.bind(this);
        this.editFinancialInstitution = this.editFinancialInstitution.bind(this);
        this.deleteFinancialInstitution = this.deleteFinancialInstitution.bind(this);
    }

    deleteFinancialInstitution(id){
        FinancialInstitutionService.deleteFinancialInstitution(id).then( res => {
            this.setState({financialInstitutions: this.state.financialInstitutions.filter(financialInstitution => financialInstitution.financialInstitutionId !== id)});
        });
    }
    viewFinancialInstitution(id){
        this.props.history.push(`/view-financialInstitution/${id}`);
    }
    editFinancialInstitution(id){
        this.props.history.push(`/add-financialInstitution/${id}`);
    }

    componentDidMount(){
        FinancialInstitutionService.getFinancialInstitutions().then((res) => {
            this.setState({ financialInstitutions: res.data});
        });
    }

    addFinancialInstitution(){
        this.props.history.push('/add-financialInstitution/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">FinancialInstitution List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addFinancialInstitution}> Add FinancialInstitution</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> LegalName </th>
                                    <th> CountryOfIncorporation </th>
                                    <th> Bic </th>
                                    <th> Website </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.financialInstitutions.map(
                                        financialInstitution => 
                                        <tr key = {financialInstitution.financialInstitutionId}>
                                             <td> { financialInstitution.name } </td>
                                             <td> { financialInstitution.legalName } </td>
                                             <td> { financialInstitution.countryOfIncorporation } </td>
                                             <td> { financialInstitution.bic } </td>
                                             <td> { financialInstitution.website } </td>
                                             <td>
                                                 <button onClick={ () => this.editFinancialInstitution(financialInstitution.financialInstitutionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteFinancialInstitution(financialInstitution.financialInstitutionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewFinancialInstitution(financialInstitution.financialInstitutionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListFinancialInstitutionComponent
