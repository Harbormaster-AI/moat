import React, { Component } from 'react'
import EncounterService from '../services/EncounterService';

class UpdateEncounterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                encounterNumber: '',
                startDateTime: '',
                endDateTime: '',
                status: '',
                encounterType: ''
        }
        this.updateEncounter = this.updateEncounter.bind(this);

        this.changeencounterNumberHandler = this.changeencounterNumberHandler.bind(this);
        this.changestartDateTimeHandler = this.changestartDateTimeHandler.bind(this);
        this.changeendDateTimeHandler = this.changeendDateTimeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changeEncounterTypeHandler = this.changeEncounterTypeHandler.bind(this);
    }

    componentDidMount(){
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

    updateEncounter = (e) => {
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
        console.log('id => ' + JSON.stringify(this.state.id));
        EncounterService.updateEncounter(encounter).then( res => {
            this.props.history.push('/encounters');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Encounter</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> encounterNumber: </label>
                                                <input placeholder="encounterNumber" name="encounterNumber" className="form-control" value={this.state.encounterNumber} onChange={this.changeencounterNumberHandler}/>

                                            <label> startDateTime: </label>
                                                <input type="time" placeholder="startDateTime" name="startDateTime" className="form-control" value={this.state.startDateTime} onChange={this.changestartDateTimeHandler}/>

                                            <label> endDateTime: </label>
                                                <input type="time" placeholder="endDateTime" name="endDateTime" className="form-control" value={this.state.endDateTime} onChange={this.changeendDateTimeHandler}/>

                                            <label> Status: </label>
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

                                            <label> EncounterType: </label>
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
                                        <button className="btn btn-success" onClick={this.updateEncounter}>Save</button>
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

export default UpdateEncounterComponent
