import React, { Component } from 'react'
import TrainingEnrollmentService from '../services/TrainingEnrollmentService';

class CreateTrainingEnrollmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                enrollmentNumber: '',
                completionDate: '',
                score: '',
                status: ''
        }
        this.changeenrollmentNumberHandler = this.changeenrollmentNumberHandler.bind(this);
        this.changecompletionDateHandler = this.changecompletionDateHandler.bind(this);
        this.changescoreHandler = this.changescoreHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateTrainingEnrollment = (e) => {
        e.preventDefault();
        let trainingEnrollment = {
                trainingEnrollmentId: this.state.id,
                enrollmentNumber: this.state.enrollmentNumber,
                completionDate: this.state.completionDate,
                score: this.state.score,
                status: this.state.status
            };
        console.log('trainingEnrollment => ' + JSON.stringify(trainingEnrollment));

        // step 5
        if(this.state.id === '_add'){
            trainingEnrollment.trainingEnrollmentId=''
            TrainingEnrollmentService.createTrainingEnrollment(trainingEnrollment).then(res =>{
                this.props.history.push('/trainingEnrollments');
            });
        }else{
            TrainingEnrollmentService.updateTrainingEnrollment(trainingEnrollment).then( res => {
                this.props.history.push('/trainingEnrollments');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add TrainingEnrollment</h3>
        }else{
            return <h3 className="text-center">Update TrainingEnrollment</h3>
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
                                            <label> enrollmentNumber:&emsp; </label>
                                                <input placeholder="enrollmentNumber" name="enrollmentNumber" className="form-control" value={this.state.enrollmentNumber} onChange={this.changeenrollmentNumberHandler}/>

                                            <label> completionDate:&emsp; </label>
                                                <input type="date" placeholder="completionDate" name="completionDate" className="form-control" value={this.state.completionDate} onChange={this.changecompletionDateHandler}/>

                                            <label> score:&emsp; </label>
                                                <input placeholder="score" name="score" className="form-control" value={this.state.score} onChange={this.changescoreHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateTrainingEnrollment}>Save</button>
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

export default CreateTrainingEnrollmentComponent
