import React, { Component } from 'react'
import LoanApplicationService from '../services/LoanApplicationService'

class ListLoanApplicationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                loanApplications: []
        }
        this.addLoanApplication = this.addLoanApplication.bind(this);
        this.editLoanApplication = this.editLoanApplication.bind(this);
        this.deleteLoanApplication = this.deleteLoanApplication.bind(this);
    }

    deleteLoanApplication(id){
        LoanApplicationService.deleteLoanApplication(id).then( res => {
            this.setState({loanApplications: this.state.loanApplications.filter(loanApplication => loanApplication.loanApplicationId !== id)});
        });
    }
    viewLoanApplication(id){
        this.props.history.push(`/view-loanApplication/${id}`);
    }
    editLoanApplication(id){
        this.props.history.push(`/add-loanApplication/${id}`);
    }

    componentDidMount(){
        LoanApplicationService.getLoanApplications().then((res) => {
            this.setState({ loanApplications: res.data});
        });
    }

    addLoanApplication(){
        this.props.history.push('/add-loanApplication/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">LoanApplication List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addLoanApplication}> Add LoanApplication</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ApplicationNumber </th>
                                    <th> AmountRequested </th>
                                    <th> TermMonths </th>
                                    <th> SubmittedAt </th>
                                    <th> Product </th>
                                    <th> Purpose </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.loanApplications.map(
                                        loanApplication => 
                                        <tr key = {loanApplication.loanApplicationId}>
                                             <td> { loanApplication.applicationNumber } </td>
                                             <td> { loanApplication.amountRequested } </td>
                                             <td> { loanApplication.termMonths } </td>
                                             <td> { loanApplication.submittedAt } </td>
                                             <td> { loanApplication.product } </td>
                                             <td> { loanApplication.purpose } </td>
                                             <td> { loanApplication.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editLoanApplication(loanApplication.loanApplicationId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteLoanApplication(loanApplication.loanApplicationId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewLoanApplication(loanApplication.loanApplicationId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListLoanApplicationComponent
