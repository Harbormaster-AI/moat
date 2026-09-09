import React, { Component } from 'react'
import DataSubjectRequestService from '../services/DataSubjectRequestService'

class ViewDataSubjectRequestComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            dataSubjectRequest: {}
        }
    }

    componentDidMount(){
        DataSubjectRequestService.getDataSubjectRequestById(this.state.id).then( res => {
            this.setState({dataSubjectRequest: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View DataSubjectRequest Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> receivedDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataSubjectRequest.receivedDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> dueDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataSubjectRequest.dueDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> requesterCountry:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataSubjectRequest.requesterCountry }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> RequestType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataSubjectRequest.requestType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataSubjectRequest.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewDataSubjectRequestComponent
