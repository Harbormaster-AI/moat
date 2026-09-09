import React, { Component } from 'react'
import ApplicationService from '../services/ApplicationService'

class ViewApplicationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            application: {}
        }
    }

    componentDidMount(){
        ApplicationService.getApplicationById(this.state.id).then( res => {
            this.setState({application: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Application Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> applicationNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.application.applicationNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> submissionDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.application.submissionDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.application.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewApplicationComponent
