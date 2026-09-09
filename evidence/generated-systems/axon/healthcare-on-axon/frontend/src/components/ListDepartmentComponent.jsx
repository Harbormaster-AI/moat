import React, { Component } from 'react'
import DepartmentService from '../services/DepartmentService'

class ListDepartmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                departments: []
        }
        this.addDepartment = this.addDepartment.bind(this);
        this.editDepartment = this.editDepartment.bind(this);
        this.deleteDepartment = this.deleteDepartment.bind(this);
    }

    deleteDepartment(id){
        DepartmentService.deleteDepartment(id).then( res => {
            this.setState({departments: this.state.departments.filter(department => department.departmentId !== id)});
        });
    }
    viewDepartment(id){
        this.props.history.push(`/view-department/${id}`);
    }
    editDepartment(id){
        this.props.history.push(`/add-department/${id}`);
    }

    componentDidMount(){
        DepartmentService.getDepartments().then((res) => {
            this.setState({ departments: res.data});
        });
    }

    addDepartment(){
        this.props.history.push('/add-department/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Department List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDepartment}> Add Department</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> DepartmentType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.departments.map(
                                        department => 
                                        <tr key = {department.departmentId}>
                                             <td> { department.name } </td>
                                             <td> { department.departmentType } </td>
                                             <td>
                                                 <button onClick={ () => this.editDepartment(department.departmentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDepartment(department.departmentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDepartment(department.departmentId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListDepartmentComponent
