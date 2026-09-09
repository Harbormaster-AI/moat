import React, { Component } from 'react'
import PricingPlanService from '../services/PricingPlanService';

class CreatePricingPlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                planCode: '',
                baseCurrency: '',
                status: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeplanCodeHandler = this.changeplanCodeHandler.bind(this);
        this.changebaseCurrencyHandler = this.changebaseCurrencyHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PricingPlanService.getPricingPlanById(this.state.id).then( (res) =>{
                let pricingPlan = res.data;
                this.setState({
                    name: pricingPlan.name,
                    planCode: pricingPlan.planCode,
                    baseCurrency: pricingPlan.baseCurrency,
                    status: pricingPlan.status
                });
            });
        }        
    }
    saveOrUpdatePricingPlan = (e) => {
        e.preventDefault();
        let pricingPlan = {
                pricingPlanId: this.state.id,
                name: this.state.name,
                planCode: this.state.planCode,
                baseCurrency: this.state.baseCurrency,
                status: this.state.status
            };
        console.log('pricingPlan => ' + JSON.stringify(pricingPlan));

        // step 5
        if(this.state.id === '_add'){
            pricingPlan.pricingPlanId=''
            PricingPlanService.createPricingPlan(pricingPlan).then(res =>{
                this.props.history.push('/pricingPlans');
            });
        }else{
            PricingPlanService.updatePricingPlan(pricingPlan).then( res => {
                this.props.history.push('/pricingPlans');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeplanCodeHandler= (event) => {
        this.setState({planCode: event.target.value});
    }
    changebaseCurrencyHandler= (event) => {
        this.setState({baseCurrency: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/pricingPlans');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add PricingPlan</h3>
        }else{
            return <h3 className="text-center">Update PricingPlan</h3>
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

                                            <label> baseCurrency:&emsp; </label>
                                                <input placeholder="baseCurrency" name="baseCurrency" className="form-control" value={this.state.baseCurrency} onChange={this.changebaseCurrencyHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Suspended
                      </option>
                      <option name="Status" className="form-control" >
                          Archived
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePricingPlan}>Save</button>
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

export default CreatePricingPlanComponent
