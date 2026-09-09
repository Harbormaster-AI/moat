import React, { Component } from 'react'
import EncounterService from '../services/EncounterService'

class ListEncounterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                encounters: []
        }
        this.addEncounter = this.addEncounter.bind(this);
        this.editEncounter = this.editEncounter.bind(this);
        this.deleteEncounter = this.deleteEncounter.bind(this);
    }

    deleteEncounter(id){
        EncounterService.deleteEncounter(id).then( res => {
            this.setState({encounters: this.state.encounters.filter(encounter => encounter.encounterId !== id)});
        });
    }
    viewEncounter(id){
        this.props.history.push(`/view-encounter/${id}`);
    }
    editEncounter(id){
        this.props.history.push(`/add-encounter/${id}`);
    }

    componentDidMount(){
        EncounterService.getEncounters().then((res) => {
            this.setState({ encounters: res.data});
        });
    }

    addEncounter(){
        this.props.history.push('/add-encounter/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Encounter List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addEncounter}> Add Encounter</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> EncounterNumber </th>
                                    <th> StartDateTime </th>
                                    <th> EndDateTime </th>
                                    <th> Status </th>
                                    <th> EncounterType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.encounters.map(
                                        encounter => 
                                        <tr key = {encounter.encounterId}>
                                             <td> { encounter.encounterNumber } </td>
                                             <td> { encounter.startDateTime } </td>
                                             <td> { encounter.endDateTime } </td>
                                             <td> { encounter.status } </td>
                                             <td> { encounter.encounterType } </td>
                                             <td>
                                                 <button onClick={ () => this.editEncounter(encounter.encounterId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteEncounter(encounter.encounterId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewEncounter(encounter.encounterId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListEncounterComponent
