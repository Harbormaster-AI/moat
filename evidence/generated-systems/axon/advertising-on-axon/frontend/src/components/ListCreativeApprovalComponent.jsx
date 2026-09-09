import React, { Component } from 'react'
import CreativeApprovalService from '../services/CreativeApprovalService'

class ListCreativeApprovalComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                creativeApprovals: []
        }
        this.addCreativeApproval = this.addCreativeApproval.bind(this);
        this.editCreativeApproval = this.editCreativeApproval.bind(this);
        this.deleteCreativeApproval = this.deleteCreativeApproval.bind(this);
    }

    deleteCreativeApproval(id){
        CreativeApprovalService.deleteCreativeApproval(id).then( res => {
            this.setState({creativeApprovals: this.state.creativeApprovals.filter(creativeApproval => creativeApproval.creativeApprovalId !== id)});
        });
    }
    viewCreativeApproval(id){
        this.props.history.push(`/view-creativeApproval/${id}`);
    }
    editCreativeApproval(id){
        this.props.history.push(`/add-creativeApproval/${id}`);
    }

    componentDidMount(){
        CreativeApprovalService.getCreativeApprovals().then((res) => {
            this.setState({ creativeApprovals: res.data});
        });
    }

    addCreativeApproval(){
        this.props.history.push('/add-creativeApproval/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CreativeApproval List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCreativeApproval}> Add CreativeApproval</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Reviewer </th>
                                    <th> ReviewedAt </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.creativeApprovals.map(
                                        creativeApproval => 
                                        <tr key = {creativeApproval.creativeApprovalId}>
                                             <td> { creativeApproval.reviewer } </td>
                                             <td> { creativeApproval.reviewedAt } </td>
                                             <td> { creativeApproval.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editCreativeApproval(creativeApproval.creativeApprovalId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCreativeApproval(creativeApproval.creativeApprovalId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCreativeApproval(creativeApproval.creativeApprovalId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCreativeApprovalComponent
