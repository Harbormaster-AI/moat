import React, { Component } from 'react'
import CompliancePolicyService from '../services/CompliancePolicyService';

class UpdateCompliancePolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                policyCode: '',
                description: '',
                status: ''
        }
        this.updateCompliancePolicy = this.updateCompliancePolicy.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changepolicyCodeHandler = this.changepolicyCodeHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        CompliancePolicyService.getCompliancePolicyById(this.state.id).then( (res) =>{
            let compliancePolicy = res.data;
            this.setState({
                name: compliancePolicy.name,
                policyCode: compliancePolicy.policyCode,
                description: compliancePolicy.description,
                status: compliancePolicy.status
            });
        });
    }

    updateCompliancePolicy = (e) => {
        e.preventDefault();
        let compliancePolicy = {
            compliancePolicyId: this.state.id,
            name: this.state.name,
            policyCode: this.state.policyCode,
            description: this.state.description,
            status: this.state.status
        };
        console.log('compliancePolicy => ' + JSON.stringify(compliancePolicy));
        console.log('id => ' + JSON.stringify(this.state.id));
        CompliancePolicyService.updateCompliancePolicy(compliancePolicy).then( res => {
            this.props.history.push('/compliancePolicys');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changepolicyCodeHandler= (event) => {
        this.setState({policyCode: event.target.value});
    }
    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/compliancePolicys');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update CompliancePolicy</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> policyCode: </label>
                                                <input placeholder="policyCode" name="policyCode" className="form-control" value={this.state.policyCode} onChange={this.changepolicyCodeHandler}/>

                                            <label> description: </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Retired
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCompliancePolicy}>Save</button>
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

export default UpdateCompliancePolicyComponent
