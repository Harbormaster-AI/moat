import React, { Component } from 'react'
import ScheduleExceptionService from '../services/ScheduleExceptionService';

class UpdateScheduleExceptionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                date: '',
                reason: '',
                hours: ''
        }
        this.updateScheduleException = this.updateScheduleException.bind(this);

        this.changedateHandler = this.changedateHandler.bind(this);
        this.changereasonHandler = this.changereasonHandler.bind(this);
        this.changehoursHandler = this.changehoursHandler.bind(this);
    }

    componentDidMount(){
        ScheduleExceptionService.getScheduleExceptionById(this.state.id).then( (res) =>{
            let scheduleException = res.data;
            this.setState({
                date: scheduleException.date,
                reason: scheduleException.reason,
                hours: scheduleException.hours
            });
        });
    }

    updateScheduleException = (e) => {
        e.preventDefault();
        let scheduleException = {
            scheduleExceptionId: this.state.id,
            date: this.state.date,
            reason: this.state.reason,
            hours: this.state.hours
        };
        console.log('scheduleException => ' + JSON.stringify(scheduleException));
        console.log('id => ' + JSON.stringify(this.state.id));
        ScheduleExceptionService.updateScheduleException(scheduleException).then( res => {
            this.props.history.push('/scheduleExceptions');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ScheduleException</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> date: </label>
                                                <input type="date" placeholder="date" name="date" className="form-control" value={this.state.date} onChange={this.changedateHandler}/>

                                            <label> reason: </label>
                                                <input placeholder="reason" name="reason" className="form-control" value={this.state.reason} onChange={this.changereasonHandler}/>

                                            <label> hours: </label>
                                                <input placeholder="hours" name="hours" className="form-control" value={this.state.hours} onChange={this.changehoursHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateScheduleException}>Save</button>
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

export default UpdateScheduleExceptionComponent
