import React, { Component } from 'react'
import InterviewService from '../services/InterviewService';

class CreateInterviewComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                interviewDate: '',
                feedback: '',
                stage: '',
                result: ''
        }
        this.changeinterviewDateHandler = this.changeinterviewDateHandler.bind(this);
        this.changefeedbackHandler = this.changefeedbackHandler.bind(this);
        this.changeStageHandler = this.changeStageHandler.bind(this);
        this.changeResultHandler = this.changeResultHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            InterviewService.getInterviewById(this.state.id).then( (res) =>{
                let interview = res.data;
                this.setState({
                    interviewDate: interview.interviewDate,
                    feedback: interview.feedback,
                    stage: interview.stage,
                    result: interview.result
                });
            });
        }        
    }
    saveOrUpdateInterview = (e) => {
        e.preventDefault();
        let interview = {
                interviewId: this.state.id,
                interviewDate: this.state.interviewDate,
                feedback: this.state.feedback,
                stage: this.state.stage,
                result: this.state.result
            };
        console.log('interview => ' + JSON.stringify(interview));

        // step 5
        if(this.state.id === '_add'){
            interview.interviewId=''
            InterviewService.createInterview(interview).then(res =>{
                this.props.history.push('/interviews');
            });
        }else{
            InterviewService.updateInterview(interview).then( res => {
                this.props.history.push('/interviews');
            });
        }
    }
    
    changeinterviewDateHandler= (event) => {
        this.setState({interviewDate: event.target.value});
    }
    changefeedbackHandler= (event) => {
        this.setState({feedback: event.target.value});
    }
    changeStageHandler= (event) => {
        this.setState({stage: event.target.value});
    }
    changeResultHandler= (event) => {
        this.setState({result: event.target.value});
    }

    cancel(){
        this.props.history.push('/interviews');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Interview</h3>
        }else{
            return <h3 className="text-center">Update Interview</h3>
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
                                            <label> interviewDate:&emsp; </label>
                                                <input type="date" placeholder="interviewDate" name="interviewDate" className="form-control" value={this.state.interviewDate} onChange={this.changeinterviewDateHandler}/>

                                            <label> feedback:&emsp; </label>
                                                <input placeholder="feedback" name="feedback" className="form-control" value={this.state.feedback} onChange={this.changefeedbackHandler}/>

                                            <label> Stage:&emsp; </label>
                                                <select value={this.state.stage} onChange={this.changeStageHandler}>
                      <option name="Stage" className="form-control" >
                          PhoneScreen
                      </option>
                      <option name="Stage" className="form-control" >
                          Technical
                      </option>
                      <option name="Stage" className="form-control" >
                          Onsite
                      </option>
                      <option name="Stage" className="form-control" >
                          Panel
                      </option>
                      <option name="Stage" className="form-control" >
                          HR
                      </option>
                      <option name="Stage" className="form-control" >
                          Executive
                      </option>
                    </select>

                                            <label> Result:&emsp; </label>
                                                <select value={this.state.result} onChange={this.changeResultHandler}>
                      <option name="Result" className="form-control" >
                          Pending
                      </option>
                      <option name="Result" className="form-control" >
                          Proceed
                      </option>
                      <option name="Result" className="form-control" >
                          Reject
                      </option>
                      <option name="Result" className="form-control" >
                          OfferRecommended
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateInterview}>Save</button>
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

export default CreateInterviewComponent
