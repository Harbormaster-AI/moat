import React, { Component } from 'react'
import IncidentService from '../services/IncidentService';

class CreateIncidentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                location: '',
                description: '',
                incidentType: ''
        }
        this.changelocationHandler = this.changelocationHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changeIncidentTypeHandler = this.changeIncidentTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            IncidentService.getIncidentById(this.state.id).then( (res) =>{
                let incident = res.data;
                this.setState({
                    location: incident.location,
                    description: incident.description,
                    incidentType: incident.incidentType
                });
            });
        }        
    }
    saveOrUpdateIncident = (e) => {
        e.preventDefault();
        let incident = {
                incidentId: this.state.id,
                location: this.state.location,
                description: this.state.description,
                incidentType: this.state.incidentType
            };
        console.log('incident => ' + JSON.stringify(incident));

        // step 5
        if(this.state.id === '_add'){
            incident.incidentId=''
            IncidentService.createIncident(incident).then(res =>{
                this.props.history.push('/incidents');
            });
        }else{
            IncidentService.updateIncident(incident).then( res => {
                this.props.history.push('/incidents');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Incident</h3>
        }else{
            return <h3 className="text-center">Update Incident</h3>
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
                                            <label> location:&emsp; </label>
                                                <input placeholder="location" name="location" className="form-control" value={this.state.location} onChange={this.changelocationHandler}/>

                                            <label> description:&emsp; </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> IncidentType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateIncident}>Save</button>
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

export default CreateIncidentComponent
