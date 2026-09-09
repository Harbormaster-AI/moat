import React, { Component } from 'react'
import InterviewService from '../services/InterviewService';

class UpdateInterviewComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                interviewDate: '',
                feedback: '',
                stage: '',
                result: ''
        }
        this.updateInterview = this.updateInterview.bind(this);

        this.changeinterviewDateHandler = this.changeinterviewDateHandler.bind(this);
        this.changefeedbackHandler = this.changefeedbackHandler.bind(this);
        this.changeStageHandler = this.changeStageHandler.bind(this);
        this.changeResultHandler = this.changeResultHandler.bind(this);
    }

    componentDidMount(){
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

    updateInterview = (e) => {
        e.preventDefault();
        let interview = {
            interviewId: this.state.id,
            interviewDate: this.state.interviewDate,
            feedback: this.state.feedback,
            stage: this.state.stage,
            result: this.state.result
        };
        console.log('interview => ' + JSON.stringify(interview));
        console.log('id => ' + JSON.stringify(this.state.id));
        InterviewService.updateInterview(interview).then( res => {
            this.props.history.push('/interviews');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Interview</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> interviewDate: </label>
                                                <input type="date" placeholder="interviewDate" name="interviewDate" className="form-control" value={this.state.interviewDate} onChange={this.changeinterviewDateHandler}/>

                                            <label> feedback: </label>
                                                <input placeholder="feedback" name="feedback" className="form-control" value={this.state.feedback} onChange={this.changefeedbackHandler}/>

                                            <label> Stage: </label>
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

                                            <label> Result: </label>
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
                                        <button className="btn btn-success" onClick={this.updateInterview}>Save</button>
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

export default UpdateInterviewComponent
