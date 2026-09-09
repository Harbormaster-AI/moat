import React, { Component } from 'react'
import ApprovalService from '../services/ApprovalService'

class ViewApprovalComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            approval: {}
        }
    }

    componentDidMount(){
        ApprovalService.getApprovalById(this.state.id).then( res => {
            this.setState({approval: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Approval Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> approverComment:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.approval.approverComment }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> actionDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.approval.actionDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.approval.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewApprovalComponent
