import React, { Component } from 'react'
import AircraftFamilyService from '../services/AircraftFamilyService'

class ListAircraftFamilyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                aircraftFamilys: []
        }
        this.addAircraftFamily = this.addAircraftFamily.bind(this);
        this.editAircraftFamily = this.editAircraftFamily.bind(this);
        this.deleteAircraftFamily = this.deleteAircraftFamily.bind(this);
    }

    deleteAircraftFamily(id){
        AircraftFamilyService.deleteAircraftFamily(id).then( res => {
            this.setState({aircraftFamilys: this.state.aircraftFamilys.filter(aircraftFamily => aircraftFamily.aircraftFamilyId !== id)});
        });
    }
    viewAircraftFamily(id){
        this.props.history.push(`/view-aircraftFamily/${id}`);
    }
    editAircraftFamily(id){
        this.props.history.push(`/add-aircraftFamily/${id}`);
    }

    componentDidMount(){
        AircraftFamilyService.getAircraftFamilys().then((res) => {
            this.setState({ aircraftFamilys: res.data});
        });
    }

    addAircraftFamily(){
        this.props.history.push('/add-aircraftFamily/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">AircraftFamily List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAircraftFamily}> Add AircraftFamily</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> FamilyCode </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.aircraftFamilys.map(
                                        aircraftFamily => 
                                        <tr key = {aircraftFamily.aircraftFamilyId}>
                                             <td> { aircraftFamily.name } </td>
                                             <td> { aircraftFamily.familyCode } </td>
                                             <td>
                                                 <button onClick={ () => this.editAircraftFamily(aircraftFamily.aircraftFamilyId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAircraftFamily(aircraftFamily.aircraftFamilyId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAircraftFamily(aircraftFamily.aircraftFamilyId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAircraftFamilyComponent
