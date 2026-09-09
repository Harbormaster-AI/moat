import React, { Component } from 'react'
import CreativeApprovalService from '../services/CreativeApprovalService';

class CreateCreativeApprovalComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                reviewer: '',
                reviewedAt: '',
                status: ''
        }
        this.changereviewerHandler = this.changereviewerHandler.bind(this);
        this.changereviewedAtHandler = this.changereviewedAtHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            CreativeApprovalService.getCreativeApprovalById(this.state.id).then( (res) =>{
                let creativeApproval = res.data;
                this.setState({
                    reviewer: creativeApproval.reviewer,
                    reviewedAt: creativeApproval.reviewedAt,
                    status: creativeApproval.status
                });
            });
        }        
    }
    saveOrUpdateCreativeApproval = (e) => {
        e.preventDefault();
        let creativeApproval = {
                creativeApprovalId: this.state.id,
                reviewer: this.state.reviewer,
                reviewedAt: this.state.reviewedAt,
                status: this.state.status
            };
        console.log('creativeApproval => ' + JSON.stringify(creativeApproval));

        // step 5
        if(this.state.id === '_add'){
            creativeApproval.creativeApprovalId=''
            CreativeApprovalService.createCreativeApproval(creativeApproval).then(res =>{
                this.props.history.push('/creativeApprovals');
            });
        }else{
            CreativeApprovalService.updateCreativeApproval(creativeApproval).then( res => {
                this.props.history.push('/creativeApprovals');
            });
        }
    }
    
    changereviewerHandler= (event) => {
        this.setState({reviewer: event.target.value});
    }
    changereviewedAtHandler= (event) => {
        this.setState({reviewedAt: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/creativeApprovals');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add CreativeApproval</h3>
        }else{
            return <h3 className="text-center">Update CreativeApproval</h3>
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
                                            <label> reviewer:&emsp; </label>
                                                <input placeholder="reviewer" name="reviewer" className="form-control" value={this.state.reviewer} onChange={this.changereviewerHandler}/>

                                            <label> reviewedAt:&emsp; </label>
                                                <input type="date" placeholder="reviewedAt" name="reviewedAt" className="form-control" value={this.state.reviewedAt} onChange={this.changereviewedAtHandler}/>

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
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCreativeApproval}>Save</button>
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

export default CreateCreativeApprovalComponent
