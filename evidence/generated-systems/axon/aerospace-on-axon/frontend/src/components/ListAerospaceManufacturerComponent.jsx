import React, { Component } from 'react'
import AerospaceManufacturerService from '../services/AerospaceManufacturerService'

class ListAerospaceManufacturerComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                aerospaceManufacturers: []
        }
        this.addAerospaceManufacturer = this.addAerospaceManufacturer.bind(this);
        this.editAerospaceManufacturer = this.editAerospaceManufacturer.bind(this);
        this.deleteAerospaceManufacturer = this.deleteAerospaceManufacturer.bind(this);
    }

    deleteAerospaceManufacturer(id){
        AerospaceManufacturerService.deleteAerospaceManufacturer(id).then( res => {
            this.setState({aerospaceManufacturers: this.state.aerospaceManufacturers.filter(aerospaceManufacturer => aerospaceManufacturer.aerospaceManufacturerId !== id)});
        });
    }
    viewAerospaceManufacturer(id){
        this.props.history.push(`/view-aerospaceManufacturer/${id}`);
    }
    editAerospaceManufacturer(id){
        this.props.history.push(`/add-aerospaceManufacturer/${id}`);
    }

    componentDidMount(){
        AerospaceManufacturerService.getAerospaceManufacturers().then((res) => {
            this.setState({ aerospaceManufacturers: res.data});
        });
    }

    addAerospaceManufacturer(){
        this.props.history.push('/add-aerospaceManufacturer/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">AerospaceManufacturer List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAerospaceManufacturer}> Add AerospaceManufacturer</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> LegalName </th>
                                    <th> HeadquartersCountry </th>
                                    <th> Website </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.aerospaceManufacturers.map(
                                        aerospaceManufacturer => 
                                        <tr key = {aerospaceManufacturer.aerospaceManufacturerId}>
                                             <td> { aerospaceManufacturer.name } </td>
                                             <td> { aerospaceManufacturer.legalName } </td>
                                             <td> { aerospaceManufacturer.headquartersCountry } </td>
                                             <td> { aerospaceManufacturer.website } </td>
                                             <td>
                                                 <button onClick={ () => this.editAerospaceManufacturer(aerospaceManufacturer.aerospaceManufacturerId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAerospaceManufacturer(aerospaceManufacturer.aerospaceManufacturerId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAerospaceManufacturer(aerospaceManufacturer.aerospaceManufacturerId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAerospaceManufacturerComponent
