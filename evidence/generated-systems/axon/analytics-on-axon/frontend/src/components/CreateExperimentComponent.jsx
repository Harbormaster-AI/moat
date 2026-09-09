import React, { Component } from 'react'
import ExperimentService from '../services/ExperimentService';

class CreateExperimentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                objective: '',
                status: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeobjectiveHandler = this.changeobjectiveHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ExperimentService.getExperimentById(this.state.id).then( (res) =>{
                let experiment = res.data;
                this.setState({
                    name: experiment.name,
                    objective: experiment.objective,
                    status: experiment.status
                });
            });
        }        
    }
    saveOrUpdateExperiment = (e) => {
        e.preventDefault();
        let experiment = {
                experimentId: this.state.id,
                name: this.state.name,
                objective: this.state.objective,
                status: this.state.status
            };
        console.log('experiment => ' + JSON.stringify(experiment));

        // step 5
        if(this.state.id === '_add'){
            experiment.experimentId=''
            ExperimentService.createExperiment(experiment).then(res =>{
                this.props.history.push('/experiments');
            });
        }else{
            ExperimentService.updateExperiment(experiment).then( res => {
                this.props.history.push('/experiments');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeobjectiveHandler= (event) => {
        this.setState({objective: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/experiments');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Experiment</h3>
        }else{
            return <h3 className="text-center">Update Experiment</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> objective:&emsp; </label>
                                                <input placeholder="objective" name="objective" className="form-control" value={this.state.objective} onChange={this.changeobjectiveHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Planned
                      </option>
                      <option name="Status" className="form-control" >
                          Running
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          Failed
                      </option>
                      <option name="Status" className="form-control" >
                          Stopped
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateExperiment}>Save</button>
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

export default CreateExperimentComponent
