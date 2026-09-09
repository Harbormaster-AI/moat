import React, { Component } from 'react'
import AircraftModelService from '../services/AircraftModelService'

class ListAircraftModelComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                aircraftModels: []
        }
        this.addAircraftModel = this.addAircraftModel.bind(this);
        this.editAircraftModel = this.editAircraftModel.bind(this);
        this.deleteAircraftModel = this.deleteAircraftModel.bind(this);
    }

    deleteAircraftModel(id){
        AircraftModelService.deleteAircraftModel(id).then( res => {
            this.setState({aircraftModels: this.state.aircraftModels.filter(aircraftModel => aircraftModel.aircraftModelId !== id)});
        });
    }
    viewAircraftModel(id){
        this.props.history.push(`/view-aircraftModel/${id}`);
    }
    editAircraftModel(id){
        this.props.history.push(`/add-aircraftModel/${id}`);
    }

    componentDidMount(){
        AircraftModelService.getAircraftModels().then((res) => {
            this.setState({ aircraftModels: res.data});
        });
    }

    addAircraftModel(){
        this.props.history.push('/add-aircraftModel/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">AircraftModel List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAircraftModel}> Add AircraftModel</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> ModelDesignation </th>
                                    <th> AircraftType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.aircraftModels.map(
                                        aircraftModel => 
                                        <tr key = {aircraftModel.aircraftModelId}>
                                             <td> { aircraftModel.name } </td>
                                             <td> { aircraftModel.modelDesignation } </td>
                                             <td> { aircraftModel.aircraftType } </td>
                                             <td>
                                                 <button onClick={ () => this.editAircraftModel(aircraftModel.aircraftModelId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAircraftModel(aircraftModel.aircraftModelId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAircraftModel(aircraftModel.aircraftModelId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAircraftModelComponent
