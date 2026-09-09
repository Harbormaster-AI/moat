import React, { Component } from 'react'
import DataPipelineService from '../services/DataPipelineService';

class UpdateDataPipelineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                schedule: '',
                triggerType: '',
                status: ''
        }
        this.updateDataPipeline = this.updateDataPipeline.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changescheduleHandler = this.changescheduleHandler.bind(this);
        this.changeTriggerTypeHandler = this.changeTriggerTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        DataPipelineService.getDataPipelineById(this.state.id).then( (res) =>{
            let dataPipeline = res.data;
            this.setState({
                name: dataPipeline.name,
                schedule: dataPipeline.schedule,
                triggerType: dataPipeline.triggerType,
                status: dataPipeline.status
            });
        });
    }

    updateDataPipeline = (e) => {
        e.preventDefault();
        let dataPipeline = {
            dataPipelineId: this.state.id,
            name: this.state.name,
            schedule: this.state.schedule,
            triggerType: this.state.triggerType,
            status: this.state.status
        };
        console.log('dataPipeline => ' + JSON.stringify(dataPipeline));
        console.log('id => ' + JSON.stringify(this.state.id));
        DataPipelineService.updateDataPipeline(dataPipeline).then( res => {
            this.props.history.push('/dataPipelines');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changescheduleHandler= (event) => {
        this.setState({schedule: event.target.value});
    }
    changeTriggerTypeHandler= (event) => {
        this.setState({triggerType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/dataPipelines');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update DataPipeline</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> schedule: </label>
                                                <input placeholder="schedule" name="schedule" className="form-control" value={this.state.schedule} onChange={this.changescheduleHandler}/>

                                            <label> TriggerType: </label>
                                                <select value={this.state.triggerType} onChange={this.changeTriggerTypeHandler}>
                      <option name="TriggerType" className="form-control" >
                          Manual
                      </option>
                      <option name="TriggerType" className="form-control" >
                          Schedule
                      </option>
                      <option name="TriggerType" className="form-control" >
                          Event
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Paused
                      </option>
                      <option name="Status" className="form-control" >
                          Failed
                      </option>
                      <option name="Status" className="form-control" >
                          Succeeded
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateDataPipeline}>Save</button>
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

export default UpdateDataPipelineComponent
