import React, { Component } from 'react'
import TrainingRunService from '../services/TrainingRunService';

class UpdateTrainingRunComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                runLabel: '',
                startedAt: '',
                completedAt: '',
                status: ''
        }
        this.updateTrainingRun = this.updateTrainingRun.bind(this);

        this.changerunLabelHandler = this.changerunLabelHandler.bind(this);
        this.changestartedAtHandler = this.changestartedAtHandler.bind(this);
        this.changecompletedAtHandler = this.changecompletedAtHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        TrainingRunService.getTrainingRunById(this.state.id).then( (res) =>{
            let trainingRun = res.data;
            this.setState({
                runLabel: trainingRun.runLabel,
                startedAt: trainingRun.startedAt,
                completedAt: trainingRun.completedAt,
                status: trainingRun.status
            });
        });
    }

    updateTrainingRun = (e) => {
        e.preventDefault();
        let trainingRun = {
            trainingRunId: this.state.id,
            runLabel: this.state.runLabel,
            startedAt: this.state.startedAt,
            completedAt: this.state.completedAt,
            status: this.state.status
        };
        console.log('trainingRun => ' + JSON.stringify(trainingRun));
        console.log('id => ' + JSON.stringify(this.state.id));
        TrainingRunService.updateTrainingRun(trainingRun).then( res => {
            this.props.history.push('/trainingRuns');
        });
    }

    changerunLabelHandler= (event) => {
        this.setState({runLabel: event.target.value});
    }
    changestartedAtHandler= (event) => {
        this.setState({startedAt: event.target.value});
    }
    changecompletedAtHandler= (event) => {
        this.setState({completedAt: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/trainingRuns');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update TrainingRun</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> runLabel: </label>
                                                <input placeholder="runLabel" name="runLabel" className="form-control" value={this.state.runLabel} onChange={this.changerunLabelHandler}/>

                                            <label> startedAt: </label>
                                                <input type="date" placeholder="startedAt" name="startedAt" className="form-control" value={this.state.startedAt} onChange={this.changestartedAtHandler}/>

                                            <label> completedAt: </label>
                                                <input type="date" placeholder="completedAt" name="completedAt" className="form-control" value={this.state.completedAt} onChange={this.changecompletedAtHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Queued
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
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateTrainingRun}>Save</button>
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

export default UpdateTrainingRunComponent
