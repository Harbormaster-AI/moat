import React, { Component } from 'react'
import PolicyService from '../services/PolicyService';

class CreatePolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                policyNumber: '',
                name: '',
                effectiveDate: '',
                description: ''
        }
        this.changepolicyNumberHandler = this.changepolicyNumberHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeeffectiveDateHandler = this.changeeffectiveDateHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
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
                    name: policy.name,
                    effectiveDate: policy.effectiveDate,
                    description: policy.description
                });
            });
        }        
    }
    saveOrUpdatePolicy = (e) => {
        e.preventDefault();
        let policy = {
                policyId: this.state.id,
                policyNumber: this.state.policyNumber,
                name: this.state.name,
                effectiveDate: this.state.effectiveDate,
                description: this.state.description
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
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeeffectiveDateHandler= (event) => {
        this.setState({effectiveDate: event.target.value});
    }
    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
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

                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> effectiveDate:&emsp; </label>
                                                <input type="date" placeholder="effectiveDate" name="effectiveDate" className="form-control" value={this.state.effectiveDate} onChange={this.changeeffectiveDateHandler}/>

                                            <label> description:&emsp; </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

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
