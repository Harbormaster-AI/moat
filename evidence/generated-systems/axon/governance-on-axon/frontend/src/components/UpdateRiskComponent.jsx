import React, { Component } from 'react'
import RiskService from '../services/RiskService';

class UpdateRiskComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                description: '',
                inherentRiskScore: '',
                residualRiskScore: '',
                category: '',
                impact: '',
                likelihood: '',
                status: ''
        }
        this.updateRisk = this.updateRisk.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changeinherentRiskScoreHandler = this.changeinherentRiskScoreHandler.bind(this);
        this.changeresidualRiskScoreHandler = this.changeresidualRiskScoreHandler.bind(this);
        this.changeCategoryHandler = this.changeCategoryHandler.bind(this);
        this.changeImpactHandler = this.changeImpactHandler.bind(this);
        this.changeLikelihoodHandler = this.changeLikelihoodHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        RiskService.getRiskById(this.state.id).then( (res) =>{
            let risk = res.data;
            this.setState({
                name: risk.name,
                description: risk.description,
                inherentRiskScore: risk.inherentRiskScore,
                residualRiskScore: risk.residualRiskScore,
                category: risk.category,
                impact: risk.impact,
                likelihood: risk.likelihood,
                status: risk.status
            });
        });
    }

    updateRisk = (e) => {
        e.preventDefault();
        let risk = {
            riskId: this.state.id,
            name: this.state.name,
            description: this.state.description,
            inherentRiskScore: this.state.inherentRiskScore,
            residualRiskScore: this.state.residualRiskScore,
            category: this.state.category,
            impact: this.state.impact,
            likelihood: this.state.likelihood,
            status: this.state.status
        };
        console.log('risk => ' + JSON.stringify(risk));
        console.log('id => ' + JSON.stringify(this.state.id));
        RiskService.updateRisk(risk).then( res => {
            this.props.history.push('/risks');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }
    changeinherentRiskScoreHandler= (event) => {
        this.setState({inherentRiskScore: event.target.value});
    }
    changeresidualRiskScoreHandler= (event) => {
        this.setState({residualRiskScore: event.target.value});
    }
    changeCategoryHandler= (event) => {
        this.setState({category: event.target.value});
    }
    changeImpactHandler= (event) => {
        this.setState({impact: event.target.value});
    }
    changeLikelihoodHandler= (event) => {
        this.setState({likelihood: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/risks');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Risk</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> description: </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> inherentRiskScore: </label>
                                                <input type="number" placeholder="inherentRiskScore" name="inherentRiskScore" className="form-control" value={this.state.inherentRiskScore} onChange={this.changeinherentRiskScoreHandler}/>

                                            <label> residualRiskScore: </label>
                                                <input type="number" placeholder="residualRiskScore" name="residualRiskScore" className="form-control" value={this.state.residualRiskScore} onChange={this.changeresidualRiskScoreHandler}/>

                                            <label> Category: </label>
                                                <select value={this.state.category} onChange={this.changeCategoryHandler}>
                      <option name="Category" className="form-control" >
                          Strategic
                      </option>
                      <option name="Category" className="form-control" >
                          Operational
                      </option>
                      <option name="Category" className="form-control" >
                          Financial
                      </option>
                      <option name="Category" className="form-control" >
                          Compliance
                      </option>
                      <option name="Category" className="form-control" >
                          Reputational
                      </option>
                      <option name="Category" className="form-control" >
                          Privacy
                      </option>
                      <option name="Category" className="form-control" >
                          Cybersecurity
                      </option>
                      <option name="Category" className="form-control" >
                          ThirdParty
                      </option>
                    </select>

                                            <label> Impact: </label>
                                                <select value={this.state.impact} onChange={this.changeImpactHandler}>
                      <option name="Impact" className="form-control" >
                          Insignificant
                      </option>
                      <option name="Impact" className="form-control" >
                          Minor
                      </option>
                      <option name="Impact" className="form-control" >
                          Moderate
                      </option>
                      <option name="Impact" className="form-control" >
                          Major
                      </option>
                      <option name="Impact" className="form-control" >
                          Severe
                      </option>
                    </select>

                                            <label> Likelihood: </label>
                                                <select value={this.state.likelihood} onChange={this.changeLikelihoodHandler}>
                      <option name="Likelihood" className="form-control" >
                          Rare
                      </option>
                      <option name="Likelihood" className="form-control" >
                          Unlikely
                      </option>
                      <option name="Likelihood" className="form-control" >
                          Possible
                      </option>
                      <option name="Likelihood" className="form-control" >
                          Likely
                      </option>
                      <option name="Likelihood" className="form-control" >
                          AlmostCertain
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Identified
                      </option>
                      <option name="Status" className="form-control" >
                          Assessed
                      </option>
                      <option name="Status" className="form-control" >
                          Mitigated
                      </option>
                      <option name="Status" className="form-control" >
                          Accepted
                      </option>
                      <option name="Status" className="form-control" >
                          Transferred
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateRisk}>Save</button>
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

export default UpdateRiskComponent
