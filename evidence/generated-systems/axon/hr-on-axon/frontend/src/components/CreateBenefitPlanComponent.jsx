import React, { Component } from 'react'
import BenefitPlanService from '../services/BenefitPlanService';

class CreateBenefitPlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                providerName: '',
                employeeContributionRate: '',
                employerContributionRate: '',
                eligibilityRules: '',
                benefitType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeproviderNameHandler = this.changeproviderNameHandler.bind(this);
        this.changeemployeeContributionRateHandler = this.changeemployeeContributionRateHandler.bind(this);
        this.changeemployerContributionRateHandler = this.changeemployerContributionRateHandler.bind(this);
        this.changeeligibilityRulesHandler = this.changeeligibilityRulesHandler.bind(this);
        this.changeBenefitTypeHandler = this.changeBenefitTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            BenefitPlanService.getBenefitPlanById(this.state.id).then( (res) =>{
                let benefitPlan = res.data;
                this.setState({
                    name: benefitPlan.name,
                    providerName: benefitPlan.providerName,
                    employeeContributionRate: benefitPlan.employeeContributionRate,
                    employerContributionRate: benefitPlan.employerContributionRate,
                    eligibilityRules: benefitPlan.eligibilityRules,
                    benefitType: benefitPlan.benefitType
                });
            });
        }        
    }
    saveOrUpdateBenefitPlan = (e) => {
        e.preventDefault();
        let benefitPlan = {
                benefitPlanId: this.state.id,
                name: this.state.name,
                providerName: this.state.providerName,
                employeeContributionRate: this.state.employeeContributionRate,
                employerContributionRate: this.state.employerContributionRate,
                eligibilityRules: this.state.eligibilityRules,
                benefitType: this.state.benefitType
            };
        console.log('benefitPlan => ' + JSON.stringify(benefitPlan));

        // step 5
        if(this.state.id === '_add'){
            benefitPlan.benefitPlanId=''
            BenefitPlanService.createBenefitPlan(benefitPlan).then(res =>{
                this.props.history.push('/benefitPlans');
            });
        }else{
            BenefitPlanService.updateBenefitPlan(benefitPlan).then( res => {
                this.props.history.push('/benefitPlans');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeproviderNameHandler= (event) => {
        this.setState({providerName: event.target.value});
    }
    changeemployeeContributionRateHandler= (event) => {
        this.setState({employeeContributionRate: event.target.value});
    }
    changeemployerContributionRateHandler= (event) => {
        this.setState({employerContributionRate: event.target.value});
    }
    changeeligibilityRulesHandler= (event) => {
        this.setState({eligibilityRules: event.target.value});
    }
    changeBenefitTypeHandler= (event) => {
        this.setState({benefitType: event.target.value});
    }

    cancel(){
        this.props.history.push('/benefitPlans');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add BenefitPlan</h3>
        }else{
            return <h3 className="text-center">Update BenefitPlan</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> providerName:&emsp; </label>
                                                <input placeholder="providerName" name="providerName" className="form-control" value={this.state.providerName} onChange={this.changeproviderNameHandler}/>

                                            <label> employeeContributionRate:&emsp; </label>
                                                <input placeholder="employeeContributionRate" name="employeeContributionRate" className="form-control" value={this.state.employeeContributionRate} onChange={this.changeemployeeContributionRateHandler}/>

                                            <label> employerContributionRate:&emsp; </label>
                                                <input placeholder="employerContributionRate" name="employerContributionRate" className="form-control" value={this.state.employerContributionRate} onChange={this.changeemployerContributionRateHandler}/>

                                            <label> eligibilityRules:&emsp; </label>
                                                <input placeholder="eligibilityRules" name="eligibilityRules" className="form-control" value={this.state.eligibilityRules} onChange={this.changeeligibilityRulesHandler}/>

                                            <label> BenefitType:&emsp; </label>
                                                <select value={this.state.benefitType} onChange={this.changeBenefitTypeHandler}>
                      <option name="BenefitType" className="form-control" >
                          Medical
                      </option>
                      <option name="BenefitType" className="form-control" >
                          Dental
                      </option>
                      <option name="BenefitType" className="form-control" >
                          Vision
                      </option>
                      <option name="BenefitType" className="form-control" >
                          LifeInsurance
                      </option>
                      <option name="BenefitType" className="form-control" >
                          Disability
                      </option>
                      <option name="BenefitType" className="form-control" >
                          Retirement
                      </option>
                      <option name="BenefitType" className="form-control" >
                          Wellness
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateBenefitPlan}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>
                   </div>
            </div>
        )
    }
}

export default CreateBenefitPlanComponent
