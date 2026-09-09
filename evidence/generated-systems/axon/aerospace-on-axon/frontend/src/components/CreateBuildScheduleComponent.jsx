import React, { Component } from 'react'
import BuildScheduleService from '../services/BuildScheduleService';

class CreateBuildScheduleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                scheduleNumber: '',
                status: ''
        }
        this.changescheduleNumberHandler = this.changescheduleNumberHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            BuildScheduleService.getBuildScheduleById(this.state.id).then( (res) =>{
                let buildSchedule = res.data;
                this.setState({
                    scheduleNumber: buildSchedule.scheduleNumber,
                    status: buildSchedule.status
                });
            });
        }        
    }
    saveOrUpdateBuildSchedule = (e) => {
        e.preventDefault();
        let buildSchedule = {
                buildScheduleId: this.state.id,
                scheduleNumber: this.state.scheduleNumber,
                status: this.state.status
            };
        console.log('buildSchedule => ' + JSON.stringify(buildSchedule));

        // step 5
        if(this.state.id === '_add'){
            buildSchedule.buildScheduleId=''
            BuildScheduleService.createBuildSchedule(buildSchedule).then(res =>{
                this.props.history.push('/buildSchedules');
            });
        }else{
            BuildScheduleService.updateBuildSchedule(buildSchedule).then( res => {
                this.props.history.push('/buildSchedules');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add BuildSchedule</h3>
        }else{
            return <h3 className="text-center">Update BuildSchedule</h3>
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
                                            <label> scheduleNumber:&emsp; </label>
                                                <input placeholder="scheduleNumber" name="scheduleNumber" className="form-control" value={this.state.scheduleNumber} onChange={this.changescheduleNumberHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateBuildSchedule}>Save</button>
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

export default CreateBuildScheduleComponent
