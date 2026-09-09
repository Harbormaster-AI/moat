import React, { Component } from 'react'
import BeneficiaryService from '../services/BeneficiaryService'

class ListBeneficiaryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                beneficiarys: []
        }
        this.addBeneficiary = this.addBeneficiary.bind(this);
        this.editBeneficiary = this.editBeneficiary.bind(this);
        this.deleteBeneficiary = this.deleteBeneficiary.bind(this);
    }

    deleteBeneficiary(id){
        BeneficiaryService.deleteBeneficiary(id).then( res => {
            this.setState({beneficiarys: this.state.beneficiarys.filter(beneficiary => beneficiary.beneficiaryId !== id)});
        });
    }
    viewBeneficiary(id){
        this.props.history.push(`/view-beneficiary/${id}`);
    }
    editBeneficiary(id){
        this.props.history.push(`/add-beneficiary/${id}`);
    }

    componentDidMount(){
        BeneficiaryService.getBeneficiarys().then((res) => {
            this.setState({ beneficiarys: res.data});
        });
    }

    addBeneficiary(){
        this.props.history.push('/add-beneficiary/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Beneficiary List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addBeneficiary}> Add Beneficiary</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> AccountIdentifier </th>
                                    <th> Iban </th>
                                    <th> Bic </th>
                                    <th> Address </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.beneficiarys.map(
                                        beneficiary => 
                                        <tr key = {beneficiary.beneficiaryId}>
                                             <td> { beneficiary.name } </td>
                                             <td> { beneficiary.accountIdentifier } </td>
                                             <td> { beneficiary.iban } </td>
                                             <td> { beneficiary.bic } </td>
                                             <td> { beneficiary.address } </td>
                                             <td>
                                                 <button onClick={ () => this.editBeneficiary(beneficiary.beneficiaryId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteBeneficiary(beneficiary.beneficiaryId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewBeneficiary(beneficiary.beneficiaryId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListBeneficiaryComponent
