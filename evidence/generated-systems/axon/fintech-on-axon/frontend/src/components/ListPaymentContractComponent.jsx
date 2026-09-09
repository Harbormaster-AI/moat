import React, { Component } from 'react'
import PaymentContractService from '../services/PaymentContractService'

class ListPaymentContractComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                paymentContracts: []
        }
        this.addPaymentContract = this.addPaymentContract.bind(this);
        this.editPaymentContract = this.editPaymentContract.bind(this);
        this.deletePaymentContract = this.deletePaymentContract.bind(this);
    }

    deletePaymentContract(id){
        PaymentContractService.deletePaymentContract(id).then( res => {
            this.setState({paymentContracts: this.state.paymentContracts.filter(paymentContract => paymentContract.paymentContractId !== id)});
        });
    }
    viewPaymentContract(id){
        this.props.history.push(`/view-paymentContract/${id}`);
    }
    editPaymentContract(id){
        this.props.history.push(`/add-paymentContract/${id}`);
    }

    componentDidMount(){
        PaymentContractService.getPaymentContracts().then((res) => {
            this.setState({ paymentContracts: res.data});
        });
    }

    addPaymentContract(){
        this.props.history.push('/add-paymentContract/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">PaymentContract List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPaymentContract}> Add PaymentContract</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ContractNumber </th>
                                    <th> PricingPlanCode </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.paymentContracts.map(
                                        paymentContract => 
                                        <tr key = {paymentContract.paymentContractId}>
                                             <td> { paymentContract.contractNumber } </td>
                                             <td> { paymentContract.pricingPlanCode } </td>
                                             <td> { paymentContract.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editPaymentContract(paymentContract.paymentContractId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePaymentContract(paymentContract.paymentContractId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPaymentContract(paymentContract.paymentContractId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPaymentContractComponent
