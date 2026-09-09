import React, { Component } from 'react'
import LeaveRequestService from '../services/LeaveRequestService'

class ViewLeaveRequestComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            leaveRequest: {}
        }
    }

    componentDidMount(){
        LeaveRequestService.getLeaveRequestById(this.state.id).then( res => {
            this.setState({leaveRequest: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View LeaveRequest Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> requestNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.leaveRequest.requestNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> startDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.leaveRequest.startDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> endDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.leaveRequest.endDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> reason:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.leaveRequest.reason }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> hours:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.leaveRequest.hours }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.leaveRequest.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewLeaveRequestComponent
