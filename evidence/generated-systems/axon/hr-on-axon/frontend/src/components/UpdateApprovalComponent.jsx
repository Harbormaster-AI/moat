import React, { Component } from 'react'
import ApprovalService from '../services/ApprovalService';

class UpdateApprovalComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                approverComment: '',
                actionDate: '',
                status: ''
        }
        this.updateApproval = this.updateApproval.bind(this);

        this.changeapproverCommentHandler = this.changeapproverCommentHandler.bind(this);
        this.changeactionDateHandler = this.changeactionDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ApprovalService.getApprovalById(this.state.id).then( (res) =>{
            let approval = res.data;
            this.setState({
                approverComment: approval.approverComment,
                actionDate: approval.actionDate,
                status: approval.status
            });
        });
    }

    updateApproval = (e) => {
        e.preventDefault();
        let approval = {
            approvalId: this.state.id,
            approverComment: this.state.approverComment,
            actionDate: this.state.actionDate,
            status: this.state.status
        };
        console.log('approval => ' + JSON.stringify(approval));
        console.log('id => ' + JSON.stringify(this.state.id));
        ApprovalService.updateApproval(approval).then( res => {
            this.props.history.push('/approvals');
        });
    }

    changeapproverCommentHandler= (event) => {
        this.setState({approverComment: event.target.value});
    }
    changeactionDateHandler= (event) => {
        this.setState({actionDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/approvals');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Approval</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> approverComment: </label>
                                                <input placeholder="approverComment" name="approverComment" className="form-control" value={this.state.approverComment} onChange={this.changeapproverCommentHandler}/>

                                            <label> actionDate: </label>
                                                <input type="date" placeholder="actionDate" name="actionDate" className="form-control" value={this.state.actionDate} onChange={this.changeactionDateHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Pending
                      </option>
                      <option name="Status" className="form-control" >
                          Approved
                      </option>
                      <option name="Status" className="form-control" >
                          Rejected
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateApproval}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateApprovalComponent
