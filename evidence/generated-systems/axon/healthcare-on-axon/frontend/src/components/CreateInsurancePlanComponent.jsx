import React, { Component } from 'react'
import InsurancePlanService from '../services/InsurancePlanService';

class CreateInsurancePlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                planCode: '',
                planType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeplanCodeHandler = this.changeplanCodeHandler.bind(this);
        this.changePlanTypeHandler = this.changePlanTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            InsurancePlanService.getInsurancePlanById(this.state.id).then( (res) =>{
                let insurancePlan = res.data;
                this.setState({
                    name: insurancePlan.name,
                    planCode: insurancePlan.planCode,
                    planType: insurancePlan.planType
                });
            });
        }        
    }
    saveOrUpdateInsurancePlan = (e) => {
        e.preventDefault();
        let insurancePlan = {
                insurancePlanId: this.state.id,
                name: this.state.name,
                planCode: this.state.planCode,
                planType: this.state.planType
            };
        console.log('insurancePlan => ' + JSON.stringify(insurancePlan));

        // step 5
        if(this.state.id === '_add'){
            insurancePlan.insurancePlanId=''
            InsurancePlanService.createInsurancePlan(insurancePlan).then(res =>{
                this.props.history.push('/insurancePlans');
            });
        }else{
            InsurancePlanService.updateInsurancePlan(insurancePlan).then( res => {
                this.props.history.push('/insurancePlans');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeplanCodeHandler= (event) => {
        this.setState({planCode: event.target.value});
    }
    changePlanTypeHandler= (event) => {
        this.setState({planType: event.target.value});
    }

    cancel(){
        this.props.history.push('/insurancePlans');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add InsurancePlan</h3>
        }else{
            return <h3 className="text-center">Update InsurancePlan</h3>
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

                                            <label> planCode:&emsp; </label>
                                                <input placeholder="planCode" name="planCode" className="form-control" value={this.state.planCode} onChange={this.changeplanCodeHandler}/>

                                            <label> PlanType:&emsp; </label>
                                                <select value={this.state.planType} onChange={this.changePlanTypeHandler}>
                      <option name="PlanType" className="form-control" >
                          HMO
                      </option>
                      <option name="PlanType" className="form-control" >
                          PPO
                      </option>
                      <option name="PlanType" className="form-control" >
                          EPO
                      </option>
                      <option name="PlanType" className="form-control" >
                          POS
                      </option>
                      <option name="PlanType" className="form-control" >
                          Indemnity
                      </option>
                      <option name="PlanType" className="form-control" >
                          MedicareAdvantage
                      </option>
                      <option name="PlanType" className="form-control" >
                          MedicaidManagedCare
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateInsurancePlan}>Save</button>
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

export default CreateInsurancePlanComponent
