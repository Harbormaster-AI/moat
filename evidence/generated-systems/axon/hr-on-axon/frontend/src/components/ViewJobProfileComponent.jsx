import React, { Component } from 'react'
import JobProfileService from '../services/JobProfileService'

class ViewJobProfileComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            jobProfile: {}
        }
    }

    componentDidMount(){
        JobProfileService.getJobProfileById(this.state.id).then( res => {
            this.setState({jobProfile: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View JobProfile Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> title:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.jobProfile.title }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> jobCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.jobProfile.jobCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> JobLevel:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.jobProfile.jobLevel }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ExemptStatus:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.jobProfile.exemptStatus }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewJobProfileComponent
