import React, { Component } from 'react'
import LoanTransactionService from '../services/LoanTransactionService'

class ViewLoanTransactionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            loanTransaction: {}
        }
    }

    componentDidMount(){
        LoanTransactionService.getLoanTransactionById(this.state.id).then( res => {
            this.setState({loanTransaction: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View LoanTransaction Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> transactionId:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.loanTransaction.transactionId }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> amount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.loanTransaction.amount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> postingDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.loanTransaction.postingDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Type:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.loanTransaction.type }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.loanTransaction.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewLoanTransactionComponent
