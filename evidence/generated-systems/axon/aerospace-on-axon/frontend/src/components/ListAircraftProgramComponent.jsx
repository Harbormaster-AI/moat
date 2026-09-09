import React, { Component } from 'react'
import AircraftProgramService from '../services/AircraftProgramService'

class ListAircraftProgramComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                aircraftPrograms: []
        }
        this.addAircraftProgram = this.addAircraftProgram.bind(this);
        this.editAircraftProgram = this.editAircraftProgram.bind(this);
        this.deleteAircraftProgram = this.deleteAircraftProgram.bind(this);
    }

    deleteAircraftProgram(id){
        AircraftProgramService.deleteAircraftProgram(id).then( res => {
            this.setState({aircraftPrograms: this.state.aircraftPrograms.filter(aircraftProgram => aircraftProgram.aircraftProgramId !== id)});
        });
    }
    viewAircraftProgram(id){
        this.props.history.push(`/view-aircraftProgram/${id}`);
    }
    editAircraftProgram(id){
        this.props.history.push(`/add-aircraftProgram/${id}`);
    }

    componentDidMount(){
        AircraftProgramService.getAircraftPrograms().then((res) => {
            this.setState({ aircraftPrograms: res.data});
        });
    }

    addAircraftProgram(){
        this.props.history.push('/add-aircraftProgram/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">AircraftProgram List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAircraftProgram}> Add AircraftProgram</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> ProgramCode </th>
                                    <th> EntryIntoServiceYear </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.aircraftPrograms.map(
                                        aircraftProgram => 
                                        <tr key = {aircraftProgram.aircraftProgramId}>
                                             <td> { aircraftProgram.name } </td>
                                             <td> { aircraftProgram.programCode } </td>
                                             <td> { aircraftProgram.entryIntoServiceYear } </td>
                                             <td> { aircraftProgram.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editAircraftProgram(aircraftProgram.aircraftProgramId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAircraftProgram(aircraftProgram.aircraftProgramId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAircraftProgram(aircraftProgram.aircraftProgramId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAircraftProgramComponent
