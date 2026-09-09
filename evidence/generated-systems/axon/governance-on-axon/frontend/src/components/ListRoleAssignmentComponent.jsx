import React, { Component } from 'react'
import RoleAssignmentService from '../services/RoleAssignmentService'

class ListRoleAssignmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                roleAssignments: []
        }
        this.addRoleAssignment = this.addRoleAssignment.bind(this);
        this.editRoleAssignment = this.editRoleAssignment.bind(this);
        this.deleteRoleAssignment = this.deleteRoleAssignment.bind(this);
    }

    deleteRoleAssignment(id){
        RoleAssignmentService.deleteRoleAssignment(id).then( res => {
            this.setState({roleAssignments: this.state.roleAssignments.filter(roleAssignment => roleAssignment.roleAssignmentId !== id)});
        });
    }
    viewRoleAssignment(id){
        this.props.history.push(`/view-roleAssignment/${id}`);
    }
    editRoleAssignment(id){
        this.props.history.push(`/add-roleAssignment/${id}`);
    }

    componentDidMount(){
        RoleAssignmentService.getRoleAssignments().then((res) => {
            this.setState({ roleAssignments: res.data});
        });
    }

    addRoleAssignment(){
        this.props.history.push('/add-roleAssignment/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">RoleAssignment List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addRoleAssignment}> Add RoleAssignment</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> EffectiveFrom </th>
                                    <th> EffectiveTo </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.roleAssignments.map(
                                        roleAssignment => 
                                        <tr key = {roleAssignment.roleAssignmentId}>
                                             <td> { roleAssignment.effectiveFrom } </td>
                                             <td> { roleAssignment.effectiveTo } </td>
                                             <td>
                                                 <button onClick={ () => this.editRoleAssignment(roleAssignment.roleAssignmentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteRoleAssignment(roleAssignment.roleAssignmentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewRoleAssignment(roleAssignment.roleAssignmentId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListRoleAssignmentComponent
