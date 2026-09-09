import React, { Component } from 'react'
import ThirdPartyAssessmentService from '../services/ThirdPartyAssessmentService';

class UpdateThirdPartyAssessmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                assessmentDate: '',
                assessor: '',
                assessmentType: '',
                result: ''
        }
        this.updateThirdPartyAssessment = this.updateThirdPartyAssessment.bind(this);

        this.changeassessmentDateHandler = this.changeassessmentDateHandler.bind(this);
        this.changeassessorHandler = this.changeassessorHandler.bind(this);
        this.changeAssessmentTypeHandler = this.changeAssessmentTypeHandler.bind(this);
        this.changeResultHandler = this.changeResultHandler.bind(this);
    }

    componentDidMount(){
        ThirdPartyAssessmentService.getThirdPartyAssessmentById(this.state.id).then( (res) =>{
            let thirdPartyAssessment = res.data;
            this.setState({
                assessmentDate: thirdPartyAssessment.assessmentDate,
                assessor: thirdPartyAssessment.assessor,
                assessmentType: thirdPartyAssessment.assessmentType,
                result: thirdPartyAssessment.result
            });
        });
    }

    updateThirdPartyAssessment = (e) => {
        e.preventDefault();
        let thirdPartyAssessment = {
            thirdPartyAssessmentId: this.state.id,
            assessmentDate: this.state.assessmentDate,
            assessor: this.state.assessor,
            assessmentType: this.state.assessmentType,
            result: this.state.result
        };
        console.log('thirdPartyAssessment => ' + JSON.stringify(thirdPartyAssessment));
        console.log('id => ' + JSON.stringify(this.state.id));
        ThirdPartyAssessmentService.updateThirdPartyAssessment(thirdPartyAssessment).then( res => {
            this.props.history.push('/thirdPartyAssessments');
        });
    }

    changeassessmentDateHandler= (event) => {
        this.setState({assessmentDate: event.target.value});
    }
    changeassessorHandler= (event) => {
        this.setState({assessor: event.target.value});
    }
    changeAssessmentTypeHandler= (event) => {
        this.setState({assessmentType: event.target.value});
    }
    changeResultHandler= (event) => {
        this.setState({result: event.target.value});
    }

    cancel(){
        this.props.history.push('/thirdPartyAssessments');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ThirdPartyAssessment</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> assessmentDate: </label>
                                                <input type="date" placeholder="assessmentDate" name="assessmentDate" className="form-control" value={this.state.assessmentDate} onChange={this.changeassessmentDateHandler}/>

                                            <label> assessor: </label>
                                                <input placeholder="assessor" name="assessor" className="form-control" value={this.state.assessor} onChange={this.changeassessorHandler}/>

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

                                            <label> Result: </label>
                                                <select value={this.state.result} onChange={this.changeResultHandler}>
                      <option name="Result" className="form-control" >
                          Pass
                      </option>
                      <option name="Result" className="form-control" >
                          ConditionalPass
                      </option>
                      <option name="Result" className="form-control" >
                          Fail
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateThirdPartyAssessment}>Save</button>
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

export default UpdateThirdPartyAssessmentComponent
