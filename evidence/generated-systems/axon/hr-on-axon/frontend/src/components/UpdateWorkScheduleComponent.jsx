import React, { Component } from 'react'
import WorkScheduleService from '../services/WorkScheduleService';

class UpdateWorkScheduleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                standardHoursPerWeek: '',
                scheduleType: ''
        }
        this.updateWorkSchedule = this.updateWorkSchedule.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changestandardHoursPerWeekHandler = this.changestandardHoursPerWeekHandler.bind(this);
        this.changeScheduleTypeHandler = this.changeScheduleTypeHandler.bind(this);
    }

    componentDidMount(){
        WorkScheduleService.getWorkScheduleById(this.state.id).then( (res) =>{
            let workSchedule = res.data;
            this.setState({
                name: workSchedule.name,
                standardHoursPerWeek: workSchedule.standardHoursPerWeek,
                scheduleType: workSchedule.scheduleType
            });
        });
    }

    updateWorkSchedule = (e) => {
        e.preventDefault();
        let workSchedule = {
            workScheduleId: this.state.id,
            name: this.state.name,
            standardHoursPerWeek: this.state.standardHoursPerWeek,
            scheduleType: this.state.scheduleType
        };
        console.log('workSchedule => ' + JSON.stringify(workSchedule));
        console.log('id => ' + JSON.stringify(this.state.id));
        WorkScheduleService.updateWorkSchedule(workSchedule).then( res => {
            this.props.history.push('/workSchedules');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changestandardHoursPerWeekHandler= (event) => {
        this.setState({standardHoursPerWeek: event.target.value});
    }
    changeScheduleTypeHandler= (event) => {
        this.setState({scheduleType: event.target.value});
    }

    cancel(){
        this.props.history.push('/workSchedules');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update WorkSchedule</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> standardHoursPerWeek: </label>
                                                <input placeholder="standardHoursPerWeek" name="standardHoursPerWeek" className="form-control" value={this.state.standardHoursPerWeek} onChange={this.changestandardHoursPerWeekHandler}/>

                                            <label> ScheduleType: </label>
                                                <select value={this.state.scheduleType} onChange={this.changeScheduleTypeHandler}>
                      <option name="ScheduleType" className="form-control" >
                          Fixed
                      </option>
                      <option name="ScheduleType" className="form-control" >
                          Flexible
                      </option>
                      <option name="ScheduleType" className="form-control" >
                          Rotating
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateWorkSchedule}>Save</button>
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

export default UpdateWorkScheduleComponent
