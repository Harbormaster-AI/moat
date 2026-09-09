import React, { Component } from 'react'
import LeavePolicyService from '../services/LeavePolicyService';

class CreateLeavePolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                accrualRate: '',
                carryoverAllowed: '',
                maxBalance: '',
                leaveCategory: '',
                accrualUnit: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeaccrualRateHandler = this.changeaccrualRateHandler.bind(this);
        this.changecarryoverAllowedHandler = this.changecarryoverAllowedHandler.bind(this);
        this.changemaxBalanceHandler = this.changemaxBalanceHandler.bind(this);
        this.changeLeaveCategoryHandler = this.changeLeaveCategoryHandler.bind(this);
        this.changeAccrualUnitHandler = this.changeAccrualUnitHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            LeavePolicyService.getLeavePolicyById(this.state.id).then( (res) =>{
                let leavePolicy = res.data;
                this.setState({
                    name: leavePolicy.name,
                    accrualRate: leavePolicy.accrualRate,
                    carryoverAllowed: leavePolicy.carryoverAllowed,
                    maxBalance: leavePolicy.maxBalance,
                    leaveCategory: leavePolicy.leaveCategory,
                    accrualUnit: leavePolicy.accrualUnit
                });
            });
        }        
    }
    saveOrUpdateLeavePolicy = (e) => {
        e.preventDefault();
        let leavePolicy = {
                leavePolicyId: this.state.id,
                name: this.state.name,
                accrualRate: this.state.accrualRate,
                carryoverAllowed: this.state.carryoverAllowed,
                maxBalance: this.state.maxBalance,
                leaveCategory: this.state.leaveCategory,
                accrualUnit: this.state.accrualUnit
            };
        console.log('leavePolicy => ' + JSON.stringify(leavePolicy));

        // step 5
        if(this.state.id === '_add'){
            leavePolicy.leavePolicyId=''
            LeavePolicyService.createLeavePolicy(leavePolicy).then(res =>{
                this.props.history.push('/leavePolicys');
            });
        }else{
            LeavePolicyService.updateLeavePolicy(leavePolicy).then( res => {
                this.props.history.push('/leavePolicys');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeaccrualRateHandler= (event) => {
        this.setState({accrualRate: event.target.value});
    }
    changecarryoverAllowedHandler= (event) => {
        this.setState({carryoverAllowed: event.target.value});
    }
    changemaxBalanceHandler= (event) => {
        this.setState({maxBalance: event.target.value});
    }
    changeLeaveCategoryHandler= (event) => {
        this.setState({leaveCategory: event.target.value});
    }
    changeAccrualUnitHandler= (event) => {
        this.setState({accrualUnit: event.target.value});
    }

    cancel(){
        this.props.history.push('/leavePolicys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add LeavePolicy</h3>
        }else{
            return <h3 className="text-center">Update LeavePolicy</h3>
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

                                            <label> accrualRate:&emsp; </label>
                                                <input placeholder="accrualRate" name="accrualRate" className="form-control" value={this.state.accrualRate} onChange={this.changeaccrualRateHandler}/>

                                            <label> carryoverAllowed:&emsp; </label>
                                                <input type="checkbox" placeholder="carryoverAllowed" name="carryoverAllowed" className="form-control" value={this.state.carryoverAllowed} onChange={this.changecarryoverAllowedHandler}/>


                                            <label> maxBalance:&emsp; </label>
                                                <input placeholder="maxBalance" name="maxBalance" className="form-control" value={this.state.maxBalance} onChange={this.changemaxBalanceHandler}/>

                                            <label> LeaveCategory:&emsp; </label>
                                                <select value={this.state.leaveCategory} onChange={this.changeLeaveCategoryHandler}>
                      <option name="LeaveCategory" className="form-control" >
                          Vacation
                      </option>
                      <option name="LeaveCategory" className="form-control" >
                          Sick
                      </option>
                      <option name="LeaveCategory" className="form-control" >
                          Parental
                      </option>
                      <option name="LeaveCategory" className="form-control" >
                          Bereavement
                      </option>
                      <option name="LeaveCategory" className="form-control" >
                          Unpaid
                      </option>
                      <option name="LeaveCategory" className="form-control" >
                          JuryDuty
                      </option>
                    </select>

                                            <label> AccrualUnit:&emsp; </label>
                                                <select value={this.state.accrualUnit} onChange={this.changeAccrualUnitHandler}>
                      <option name="AccrualUnit" className="form-control" >
                          Hours
                      </option>
                      <option name="AccrualUnit" className="form-control" >
                          Days
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateLeavePolicy}>Save</button>
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

export default CreateLeavePolicyComponent
