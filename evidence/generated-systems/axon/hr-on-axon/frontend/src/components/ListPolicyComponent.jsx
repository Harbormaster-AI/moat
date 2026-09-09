import React, { Component } from 'react'
import PolicyService from '../services/PolicyService'

class ListPolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                policys: []
        }
        this.addPolicy = this.addPolicy.bind(this);
        this.editPolicy = this.editPolicy.bind(this);
        this.deletePolicy = this.deletePolicy.bind(this);
    }

    deletePolicy(id){
        PolicyService.deletePolicy(id).then( res => {
            this.setState({policys: this.state.policys.filter(policy => policy.policyId !== id)});
        });
    }
    viewPolicy(id){
        this.props.history.push(`/view-policy/${id}`);
    }
    editPolicy(id){
        this.props.history.push(`/add-policy/${id}`);
    }

    componentDidMount(){
        PolicyService.getPolicys().then((res) => {
            this.setState({ policys: res.data});
        });
    }

    addPolicy(){
        this.props.history.push('/add-policy/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Policy List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPolicy}> Add Policy</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> PolicyNumber </th>
                                    <th> Name </th>
                                    <th> EffectiveDate </th>
                                    <th> Description </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.policys.map(
                                        policy => 
                                        <tr key = {policy.policyId}>
                                             <td> { policy.policyNumber } </td>
                                             <td> { policy.name } </td>
                                             <td> { policy.effectiveDate } </td>
                                             <td> { policy.description } </td>
                                             <td>
                                                 <button onClick={ () => this.editPolicy(policy.policyId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePolicy(policy.policyId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPolicy(policy.policyId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPolicyComponent
