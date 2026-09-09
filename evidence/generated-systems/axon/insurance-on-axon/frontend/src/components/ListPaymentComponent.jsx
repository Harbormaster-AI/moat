import React, { Component } from 'react'
import PaymentService from '../services/PaymentService'

class ListPaymentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                payments: []
        }
        this.addPayment = this.addPayment.bind(this);
        this.editPayment = this.editPayment.bind(this);
        this.deletePayment = this.deletePayment.bind(this);
    }

    deletePayment(id){
        PaymentService.deletePayment(id).then( res => {
            this.setState({payments: this.state.payments.filter(payment => payment.paymentId !== id)});
        });
    }
    viewPayment(id){
        this.props.history.push(`/view-payment/${id}`);
    }
    editPayment(id){
        this.props.history.push(`/add-payment/${id}`);
    }

    componentDidMount(){
        PaymentService.getPayments().then((res) => {
            this.setState({ payments: res.data});
        });
    }

    addPayment(){
        this.props.history.push('/add-payment/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Payment List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPayment}> Add Payment</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> PaymentReference </th>
                                    <th> Amount </th>
                                    <th> PaymentDate </th>
                                    <th> Method </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.payments.map(
                                        payment => 
                                        <tr key = {payment.paymentId}>
                                             <td> { payment.paymentReference } </td>
                                             <td> { payment.amount } </td>
                                             <td> { payment.paymentDate } </td>
                                             <td> { payment.method } </td>
                                             <td> { payment.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editPayment(payment.paymentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePayment(payment.paymentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPayment(payment.paymentId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPaymentComponent
