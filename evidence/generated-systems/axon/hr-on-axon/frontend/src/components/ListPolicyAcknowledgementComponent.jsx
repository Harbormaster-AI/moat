import React, { Component } from 'react'
import PolicyAcknowledgementService from '../services/PolicyAcknowledgementService'

class ListPolicyAcknowledgementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                policyAcknowledgements: []
        }
        this.addPolicyAcknowledgement = this.addPolicyAcknowledgement.bind(this);
        this.editPolicyAcknowledgement = this.editPolicyAcknowledgement.bind(this);
        this.deletePolicyAcknowledgement = this.deletePolicyAcknowledgement.bind(this);
    }

    deletePolicyAcknowledgement(id){
        PolicyAcknowledgementService.deletePolicyAcknowledgement(id).then( res => {
            this.setState({policyAcknowledgements: this.state.policyAcknowledgements.filter(policyAcknowledgement => policyAcknowledgement.policyAcknowledgementId !== id)});
        });
    }
    viewPolicyAcknowledgement(id){
        this.props.history.push(`/view-policyAcknowledgement/${id}`);
    }
    editPolicyAcknowledgement(id){
        this.props.history.push(`/add-policyAcknowledgement/${id}`);
    }

    componentDidMount(){
        PolicyAcknowledgementService.getPolicyAcknowledgements().then((res) => {
            this.setState({ policyAcknowledgements: res.data});
        });
    }

    addPolicyAcknowledgement(){
        this.props.history.push('/add-policyAcknowledgement/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">PolicyAcknowledgement List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPolicyAcknowledgement}> Add PolicyAcknowledgement</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> AcknowledgementDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.policyAcknowledgements.map(
                                        policyAcknowledgement => 
                                        <tr key = {policyAcknowledgement.policyAcknowledgementId}>
                                             <td> { policyAcknowledgement.acknowledgementDate } </td>
                                             <td> { policyAcknowledgement.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editPolicyAcknowledgement(policyAcknowledgement.policyAcknowledgementId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePolicyAcknowledgement(policyAcknowledgement.policyAcknowledgementId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPolicyAcknowledgement(policyAcknowledgement.policyAcknowledgementId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPolicyAcknowledgementComponent
