import React, { Component } from 'react'
import ObservationService from '../services/ObservationService'

class ListObservationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                observations: []
        }
        this.addObservation = this.addObservation.bind(this);
        this.editObservation = this.editObservation.bind(this);
        this.deleteObservation = this.deleteObservation.bind(this);
    }

    deleteObservation(id){
        ObservationService.deleteObservation(id).then( res => {
            this.setState({observations: this.state.observations.filter(observation => observation.observationId !== id)});
        });
    }
    viewObservation(id){
        this.props.history.push(`/view-observation/${id}`);
    }
    editObservation(id){
        this.props.history.push(`/add-observation/${id}`);
    }

    componentDidMount(){
        ObservationService.getObservations().then((res) => {
            this.setState({ observations: res.data});
        });
    }

    addObservation(){
        this.props.history.push('/add-observation/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Observation List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addObservation}> Add Observation</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Code </th>
                                    <th> Value </th>
                                    <th> Unit </th>
                                    <th> EffectiveDateTime </th>
                                    <th> Interpretation </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.observations.map(
                                        observation => 
                                        <tr key = {observation.observationId}>
                                             <td> { observation.code } </td>
                                             <td> { observation.value } </td>
                                             <td> { observation.unit } </td>
                                             <td> { observation.effectiveDateTime } </td>
                                             <td> { observation.interpretation } </td>
                                             <td>
                                                 <button onClick={ () => this.editObservation(observation.observationId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteObservation(observation.observationId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewObservation(observation.observationId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListObservationComponent
