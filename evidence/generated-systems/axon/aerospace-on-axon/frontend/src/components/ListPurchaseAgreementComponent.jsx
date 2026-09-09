import React, { Component } from 'react'
import PurchaseAgreementService from '../services/PurchaseAgreementService'

class ListPurchaseAgreementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                purchaseAgreements: []
        }
        this.addPurchaseAgreement = this.addPurchaseAgreement.bind(this);
        this.editPurchaseAgreement = this.editPurchaseAgreement.bind(this);
        this.deletePurchaseAgreement = this.deletePurchaseAgreement.bind(this);
    }

    deletePurchaseAgreement(id){
        PurchaseAgreementService.deletePurchaseAgreement(id).then( res => {
            this.setState({purchaseAgreements: this.state.purchaseAgreements.filter(purchaseAgreement => purchaseAgreement.purchaseAgreementId !== id)});
        });
    }
    viewPurchaseAgreement(id){
        this.props.history.push(`/view-purchaseAgreement/${id}`);
    }
    editPurchaseAgreement(id){
        this.props.history.push(`/add-purchaseAgreement/${id}`);
    }

    componentDidMount(){
        PurchaseAgreementService.getPurchaseAgreements().then((res) => {
            this.setState({ purchaseAgreements: res.data});
        });
    }

    addPurchaseAgreement(){
        this.props.history.push('/add-purchaseAgreement/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">PurchaseAgreement List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPurchaseAgreement}> Add PurchaseAgreement</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> AgreementNumber </th>
                                    <th> EffectiveDate </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.purchaseAgreements.map(
                                        purchaseAgreement => 
                                        <tr key = {purchaseAgreement.purchaseAgreementId}>
                                             <td> { purchaseAgreement.agreementNumber } </td>
                                             <td> { purchaseAgreement.effectiveDate } </td>
                                             <td>
                                                 <button onClick={ () => this.editPurchaseAgreement(purchaseAgreement.purchaseAgreementId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePurchaseAgreement(purchaseAgreement.purchaseAgreementId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPurchaseAgreement(purchaseAgreement.purchaseAgreementId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPurchaseAgreementComponent
