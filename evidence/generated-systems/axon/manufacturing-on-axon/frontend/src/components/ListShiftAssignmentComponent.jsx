import React, { Component } from 'react'
import ShiftAssignmentService from '../services/ShiftAssignmentService'

class ListShiftAssignmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                shiftAssignments: []
        }
        this.addShiftAssignment = this.addShiftAssignment.bind(this);
        this.editShiftAssignment = this.editShiftAssignment.bind(this);
        this.deleteShiftAssignment = this.deleteShiftAssignment.bind(this);
    }

    deleteShiftAssignment(id){
        ShiftAssignmentService.deleteShiftAssignment(id).then( res => {
            this.setState({shiftAssignments: this.state.shiftAssignments.filter(shiftAssignment => shiftAssignment.shiftAssignmentId !== id)});
        });
    }
    viewShiftAssignment(id){
        this.props.history.push(`/view-shiftAssignment/${id}`);
    }
    editShiftAssignment(id){
        this.props.history.push(`/add-shiftAssignment/${id}`);
    }

    componentDidMount(){
        ShiftAssignmentService.getShiftAssignments().then((res) => {
            this.setState({ shiftAssignments: res.data});
        });
    }

    addShiftAssignment(){
        this.props.history.push('/add-shiftAssignment/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ShiftAssignment List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addShiftAssignment}> Add ShiftAssignment</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> AssignmentDate </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.shiftAssignments.map(
                                        shiftAssignment => 
                                        <tr key = {shiftAssignment.shiftAssignmentId}>
                                             <td> { shiftAssignment.assignmentDate } </td>
                                             <td>
                                                 <button onClick={ () => this.editShiftAssignment(shiftAssignment.shiftAssignmentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteShiftAssignment(shiftAssignment.shiftAssignmentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewShiftAssignment(shiftAssignment.shiftAssignmentId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListShiftAssignmentComponent
