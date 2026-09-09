import React, { Component } from 'react'
import ThirdPartyAssessmentService from '../services/ThirdPartyAssessmentService'

class ViewThirdPartyAssessmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            thirdPartyAssessment: {}
        }
    }

    componentDidMount(){
        ThirdPartyAssessmentService.getThirdPartyAssessmentById(this.state.id).then( res => {
            this.setState({thirdPartyAssessment: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ThirdPartyAssessment Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> assessmentDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.thirdPartyAssessment.assessmentDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> assessor:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.thirdPartyAssessment.assessor }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> AssessmentType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.thirdPartyAssessment.assessmentType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Result:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.thirdPartyAssessment.result }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewThirdPartyAssessmentComponent
