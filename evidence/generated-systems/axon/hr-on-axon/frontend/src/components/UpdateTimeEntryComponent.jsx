import React, { Component } from 'react'
import TimeEntryService from '../services/TimeEntryService';

class UpdateTimeEntryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                entryDate: '',
                hoursWorked: '',
                entryType: ''
        }
        this.updateTimeEntry = this.updateTimeEntry.bind(this);

        this.changeentryDateHandler = this.changeentryDateHandler.bind(this);
        this.changehoursWorkedHandler = this.changehoursWorkedHandler.bind(this);
        this.changeEntryTypeHandler = this.changeEntryTypeHandler.bind(this);
    }

    componentDidMount(){
        TimeEntryService.getTimeEntryById(this.state.id).then( (res) =>{
            let timeEntry = res.data;
            this.setState({
                entryDate: timeEntry.entryDate,
                hoursWorked: timeEntry.hoursWorked,
                entryType: timeEntry.entryType
            });
        });
    }

    updateTimeEntry = (e) => {
        e.preventDefault();
        let timeEntry = {
            timeEntryId: this.state.id,
            entryDate: this.state.entryDate,
            hoursWorked: this.state.hoursWorked,
            entryType: this.state.entryType
        };
        console.log('timeEntry => ' + JSON.stringify(timeEntry));
        console.log('id => ' + JSON.stringify(this.state.id));
        TimeEntryService.updateTimeEntry(timeEntry).then( res => {
            this.props.history.push('/timeEntrys');
        });
    }

    changeentryDateHandler= (event) => {
        this.setState({entryDate: event.target.value});
    }
    changehoursWorkedHandler= (event) => {
        this.setState({hoursWorked: event.target.value});
    }
    changeEntryTypeHandler= (event) => {
        this.setState({entryType: event.target.value});
    }

    cancel(){
        this.props.history.push('/timeEntrys');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update TimeEntry</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> entryDate: </label>
                                                <input type="date" placeholder="entryDate" name="entryDate" className="form-control" value={this.state.entryDate} onChange={this.changeentryDateHandler}/>

                                            <label> hoursWorked: </label>
                                                <input placeholder="hoursWorked" name="hoursWorked" className="form-control" value={this.state.hoursWorked} onChange={this.changehoursWorkedHandler}/>

                                            <label> EntryType: </label>
                                                <select value={this.state.entryType} onChange={this.changeEntryTypeHandler}>
                      <option name="EntryType" className="form-control" >
                          Regular
                      </option>
                      <option name="EntryType" className="form-control" >
                          Overtime
                      </option>
                      <option name="EntryType" className="form-control" >
                          Sick
                      </option>
                      <option name="EntryType" className="form-control" >
                          Vacation
                      </option>
                      <option name="EntryType" className="form-control" >
                          Unpaid
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateTimeEntry}>Save</button>
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

export default UpdateTimeEntryComponent
