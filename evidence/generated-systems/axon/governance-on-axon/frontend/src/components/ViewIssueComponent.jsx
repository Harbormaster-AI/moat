import React, { Component } from 'react'
import IssueService from '../services/IssueService'

class ViewIssueComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            issue: {}
        }
    }

    componentDidMount(){
        IssueService.getIssueById(this.state.id).then( res => {
            this.setState({issue: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Issue Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> title:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.issue.title }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> openedDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.issue.openedDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> closedDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.issue.closedDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> IssueType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.issue.issueType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Priority:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.issue.priority }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.issue.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewIssueComponent
