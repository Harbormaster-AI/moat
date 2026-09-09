import React, { Component } from 'react'
import InsurancePlanService from '../services/InsurancePlanService';

class UpdateInsurancePlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                planCode: '',
                planType: ''
        }
        this.updateInsurancePlan = this.updateInsurancePlan.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeplanCodeHandler = this.changeplanCodeHandler.bind(this);
        this.changePlanTypeHandler = this.changePlanTypeHandler.bind(this);
    }

    componentDidMount(){
        InsurancePlanService.getInsurancePlanById(this.state.id).then( (res) =>{
            let insurancePlan = res.data;
            this.setState({
                name: insurancePlan.name,
                planCode: insurancePlan.planCode,
                planType: insurancePlan.planType
            });
        });
    }

    updateInsurancePlan = (e) => {
        e.preventDefault();
        let insurancePlan = {
            insurancePlanId: this.state.id,
            name: this.state.name,
            planCode: this.state.planCode,
            planType: this.state.planType
        };
        console.log('insurancePlan => ' + JSON.stringify(insurancePlan));
        console.log('id => ' + JSON.stringify(this.state.id));
        InsurancePlanService.updateInsurancePlan(insurancePlan).then( res => {
            this.props.history.push('/insurancePlans');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update InsurancePlan</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> planCode: </label>
                                                <input placeholder="planCode" name="planCode" className="form-control" value={this.state.planCode} onChange={this.changeplanCodeHandler}/>

                                            <label> PlanType: </label>
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
                                        <button className="btn btn-success" onClick={this.updateInsurancePlan}>Save</button>
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

export default UpdateInsurancePlanComponent
