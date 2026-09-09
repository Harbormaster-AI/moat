import React, { Component } from 'react'
import ExpirationPolicyService from '../services/ExpirationPolicyService'

class ListExpirationPolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                expirationPolicys: []
        }
        this.addExpirationPolicy = this.addExpirationPolicy.bind(this);
        this.editExpirationPolicy = this.editExpirationPolicy.bind(this);
        this.deleteExpirationPolicy = this.deleteExpirationPolicy.bind(this);
    }

    deleteExpirationPolicy(id){
        ExpirationPolicyService.deleteExpirationPolicy(id).then( res => {
            this.setState({expirationPolicys: this.state.expirationPolicys.filter(expirationPolicy => expirationPolicy.expirationPolicyId !== id)});
        });
    }
    viewExpirationPolicy(id){
        this.props.history.push(`/view-expirationPolicy/${id}`);
    }
    editExpirationPolicy(id){
        this.props.history.push(`/add-expirationPolicy/${id}`);
    }

    componentDidMount(){
        ExpirationPolicyService.getExpirationPolicys().then((res) => {
            this.setState({ expirationPolicys: res.data});
        });
    }

    addExpirationPolicy(){
        this.props.history.push('/add-expirationPolicy/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ExpirationPolicy List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addExpirationPolicy}> Add ExpirationPolicy</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> RejectIfDaysToExpireLessThan </th>
                                    <th> AutoQuarantineDaysToExpire </th>
                                    <th> RotationMethod </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.expirationPolicys.map(
                                        expirationPolicy => 
                                        <tr key = {expirationPolicy.expirationPolicyId}>
                                             <td> { expirationPolicy.rejectIfDaysToExpireLessThan } </td>
                                             <td> { expirationPolicy.autoQuarantineDaysToExpire } </td>
                                             <td> { expirationPolicy.rotationMethod } </td>
                                             <td>
                                                 <button onClick={ () => this.editExpirationPolicy(expirationPolicy.expirationPolicyId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteExpirationPolicy(expirationPolicy.expirationPolicyId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewExpirationPolicy(expirationPolicy.expirationPolicyId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListExpirationPolicyComponent
