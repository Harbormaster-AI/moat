import React, { Component } from 'react'
import PlacementService from '../services/PlacementService'

class ListPlacementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                placements: []
        }
        this.addPlacement = this.addPlacement.bind(this);
        this.editPlacement = this.editPlacement.bind(this);
        this.deletePlacement = this.deletePlacement.bind(this);
    }

    deletePlacement(id){
        PlacementService.deletePlacement(id).then( res => {
            this.setState({placements: this.state.placements.filter(placement => placement.placementId !== id)});
        });
    }
    viewPlacement(id){
        this.props.history.push(`/view-placement/${id}`);
    }
    editPlacement(id){
        this.props.history.push(`/add-placement/${id}`);
    }

    componentDidMount(){
        PlacementService.getPlacements().then((res) => {
            this.setState({ placements: res.data});
        });
    }

    addPlacement(){
        this.props.history.push('/add-placement/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Placement List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPlacement}> Add Placement</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Flight </th>
                                    <th> GoalImpressions </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.placements.map(
                                        placement => 
                                        <tr key = {placement.placementId}>
                                             <td> { placement.name } </td>
                                             <td> { placement.flight } </td>
                                             <td> { placement.goalImpressions } </td>
                                             <td>
                                                 <button onClick={ () => this.editPlacement(placement.placementId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePlacement(placement.placementId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPlacement(placement.placementId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPlacementComponent
