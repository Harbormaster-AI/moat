import React, { Component } from 'react'
import IssueService from '../services/IssueService';

class CreateIssueComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                title: '',
                openedDate: '',
                closedDate: '',
                issueType: '',
                priority: '',
                status: ''
        }
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changeopenedDateHandler = this.changeopenedDateHandler.bind(this);
        this.changeclosedDateHandler = this.changeclosedDateHandler.bind(this);
        this.changeIssueTypeHandler = this.changeIssueTypeHandler.bind(this);
        this.changePriorityHandler = this.changePriorityHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            IssueService.getIssueById(this.state.id).then( (res) =>{
                let issue = res.data;
                this.setState({
                    title: issue.title,
                    openedDate: issue.openedDate,
                    closedDate: issue.closedDate,
                    issueType: issue.issueType,
                    priority: issue.priority,
                    status: issue.status
                });
            });
        }        
    }
    saveOrUpdateIssue = (e) => {
        e.preventDefault();
        let issue = {
                issueId: this.state.id,
                title: this.state.title,
                openedDate: this.state.openedDate,
                closedDate: this.state.closedDate,
                issueType: this.state.issueType,
                priority: this.state.priority,
                status: this.state.status
            };
        console.log('issue => ' + JSON.stringify(issue));

        // step 5
        if(this.state.id === '_add'){
            issue.issueId=''
            IssueService.createIssue(issue).then(res =>{
                this.props.history.push('/issues');
            });
        }else{
            IssueService.updateIssue(issue).then( res => {
                this.props.history.push('/issues');
            });
        }
    }
    
    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changeopenedDateHandler= (event) => {
        this.setState({openedDate: event.target.value});
    }
    changeclosedDateHandler= (event) => {
        this.setState({closedDate: event.target.value});
    }
    changeIssueTypeHandler= (event) => {
        this.setState({issueType: event.target.value});
    }
    changePriorityHandler= (event) => {
        this.setState({priority: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/issues');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Issue</h3>
        }else{
            return <h3 className="text-center">Update Issue</h3>
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
                                            <label> title:&emsp; </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> openedDate:&emsp; </label>
                                                <input type="date" placeholder="openedDate" name="openedDate" className="form-control" value={this.state.openedDate} onChange={this.changeopenedDateHandler}/>

                                            <label> closedDate:&emsp; </label>
                                                <input type="date" placeholder="closedDate" name="closedDate" className="form-control" value={this.state.closedDate} onChange={this.changeclosedDateHandler}/>

                                            <label> IssueType:&emsp; </label>
                                                <select value={this.state.issueType} onChange={this.changeIssueTypeHandler}>
                      <option name="IssueType" className="form-control" >
                          ControlDeficiency
                      </option>
                      <option name="IssueType" className="form-control" >
                          ProcessGap
                      </option>
                      <option name="IssueType" className="form-control" >
                          ComplianceBreach
                      </option>
                      <option name="IssueType" className="form-control" >
                          SecurityIncident
                      </option>
                      <option name="IssueType" className="form-control" >
                          DataQualityIssue
                      </option>
                      <option name="IssueType" className="form-control" >
                          ThirdPartyIssue
                      </option>
                    </select>

                                            <label> Priority:&emsp; </label>
                                                <select value={this.state.priority} onChange={this.changePriorityHandler}>
                      <option name="Priority" className="form-control" >
                          Low
                      </option>
                      <option name="Priority" className="form-control" >
                          Medium
                      </option>
                      <option name="Priority" className="form-control" >
                          High
                      </option>
                      <option name="Priority" className="form-control" >
                          Urgent
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Open
                      </option>
                      <option name="Status" className="form-control" >
                          Investigating
                      </option>
                      <option name="Status" className="form-control" >
                          RemediationPlanned
                      </option>
                      <option name="Status" className="form-control" >
                          RemediationInProgress
                      </option>
                      <option name="Status" className="form-control" >
                          Verified
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateIssue}>Save</button>
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

export default CreateIssueComponent
