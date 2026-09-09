import React, { Component } from 'react'
import JobApplicationService from '../services/JobApplicationService'

class ViewJobApplicationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            jobApplication: {}
        }
    }

    componentDidMount(){
        JobApplicationService.getJobApplicationById(this.state.id).then( res => {
            this.setState({jobApplication: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View JobApplication Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> applicationNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.jobApplication.applicationNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> appliedDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.jobApplication.appliedDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> resumeUrl:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.jobApplication.resumeUrl }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.jobApplication.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewJobApplicationComponent
