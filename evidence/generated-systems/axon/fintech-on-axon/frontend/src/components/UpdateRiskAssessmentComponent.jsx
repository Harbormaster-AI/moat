import React, { Component } from 'react'
import RiskAssessmentService from '../services/RiskAssessmentService';

class UpdateRiskAssessmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                score: '',
                assessedAt: '',
                modelVersion: '',
                notes: '',
                decision: ''
        }
        this.updateRiskAssessment = this.updateRiskAssessment.bind(this);

        this.changescoreHandler = this.changescoreHandler.bind(this);
        this.changeassessedAtHandler = this.changeassessedAtHandler.bind(this);
        this.changemodelVersionHandler = this.changemodelVersionHandler.bind(this);
        this.changenotesHandler = this.changenotesHandler.bind(this);
        this.changeDecisionHandler = this.changeDecisionHandler.bind(this);
    }

    componentDidMount(){
        RiskAssessmentService.getRiskAssessmentById(this.state.id).then( (res) =>{
            let riskAssessment = res.data;
            this.setState({
                score: riskAssessment.score,
                assessedAt: riskAssessment.assessedAt,
                modelVersion: riskAssessment.modelVersion,
                notes: riskAssessment.notes,
                decision: riskAssessment.decision
            });
        });
    }

    updateRiskAssessment = (e) => {
        e.preventDefault();
        let riskAssessment = {
            riskAssessmentId: this.state.id,
            score: this.state.score,
            assessedAt: this.state.assessedAt,
            modelVersion: this.state.modelVersion,
            notes: this.state.notes,
            decision: this.state.decision
        };
        console.log('riskAssessment => ' + JSON.stringify(riskAssessment));
        console.log('id => ' + JSON.stringify(this.state.id));
        RiskAssessmentService.updateRiskAssessment(riskAssessment).then( res => {
            this.props.history.push('/riskAssessments');
        });
    }

    changescoreHandler= (event) => {
        this.setState({score: event.target.value});
    }
    changeassessedAtHandler= (event) => {
        this.setState({assessedAt: event.target.value});
    }
    changemodelVersionHandler= (event) => {
        this.setState({modelVersion: event.target.value});
    }
    changenotesHandler= (event) => {
        this.setState({notes: event.target.value});
    }
    changeDecisionHandler= (event) => {
        this.setState({decision: event.target.value});
    }

    cancel(){
        this.props.history.push('/riskAssessments');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update RiskAssessment</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> score: </label>
                                                <input placeholder="score" name="score" className="form-control" value={this.state.score} onChange={this.changescoreHandler}/>

                                            <label> assessedAt: </label>
                                                <input type="time" placeholder="assessedAt" name="assessedAt" className="form-control" value={this.state.assessedAt} onChange={this.changeassessedAtHandler}/>

                                            <label> modelVersion: </label>
                                                <input placeholder="modelVersion" name="modelVersion" className="form-control" value={this.state.modelVersion} onChange={this.changemodelVersionHandler}/>

                                            <label> notes: </label>
                                                <input placeholder="notes" name="notes" className="form-control" value={this.state.notes} onChange={this.changenotesHandler}/>

                                            <label> Decision: </label>
                                                <select value={this.state.decision} onChange={this.changeDecisionHandler}>
                      <option name="Decision" className="form-control" >
                          Approve
                      </option>
                      <option name="Decision" className="form-control" >
                          Decline
                      </option>
                      <option name="Decision" className="form-control" >
                          Refer
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateRiskAssessment}>Save</button>
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

export default UpdateRiskAssessmentComponent
