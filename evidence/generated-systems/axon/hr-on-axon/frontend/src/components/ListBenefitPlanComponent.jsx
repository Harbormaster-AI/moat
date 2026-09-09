import React, { Component } from 'react'
import BenefitPlanService from '../services/BenefitPlanService'

class ListBenefitPlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                benefitPlans: []
        }
        this.addBenefitPlan = this.addBenefitPlan.bind(this);
        this.editBenefitPlan = this.editBenefitPlan.bind(this);
        this.deleteBenefitPlan = this.deleteBenefitPlan.bind(this);
    }

    deleteBenefitPlan(id){
        BenefitPlanService.deleteBenefitPlan(id).then( res => {
            this.setState({benefitPlans: this.state.benefitPlans.filter(benefitPlan => benefitPlan.benefitPlanId !== id)});
        });
    }
    viewBenefitPlan(id){
        this.props.history.push(`/view-benefitPlan/${id}`);
    }
    editBenefitPlan(id){
        this.props.history.push(`/add-benefitPlan/${id}`);
    }

    componentDidMount(){
        BenefitPlanService.getBenefitPlans().then((res) => {
            this.setState({ benefitPlans: res.data});
        });
    }

    addBenefitPlan(){
        this.props.history.push('/add-benefitPlan/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">BenefitPlan List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addBenefitPlan}> Add BenefitPlan</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> ProviderName </th>
                                    <th> EmployeeContributionRate </th>
                                    <th> EmployerContributionRate </th>
                                    <th> EligibilityRules </th>
                                    <th> BenefitType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.benefitPlans.map(
                                        benefitPlan => 
                                        <tr key = {benefitPlan.benefitPlanId}>
                                             <td> { benefitPlan.name } </td>
                                             <td> { benefitPlan.providerName } </td>
                                             <td> { benefitPlan.employeeContributionRate } </td>
                                             <td> { benefitPlan.employerContributionRate } </td>
                                             <td> { benefitPlan.eligibilityRules } </td>
                                             <td> { benefitPlan.benefitType } </td>
                                             <td>
                                                 <button onClick={ () => this.editBenefitPlan(benefitPlan.benefitPlanId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteBenefitPlan(benefitPlan.benefitPlanId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewBenefitPlan(benefitPlan.benefitPlanId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListBenefitPlanComponent
