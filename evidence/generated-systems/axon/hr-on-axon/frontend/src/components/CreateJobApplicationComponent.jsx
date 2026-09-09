import React, { Component } from 'react'
import JobApplicationService from '../services/JobApplicationService';

class CreateJobApplicationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                applicationNumber: '',
                appliedDate: '',
                resumeUrl: '',
                status: ''
        }
        this.changeapplicationNumberHandler = this.changeapplicationNumberHandler.bind(this);
        this.changeappliedDateHandler = this.changeappliedDateHandler.bind(this);
        this.changeresumeUrlHandler = this.changeresumeUrlHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            JobApplicationService.getJobApplicationById(this.state.id).then( (res) =>{
                let jobApplication = res.data;
                this.setState({
                    applicationNumber: jobApplication.applicationNumber,
                    appliedDate: jobApplication.appliedDate,
                    resumeUrl: jobApplication.resumeUrl,
                    status: jobApplication.status
                });
            });
        }        
    }
    saveOrUpdateJobApplication = (e) => {
        e.preventDefault();
        let jobApplication = {
                jobApplicationId: this.state.id,
                applicationNumber: this.state.applicationNumber,
                appliedDate: this.state.appliedDate,
                resumeUrl: this.state.resumeUrl,
                status: this.state.status
            };
        console.log('jobApplication => ' + JSON.stringify(jobApplication));

        // step 5
        if(this.state.id === '_add'){
            jobApplication.jobApplicationId=''
            JobApplicationService.createJobApplication(jobApplication).then(res =>{
                this.props.history.push('/jobApplications');
            });
        }else{
            JobApplicationService.updateJobApplication(jobApplication).then( res => {
                this.props.history.push('/jobApplications');
            });
        }
    }
    
    changeapplicationNumberHandler= (event) => {
        this.setState({applicationNumber: event.target.value});
    }
    changeappliedDateHandler= (event) => {
        this.setState({appliedDate: event.target.value});
    }
    changeresumeUrlHandler= (event) => {
        this.setState({resumeUrl: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/jobApplications');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add JobApplication</h3>
        }else{
            return <h3 className="text-center">Update JobApplication</h3>
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
                                            <label> applicationNumber:&emsp; </label>
                                                <input placeholder="applicationNumber" name="applicationNumber" className="form-control" value={this.state.applicationNumber} onChange={this.changeapplicationNumberHandler}/>

                                            <label> appliedDate:&emsp; </label>
                                                <input type="date" placeholder="appliedDate" name="appliedDate" className="form-control" value={this.state.appliedDate} onChange={this.changeappliedDateHandler}/>

                                            <label> resumeUrl:&emsp; </label>
                                                <input placeholder="resumeUrl" name="resumeUrl" className="form-control" value={this.state.resumeUrl} onChange={this.changeresumeUrlHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          New
                      </option>
                      <option name="Status" className="form-control" >
                          Screening
                      </option>
                      <option name="Status" className="form-control" >
                          Interview
                      </option>
                      <option name="Status" className="form-control" >
                          Offer
                      </option>
                      <option name="Status" className="form-control" >
                          Hired
                      </option>
                      <option name="Status" className="form-control" >
                          Rejected
                      </option>
                      <option name="Status" className="form-control" >
                          Withdrawn
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateJobApplication}>Save</button>
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

export default CreateJobApplicationComponent
