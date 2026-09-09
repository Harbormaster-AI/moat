import React, { Component } from 'react'
import InterviewService from '../services/InterviewService'

class ViewInterviewComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            interview: {}
        }
    }

    componentDidMount(){
        InterviewService.getInterviewById(this.state.id).then( res => {
            this.setState({interview: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Interview Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> interviewDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.interview.interviewDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> feedback:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.interview.feedback }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Stage:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.interview.stage }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Result:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.interview.result }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewInterviewComponent
