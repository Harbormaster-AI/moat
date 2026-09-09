import React, { Component } from 'react'
import LoanService from '../services/LoanService'

class ListLoanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                loans: []
        }
        this.addLoan = this.addLoan.bind(this);
        this.editLoan = this.editLoan.bind(this);
        this.deleteLoan = this.deleteLoan.bind(this);
    }

    deleteLoan(id){
        LoanService.deleteLoan(id).then( res => {
            this.setState({loans: this.state.loans.filter(loan => loan.loanId !== id)});
        });
    }
    viewLoan(id){
        this.props.history.push(`/view-loan/${id}`);
    }
    editLoan(id){
        this.props.history.push(`/add-loan/${id}`);
    }

    componentDidMount(){
        LoanService.getLoans().then((res) => {
            this.setState({ loans: res.data});
        });
    }

    addLoan(){
        this.props.history.push('/add-loan/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Loan List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addLoan}> Add Loan</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> LoanNumber </th>
                                    <th> Principal </th>
                                    <th> InterestRate </th>
                                    <th> OriginationDate </th>
                                    <th> MaturityDate </th>
                                    <th> RateType </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.loans.map(
                                        loan => 
                                        <tr key = {loan.loanId}>
                                             <td> { loan.loanNumber } </td>
                                             <td> { loan.principal } </td>
                                             <td> { loan.interestRate } </td>
                                             <td> { loan.originationDate } </td>
                                             <td> { loan.maturityDate } </td>
                                             <td> { loan.rateType } </td>
                                             <td> { loan.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editLoan(loan.loanId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteLoan(loan.loanId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewLoan(loan.loanId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListLoanComponent
