import React, { Component } from 'react'
import JobProfileService from '../services/JobProfileService';

class UpdateJobProfileComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                title: '',
                jobCode: '',
                jobLevel: '',
                exemptStatus: ''
        }
        this.updateJobProfile = this.updateJobProfile.bind(this);

        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changejobCodeHandler = this.changejobCodeHandler.bind(this);
        this.changeJobLevelHandler = this.changeJobLevelHandler.bind(this);
        this.changeExemptStatusHandler = this.changeExemptStatusHandler.bind(this);
    }

    componentDidMount(){
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

    updateJobProfile = (e) => {
        e.preventDefault();
        let jobProfile = {
            jobProfileId: this.state.id,
            title: this.state.title,
            jobCode: this.state.jobCode,
            jobLevel: this.state.jobLevel,
            exemptStatus: this.state.exemptStatus
        };
        console.log('jobProfile => ' + JSON.stringify(jobProfile));
        console.log('id => ' + JSON.stringify(this.state.id));
        JobProfileService.updateJobProfile(jobProfile).then( res => {
            this.props.history.push('/jobProfiles');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update JobProfile</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> title: </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> jobCode: </label>
                                                <input placeholder="jobCode" name="jobCode" className="form-control" value={this.state.jobCode} onChange={this.changejobCodeHandler}/>

                                            <label> JobLevel: </label>
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

                                            <label> ExemptStatus: </label>
                                                <select value={this.state.exemptStatus} onChange={this.changeExemptStatusHandler}>
                      <option name="ExemptStatus" className="form-control" >
                          Exempt
                      </option>
                      <option name="ExemptStatus" className="form-control" >
                          NonExempt
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateJobProfile}>Save</button>
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

export default UpdateJobProfileComponent
