import React, { Component } from 'react'
import LandingGearService from '../services/LandingGearService'

class ListLandingGearComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                landingGears: []
        }
        this.addLandingGear = this.addLandingGear.bind(this);
        this.editLandingGear = this.editLandingGear.bind(this);
        this.deleteLandingGear = this.deleteLandingGear.bind(this);
    }

    deleteLandingGear(id){
        LandingGearService.deleteLandingGear(id).then( res => {
            this.setState({landingGears: this.state.landingGears.filter(landingGear => landingGear.landingGearId !== id)});
        });
    }
    viewLandingGear(id){
        this.props.history.push(`/view-landingGear/${id}`);
    }
    editLandingGear(id){
        this.props.history.push(`/add-landingGear/${id}`);
    }

    componentDidMount(){
        LandingGearService.getLandingGears().then((res) => {
            this.setState({ landingGears: res.data});
        });
    }

    addLandingGear(){
        this.props.history.push('/add-landingGear/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">LandingGear List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addLandingGear}> Add LandingGear</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> SupplierPartNumber </th>
                                    <th> GearType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.landingGears.map(
                                        landingGear => 
                                        <tr key = {landingGear.landingGearId}>
                                             <td> { landingGear.supplierPartNumber } </td>
                                             <td> { landingGear.gearType } </td>
                                             <td>
                                                 <button onClick={ () => this.editLandingGear(landingGear.landingGearId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteLandingGear(landingGear.landingGearId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewLandingGear(landingGear.landingGearId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListLandingGearComponent
