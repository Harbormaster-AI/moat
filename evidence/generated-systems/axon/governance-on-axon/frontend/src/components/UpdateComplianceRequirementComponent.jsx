import React, { Component } from 'react'
import ComplianceRequirementService from '../services/ComplianceRequirementService';

class UpdateComplianceRequirementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                source: '',
                citation: '',
                applicability: '',
                status: ''
        }
        this.updateComplianceRequirement = this.updateComplianceRequirement.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changesourceHandler = this.changesourceHandler.bind(this);
        this.changecitationHandler = this.changecitationHandler.bind(this);
        this.changeApplicabilityHandler = this.changeApplicabilityHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ComplianceRequirementService.getComplianceRequirementById(this.state.id).then( (res) =>{
            let complianceRequirement = res.data;
            this.setState({
                name: complianceRequirement.name,
                source: complianceRequirement.source,
                citation: complianceRequirement.citation,
                applicability: complianceRequirement.applicability,
                status: complianceRequirement.status
            });
        });
    }

    updateComplianceRequirement = (e) => {
        e.preventDefault();
        let complianceRequirement = {
            complianceRequirementId: this.state.id,
            name: this.state.name,
            source: this.state.source,
            citation: this.state.citation,
            applicability: this.state.applicability,
            status: this.state.status
        };
        console.log('complianceRequirement => ' + JSON.stringify(complianceRequirement));
        console.log('id => ' + JSON.stringify(this.state.id));
        ComplianceRequirementService.updateComplianceRequirement(complianceRequirement).then( res => {
            this.props.history.push('/complianceRequirements');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changesourceHandler= (event) => {
        this.setState({source: event.target.value});
    }
    changecitationHandler= (event) => {
        this.setState({citation: event.target.value});
    }
    changeApplicabilityHandler= (event) => {
        this.setState({applicability: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/complianceRequirements');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ComplianceRequirement</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> source: </label>
                                                <input placeholder="source" name="source" className="form-control" value={this.state.source} onChange={this.changesourceHandler}/>

                                            <label> citation: </label>
                                                <input placeholder="citation" name="citation" className="form-control" value={this.state.citation} onChange={this.changecitationHandler}/>

                                            <label> Applicability: </label>
                                                <select value={this.state.applicability} onChange={this.changeApplicabilityHandler}>
                      <option name="Applicability" className="form-control" >
                          Mandatory
                      </option>
                      <option name="Applicability" className="form-control" >
                          Recommended
                      </option>
                      <option name="Applicability" className="form-control" >
                          NotApplicable
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          NotStarted
                      </option>
                      <option name="Status" className="form-control" >
                          InProgress
                      </option>
                      <option name="Status" className="form-control" >
                          Compliant
                      </option>
                      <option name="Status" className="form-control" >
                          NonCompliant
                      </option>
                      <option name="Status" className="form-control" >
                          Waived
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateComplianceRequirement}>Save</button>
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

export default UpdateComplianceRequirementComponent
