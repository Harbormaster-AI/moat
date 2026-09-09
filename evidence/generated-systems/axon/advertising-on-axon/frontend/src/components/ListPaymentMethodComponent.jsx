import React, { Component } from 'react'
import PaymentMethodService from '../services/PaymentMethodService'

class ListPaymentMethodComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                paymentMethods: []
        }
        this.addPaymentMethod = this.addPaymentMethod.bind(this);
        this.editPaymentMethod = this.editPaymentMethod.bind(this);
        this.deletePaymentMethod = this.deletePaymentMethod.bind(this);
    }

    deletePaymentMethod(id){
        PaymentMethodService.deletePaymentMethod(id).then( res => {
            this.setState({paymentMethods: this.state.paymentMethods.filter(paymentMethod => paymentMethod.paymentMethodId !== id)});
        });
    }
    viewPaymentMethod(id){
        this.props.history.push(`/view-paymentMethod/${id}`);
    }
    editPaymentMethod(id){
        this.props.history.push(`/add-paymentMethod/${id}`);
    }

    componentDidMount(){
        PaymentMethodService.getPaymentMethods().then((res) => {
            this.setState({ paymentMethods: res.data});
        });
    }

    addPaymentMethod(){
        this.props.history.push('/add-paymentMethod/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">PaymentMethod List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPaymentMethod}> Add PaymentMethod</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Last4 </th>
                                    <th> CardholderName </th>
                                    <th> BillingAddress </th>
                                    <th> MethodType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.paymentMethods.map(
                                        paymentMethod => 
                                        <tr key = {paymentMethod.paymentMethodId}>
                                             <td> { paymentMethod.last4 } </td>
                                             <td> { paymentMethod.cardholderName } </td>
                                             <td> { paymentMethod.billingAddress } </td>
                                             <td> { paymentMethod.methodType } </td>
                                             <td>
                                                 <button onClick={ () => this.editPaymentMethod(paymentMethod.paymentMethodId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePaymentMethod(paymentMethod.paymentMethodId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPaymentMethod(paymentMethod.paymentMethodId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPaymentMethodComponent
