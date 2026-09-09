import React, { Component } from 'react'
import BenefitEnrollmentService from '../services/BenefitEnrollmentService'

class ViewBenefitEnrollmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            benefitEnrollment: {}
        }
    }

    componentDidMount(){
        BenefitEnrollmentService.getBenefitEnrollmentById(this.state.id).then( res => {
            this.setState({benefitEnrollment: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View BenefitEnrollment Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> enrollmentId:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.benefitEnrollment.enrollmentId }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> effectiveFrom:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.benefitEnrollment.effectiveFrom }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> effectiveTo:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.benefitEnrollment.effectiveTo }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.benefitEnrollment.status }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> CoverageLevel:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.benefitEnrollment.coverageLevel }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewBenefitEnrollmentComponent
