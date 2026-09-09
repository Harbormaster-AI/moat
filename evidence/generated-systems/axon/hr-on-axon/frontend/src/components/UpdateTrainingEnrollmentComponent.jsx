import React, { Component } from 'react'
import TrainingEnrollmentService from '../services/TrainingEnrollmentService';

class UpdateTrainingEnrollmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                enrollmentNumber: '',
                completionDate: '',
                score: '',
                status: ''
        }
        this.updateTrainingEnrollment = this.updateTrainingEnrollment.bind(this);

        this.changeenrollmentNumberHandler = this.changeenrollmentNumberHandler.bind(this);
        this.changecompletionDateHandler = this.changecompletionDateHandler.bind(this);
        this.changescoreHandler = this.changescoreHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        TrainingEnrollmentService.getTrainingEnrollmentById(this.state.id).then( (res) =>{
            let trainingEnrollment = res.data;
            this.setState({
                enrollmentNumber: trainingEnrollment.enrollmentNumber,
                completionDate: trainingEnrollment.completionDate,
                score: trainingEnrollment.score,
                status: trainingEnrollment.status
            });
        });
    }

    updateTrainingEnrollment = (e) => {
        e.preventDefault();
        let trainingEnrollment = {
            trainingEnrollmentId: this.state.id,
            enrollmentNumber: this.state.enrollmentNumber,
            completionDate: this.state.completionDate,
            score: this.state.score,
            status: this.state.status
        };
        console.log('trainingEnrollment => ' + JSON.stringify(trainingEnrollment));
        console.log('id => ' + JSON.stringify(this.state.id));
        TrainingEnrollmentService.updateTrainingEnrollment(trainingEnrollment).then( res => {
            this.props.history.push('/trainingEnrollments');
        });
    }

    changeenrollmentNumberHandler= (event) => {
        this.setState({enrollmentNumber: event.target.value});
    }
    changecompletionDateHandler= (event) => {
        this.setState({completionDate: event.target.value});
    }
    changescoreHandler= (event) => {
        this.setState({score: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/trainingEnrollments');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update TrainingEnrollment</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> enrollmentNumber: </label>
                                                <input placeholder="enrollmentNumber" name="enrollmentNumber" className="form-control" value={this.state.enrollmentNumber} onChange={this.changeenrollmentNumberHandler}/>

                                            <label> completionDate: </label>
                                                <input type="date" placeholder="completionDate" name="completionDate" className="form-control" value={this.state.completionDate} onChange={this.changecompletionDateHandler}/>

                                            <label> score: </label>
                                                <input placeholder="score" name="score" className="form-control" value={this.state.score} onChange={this.changescoreHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Enrolled
                      </option>
                      <option name="Status" className="form-control" >
                          InProgress
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          Failed
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateTrainingEnrollment}>Save</button>
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

export default UpdateTrainingEnrollmentComponent
