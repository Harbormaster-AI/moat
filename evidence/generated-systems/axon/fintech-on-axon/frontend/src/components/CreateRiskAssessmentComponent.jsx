import React, { Component } from 'react'
import RiskAssessmentService from '../services/RiskAssessmentService';

class CreateRiskAssessmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                score: '',
                assessedAt: '',
                modelVersion: '',
                notes: '',
                decision: ''
        }
        this.changescoreHandler = this.changescoreHandler.bind(this);
        this.changeassessedAtHandler = this.changeassessedAtHandler.bind(this);
        this.changemodelVersionHandler = this.changemodelVersionHandler.bind(this);
        this.changenotesHandler = this.changenotesHandler.bind(this);
        this.changeDecisionHandler = this.changeDecisionHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateRiskAssessment = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            riskAssessment.riskAssessmentId=''
            RiskAssessmentService.createRiskAssessment(riskAssessment).then(res =>{
                this.props.history.push('/riskAssessments');
            });
        }else{
            RiskAssessmentService.updateRiskAssessment(riskAssessment).then( res => {
                this.props.history.push('/riskAssessments');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add RiskAssessment</h3>
        }else{
            return <h3 className="text-center">Update RiskAssessment</h3>
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
                                            <label> score:&emsp; </label>
                                                <input placeholder="score" name="score" className="form-control" value={this.state.score} onChange={this.changescoreHandler}/>

                                            <label> assessedAt:&emsp; </label>
                                                <input type="time" placeholder="assessedAt" name="assessedAt" className="form-control" value={this.state.assessedAt} onChange={this.changeassessedAtHandler}/>

                                            <label> modelVersion:&emsp; </label>
                                                <input placeholder="modelVersion" name="modelVersion" className="form-control" value={this.state.modelVersion} onChange={this.changemodelVersionHandler}/>

                                            <label> notes:&emsp; </label>
                                                <input placeholder="notes" name="notes" className="form-control" value={this.state.notes} onChange={this.changenotesHandler}/>

                                            <label> Decision:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateRiskAssessment}>Save</button>
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

export default CreateRiskAssessmentComponent
