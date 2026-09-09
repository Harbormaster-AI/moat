import React, { Component } from 'react'
import EmploymentContractService from '../services/EmploymentContractService'

class ListEmploymentContractComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                employmentContracts: []
        }
        this.addEmploymentContract = this.addEmploymentContract.bind(this);
        this.editEmploymentContract = this.editEmploymentContract.bind(this);
        this.deleteEmploymentContract = this.deleteEmploymentContract.bind(this);
    }

    deleteEmploymentContract(id){
        EmploymentContractService.deleteEmploymentContract(id).then( res => {
            this.setState({employmentContracts: this.state.employmentContracts.filter(employmentContract => employmentContract.employmentContractId !== id)});
        });
    }
    viewEmploymentContract(id){
        this.props.history.push(`/view-employmentContract/${id}`);
    }
    editEmploymentContract(id){
        this.props.history.push(`/add-employmentContract/${id}`);
    }

    componentDidMount(){
        EmploymentContractService.getEmploymentContracts().then((res) => {
            this.setState({ employmentContracts: res.data});
        });
    }

    addEmploymentContract(){
        this.props.history.push('/add-employmentContract/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">EmploymentContract List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addEmploymentContract}> Add EmploymentContract</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ContractNumber </th>
                                    <th> StartDate </th>
                                    <th> EndDate </th>
                                    <th> WorkHoursPerWeek </th>
                                    <th> EmploymentType </th>
                                    <th> Status </th>
                                    <th> PayFrequency </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.employmentContracts.map(
                                        employmentContract => 
                                        <tr key = {employmentContract.employmentContractId}>
                                             <td> { employmentContract.contractNumber } </td>
                                             <td> { employmentContract.startDate } </td>
                                             <td> { employmentContract.endDate } </td>
                                             <td> { employmentContract.workHoursPerWeek } </td>
                                             <td> { employmentContract.employmentType } </td>
                                             <td> { employmentContract.status } </td>
                                             <td> { employmentContract.payFrequency } </td>
                                             <td>
                                                 <button onClick={ () => this.editEmploymentContract(employmentContract.employmentContractId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteEmploymentContract(employmentContract.employmentContractId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewEmploymentContract(employmentContract.employmentContractId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListEmploymentContractComponent
