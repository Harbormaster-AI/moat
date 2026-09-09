import React, { Component } from 'react'
import PolicyService from '../services/PolicyService';

class UpdatePolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                policyNumber: '',
                name: '',
                effectiveDate: '',
                description: ''
        }
        this.updatePolicy = this.updatePolicy.bind(this);

        this.changepolicyNumberHandler = this.changepolicyNumberHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeeffectiveDateHandler = this.changeeffectiveDateHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
    }

    componentDidMount(){
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

    updatePolicy = (e) => {
        e.preventDefault();
        let policy = {
            policyId: this.state.id,
            policyNumber: this.state.policyNumber,
            name: this.state.name,
            effectiveDate: this.state.effectiveDate,
            description: this.state.description
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

                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> effectiveDate: </label>
                                                <input type="date" placeholder="effectiveDate" name="effectiveDate" className="form-control" value={this.state.effectiveDate} onChange={this.changeeffectiveDateHandler}/>

                                            <label> description: </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

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
