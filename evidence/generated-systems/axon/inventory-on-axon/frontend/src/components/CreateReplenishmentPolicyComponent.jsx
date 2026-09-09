import React, { Component } from 'react'
import ReplenishmentPolicyService from '../services/ReplenishmentPolicyService';

class CreateReplenishmentPolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                minLevel: '',
                maxLevel: '',
                reorderPoint: '',
                reorderQuantity: '',
                leadTimeDays: '',
                reviewPeriodDays: '',
                policyType: ''
        }
        this.changeminLevelHandler = this.changeminLevelHandler.bind(this);
        this.changemaxLevelHandler = this.changemaxLevelHandler.bind(this);
        this.changereorderPointHandler = this.changereorderPointHandler.bind(this);
        this.changereorderQuantityHandler = this.changereorderQuantityHandler.bind(this);
        this.changeleadTimeDaysHandler = this.changeleadTimeDaysHandler.bind(this);
        this.changereviewPeriodDaysHandler = this.changereviewPeriodDaysHandler.bind(this);
        this.changePolicyTypeHandler = this.changePolicyTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ReplenishmentPolicyService.getReplenishmentPolicyById(this.state.id).then( (res) =>{
                let replenishmentPolicy = res.data;
                this.setState({
                    minLevel: replenishmentPolicy.minLevel,
                    maxLevel: replenishmentPolicy.maxLevel,
                    reorderPoint: replenishmentPolicy.reorderPoint,
                    reorderQuantity: replenishmentPolicy.reorderQuantity,
                    leadTimeDays: replenishmentPolicy.leadTimeDays,
                    reviewPeriodDays: replenishmentPolicy.reviewPeriodDays,
                    policyType: replenishmentPolicy.policyType
                });
            });
        }        
    }
    saveOrUpdateReplenishmentPolicy = (e) => {
        e.preventDefault();
        let replenishmentPolicy = {
                replenishmentPolicyId: this.state.id,
                minLevel: this.state.minLevel,
                maxLevel: this.state.maxLevel,
                reorderPoint: this.state.reorderPoint,
                reorderQuantity: this.state.reorderQuantity,
                leadTimeDays: this.state.leadTimeDays,
                reviewPeriodDays: this.state.reviewPeriodDays,
                policyType: this.state.policyType
            };
        console.log('replenishmentPolicy => ' + JSON.stringify(replenishmentPolicy));

        // step 5
        if(this.state.id === '_add'){
            replenishmentPolicy.replenishmentPolicyId=''
            ReplenishmentPolicyService.createReplenishmentPolicy(replenishmentPolicy).then(res =>{
                this.props.history.push('/replenishmentPolicys');
            });
        }else{
            ReplenishmentPolicyService.updateReplenishmentPolicy(replenishmentPolicy).then( res => {
                this.props.history.push('/replenishmentPolicys');
            });
        }
    }
    
    changeminLevelHandler= (event) => {
        this.setState({minLevel: event.target.value});
    }
    changemaxLevelHandler= (event) => {
        this.setState({maxLevel: event.target.value});
    }
    changereorderPointHandler= (event) => {
        this.setState({reorderPoint: event.target.value});
    }
    changereorderQuantityHandler= (event) => {
        this.setState({reorderQuantity: event.target.value});
    }
    changeleadTimeDaysHandler= (event) => {
        this.setState({leadTimeDays: event.target.value});
    }
    changereviewPeriodDaysHandler= (event) => {
        this.setState({reviewPeriodDays: event.target.value});
    }
    changePolicyTypeHandler= (event) => {
        this.setState({policyType: event.target.value});
    }

    cancel(){
        this.props.history.push('/replenishmentPolicys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ReplenishmentPolicy</h3>
        }else{
            return <h3 className="text-center">Update ReplenishmentPolicy</h3>
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
                                            <label> minLevel:&emsp; </label>
                                                <input placeholder="minLevel" name="minLevel" className="form-control" value={this.state.minLevel} onChange={this.changeminLevelHandler}/>

                                            <label> maxLevel:&emsp; </label>
                                                <input placeholder="maxLevel" name="maxLevel" className="form-control" value={this.state.maxLevel} onChange={this.changemaxLevelHandler}/>

                                            <label> reorderPoint:&emsp; </label>
                                                <input placeholder="reorderPoint" name="reorderPoint" className="form-control" value={this.state.reorderPoint} onChange={this.changereorderPointHandler}/>

                                            <label> reorderQuantity:&emsp; </label>
                                                <input placeholder="reorderQuantity" name="reorderQuantity" className="form-control" value={this.state.reorderQuantity} onChange={this.changereorderQuantityHandler}/>

                                            <label> leadTimeDays:&emsp; </label>
                                                <input type="number" placeholder="leadTimeDays" name="leadTimeDays" className="form-control" value={this.state.leadTimeDays} onChange={this.changeleadTimeDaysHandler}/>

                                            <label> reviewPeriodDays:&emsp; </label>
                                                <input type="number" placeholder="reviewPeriodDays" name="reviewPeriodDays" className="form-control" value={this.state.reviewPeriodDays} onChange={this.changereviewPeriodDaysHandler}/>

                                            <label> PolicyType:&emsp; </label>
                                                <select value={this.state.policyType} onChange={this.changePolicyTypeHandler}>
                      <option name="PolicyType" className="form-control" >
                          MinMax
                      </option>
                      <option name="PolicyType" className="form-control" >
                          ReorderPoint
                      </option>
                      <option name="PolicyType" className="form-control" >
                          EOQ
                      </option>
                      <option name="PolicyType" className="form-control" >
                          Kanban
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateReplenishmentPolicy}>Save</button>
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

export default CreateReplenishmentPolicyComponent
