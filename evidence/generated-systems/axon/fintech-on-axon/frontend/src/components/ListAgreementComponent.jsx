import React, { Component } from 'react'
import AgreementService from '../services/AgreementService'

class ListAgreementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                agreements: []
        }
        this.addAgreement = this.addAgreement.bind(this);
        this.editAgreement = this.editAgreement.bind(this);
        this.deleteAgreement = this.deleteAgreement.bind(this);
    }

    deleteAgreement(id){
        AgreementService.deleteAgreement(id).then( res => {
            this.setState({agreements: this.state.agreements.filter(agreement => agreement.agreementId !== id)});
        });
    }
    viewAgreement(id){
        this.props.history.push(`/view-agreement/${id}`);
    }
    editAgreement(id){
        this.props.history.push(`/add-agreement/${id}`);
    }

    componentDidMount(){
        AgreementService.getAgreements().then((res) => {
            this.setState({ agreements: res.data});
        });
    }

    addAgreement(){
        this.props.history.push('/add-agreement/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Agreement List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAgreement}> Add Agreement</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> AgreementNumber </th>
                                    <th> EffectiveDate </th>
                                    <th> AgreementType </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.agreements.map(
                                        agreement => 
                                        <tr key = {agreement.agreementId}>
                                             <td> { agreement.agreementNumber } </td>
                                             <td> { agreement.effectiveDate } </td>
                                             <td> { agreement.agreementType } </td>
                                             <td> { agreement.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editAgreement(agreement.agreementId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAgreement(agreement.agreementId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAgreement(agreement.agreementId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAgreementComponent
