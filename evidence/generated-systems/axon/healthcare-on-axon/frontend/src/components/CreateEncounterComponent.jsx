import React, { Component } from 'react'
import EncounterService from '../services/EncounterService';

class CreateEncounterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                encounterNumber: '',
                startDateTime: '',
                endDateTime: '',
                status: '',
                encounterType: ''
        }
        this.changeencounterNumberHandler = this.changeencounterNumberHandler.bind(this);
        this.changestartDateTimeHandler = this.changestartDateTimeHandler.bind(this);
        this.changeendDateTimeHandler = this.changeendDateTimeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changeEncounterTypeHandler = this.changeEncounterTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            EncounterService.getEncounterById(this.state.id).then( (res) =>{
                let encounter = res.data;
                this.setState({
                    encounterNumber: encounter.encounterNumber,
                    startDateTime: encounter.startDateTime,
                    endDateTime: encounter.endDateTime,
                    status: encounter.status,
                    encounterType: encounter.encounterType
                });
            });
        }        
    }
    saveOrUpdateEncounter = (e) => {
        e.preventDefault();
        let encounter = {
                encounterId: this.state.id,
                encounterNumber: this.state.encounterNumber,
                startDateTime: this.state.startDateTime,
                endDateTime: this.state.endDateTime,
                status: this.state.status,
                encounterType: this.state.encounterType
            };
        console.log('encounter => ' + JSON.stringify(encounter));

        // step 5
        if(this.state.id === '_add'){
            encounter.encounterId=''
            EncounterService.createEncounter(encounter).then(res =>{
                this.props.history.push('/encounters');
            });
        }else{
            EncounterService.updateEncounter(encounter).then( res => {
                this.props.history.push('/encounters');
            });
        }
    }
    
    changeencounterNumberHandler= (event) => {
        this.setState({encounterNumber: event.target.value});
    }
    changestartDateTimeHandler= (event) => {
        this.setState({startDateTime: event.target.value});
    }
    changeendDateTimeHandler= (event) => {
        this.setState({endDateTime: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changeEncounterTypeHandler= (event) => {
        this.setState({encounterType: event.target.value});
    }

    cancel(){
        this.props.history.push('/encounters');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Encounter</h3>
        }else{
            return <h3 className="text-center">Update Encounter</h3>
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
                                            <label> encounterNumber:&emsp; </label>
                                                <input placeholder="encounterNumber" name="encounterNumber" className="form-control" value={this.state.encounterNumber} onChange={this.changeencounterNumberHandler}/>

                                            <label> startDateTime:&emsp; </label>
                                                <input type="time" placeholder="startDateTime" name="startDateTime" className="form-control" value={this.state.startDateTime} onChange={this.changestartDateTimeHandler}/>

                                            <label> endDateTime:&emsp; </label>
                                                <input type="time" placeholder="endDateTime" name="endDateTime" className="form-control" value={this.state.endDateTime} onChange={this.changeendDateTimeHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Planned
                      </option>
                      <option name="Status" className="form-control" >
                          InProgress
                      </option>
                      <option name="Status" className="form-control" >
                          OnHold
                      </option>
                      <option name="Status" className="form-control" >
                          Discharged
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                            <label> EncounterType:&emsp; </label>
                                                <select value={this.state.encounterType} onChange={this.changeEncounterTypeHandler}>
                      <option name="EncounterType" className="form-control" >
                          Inpatient
                      </option>
                      <option name="EncounterType" className="form-control" >
                          Outpatient
                      </option>
                      <option name="EncounterType" className="form-control" >
                          Emergency
                      </option>
                      <option name="EncounterType" className="form-control" >
                          Observation
                      </option>
                      <option name="EncounterType" className="form-control" >
                          Telemedicine
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateEncounter}>Save</button>
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

export default CreateEncounterComponent
