import React, { Component } from 'react'
import ExperimentService from '../services/ExperimentService';

class UpdateExperimentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                objective: '',
                status: ''
        }
        this.updateExperiment = this.updateExperiment.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeobjectiveHandler = this.changeobjectiveHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ExperimentService.getExperimentById(this.state.id).then( (res) =>{
            let experiment = res.data;
            this.setState({
                name: experiment.name,
                objective: experiment.objective,
                status: experiment.status
            });
        });
    }

    updateExperiment = (e) => {
        e.preventDefault();
        let experiment = {
            experimentId: this.state.id,
            name: this.state.name,
            objective: this.state.objective,
            status: this.state.status
        };
        console.log('experiment => ' + JSON.stringify(experiment));
        console.log('id => ' + JSON.stringify(this.state.id));
        ExperimentService.updateExperiment(experiment).then( res => {
            this.props.history.push('/experiments');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Experiment</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> objective: </label>
                                                <input placeholder="objective" name="objective" className="form-control" value={this.state.objective} onChange={this.changeobjectiveHandler}/>

                                            <label> Status: </label>
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
                                        <button className="btn btn-success" onClick={this.updateExperiment}>Save</button>
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

export default UpdateExperimentComponent
