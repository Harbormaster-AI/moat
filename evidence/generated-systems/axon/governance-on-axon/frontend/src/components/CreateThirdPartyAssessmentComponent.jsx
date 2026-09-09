import React, { Component } from 'react'
import ThirdPartyAssessmentService from '../services/ThirdPartyAssessmentService';

class CreateThirdPartyAssessmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                assessmentDate: '',
                assessor: '',
                assessmentType: '',
                result: ''
        }
        this.changeassessmentDateHandler = this.changeassessmentDateHandler.bind(this);
        this.changeassessorHandler = this.changeassessorHandler.bind(this);
        this.changeAssessmentTypeHandler = this.changeAssessmentTypeHandler.bind(this);
        this.changeResultHandler = this.changeResultHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateThirdPartyAssessment = (e) => {
        e.preventDefault();
        let thirdPartyAssessment = {
                thirdPartyAssessmentId: this.state.id,
                assessmentDate: this.state.assessmentDate,
                assessor: this.state.assessor,
                assessmentType: this.state.assessmentType,
                result: this.state.result
            };
        console.log('thirdPartyAssessment => ' + JSON.stringify(thirdPartyAssessment));

        // step 5
        if(this.state.id === '_add'){
            thirdPartyAssessment.thirdPartyAssessmentId=''
            ThirdPartyAssessmentService.createThirdPartyAssessment(thirdPartyAssessment).then(res =>{
                this.props.history.push('/thirdPartyAssessments');
            });
        }else{
            ThirdPartyAssessmentService.updateThirdPartyAssessment(thirdPartyAssessment).then( res => {
                this.props.history.push('/thirdPartyAssessments');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ThirdPartyAssessment</h3>
        }else{
            return <h3 className="text-center">Update ThirdPartyAssessment</h3>
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
                                            <label> assessmentDate:&emsp; </label>
                                                <input type="date" placeholder="assessmentDate" name="assessmentDate" className="form-control" value={this.state.assessmentDate} onChange={this.changeassessmentDateHandler}/>

                                            <label> assessor:&emsp; </label>
                                                <input placeholder="assessor" name="assessor" className="form-control" value={this.state.assessor} onChange={this.changeassessorHandler}/>

                                            <label> AssessmentType:&emsp; </label>
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

                                            <label> Result:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateThirdPartyAssessment}>Save</button>
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

export default CreateThirdPartyAssessmentComponent
