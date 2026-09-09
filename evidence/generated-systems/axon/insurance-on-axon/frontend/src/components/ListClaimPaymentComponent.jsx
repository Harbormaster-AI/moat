import React, { Component } from 'react'
import ClaimPaymentService from '../services/ClaimPaymentService'

class ListClaimPaymentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                claimPayments: []
        }
        this.addClaimPayment = this.addClaimPayment.bind(this);
        this.editClaimPayment = this.editClaimPayment.bind(this);
        this.deleteClaimPayment = this.deleteClaimPayment.bind(this);
    }

    deleteClaimPayment(id){
        ClaimPaymentService.deleteClaimPayment(id).then( res => {
            this.setState({claimPayments: this.state.claimPayments.filter(claimPayment => claimPayment.claimPaymentId !== id)});
        });
    }
    viewClaimPayment(id){
        this.props.history.push(`/view-claimPayment/${id}`);
    }
    editClaimPayment(id){
        this.props.history.push(`/add-claimPayment/${id}`);
    }

    componentDidMount(){
        ClaimPaymentService.getClaimPayments().then((res) => {
            this.setState({ claimPayments: res.data});
        });
    }

    addClaimPayment(){
        this.props.history.push('/add-claimPayment/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ClaimPayment List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addClaimPayment}> Add ClaimPayment</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> PaymentNumber </th>
                                    <th> Amount </th>
                                    <th> PaymentDate </th>
                                    <th> PayeeType </th>
                                    <th> Method </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.claimPayments.map(
                                        claimPayment => 
                                        <tr key = {claimPayment.claimPaymentId}>
                                             <td> { claimPayment.paymentNumber } </td>
                                             <td> { claimPayment.amount } </td>
                                             <td> { claimPayment.paymentDate } </td>
                                             <td> { claimPayment.payeeType } </td>
                                             <td> { claimPayment.method } </td>
                                             <td> { claimPayment.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editClaimPayment(claimPayment.claimPaymentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteClaimPayment(claimPayment.claimPaymentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewClaimPayment(claimPayment.claimPaymentId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListClaimPaymentComponent
