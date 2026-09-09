import React, { Component } from 'react'
import ScheduleExceptionService from '../services/ScheduleExceptionService';

class CreateScheduleExceptionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                date: '',
                reason: '',
                hours: ''
        }
        this.changedateHandler = this.changedateHandler.bind(this);
        this.changereasonHandler = this.changereasonHandler.bind(this);
        this.changehoursHandler = this.changehoursHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ScheduleExceptionService.getScheduleExceptionById(this.state.id).then( (res) =>{
                let scheduleException = res.data;
                this.setState({
                    date: scheduleException.date,
                    reason: scheduleException.reason,
                    hours: scheduleException.hours
                });
            });
        }        
    }
    saveOrUpdateScheduleException = (e) => {
        e.preventDefault();
        let scheduleException = {
                scheduleExceptionId: this.state.id,
                date: this.state.date,
                reason: this.state.reason,
                hours: this.state.hours
            };
        console.log('scheduleException => ' + JSON.stringify(scheduleException));

        // step 5
        if(this.state.id === '_add'){
            scheduleException.scheduleExceptionId=''
            ScheduleExceptionService.createScheduleException(scheduleException).then(res =>{
                this.props.history.push('/scheduleExceptions');
            });
        }else{
            ScheduleExceptionService.updateScheduleException(scheduleException).then( res => {
                this.props.history.push('/scheduleExceptions');
            });
        }
    }
    
    changedateHandler= (event) => {
        this.setState({date: event.target.value});
    }
    changereasonHandler= (event) => {
        this.setState({reason: event.target.value});
    }
    changehoursHandler= (event) => {
        this.setState({hours: event.target.value});
    }

    cancel(){
        this.props.history.push('/scheduleExceptions');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ScheduleException</h3>
        }else{
            return <h3 className="text-center">Update ScheduleException</h3>
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
                                            <label> date:&emsp; </label>
                                                <input type="date" placeholder="date" name="date" className="form-control" value={this.state.date} onChange={this.changedateHandler}/>

                                            <label> reason:&emsp; </label>
                                                <input placeholder="reason" name="reason" className="form-control" value={this.state.reason} onChange={this.changereasonHandler}/>

                                            <label> hours:&emsp; </label>
                                                <input placeholder="hours" name="hours" className="form-control" value={this.state.hours} onChange={this.changehoursHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateScheduleException}>Save</button>
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

export default CreateScheduleExceptionComponent
