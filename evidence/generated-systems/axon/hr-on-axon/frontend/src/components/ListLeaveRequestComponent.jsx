import React, { Component } from 'react'
import LeaveRequestService from '../services/LeaveRequestService'

class ListLeaveRequestComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                leaveRequests: []
        }
        this.addLeaveRequest = this.addLeaveRequest.bind(this);
        this.editLeaveRequest = this.editLeaveRequest.bind(this);
        this.deleteLeaveRequest = this.deleteLeaveRequest.bind(this);
    }

    deleteLeaveRequest(id){
        LeaveRequestService.deleteLeaveRequest(id).then( res => {
            this.setState({leaveRequests: this.state.leaveRequests.filter(leaveRequest => leaveRequest.leaveRequestId !== id)});
        });
    }
    viewLeaveRequest(id){
        this.props.history.push(`/view-leaveRequest/${id}`);
    }
    editLeaveRequest(id){
        this.props.history.push(`/add-leaveRequest/${id}`);
    }

    componentDidMount(){
        LeaveRequestService.getLeaveRequests().then((res) => {
            this.setState({ leaveRequests: res.data});
        });
    }

    addLeaveRequest(){
        this.props.history.push('/add-leaveRequest/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">LeaveRequest List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addLeaveRequest}> Add LeaveRequest</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> RequestNumber </th>
                                    <th> StartDate </th>
                                    <th> EndDate </th>
                                    <th> Reason </th>
                                    <th> Hours </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.leaveRequests.map(
                                        leaveRequest => 
                                        <tr key = {leaveRequest.leaveRequestId}>
                                             <td> { leaveRequest.requestNumber } </td>
                                             <td> { leaveRequest.startDate } </td>
                                             <td> { leaveRequest.endDate } </td>
                                             <td> { leaveRequest.reason } </td>
                                             <td> { leaveRequest.hours } </td>
                                             <td> { leaveRequest.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editLeaveRequest(leaveRequest.leaveRequestId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteLeaveRequest(leaveRequest.leaveRequestId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewLeaveRequest(leaveRequest.leaveRequestId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListLeaveRequestComponent
