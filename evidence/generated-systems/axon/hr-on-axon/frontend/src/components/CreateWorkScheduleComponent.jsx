import React, { Component } from 'react'
import WorkScheduleService from '../services/WorkScheduleService';

class CreateWorkScheduleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                standardHoursPerWeek: '',
                scheduleType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changestandardHoursPerWeekHandler = this.changestandardHoursPerWeekHandler.bind(this);
        this.changeScheduleTypeHandler = this.changeScheduleTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            WorkScheduleService.getWorkScheduleById(this.state.id).then( (res) =>{
                let workSchedule = res.data;
                this.setState({
                    name: workSchedule.name,
                    standardHoursPerWeek: workSchedule.standardHoursPerWeek,
                    scheduleType: workSchedule.scheduleType
                });
            });
        }        
    }
    saveOrUpdateWorkSchedule = (e) => {
        e.preventDefault();
        let workSchedule = {
                workScheduleId: this.state.id,
                name: this.state.name,
                standardHoursPerWeek: this.state.standardHoursPerWeek,
                scheduleType: this.state.scheduleType
            };
        console.log('workSchedule => ' + JSON.stringify(workSchedule));

        // step 5
        if(this.state.id === '_add'){
            workSchedule.workScheduleId=''
            WorkScheduleService.createWorkSchedule(workSchedule).then(res =>{
                this.props.history.push('/workSchedules');
            });
        }else{
            WorkScheduleService.updateWorkSchedule(workSchedule).then( res => {
                this.props.history.push('/workSchedules');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add WorkSchedule</h3>
        }else{
            return <h3 className="text-center">Update WorkSchedule</h3>
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

                                            <label> standardHoursPerWeek:&emsp; </label>
                                                <input placeholder="standardHoursPerWeek" name="standardHoursPerWeek" className="form-control" value={this.state.standardHoursPerWeek} onChange={this.changestandardHoursPerWeekHandler}/>

                                            <label> ScheduleType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateWorkSchedule}>Save</button>
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

export default CreateWorkScheduleComponent
