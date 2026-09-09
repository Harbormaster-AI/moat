import React, { Component } from 'react'
import JobProfileService from '../services/JobProfileService';

class CreateJobProfileComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                title: '',
                jobCode: '',
                jobLevel: '',
                exemptStatus: ''
        }
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changejobCodeHandler = this.changejobCodeHandler.bind(this);
        this.changeJobLevelHandler = this.changeJobLevelHandler.bind(this);
        this.changeExemptStatusHandler = this.changeExemptStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            JobProfileService.getJobProfileById(this.state.id).then( (res) =>{
                let jobProfile = res.data;
                this.setState({
                    title: jobProfile.title,
                    jobCode: jobProfile.jobCode,
                    jobLevel: jobProfile.jobLevel,
                    exemptStatus: jobProfile.exemptStatus
                });
            });
        }        
    }
    saveOrUpdateJobProfile = (e) => {
        e.preventDefault();
        let jobProfile = {
                jobProfileId: this.state.id,
                title: this.state.title,
                jobCode: this.state.jobCode,
                jobLevel: this.state.jobLevel,
                exemptStatus: this.state.exemptStatus
            };
        console.log('jobProfile => ' + JSON.stringify(jobProfile));

        // step 5
        if(this.state.id === '_add'){
            jobProfile.jobProfileId=''
            JobProfileService.createJobProfile(jobProfile).then(res =>{
                this.props.history.push('/jobProfiles');
            });
        }else{
            JobProfileService.updateJobProfile(jobProfile).then( res => {
                this.props.history.push('/jobProfiles');
            });
        }
    }
    
    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changejobCodeHandler= (event) => {
        this.setState({jobCode: event.target.value});
    }
    changeJobLevelHandler= (event) => {
        this.setState({jobLevel: event.target.value});
    }
    changeExemptStatusHandler= (event) => {
        this.setState({exemptStatus: event.target.value});
    }

    cancel(){
        this.props.history.push('/jobProfiles');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add JobProfile</h3>
        }else{
            return <h3 className="text-center">Update JobProfile</h3>
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
                                            <label> title:&emsp; </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> jobCode:&emsp; </label>
                                                <input placeholder="jobCode" name="jobCode" className="form-control" value={this.state.jobCode} onChange={this.changejobCodeHandler}/>

                                            <label> JobLevel:&emsp; </label>
                                                <select value={this.state.jobLevel} onChange={this.changeJobLevelHandler}>
                      <option name="JobLevel" className="form-control" >
                          Entry
                      </option>
                      <option name="JobLevel" className="form-control" >
                          Intermediate
                      </option>
                      <option name="JobLevel" className="form-control" >
                          Senior
                      </option>
                      <option name="JobLevel" className="form-control" >
                          Lead
                      </option>
                      <option name="JobLevel" className="form-control" >
                          Manager
                      </option>
                      <option name="JobLevel" className="form-control" >
                          Director
                      </option>
                      <option name="JobLevel" className="form-control" >
                          Executive
                      </option>
                    </select>

                                            <label> ExemptStatus:&emsp; </label>
                                                <select value={this.state.exemptStatus} onChange={this.changeExemptStatusHandler}>
                      <option name="ExemptStatus" className="form-control" >
                          Exempt
                      </option>
                      <option name="ExemptStatus" className="form-control" >
                          NonExempt
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateJobProfile}>Save</button>
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

export default CreateJobProfileComponent
