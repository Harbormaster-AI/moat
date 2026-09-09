import React, { Component } from 'react'
import LoanTransactionService from '../services/LoanTransactionService';

class UpdateLoanTransactionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                transactionId: '',
                amount: '',
                postingDate: '',
                type: '',
                status: ''
        }
        this.updateLoanTransaction = this.updateLoanTransaction.bind(this);

        this.changetransactionIdHandler = this.changetransactionIdHandler.bind(this);
        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changepostingDateHandler = this.changepostingDateHandler.bind(this);
        this.changeTypeHandler = this.changeTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        LoanTransactionService.getLoanTransactionById(this.state.id).then( (res) =>{
            let loanTransaction = res.data;
            this.setState({
                transactionId: loanTransaction.transactionId,
                amount: loanTransaction.amount,
                postingDate: loanTransaction.postingDate,
                type: loanTransaction.type,
                status: loanTransaction.status
            });
        });
    }

    updateLoanTransaction = (e) => {
        e.preventDefault();
        let loanTransaction = {
            loanTransactionId: this.state.id,
            transactionId: this.state.transactionId,
            amount: this.state.amount,
            postingDate: this.state.postingDate,
            type: this.state.type,
            status: this.state.status
        };
        console.log('loanTransaction => ' + JSON.stringify(loanTransaction));
        console.log('id => ' + JSON.stringify(this.state.id));
        LoanTransactionService.updateLoanTransaction(loanTransaction).then( res => {
            this.props.history.push('/loanTransactions');
        });
    }

    changetransactionIdHandler= (event) => {
        this.setState({transactionId: event.target.value});
    }
    changeamountHandler= (event) => {
        this.setState({amount: event.target.value});
    }
    changepostingDateHandler= (event) => {
        this.setState({postingDate: event.target.value});
    }
    changeTypeHandler= (event) => {
        this.setState({type: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/loanTransactions');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update LoanTransaction</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> transactionId: </label>
                                                <input placeholder="transactionId" name="transactionId" className="form-control" value={this.state.transactionId} onChange={this.changetransactionIdHandler}/>

                                            <label> amount: </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> postingDate: </label>
                                                <input type="date" placeholder="postingDate" name="postingDate" className="form-control" value={this.state.postingDate} onChange={this.changepostingDateHandler}/>

                                            <label> Type: </label>
                                                <select value={this.state.type} onChange={this.changeTypeHandler}>
                      <option name="Type" className="form-control" >
                          Disbursement
                      </option>
                      <option name="Type" className="form-control" >
                          Repayment
                      </option>
                      <option name="Type" className="form-control" >
                          Interest
                      </option>
                      <option name="Type" className="form-control" >
                          Fee
                      </option>
                      <option name="Type" className="form-control" >
                          Reversal
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Pending
                      </option>
                      <option name="Status" className="form-control" >
                          Posted
                      </option>
                      <option name="Status" className="form-control" >
                          Reversed
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateLoanTransaction}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateLoanTransactionComponent
