import React, { Component } from 'react'
import AircraftVariantService from '../services/AircraftVariantService'

class ListAircraftVariantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                aircraftVariants: []
        }
        this.addAircraftVariant = this.addAircraftVariant.bind(this);
        this.editAircraftVariant = this.editAircraftVariant.bind(this);
        this.deleteAircraftVariant = this.deleteAircraftVariant.bind(this);
    }

    deleteAircraftVariant(id){
        AircraftVariantService.deleteAircraftVariant(id).then( res => {
            this.setState({aircraftVariants: this.state.aircraftVariants.filter(aircraftVariant => aircraftVariant.aircraftVariantId !== id)});
        });
    }
    viewAircraftVariant(id){
        this.props.history.push(`/view-aircraftVariant/${id}`);
    }
    editAircraftVariant(id){
        this.props.history.push(`/add-aircraftVariant/${id}`);
    }

    componentDidMount(){
        AircraftVariantService.getAircraftVariants().then((res) => {
            this.setState({ aircraftVariants: res.data});
        });
    }

    addAircraftVariant(){
        this.props.history.push('/add-aircraftVariant/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">AircraftVariant List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAircraftVariant}> Add AircraftVariant</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> VariantCode </th>
                                    <th> RangeNm </th>
                                    <th> MaxTakeoffWeightKg </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.aircraftVariants.map(
                                        aircraftVariant => 
                                        <tr key = {aircraftVariant.aircraftVariantId}>
                                             <td> { aircraftVariant.variantCode } </td>
                                             <td> { aircraftVariant.rangeNm } </td>
                                             <td> { aircraftVariant.maxTakeoffWeightKg } </td>
                                             <td>
                                                 <button onClick={ () => this.editAircraftVariant(aircraftVariant.aircraftVariantId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAircraftVariant(aircraftVariant.aircraftVariantId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAircraftVariant(aircraftVariant.aircraftVariantId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAircraftVariantComponent
