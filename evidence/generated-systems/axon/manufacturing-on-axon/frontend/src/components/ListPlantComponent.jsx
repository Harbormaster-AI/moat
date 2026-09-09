import React, { Component } from 'react'
import PlantService from '../services/PlantService'

class ListPlantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                plants: []
        }
        this.addPlant = this.addPlant.bind(this);
        this.editPlant = this.editPlant.bind(this);
        this.deletePlant = this.deletePlant.bind(this);
    }

    deletePlant(id){
        PlantService.deletePlant(id).then( res => {
            this.setState({plants: this.state.plants.filter(plant => plant.plantId !== id)});
        });
    }
    viewPlant(id){
        this.props.history.push(`/view-plant/${id}`);
    }
    editPlant(id){
        this.props.history.push(`/add-plant/${id}`);
    }

    componentDidMount(){
        PlantService.getPlants().then((res) => {
            this.setState({ plants: res.data});
        });
    }

    addPlant(){
        this.props.history.push('/add-plant/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Plant List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPlant}> Add Plant</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> PlantCode </th>
                                    <th> Address </th>
                                    <th> TimeZone </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.plants.map(
                                        plant => 
                                        <tr key = {plant.plantId}>
                                             <td> { plant.name } </td>
                                             <td> { plant.plantCode } </td>
                                             <td> { plant.address } </td>
                                             <td> { plant.timeZone } </td>
                                             <td>
                                                 <button onClick={ () => this.editPlant(plant.plantId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePlant(plant.plantId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPlant(plant.plantId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPlantComponent
