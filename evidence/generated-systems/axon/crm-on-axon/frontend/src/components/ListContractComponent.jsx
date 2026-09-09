import React, { Component } from 'react'
import ContractService from '../services/ContractService'

class ListContractComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                contracts: []
        }
        this.addContract = this.addContract.bind(this);
        this.editContract = this.editContract.bind(this);
        this.deleteContract = this.deleteContract.bind(this);
    }

    deleteContract(id){
        ContractService.deleteContract(id).then( res => {
            this.setState({contracts: this.state.contracts.filter(contract => contract.contractId !== id)});
        });
    }
    viewContract(id){
        this.props.history.push(`/view-contract/${id}`);
    }
    editContract(id){
        this.props.history.push(`/add-contract/${id}`);
    }

    componentDidMount(){
        ContractService.getContracts().then((res) => {
            this.setState({ contracts: res.data});
        });
    }

    addContract(){
        this.props.history.push('/add-contract/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Contract List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addContract}> Add Contract</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ContractNumber </th>
                                    <th> StartDate </th>
                                    <th> EndDate </th>
                                    <th> RenewalTermMonths </th>
                                    <th> AutoRenew </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.contracts.map(
                                        contract => 
                                        <tr key = {contract.contractId}>
                                             <td> { contract.contractNumber } </td>
                                             <td> { contract.startDate } </td>
                                             <td> { contract.endDate } </td>
                                             <td> { contract.renewalTermMonths } </td>
                                             <td> { contract.autoRenew } </td>
                                             <td> { contract.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editContract(contract.contractId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteContract(contract.contractId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewContract(contract.contractId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListContractComponent
