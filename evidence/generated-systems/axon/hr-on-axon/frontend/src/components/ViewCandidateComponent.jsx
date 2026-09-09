import React, { Component } from 'react'
import CandidateService from '../services/CandidateService'

class ViewCandidateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            candidate: {}
        }
    }

    componentDidMount(){
        CandidateService.getCandidateById(this.state.id).then( res => {
            this.setState({candidate: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Candidate Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.candidate.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> email:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.candidate.email }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> phone:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.candidate.phone }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Source:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.candidate.source }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCandidateComponent
