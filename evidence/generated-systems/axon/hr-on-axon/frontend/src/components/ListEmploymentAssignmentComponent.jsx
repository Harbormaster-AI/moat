import React, { Component } from 'react'
import EmploymentAssignmentService from '../services/EmploymentAssignmentService'

class ListEmploymentAssignmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                employmentAssignments: []
        }
        this.addEmploymentAssignment = this.addEmploymentAssignment.bind(this);
        this.editEmploymentAssignment = this.editEmploymentAssignment.bind(this);
        this.deleteEmploymentAssignment = this.deleteEmploymentAssignment.bind(this);
    }

    deleteEmploymentAssignment(id){
        EmploymentAssignmentService.deleteEmploymentAssignment(id).then( res => {
            this.setState({employmentAssignments: this.state.employmentAssignments.filter(employmentAssignment => employmentAssignment.employmentAssignmentId !== id)});
        });
    }
    viewEmploymentAssignment(id){
        this.props.history.push(`/view-employmentAssignment/${id}`);
    }
    editEmploymentAssignment(id){
        this.props.history.push(`/add-employmentAssignment/${id}`);
    }

    componentDidMount(){
        EmploymentAssignmentService.getEmploymentAssignments().then((res) => {
            this.setState({ employmentAssignments: res.data});
        });
    }

    addEmploymentAssignment(){
        this.props.history.push('/add-employmentAssignment/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">EmploymentAssignment List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addEmploymentAssignment}> Add EmploymentAssignment</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> StartDate </th>
                                    <th> EndDate </th>
                                    <th> Primary </th>
                                    <th> AssignmentType </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.employmentAssignments.map(
                                        employmentAssignment => 
                                        <tr key = {employmentAssignment.employmentAssignmentId}>
                                             <td> { employmentAssignment.startDate } </td>
                                             <td> { employmentAssignment.endDate } </td>
                                             <td> { employmentAssignment.primary } </td>
                                             <td> { employmentAssignment.assignmentType } </td>
                                             <td> { employmentAssignment.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editEmploymentAssignment(employmentAssignment.employmentAssignmentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteEmploymentAssignment(employmentAssignment.employmentAssignmentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewEmploymentAssignment(employmentAssignment.employmentAssignmentId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListEmploymentAssignmentComponent
