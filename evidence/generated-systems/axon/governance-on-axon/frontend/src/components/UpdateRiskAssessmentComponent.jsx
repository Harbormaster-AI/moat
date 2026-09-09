import React, { Component } from 'react'
import RiskAssessmentService from '../services/RiskAssessmentService';

class UpdateRiskAssessmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                assessmentDate: '',
                assessor: '',
                summary: '',
                assessmentType: ''
        }
        this.updateRiskAssessment = this.updateRiskAssessment.bind(this);

        this.changeassessmentDateHandler = this.changeassessmentDateHandler.bind(this);
        this.changeassessorHandler = this.changeassessorHandler.bind(this);
        this.changesummaryHandler = this.changesummaryHandler.bind(this);
        this.changeAssessmentTypeHandler = this.changeAssessmentTypeHandler.bind(this);
    }

    componentDidMount(){
        RiskAssessmentService.getRiskAssessmentById(this.state.id).then( (res) =>{
            let riskAssessment = res.data;
            this.setState({
                assessmentDate: riskAssessment.assessmentDate,
                assessor: riskAssessment.assessor,
                summary: riskAssessment.summary,
                assessmentType: riskAssessment.assessmentType
            });
        });
    }

    updateRiskAssessment = (e) => {
        e.preventDefault();
        let riskAssessment = {
            riskAssessmentId: this.state.id,
            assessmentDate: this.state.assessmentDate,
            assessor: this.state.assessor,
            summary: this.state.summary,
            assessmentType: this.state.assessmentType
        };
        console.log('riskAssessment => ' + JSON.stringify(riskAssessment));
        console.log('id => ' + JSON.stringify(this.state.id));
        RiskAssessmentService.updateRiskAssessment(riskAssessment).then( res => {
            this.props.history.push('/riskAssessments');
        });
    }

    changeassessmentDateHandler= (event) => {
        this.setState({assessmentDate: event.target.value});
    }
    changeassessorHandler= (event) => {
        this.setState({assessor: event.target.value});
    }
    changesummaryHandler= (event) => {
        this.setState({summary: event.target.value});
    }
    changeAssessmentTypeHandler= (event) => {
        this.setState({assessmentType: event.target.value});
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
                                            <label> assessmentDate: </label>
                                                <input type="date" placeholder="assessmentDate" name="assessmentDate" className="form-control" value={this.state.assessmentDate} onChange={this.changeassessmentDateHandler}/>

                                            <label> assessor: </label>
                                                <input placeholder="assessor" name="assessor" className="form-control" value={this.state.assessor} onChange={this.changeassessorHandler}/>

                                            <label> summary: </label>
                                                <input placeholder="summary" name="summary" className="form-control" value={this.state.summary} onChange={this.changesummaryHandler}/>

                                            <label> AssessmentType: </label>
                                                <select value={this.state.assessmentType} onChange={this.changeAssessmentTypeHandler}>
                      <option name="AssessmentType" className="form-control" >
                          SelfAssessment
                      </option>
                      <option name="AssessmentType" className="form-control" >
                          InternalAssessment
                      </option>
                      <option name="AssessmentType" className="form-control" >
                          ExternalAssessment
                      </option>
                      <option name="AssessmentType" className="form-control" >
                          ReadinessReview
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
