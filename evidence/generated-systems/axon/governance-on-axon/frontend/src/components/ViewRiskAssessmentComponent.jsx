import React, { Component } from 'react'
import RiskAssessmentService from '../services/RiskAssessmentService'

class ViewRiskAssessmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            riskAssessment: {}
        }
    }

    componentDidMount(){
        RiskAssessmentService.getRiskAssessmentById(this.state.id).then( res => {
            this.setState({riskAssessment: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View RiskAssessment Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> assessmentDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.riskAssessment.assessmentDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> assessor:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.riskAssessment.assessor }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> summary:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.riskAssessment.summary }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> AssessmentType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.riskAssessment.assessmentType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewRiskAssessmentComponent
