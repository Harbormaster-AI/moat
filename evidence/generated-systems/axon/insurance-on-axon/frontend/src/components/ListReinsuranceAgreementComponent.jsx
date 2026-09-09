import React, { Component } from 'react'
import ReinsuranceAgreementService from '../services/ReinsuranceAgreementService'

class ListReinsuranceAgreementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                reinsuranceAgreements: []
        }
        this.addReinsuranceAgreement = this.addReinsuranceAgreement.bind(this);
        this.editReinsuranceAgreement = this.editReinsuranceAgreement.bind(this);
        this.deleteReinsuranceAgreement = this.deleteReinsuranceAgreement.bind(this);
    }

    deleteReinsuranceAgreement(id){
        ReinsuranceAgreementService.deleteReinsuranceAgreement(id).then( res => {
            this.setState({reinsuranceAgreements: this.state.reinsuranceAgreements.filter(reinsuranceAgreement => reinsuranceAgreement.reinsuranceAgreementId !== id)});
        });
    }
    viewReinsuranceAgreement(id){
        this.props.history.push(`/view-reinsuranceAgreement/${id}`);
    }
    editReinsuranceAgreement(id){
        this.props.history.push(`/add-reinsuranceAgreement/${id}`);
    }

    componentDidMount(){
        ReinsuranceAgreementService.getReinsuranceAgreements().then((res) => {
            this.setState({ reinsuranceAgreements: res.data});
        });
    }

    addReinsuranceAgreement(){
        this.props.history.push('/add-reinsuranceAgreement/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ReinsuranceAgreement List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addReinsuranceAgreement}> Add ReinsuranceAgreement</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> AgreementNumber </th>
                                    <th> EffectivePeriod </th>
                                    <th> Retention </th>
                                    <th> Limit </th>
                                    <th> CessionPercentage </th>
                                    <th> ReinsuranceType </th>
                                    <th> TreatyType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.reinsuranceAgreements.map(
                                        reinsuranceAgreement => 
                                        <tr key = {reinsuranceAgreement.reinsuranceAgreementId}>
                                             <td> { reinsuranceAgreement.agreementNumber } </td>
                                             <td> { reinsuranceAgreement.effectivePeriod } </td>
                                             <td> { reinsuranceAgreement.retention } </td>
                                             <td> { reinsuranceAgreement.limit } </td>
                                             <td> { reinsuranceAgreement.cessionPercentage } </td>
                                             <td> { reinsuranceAgreement.reinsuranceType } </td>
                                             <td> { reinsuranceAgreement.treatyType } </td>
                                             <td>
                                                 <button onClick={ () => this.editReinsuranceAgreement(reinsuranceAgreement.reinsuranceAgreementId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteReinsuranceAgreement(reinsuranceAgreement.reinsuranceAgreementId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewReinsuranceAgreement(reinsuranceAgreement.reinsuranceAgreementId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListReinsuranceAgreementComponent
