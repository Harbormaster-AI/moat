import React, { Component } from 'react'
import LoanTransactionService from '../services/LoanTransactionService';

class CreateLoanTransactionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                transactionId: '',
                amount: '',
                postingDate: '',
                type: '',
                status: ''
        }
        this.changetransactionIdHandler = this.changetransactionIdHandler.bind(this);
        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changepostingDateHandler = this.changepostingDateHandler.bind(this);
        this.changeTypeHandler = this.changeTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateLoanTransaction = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            loanTransaction.loanTransactionId=''
            LoanTransactionService.createLoanTransaction(loanTransaction).then(res =>{
                this.props.history.push('/loanTransactions');
            });
        }else{
            LoanTransactionService.updateLoanTransaction(loanTransaction).then( res => {
                this.props.history.push('/loanTransactions');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add LoanTransaction</h3>
        }else{
            return <h3 className="text-center">Update LoanTransaction</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> transactionId:&emsp; </label>
                                                <input placeholder="transactionId" name="transactionId" className="form-control" value={this.state.transactionId} onChange={this.changetransactionIdHandler}/>

                                            <label> amount:&emsp; </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> postingDate:&emsp; </label>
                                                <input type="date" placeholder="postingDate" name="postingDate" className="form-control" value={this.state.postingDate} onChange={this.changepostingDateHandler}/>

                                            <label> Type:&emsp; </label>
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

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateLoanTransaction}>Save</button>
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

export default CreateLoanTransactionComponent
