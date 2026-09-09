import React, { Component } from 'react'
import LeavePolicyService from '../services/LeavePolicyService'

class ListLeavePolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                leavePolicys: []
        }
        this.addLeavePolicy = this.addLeavePolicy.bind(this);
        this.editLeavePolicy = this.editLeavePolicy.bind(this);
        this.deleteLeavePolicy = this.deleteLeavePolicy.bind(this);
    }

    deleteLeavePolicy(id){
        LeavePolicyService.deleteLeavePolicy(id).then( res => {
            this.setState({leavePolicys: this.state.leavePolicys.filter(leavePolicy => leavePolicy.leavePolicyId !== id)});
        });
    }
    viewLeavePolicy(id){
        this.props.history.push(`/view-leavePolicy/${id}`);
    }
    editLeavePolicy(id){
        this.props.history.push(`/add-leavePolicy/${id}`);
    }

    componentDidMount(){
        LeavePolicyService.getLeavePolicys().then((res) => {
            this.setState({ leavePolicys: res.data});
        });
    }

    addLeavePolicy(){
        this.props.history.push('/add-leavePolicy/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">LeavePolicy List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addLeavePolicy}> Add LeavePolicy</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> AccrualRate </th>
                                    <th> CarryoverAllowed </th>
                                    <th> MaxBalance </th>
                                    <th> LeaveCategory </th>
                                    <th> AccrualUnit </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.leavePolicys.map(
                                        leavePolicy => 
                                        <tr key = {leavePolicy.leavePolicyId}>
                                             <td> { leavePolicy.name } </td>
                                             <td> { leavePolicy.accrualRate } </td>
                                             <td> { leavePolicy.carryoverAllowed } </td>
                                             <td> { leavePolicy.maxBalance } </td>
                                             <td> { leavePolicy.leaveCategory } </td>
                                             <td> { leavePolicy.accrualUnit } </td>
                                             <td>
                                                 <button onClick={ () => this.editLeavePolicy(leavePolicy.leavePolicyId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteLeavePolicy(leavePolicy.leavePolicyId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewLeavePolicy(leavePolicy.leavePolicyId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListLeavePolicyComponent
