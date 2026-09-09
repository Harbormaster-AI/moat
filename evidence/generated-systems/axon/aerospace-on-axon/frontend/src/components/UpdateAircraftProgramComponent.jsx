import React, { Component } from 'react'
import AircraftProgramService from '../services/AircraftProgramService';

class UpdateAircraftProgramComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                programCode: '',
                entryIntoServiceYear: '',
                status: ''
        }
        this.updateAircraftProgram = this.updateAircraftProgram.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeprogramCodeHandler = this.changeprogramCodeHandler.bind(this);
        this.changeentryIntoServiceYearHandler = this.changeentryIntoServiceYearHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        AircraftProgramService.getAircraftProgramById(this.state.id).then( (res) =>{
            let aircraftProgram = res.data;
            this.setState({
                name: aircraftProgram.name,
                programCode: aircraftProgram.programCode,
                entryIntoServiceYear: aircraftProgram.entryIntoServiceYear,
                status: aircraftProgram.status
            });
        });
    }

    updateAircraftProgram = (e) => {
        e.preventDefault();
        let aircraftProgram = {
            aircraftProgramId: this.state.id,
            name: this.state.name,
            programCode: this.state.programCode,
            entryIntoServiceYear: this.state.entryIntoServiceYear,
            status: this.state.status
        };
        console.log('aircraftProgram => ' + JSON.stringify(aircraftProgram));
        console.log('id => ' + JSON.stringify(this.state.id));
        AircraftProgramService.updateAircraftProgram(aircraftProgram).then( res => {
            this.props.history.push('/aircraftPrograms');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeprogramCodeHandler= (event) => {
        this.setState({programCode: event.target.value});
    }
    changeentryIntoServiceYearHandler= (event) => {
        this.setState({entryIntoServiceYear: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/aircraftPrograms');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update AircraftProgram</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> programCode: </label>
                                                <input placeholder="programCode" name="programCode" className="form-control" value={this.state.programCode} onChange={this.changeprogramCodeHandler}/>

                                            <label> entryIntoServiceYear: </label>
                                                <input type="number" placeholder="entryIntoServiceYear" name="entryIntoServiceYear" className="form-control" value={this.state.entryIntoServiceYear} onChange={this.changeentryIntoServiceYearHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Concept
                      </option>
                      <option name="Status" className="form-control" >
                          Development
                      </option>
                      <option name="Status" className="form-control" >
                          Certification
                      </option>
                      <option name="Status" className="form-control" >
                          Production
                      </option>
                      <option name="Status" className="form-control" >
                          InService
                      </option>
                      <option name="Status" className="form-control" >
                          Sunset
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateAircraftProgram}>Save</button>
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

export default UpdateAircraftProgramComponent
