import React, { Component } from 'react'
import TransactionService from '../services/TransactionService'

class ListTransactionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                transactions: []
        }
        this.addTransaction = this.addTransaction.bind(this);
        this.editTransaction = this.editTransaction.bind(this);
        this.deleteTransaction = this.deleteTransaction.bind(this);
    }

    deleteTransaction(id){
        TransactionService.deleteTransaction(id).then( res => {
            this.setState({transactions: this.state.transactions.filter(transaction => transaction.transactionId !== id)});
        });
    }
    viewTransaction(id){
        this.props.history.push(`/view-transaction/${id}`);
    }
    editTransaction(id){
        this.props.history.push(`/add-transaction/${id}`);
    }

    componentDidMount(){
        TransactionService.getTransactions().then((res) => {
            this.setState({ transactions: res.data});
        });
    }

    addTransaction(){
        this.props.history.push('/add-transaction/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Transaction List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addTransaction}> Add Transaction</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Amount </th>
                                    <th> Fee </th>
                                    <th> ExchangeRate </th>
                                    <th> CreatedAt </th>
                                    <th> CompletedAt </th>
                                    <th> Narrative </th>
                                    <th> TransactionType </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.transactions.map(
                                        transaction => 
                                        <tr key = {transaction.transactionId}>
                                             <td> { transaction.amount } </td>
                                             <td> { transaction.fee } </td>
                                             <td> { transaction.exchangeRate } </td>
                                             <td> { transaction.createdAt } </td>
                                             <td> { transaction.completedAt } </td>
                                             <td> { transaction.narrative } </td>
                                             <td> { transaction.transactionType } </td>
                                             <td> { transaction.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editTransaction(transaction.transactionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteTransaction(transaction.transactionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewTransaction(transaction.transactionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListTransactionComponent
