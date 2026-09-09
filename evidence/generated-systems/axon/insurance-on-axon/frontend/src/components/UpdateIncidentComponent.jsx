import React, { Component } from 'react'
import IncidentService from '../services/IncidentService';

class UpdateIncidentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                location: '',
                description: '',
                incidentType: ''
        }
        this.updateIncident = this.updateIncident.bind(this);

        this.changelocationHandler = this.changelocationHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changeIncidentTypeHandler = this.changeIncidentTypeHandler.bind(this);
    }

    componentDidMount(){
        IncidentService.getIncidentById(this.state.id).then( (res) =>{
            let incident = res.data;
            this.setState({
                location: incident.location,
                description: incident.description,
                incidentType: incident.incidentType
            });
        });
    }

    updateIncident = (e) => {
        e.preventDefault();
        let incident = {
            incidentId: this.state.id,
            location: this.state.location,
            description: this.state.description,
            incidentType: this.state.incidentType
        };
        console.log('incident => ' + JSON.stringify(incident));
        console.log('id => ' + JSON.stringify(this.state.id));
        IncidentService.updateIncident(incident).then( res => {
            this.props.history.push('/incidents');
        });
    }

    changelocationHandler= (event) => {
        this.setState({location: event.target.value});
    }
    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }
    changeIncidentTypeHandler= (event) => {
        this.setState({incidentType: event.target.value});
    }

    cancel(){
        this.props.history.push('/incidents');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Incident</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> location: </label>
                                                <input placeholder="location" name="location" className="form-control" value={this.state.location} onChange={this.changelocationHandler}/>

                                            <label> description: </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> IncidentType: </label>
                                                <select value={this.state.incidentType} onChange={this.changeIncidentTypeHandler}>
                      <option name="IncidentType" className="form-control" >
                          AutoAccident
                      </option>
                      <option name="IncidentType" className="form-control" >
                          Fire
                      </option>
                      <option name="IncidentType" className="form-control" >
                          Theft
                      </option>
                      <option name="IncidentType" className="form-control" >
                          Windstorm
                      </option>
                      <option name="IncidentType" className="form-control" >
                          Flood
                      </option>
                      <option name="IncidentType" className="form-control" >
                          Hail
                      </option>
                      <option name="IncidentType" className="form-control" >
                          Earthquake
                      </option>
                      <option name="IncidentType" className="form-control" >
                          Vandalism
                      </option>
                      <option name="IncidentType" className="form-control" >
                          Injury
                      </option>
                      <option name="IncidentType" className="form-control" >
                          Death
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateIncident}>Save</button>
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

export default UpdateIncidentComponent
