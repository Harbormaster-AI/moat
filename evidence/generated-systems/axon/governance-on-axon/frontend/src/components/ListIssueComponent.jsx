import React, { Component } from 'react'
import IssueService from '../services/IssueService'

class ListIssueComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                issues: []
        }
        this.addIssue = this.addIssue.bind(this);
        this.editIssue = this.editIssue.bind(this);
        this.deleteIssue = this.deleteIssue.bind(this);
    }

    deleteIssue(id){
        IssueService.deleteIssue(id).then( res => {
            this.setState({issues: this.state.issues.filter(issue => issue.issueId !== id)});
        });
    }
    viewIssue(id){
        this.props.history.push(`/view-issue/${id}`);
    }
    editIssue(id){
        this.props.history.push(`/add-issue/${id}`);
    }

    componentDidMount(){
        IssueService.getIssues().then((res) => {
            this.setState({ issues: res.data});
        });
    }

    addIssue(){
        this.props.history.push('/add-issue/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Issue List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addIssue}> Add Issue</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Title </th>
                                    <th> OpenedDate </th>
                                    <th> ClosedDate </th>
                                    <th> IssueType </th>
                                    <th> Priority </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.issues.map(
                                        issue => 
                                        <tr key = {issue.issueId}>
                                             <td> { issue.title } </td>
                                             <td> { issue.openedDate } </td>
                                             <td> { issue.closedDate } </td>
                                             <td> { issue.issueType } </td>
                                             <td> { issue.priority } </td>
                                             <td> { issue.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editIssue(issue.issueId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteIssue(issue.issueId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewIssue(issue.issueId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListIssueComponent
