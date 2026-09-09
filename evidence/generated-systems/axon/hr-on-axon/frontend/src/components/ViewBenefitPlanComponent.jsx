import React, { Component } from 'react'
import BenefitPlanService from '../services/BenefitPlanService'

class ViewBenefitPlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            benefitPlan: {}
        }
    }

    componentDidMount(){
        BenefitPlanService.getBenefitPlanById(this.state.id).then( res => {
            this.setState({benefitPlan: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View BenefitPlan Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.benefitPlan.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> providerName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.benefitPlan.providerName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> employeeContributionRate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.benefitPlan.employeeContributionRate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> employerContributionRate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.benefitPlan.employerContributionRate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> eligibilityRules:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.benefitPlan.eligibilityRules }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> BenefitType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.benefitPlan.benefitType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewBenefitPlanComponent
