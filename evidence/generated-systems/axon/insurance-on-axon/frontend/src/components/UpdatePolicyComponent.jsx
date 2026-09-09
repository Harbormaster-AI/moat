import React, { Component } from 'react'
import PolicyService from '../services/PolicyService';

class UpdatePolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                policyNumber: '',
                effectivePeriod: '',
                totalPremium: '',
                status: '',
                paymentPlan: ''
        }
        this.updatePolicy = this.updatePolicy.bind(this);

        this.changepolicyNumberHandler = this.changepolicyNumberHandler.bind(this);
        this.changeeffectivePeriodHandler = this.changeeffectivePeriodHandler.bind(this);
        this.changetotalPremiumHandler = this.changetotalPremiumHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changePaymentPlanHandler = this.changePaymentPlanHandler.bind(this);
    }

    componentDidMount(){
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

    updatePolicy = (e) => {
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
        console.log('id => ' + JSON.stringify(this.state.id));
        PolicyService.updatePolicy(policy).then( res => {
            this.props.history.push('/policys');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Policy</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> policyNumber: </label>
                                                <input placeholder="policyNumber" name="policyNumber" className="form-control" value={this.state.policyNumber} onChange={this.changepolicyNumberHandler}/>

                                            <label> effectivePeriod: </label>
                                                <input placeholder="effectivePeriod" name="effectivePeriod" className="form-control" value={this.state.effectivePeriod} onChange={this.changeeffectivePeriodHandler}/>

                                            <label> totalPremium: </label>
                                                <input placeholder="totalPremium" name="totalPremium" className="form-control" value={this.state.totalPremium} onChange={this.changetotalPremiumHandler}/>

                                            <label> Status: </label>
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

                                            <label> PaymentPlan: </label>
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
                                        <button className="btn btn-success" onClick={this.updatePolicy}>Save</button>
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

export default UpdatePolicyComponent
