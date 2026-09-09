import React, { Component } from 'react'
import TransactionService from '../services/TransactionService';

class CreateTransactionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                amount: '',
                fee: '',
                exchangeRate: '',
                createdAt: '',
                completedAt: '',
                narrative: '',
                transactionType: '',
                status: ''
        }
        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changefeeHandler = this.changefeeHandler.bind(this);
        this.changeexchangeRateHandler = this.changeexchangeRateHandler.bind(this);
        this.changecreatedAtHandler = this.changecreatedAtHandler.bind(this);
        this.changecompletedAtHandler = this.changecompletedAtHandler.bind(this);
        this.changenarrativeHandler = this.changenarrativeHandler.bind(this);
        this.changeTransactionTypeHandler = this.changeTransactionTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            TransactionService.getTransactionById(this.state.id).then( (res) =>{
                let transaction = res.data;
                this.setState({
                    amount: transaction.amount,
                    fee: transaction.fee,
                    exchangeRate: transaction.exchangeRate,
                    createdAt: transaction.createdAt,
                    completedAt: transaction.completedAt,
                    narrative: transaction.narrative,
                    transactionType: transaction.transactionType,
                    status: transaction.status
                });
            });
        }        
    }
    saveOrUpdateTransaction = (e) => {
        e.preventDefault();
        let transaction = {
                transactionId: this.state.id,
                amount: this.state.amount,
                fee: this.state.fee,
                exchangeRate: this.state.exchangeRate,
                createdAt: this.state.createdAt,
                completedAt: this.state.completedAt,
                narrative: this.state.narrative,
                transactionType: this.state.transactionType,
                status: this.state.status
            };
        console.log('transaction => ' + JSON.stringify(transaction));

        // step 5
        if(this.state.id === '_add'){
            transaction.transactionId=''
            TransactionService.createTransaction(transaction).then(res =>{
                this.props.history.push('/transactions');
            });
        }else{
            TransactionService.updateTransaction(transaction).then( res => {
                this.props.history.push('/transactions');
            });
        }
    }
    
    changeamountHandler= (event) => {
        this.setState({amount: event.target.value});
    }
    changefeeHandler= (event) => {
        this.setState({fee: event.target.value});
    }
    changeexchangeRateHandler= (event) => {
        this.setState({exchangeRate: event.target.value});
    }
    changecreatedAtHandler= (event) => {
        this.setState({createdAt: event.target.value});
    }
    changecompletedAtHandler= (event) => {
        this.setState({completedAt: event.target.value});
    }
    changenarrativeHandler= (event) => {
        this.setState({narrative: event.target.value});
    }
    changeTransactionTypeHandler= (event) => {
        this.setState({transactionType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/transactions');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Transaction</h3>
        }else{
            return <h3 className="text-center">Update Transaction</h3>
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
                                            <label> amount:&emsp; </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> fee:&emsp; </label>
                                                <input placeholder="fee" name="fee" className="form-control" value={this.state.fee} onChange={this.changefeeHandler}/>

                                            <label> exchangeRate:&emsp; </label>
                                                <input placeholder="exchangeRate" name="exchangeRate" className="form-control" value={this.state.exchangeRate} onChange={this.changeexchangeRateHandler}/>

                                            <label> createdAt:&emsp; </label>
                                                <input type="time" placeholder="createdAt" name="createdAt" className="form-control" value={this.state.createdAt} onChange={this.changecreatedAtHandler}/>

                                            <label> completedAt:&emsp; </label>
                                                <input type="time" placeholder="completedAt" name="completedAt" className="form-control" value={this.state.completedAt} onChange={this.changecompletedAtHandler}/>

                                            <label> narrative:&emsp; </label>
                                                <input placeholder="narrative" name="narrative" className="form-control" value={this.state.narrative} onChange={this.changenarrativeHandler}/>

                                            <label> TransactionType:&emsp; </label>
                                                <select value={this.state.transactionType} onChange={this.changeTransactionTypeHandler}>
                      <option name="TransactionType" className="form-control" >
                          Deposit
                      </option>
                      <option name="TransactionType" className="form-control" >
                          Withdrawal
                      </option>
                      <option name="TransactionType" className="form-control" >
                          Transfer
                      </option>
                      <option name="TransactionType" className="form-control" >
                          Payment
                      </option>
                      <option name="TransactionType" className="form-control" >
                          Refund
                      </option>
                      <option name="TransactionType" className="form-control" >
                          Fee
                      </option>
                      <option name="TransactionType" className="form-control" >
                          Interest
                      </option>
                      <option name="TransactionType" className="form-control" >
                          FXConversion
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Pending
                      </option>
                      <option name="Status" className="form-control" >
                          Authorized
                      </option>
                      <option name="Status" className="form-control" >
                          Posted
                      </option>
                      <option name="Status" className="form-control" >
                          Settled
                      </option>
                      <option name="Status" className="form-control" >
                          Reversed
                      </option>
                      <option name="Status" className="form-control" >
                          Failed
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateTransaction}>Save</button>
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

export default CreateTransactionComponent
