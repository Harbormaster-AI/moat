import React, { Component } from 'react'
import LoanService from '../services/LoanService'

class ViewLoanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            loan: {}
        }
    }

    componentDidMount(){
        LoanService.getLoanById(this.state.id).then( res => {
            this.setState({loan: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Loan Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> loanNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.loan.loanNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> principal:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.loan.principal }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> interestRate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.loan.interestRate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> originationDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.loan.originationDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> maturityDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.loan.maturityDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> RateType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.loan.rateType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.loan.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewLoanComponent
