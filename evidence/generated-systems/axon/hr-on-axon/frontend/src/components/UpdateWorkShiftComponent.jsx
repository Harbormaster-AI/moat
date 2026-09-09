import React, { Component } from 'react'
import WorkShiftService from '../services/WorkShiftService';

class UpdateWorkShiftComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                startTime: '',
                endTime: '',
                breakMinutes: '',
                dayOfWeek: ''
        }
        this.updateWorkShift = this.updateWorkShift.bind(this);

        this.changestartTimeHandler = this.changestartTimeHandler.bind(this);
        this.changeendTimeHandler = this.changeendTimeHandler.bind(this);
        this.changebreakMinutesHandler = this.changebreakMinutesHandler.bind(this);
        this.changeDayOfWeekHandler = this.changeDayOfWeekHandler.bind(this);
    }

    componentDidMount(){
        WorkShiftService.getWorkShiftById(this.state.id).then( (res) =>{
            let workShift = res.data;
            this.setState({
                startTime: workShift.startTime,
                endTime: workShift.endTime,
                breakMinutes: workShift.breakMinutes,
                dayOfWeek: workShift.dayOfWeek
            });
        });
    }

    updateWorkShift = (e) => {
        e.preventDefault();
        let workShift = {
            workShiftId: this.state.id,
            startTime: this.state.startTime,
            endTime: this.state.endTime,
            breakMinutes: this.state.breakMinutes,
            dayOfWeek: this.state.dayOfWeek
        };
        console.log('workShift => ' + JSON.stringify(workShift));
        console.log('id => ' + JSON.stringify(this.state.id));
        WorkShiftService.updateWorkShift(workShift).then( res => {
            this.props.history.push('/workShifts');
        });
    }

    changestartTimeHandler= (event) => {
        this.setState({startTime: event.target.value});
    }
    changeendTimeHandler= (event) => {
        this.setState({endTime: event.target.value});
    }
    changebreakMinutesHandler= (event) => {
        this.setState({breakMinutes: event.target.value});
    }
    changeDayOfWeekHandler= (event) => {
        this.setState({dayOfWeek: event.target.value});
    }

    cancel(){
        this.props.history.push('/workShifts');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update WorkShift</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> startTime: </label>
                                                <input type="time" placeholder="startTime" name="startTime" className="form-control" value={this.state.startTime} onChange={this.changestartTimeHandler}/>

                                            <label> endTime: </label>
                                                <input type="time" placeholder="endTime" name="endTime" className="form-control" value={this.state.endTime} onChange={this.changeendTimeHandler}/>

                                            <label> breakMinutes: </label>
                                                <input type="number" placeholder="breakMinutes" name="breakMinutes" className="form-control" value={this.state.breakMinutes} onChange={this.changebreakMinutesHandler}/>

                                            <label> DayOfWeek: </label>
                                                <select value={this.state.dayOfWeek} onChange={this.changeDayOfWeekHandler}>
                      <option name="DayOfWeek" className="form-control" >
                          Monday
                      </option>
                      <option name="DayOfWeek" className="form-control" >
                          Tuesday
                      </option>
                      <option name="DayOfWeek" className="form-control" >
                          Wednesday
                      </option>
                      <option name="DayOfWeek" className="form-control" >
                          Thursday
                      </option>
                      <option name="DayOfWeek" className="form-control" >
                          Friday
                      </option>
                      <option name="DayOfWeek" className="form-control" >
                          Saturday
                      </option>
                      <option name="DayOfWeek" className="form-control" >
                          Sunday
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateWorkShift}>Save</button>
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

export default UpdateWorkShiftComponent
