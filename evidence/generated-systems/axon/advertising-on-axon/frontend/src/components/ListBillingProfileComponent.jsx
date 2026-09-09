import React, { Component } from 'react'
import BillingProfileService from '../services/BillingProfileService'

class ListBillingProfileComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                billingProfiles: []
        }
        this.addBillingProfile = this.addBillingProfile.bind(this);
        this.editBillingProfile = this.editBillingProfile.bind(this);
        this.deleteBillingProfile = this.deleteBillingProfile.bind(this);
    }

    deleteBillingProfile(id){
        BillingProfileService.deleteBillingProfile(id).then( res => {
            this.setState({billingProfiles: this.state.billingProfiles.filter(billingProfile => billingProfile.billingProfileId !== id)});
        });
    }
    viewBillingProfile(id){
        this.props.history.push(`/view-billingProfile/${id}`);
    }
    editBillingProfile(id){
        this.props.history.push(`/add-billingProfile/${id}`);
    }

    componentDidMount(){
        BillingProfileService.getBillingProfiles().then((res) => {
            this.setState({ billingProfiles: res.data});
        });
    }

    addBillingProfile(){
        this.props.history.push('/add-billingProfile/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">BillingProfile List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addBillingProfile}> Add BillingProfile</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> BillingName </th>
                                    <th> TaxId </th>
                                    <th> BillingAddress </th>
                                    <th> PaymentTerms </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.billingProfiles.map(
                                        billingProfile => 
                                        <tr key = {billingProfile.billingProfileId}>
                                             <td> { billingProfile.billingName } </td>
                                             <td> { billingProfile.taxId } </td>
                                             <td> { billingProfile.billingAddress } </td>
                                             <td> { billingProfile.paymentTerms } </td>
                                             <td>
                                                 <button onClick={ () => this.editBillingProfile(billingProfile.billingProfileId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteBillingProfile(billingProfile.billingProfileId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewBillingProfile(billingProfile.billingProfileId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListBillingProfileComponent
