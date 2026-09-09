import React, { Component } from 'react'
import WorkShiftService from '../services/WorkShiftService';

class CreateWorkShiftComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                startTime: '',
                endTime: '',
                breakMinutes: '',
                dayOfWeek: ''
        }
        this.changestartTimeHandler = this.changestartTimeHandler.bind(this);
        this.changeendTimeHandler = this.changeendTimeHandler.bind(this);
        this.changebreakMinutesHandler = this.changebreakMinutesHandler.bind(this);
        this.changeDayOfWeekHandler = this.changeDayOfWeekHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateWorkShift = (e) => {
        e.preventDefault();
        let workShift = {
                workShiftId: this.state.id,
                startTime: this.state.startTime,
                endTime: this.state.endTime,
                breakMinutes: this.state.breakMinutes,
                dayOfWeek: this.state.dayOfWeek
            };
        console.log('workShift => ' + JSON.stringify(workShift));

        // step 5
        if(this.state.id === '_add'){
            workShift.workShiftId=''
            WorkShiftService.createWorkShift(workShift).then(res =>{
                this.props.history.push('/workShifts');
            });
        }else{
            WorkShiftService.updateWorkShift(workShift).then( res => {
                this.props.history.push('/workShifts');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add WorkShift</h3>
        }else{
            return <h3 className="text-center">Update WorkShift</h3>
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
                                            <label> startTime:&emsp; </label>
                                                <input type="time" placeholder="startTime" name="startTime" className="form-control" value={this.state.startTime} onChange={this.changestartTimeHandler}/>

                                            <label> endTime:&emsp; </label>
                                                <input type="time" placeholder="endTime" name="endTime" className="form-control" value={this.state.endTime} onChange={this.changeendTimeHandler}/>

                                            <label> breakMinutes:&emsp; </label>
                                                <input type="number" placeholder="breakMinutes" name="breakMinutes" className="form-control" value={this.state.breakMinutes} onChange={this.changebreakMinutesHandler}/>

                                            <label> DayOfWeek:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateWorkShift}>Save</button>
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

export default CreateWorkShiftComponent
