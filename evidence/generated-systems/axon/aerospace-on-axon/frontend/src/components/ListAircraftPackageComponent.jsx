import React, { Component } from 'react'
import AircraftPackageService from '../services/AircraftPackageService'

class ListAircraftPackageComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                aircraftPackages: []
        }
        this.addAircraftPackage = this.addAircraftPackage.bind(this);
        this.editAircraftPackage = this.editAircraftPackage.bind(this);
        this.deleteAircraftPackage = this.deleteAircraftPackage.bind(this);
    }

    deleteAircraftPackage(id){
        AircraftPackageService.deleteAircraftPackage(id).then( res => {
            this.setState({aircraftPackages: this.state.aircraftPackages.filter(aircraftPackage => aircraftPackage.aircraftPackageId !== id)});
        });
    }
    viewAircraftPackage(id){
        this.props.history.push(`/view-aircraftPackage/${id}`);
    }
    editAircraftPackage(id){
        this.props.history.push(`/add-aircraftPackage/${id}`);
    }

    componentDidMount(){
        AircraftPackageService.getAircraftPackages().then((res) => {
            this.setState({ aircraftPackages: res.data});
        });
    }

    addAircraftPackage(){
        this.props.history.push('/add-aircraftPackage/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">AircraftPackage List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAircraftPackage}> Add AircraftPackage</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> PackageType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.aircraftPackages.map(
                                        aircraftPackage => 
                                        <tr key = {aircraftPackage.aircraftPackageId}>
                                             <td> { aircraftPackage.name } </td>
                                             <td> { aircraftPackage.packageType } </td>
                                             <td>
                                                 <button onClick={ () => this.editAircraftPackage(aircraftPackage.aircraftPackageId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAircraftPackage(aircraftPackage.aircraftPackageId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAircraftPackage(aircraftPackage.aircraftPackageId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAircraftPackageComponent
