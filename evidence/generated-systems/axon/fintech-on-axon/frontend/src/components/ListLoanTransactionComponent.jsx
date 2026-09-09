import React, { Component } from 'react'
import LoanTransactionService from '../services/LoanTransactionService'

class ListLoanTransactionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                loanTransactions: []
        }
        this.addLoanTransaction = this.addLoanTransaction.bind(this);
        this.editLoanTransaction = this.editLoanTransaction.bind(this);
        this.deleteLoanTransaction = this.deleteLoanTransaction.bind(this);
    }

    deleteLoanTransaction(id){
        LoanTransactionService.deleteLoanTransaction(id).then( res => {
            this.setState({loanTransactions: this.state.loanTransactions.filter(loanTransaction => loanTransaction.loanTransactionId !== id)});
        });
    }
    viewLoanTransaction(id){
        this.props.history.push(`/view-loanTransaction/${id}`);
    }
    editLoanTransaction(id){
        this.props.history.push(`/add-loanTransaction/${id}`);
    }

    componentDidMount(){
        LoanTransactionService.getLoanTransactions().then((res) => {
            this.setState({ loanTransactions: res.data});
        });
    }

    addLoanTransaction(){
        this.props.history.push('/add-loanTransaction/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">LoanTransaction List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addLoanTransaction}> Add LoanTransaction</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> TransactionId </th>
                                    <th> Amount </th>
                                    <th> PostingDate </th>
                                    <th> Type </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.loanTransactions.map(
                                        loanTransaction => 
                                        <tr key = {loanTransaction.loanTransactionId}>
                                             <td> { loanTransaction.transactionId } </td>
                                             <td> { loanTransaction.amount } </td>
                                             <td> { loanTransaction.postingDate } </td>
                                             <td> { loanTransaction.type } </td>
                                             <td> { loanTransaction.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editLoanTransaction(loanTransaction.loanTransactionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteLoanTransaction(loanTransaction.loanTransactionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewLoanTransaction(loanTransaction.loanTransactionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListLoanTransactionComponent
