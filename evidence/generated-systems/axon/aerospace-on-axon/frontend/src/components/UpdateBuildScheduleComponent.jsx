import React, { Component } from 'react'
import BuildScheduleService from '../services/BuildScheduleService';

class UpdateBuildScheduleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                scheduleNumber: '',
                status: ''
        }
        this.updateBuildSchedule = this.updateBuildSchedule.bind(this);

        this.changescheduleNumberHandler = this.changescheduleNumberHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        BuildScheduleService.getBuildScheduleById(this.state.id).then( (res) =>{
            let buildSchedule = res.data;
            this.setState({
                scheduleNumber: buildSchedule.scheduleNumber,
                status: buildSchedule.status
            });
        });
    }

    updateBuildSchedule = (e) => {
        e.preventDefault();
        let buildSchedule = {
            buildScheduleId: this.state.id,
            scheduleNumber: this.state.scheduleNumber,
            status: this.state.status
        };
        console.log('buildSchedule => ' + JSON.stringify(buildSchedule));
        console.log('id => ' + JSON.stringify(this.state.id));
        BuildScheduleService.updateBuildSchedule(buildSchedule).then( res => {
            this.props.history.push('/buildSchedules');
        });
    }

    changescheduleNumberHandler= (event) => {
        this.setState({scheduleNumber: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/buildSchedules');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update BuildSchedule</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> scheduleNumber: </label>
                                                <input placeholder="scheduleNumber" name="scheduleNumber" className="form-control" value={this.state.scheduleNumber} onChange={this.changescheduleNumberHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Published
                      </option>
                      <option name="Status" className="form-control" >
                          Revised
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateBuildSchedule}>Save</button>
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

export default UpdateBuildScheduleComponent
