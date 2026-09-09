import React, { Component } from 'react'
import ApprovalService from '../services/ApprovalService'

class ListApprovalComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                approvals: []
        }
        this.addApproval = this.addApproval.bind(this);
        this.editApproval = this.editApproval.bind(this);
        this.deleteApproval = this.deleteApproval.bind(this);
    }

    deleteApproval(id){
        ApprovalService.deleteApproval(id).then( res => {
            this.setState({approvals: this.state.approvals.filter(approval => approval.approvalId !== id)});
        });
    }
    viewApproval(id){
        this.props.history.push(`/view-approval/${id}`);
    }
    editApproval(id){
        this.props.history.push(`/add-approval/${id}`);
    }

    componentDidMount(){
        ApprovalService.getApprovals().then((res) => {
            this.setState({ approvals: res.data});
        });
    }

    addApproval(){
        this.props.history.push('/add-approval/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Approval List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addApproval}> Add Approval</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ApproverComment </th>
                                    <th> ActionDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.approvals.map(
                                        approval => 
                                        <tr key = {approval.approvalId}>
                                             <td> { approval.approverComment } </td>
                                             <td> { approval.actionDate } </td>
                                             <td> { approval.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editApproval(approval.approvalId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteApproval(approval.approvalId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewApproval(approval.approvalId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListApprovalComponent
