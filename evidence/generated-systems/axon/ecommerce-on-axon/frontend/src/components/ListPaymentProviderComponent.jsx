import React, { Component } from 'react'
import PaymentProviderService from '../services/PaymentProviderService'

class ListPaymentProviderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                paymentProviders: []
        }
        this.addPaymentProvider = this.addPaymentProvider.bind(this);
        this.editPaymentProvider = this.editPaymentProvider.bind(this);
        this.deletePaymentProvider = this.deletePaymentProvider.bind(this);
    }

    deletePaymentProvider(id){
        PaymentProviderService.deletePaymentProvider(id).then( res => {
            this.setState({paymentProviders: this.state.paymentProviders.filter(paymentProvider => paymentProvider.paymentProviderId !== id)});
        });
    }
    viewPaymentProvider(id){
        this.props.history.push(`/view-paymentProvider/${id}`);
    }
    editPaymentProvider(id){
        this.props.history.push(`/add-paymentProvider/${id}`);
    }

    componentDidMount(){
        PaymentProviderService.getPaymentProviders().then((res) => {
            this.setState({ paymentProviders: res.data});
        });
    }

    addPaymentProvider(){
        this.props.history.push('/add-paymentProvider/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">PaymentProvider List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPaymentProvider}> Add PaymentProvider</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Enabled </th>
                                    <th> MerchantAccountId </th>
                                    <th> ProviderType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.paymentProviders.map(
                                        paymentProvider => 
                                        <tr key = {paymentProvider.paymentProviderId}>
                                             <td> { paymentProvider.name } </td>
                                             <td> { paymentProvider.enabled } </td>
                                             <td> { paymentProvider.merchantAccountId } </td>
                                             <td> { paymentProvider.providerType } </td>
                                             <td>
                                                 <button onClick={ () => this.editPaymentProvider(paymentProvider.paymentProviderId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePaymentProvider(paymentProvider.paymentProviderId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPaymentProvider(paymentProvider.paymentProviderId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPaymentProviderComponent
