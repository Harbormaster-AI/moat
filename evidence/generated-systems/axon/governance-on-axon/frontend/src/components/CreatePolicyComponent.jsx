import React, { Component } from 'react'
import PolicyService from '../services/PolicyService';

class CreatePolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                title: '',
                versionLabel: '',
                approvalDate: '',
                nextReviewDate: '',
                documentUrl: '',
                policyType: '',
                status: ''
        }
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changeversionLabelHandler = this.changeversionLabelHandler.bind(this);
        this.changeapprovalDateHandler = this.changeapprovalDateHandler.bind(this);
        this.changenextReviewDateHandler = this.changenextReviewDateHandler.bind(this);
        this.changedocumentUrlHandler = this.changedocumentUrlHandler.bind(this);
        this.changePolicyTypeHandler = this.changePolicyTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
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
                    title: policy.title,
                    versionLabel: policy.versionLabel,
                    approvalDate: policy.approvalDate,
                    nextReviewDate: policy.nextReviewDate,
                    documentUrl: policy.documentUrl,
                    policyType: policy.policyType,
                    status: policy.status
                });
            });
        }        
    }
    saveOrUpdatePolicy = (e) => {
        e.preventDefault();
        let policy = {
                policyId: this.state.id,
                title: this.state.title,
                versionLabel: this.state.versionLabel,
                approvalDate: this.state.approvalDate,
                nextReviewDate: this.state.nextReviewDate,
                documentUrl: this.state.documentUrl,
                policyType: this.state.policyType,
                status: this.state.status
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
    
    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changeversionLabelHandler= (event) => {
        this.setState({versionLabel: event.target.value});
    }
    changeapprovalDateHandler= (event) => {
        this.setState({approvalDate: event.target.value});
    }
    changenextReviewDateHandler= (event) => {
        this.setState({nextReviewDate: event.target.value});
    }
    changedocumentUrlHandler= (event) => {
        this.setState({documentUrl: event.target.value});
    }
    changePolicyTypeHandler= (event) => {
        this.setState({policyType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
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
                                            <label> title:&emsp; </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> versionLabel:&emsp; </label>
                                                <input placeholder="versionLabel" name="versionLabel" className="form-control" value={this.state.versionLabel} onChange={this.changeversionLabelHandler}/>

                                            <label> approvalDate:&emsp; </label>
                                                <input type="date" placeholder="approvalDate" name="approvalDate" className="form-control" value={this.state.approvalDate} onChange={this.changeapprovalDateHandler}/>

                                            <label> nextReviewDate:&emsp; </label>
                                                <input type="date" placeholder="nextReviewDate" name="nextReviewDate" className="form-control" value={this.state.nextReviewDate} onChange={this.changenextReviewDateHandler}/>

                                            <label> documentUrl:&emsp; </label>
                                                <input placeholder="documentUrl" name="documentUrl" className="form-control" value={this.state.documentUrl} onChange={this.changedocumentUrlHandler}/>

                                            <label> PolicyType:&emsp; </label>
                                                <select value={this.state.policyType} onChange={this.changePolicyTypeHandler}>
                      <option name="PolicyType" className="form-control" >
                          InformationSecurity
                      </option>
                      <option name="PolicyType" className="form-control" >
                          DataProtection
                      </option>
                      <option name="PolicyType" className="form-control" >
                          Ethics
                      </option>
                      <option name="PolicyType" className="form-control" >
                          RecordsManagement
                      </option>
                      <option name="PolicyType" className="form-control" >
                          RiskManagement
                      </option>
                      <option name="PolicyType" className="form-control" >
                          Compliance
                      </option>
                      <option name="PolicyType" className="form-control" >
                          Privacy
                      </option>
                      <option name="PolicyType" className="form-control" >
                          AcceptableUse
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          InReview
                      </option>
                      <option name="Status" className="form-control" >
                          Approved
                      </option>
                      <option name="Status" className="form-control" >
                          Retired
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
