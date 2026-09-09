import React, { Component } from 'react'
import ApprovalService from '../services/ApprovalService';

class CreateApprovalComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                approverComment: '',
                actionDate: '',
                status: ''
        }
        this.changeapproverCommentHandler = this.changeapproverCommentHandler.bind(this);
        this.changeactionDateHandler = this.changeactionDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ApprovalService.getApprovalById(this.state.id).then( (res) =>{
                let approval = res.data;
                this.setState({
                    approverComment: approval.approverComment,
                    actionDate: approval.actionDate,
                    status: approval.status
                });
            });
        }        
    }
    saveOrUpdateApproval = (e) => {
        e.preventDefault();
        let approval = {
                approvalId: this.state.id,
                approverComment: this.state.approverComment,
                actionDate: this.state.actionDate,
                status: this.state.status
            };
        console.log('approval => ' + JSON.stringify(approval));

        // step 5
        if(this.state.id === '_add'){
            approval.approvalId=''
            ApprovalService.createApproval(approval).then(res =>{
                this.props.history.push('/approvals');
            });
        }else{
            ApprovalService.updateApproval(approval).then( res => {
                this.props.history.push('/approvals');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Approval</h3>
        }else{
            return <h3 className="text-center">Update Approval</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> approverComment:&emsp; </label>
                                                <input placeholder="approverComment" name="approverComment" className="form-control" value={this.state.approverComment} onChange={this.changeapproverCommentHandler}/>

                                            <label> actionDate:&emsp; </label>
                                                <input type="date" placeholder="actionDate" name="actionDate" className="form-control" value={this.state.actionDate} onChange={this.changeactionDateHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateApproval}>Save</button>
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

export default CreateApprovalComponent
