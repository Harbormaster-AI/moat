import React, { Component } from 'react'
import AuditEngagementService from '../services/AuditEngagementService'

class ViewAuditEngagementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            auditEngagement: {}
        }
    }

    componentDidMount(){
        AuditEngagementService.getAuditEngagementById(this.state.id).then( res => {
            this.setState({auditEngagement: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View AuditEngagement Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> title:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.auditEngagement.title }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> startDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.auditEngagement.startDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> endDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.auditEngagement.endDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.auditEngagement.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAuditEngagementComponent
