import React, { Component } from 'react'
import GeoRegionService from '../services/GeoRegionService'

class ListGeoRegionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                geoRegions: []
        }
        this.addGeoRegion = this.addGeoRegion.bind(this);
        this.editGeoRegion = this.editGeoRegion.bind(this);
        this.deleteGeoRegion = this.deleteGeoRegion.bind(this);
    }

    deleteGeoRegion(id){
        GeoRegionService.deleteGeoRegion(id).then( res => {
            this.setState({geoRegions: this.state.geoRegions.filter(geoRegion => geoRegion.geoRegionId !== id)});
        });
    }
    viewGeoRegion(id){
        this.props.history.push(`/view-geoRegion/${id}`);
    }
    editGeoRegion(id){
        this.props.history.push(`/add-geoRegion/${id}`);
    }

    componentDidMount(){
        GeoRegionService.getGeoRegions().then((res) => {
            this.setState({ geoRegions: res.data});
        });
    }

    addGeoRegion(){
        this.props.history.push('/add-geoRegion/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">GeoRegion List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addGeoRegion}> Add GeoRegion</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Code </th>
                                    <th> Name </th>
                                    <th> RegionType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.geoRegions.map(
                                        geoRegion => 
                                        <tr key = {geoRegion.geoRegionId}>
                                             <td> { geoRegion.code } </td>
                                             <td> { geoRegion.name } </td>
                                             <td> { geoRegion.regionType } </td>
                                             <td>
                                                 <button onClick={ () => this.editGeoRegion(geoRegion.geoRegionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteGeoRegion(geoRegion.geoRegionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewGeoRegion(geoRegion.geoRegionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListGeoRegionComponent
