import React, { Component } from 'react'
import PolicyService from '../services/PolicyService';

class CreatePolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                policyNumber: '',
                effectivePeriod: '',
                totalPremium: '',
                status: '',
                paymentPlan: ''
        }
        this.changepolicyNumberHandler = this.changepolicyNumberHandler.bind(this);
        this.changeeffectivePeriodHandler = this.changeeffectivePeriodHandler.bind(this);
        this.changetotalPremiumHandler = this.changetotalPremiumHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changePaymentPlanHandler = this.changePaymentPlanHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PolicyService.getPolicyById(this.state.id).then( (res) =>{
                let policy = res.data;
                this.setState({
                    policyNumber: policy.policyNumber,
                    effectivePeriod: policy.effectivePeriod,
                    totalPremium: policy.totalPremium,
                    status: policy.status,
                    paymentPlan: policy.paymentPlan
                });
            });
        }        
    }
    saveOrUpdatePolicy = (e) => {
        e.preventDefault();
        let policy = {
                policyId: this.state.id,
                policyNumber: this.state.policyNumber,
                effectivePeriod: this.state.effectivePeriod,
                totalPremium: this.state.totalPremium,
                status: this.state.status,
                paymentPlan: this.state.paymentPlan
            };
        console.log('policy => ' + JSON.stringify(policy));

        // step 5
        if(this.state.id === '_add'){
            policy.policyId=''
            PolicyService.createPolicy(policy).then(res =>{
                this.props.history.push('/policys');
            });
        }else{
            PolicyService.updatePolicy(policy).then( res => {
                this.props.history.push('/policys');
            });
        }
    }
    
    changepolicyNumberHandler= (event) => {
        this.setState({policyNumber: event.target.value});
    }
    changeeffectivePeriodHandler= (event) => {
        this.setState({effectivePeriod: event.target.value});
    }
    changetotalPremiumHandler= (event) => {
        this.setState({totalPremium: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changePaymentPlanHandler= (event) => {
        this.setState({paymentPlan: event.target.value});
    }

    cancel(){
        this.props.history.push('/policys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Policy</h3>
        }else{
            return <h3 className="text-center">Update Policy</h3>
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
                                            <label> policyNumber:&emsp; </label>
                                                <input placeholder="policyNumber" name="policyNumber" className="form-control" value={this.state.policyNumber} onChange={this.changepolicyNumberHandler}/>

                                            <label> effectivePeriod:&emsp; </label>
                                                <input placeholder="effectivePeriod" name="effectivePeriod" className="form-control" value={this.state.effectivePeriod} onChange={this.changeeffectivePeriodHandler}/>

                                            <label> totalPremium:&emsp; </label>
                                                <input placeholder="totalPremium" name="totalPremium" className="form-control" value={this.state.totalPremium} onChange={this.changetotalPremiumHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Quoted
                      </option>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Lapsed
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                      <option name="Status" className="form-control" >
                          Expired
                      </option>
                      <option name="Status" className="form-control" >
                          PendingCancel
                      </option>
                      <option name="Status" className="form-control" >
                          PendingReinstatement
                      </option>
                    </select>

                                            <label> PaymentPlan:&emsp; </label>
                                                <select value={this.state.paymentPlan} onChange={this.changePaymentPlanHandler}>
                      <option name="PaymentPlan" className="form-control" >
                          Annual
                      </option>
                      <option name="PaymentPlan" className="form-control" >
                          SemiAnnual
                      </option>
                      <option name="PaymentPlan" className="form-control" >
                          Quarterly
                      </option>
                      <option name="PaymentPlan" className="form-control" >
                          Monthly
                      </option>
                      <option name="PaymentPlan" className="form-control" >
                          PayInFull
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePolicy}>Save</button>
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

export default CreatePolicyComponent
