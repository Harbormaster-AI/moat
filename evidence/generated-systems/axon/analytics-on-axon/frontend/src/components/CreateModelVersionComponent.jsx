import React, { Component } from 'react'
import ModelVersionService from '../services/ModelVersionService';

class CreateModelVersionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                version: '',
                lifecycle: '',
                trainingStatus: ''
        }
        this.changeversionHandler = this.changeversionHandler.bind(this);
        this.changeLifecycleHandler = this.changeLifecycleHandler.bind(this);
        this.changeTrainingStatusHandler = this.changeTrainingStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ModelVersionService.getModelVersionById(this.state.id).then( (res) =>{
                let modelVersion = res.data;
                this.setState({
                    version: modelVersion.version,
                    lifecycle: modelVersion.lifecycle,
                    trainingStatus: modelVersion.trainingStatus
                });
            });
        }        
    }
    saveOrUpdateModelVersion = (e) => {
        e.preventDefault();
        let modelVersion = {
                modelVersionId: this.state.id,
                version: this.state.version,
                lifecycle: this.state.lifecycle,
                trainingStatus: this.state.trainingStatus
            };
        console.log('modelVersion => ' + JSON.stringify(modelVersion));

        // step 5
        if(this.state.id === '_add'){
            modelVersion.modelVersionId=''
            ModelVersionService.createModelVersion(modelVersion).then(res =>{
                this.props.history.push('/modelVersions');
            });
        }else{
            ModelVersionService.updateModelVersion(modelVersion).then( res => {
                this.props.history.push('/modelVersions');
            });
        }
    }
    
    changeversionHandler= (event) => {
        this.setState({version: event.target.value});
    }
    changeLifecycleHandler= (event) => {
        this.setState({lifecycle: event.target.value});
    }
    changeTrainingStatusHandler= (event) => {
        this.setState({trainingStatus: event.target.value});
    }

    cancel(){
        this.props.history.push('/modelVersions');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ModelVersion</h3>
        }else{
            return <h3 className="text-center">Update ModelVersion</h3>
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
                                            <label> version:&emsp; </label>
                                                <input placeholder="version" name="version" className="form-control" value={this.state.version} onChange={this.changeversionHandler}/>

                                            <label> Lifecycle:&emsp; </label>
                                                <select value={this.state.lifecycle} onChange={this.changeLifecycleHandler}>
                      <option name="Lifecycle" className="form-control" >
                          Draft
                      </option>
                      <option name="Lifecycle" className="form-control" >
                          Staging
                      </option>
                      <option name="Lifecycle" className="form-control" >
                          Production
                      </option>
                      <option name="Lifecycle" className="form-control" >
                          Archived
                      </option>
                    </select>

                                            <label> TrainingStatus:&emsp; </label>
                                                <select value={this.state.trainingStatus} onChange={this.changeTrainingStatusHandler}>
                      <option name="TrainingStatus" className="form-control" >
                          Queued
                      </option>
                      <option name="TrainingStatus" className="form-control" >
                          Running
                      </option>
                      <option name="TrainingStatus" className="form-control" >
                          Completed
                      </option>
                      <option name="TrainingStatus" className="form-control" >
                          Failed
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateModelVersion}>Save</button>
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

export default CreateModelVersionComponent
