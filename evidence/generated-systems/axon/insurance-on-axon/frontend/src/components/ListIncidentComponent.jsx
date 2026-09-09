import React, { Component } from 'react'
import IncidentService from '../services/IncidentService'

class ListIncidentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                incidents: []
        }
        this.addIncident = this.addIncident.bind(this);
        this.editIncident = this.editIncident.bind(this);
        this.deleteIncident = this.deleteIncident.bind(this);
    }

    deleteIncident(id){
        IncidentService.deleteIncident(id).then( res => {
            this.setState({incidents: this.state.incidents.filter(incident => incident.incidentId !== id)});
        });
    }
    viewIncident(id){
        this.props.history.push(`/view-incident/${id}`);
    }
    editIncident(id){
        this.props.history.push(`/add-incident/${id}`);
    }

    componentDidMount(){
        IncidentService.getIncidents().then((res) => {
            this.setState({ incidents: res.data});
        });
    }

    addIncident(){
        this.props.history.push('/add-incident/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Incident List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addIncident}> Add Incident</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Location </th>
                                    <th> Description </th>
                                    <th> IncidentType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.incidents.map(
                                        incident => 
                                        <tr key = {incident.incidentId}>
                                             <td> { incident.location } </td>
                                             <td> { incident.description } </td>
                                             <td> { incident.incidentType } </td>
                                             <td>
                                                 <button onClick={ () => this.editIncident(incident.incidentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteIncident(incident.incidentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewIncident(incident.incidentId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListIncidentComponent
