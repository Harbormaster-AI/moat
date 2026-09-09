import React, { Component } from 'react'
import TimeEntryService from '../services/TimeEntryService';

class CreateTimeEntryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                entryDate: '',
                hoursWorked: '',
                entryType: ''
        }
        this.changeentryDateHandler = this.changeentryDateHandler.bind(this);
        this.changehoursWorkedHandler = this.changehoursWorkedHandler.bind(this);
        this.changeEntryTypeHandler = this.changeEntryTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            TimeEntryService.getTimeEntryById(this.state.id).then( (res) =>{
                let timeEntry = res.data;
                this.setState({
                    entryDate: timeEntry.entryDate,
                    hoursWorked: timeEntry.hoursWorked,
                    entryType: timeEntry.entryType
                });
            });
        }        
    }
    saveOrUpdateTimeEntry = (e) => {
        e.preventDefault();
        let timeEntry = {
                timeEntryId: this.state.id,
                entryDate: this.state.entryDate,
                hoursWorked: this.state.hoursWorked,
                entryType: this.state.entryType
            };
        console.log('timeEntry => ' + JSON.stringify(timeEntry));

        // step 5
        if(this.state.id === '_add'){
            timeEntry.timeEntryId=''
            TimeEntryService.createTimeEntry(timeEntry).then(res =>{
                this.props.history.push('/timeEntrys');
            });
        }else{
            TimeEntryService.updateTimeEntry(timeEntry).then( res => {
                this.props.history.push('/timeEntrys');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add TimeEntry</h3>
        }else{
            return <h3 className="text-center">Update TimeEntry</h3>
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
                                            <label> entryDate:&emsp; </label>
                                                <input type="date" placeholder="entryDate" name="entryDate" className="form-control" value={this.state.entryDate} onChange={this.changeentryDateHandler}/>

                                            <label> hoursWorked:&emsp; </label>
                                                <input placeholder="hoursWorked" name="hoursWorked" className="form-control" value={this.state.hoursWorked} onChange={this.changehoursWorkedHandler}/>

                                            <label> EntryType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateTimeEntry}>Save</button>
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

export default CreateTimeEntryComponent
