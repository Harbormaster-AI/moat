import React, { Component } from 'react'
import LoanApplicationService from '../services/LoanApplicationService'

class ViewLoanApplicationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            loanApplication: {}
        }
    }

    componentDidMount(){
        LoanApplicationService.getLoanApplicationById(this.state.id).then( res => {
            this.setState({loanApplication: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View LoanApplication Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> applicationNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.loanApplication.applicationNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> amountRequested:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.loanApplication.amountRequested }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> termMonths:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.loanApplication.termMonths }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> submittedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.loanApplication.submittedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Product:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.loanApplication.product }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Purpose:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.loanApplication.purpose }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.loanApplication.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewLoanApplicationComponent
