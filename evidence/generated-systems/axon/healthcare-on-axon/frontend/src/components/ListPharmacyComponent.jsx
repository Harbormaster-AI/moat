import React, { Component } from 'react'
import PharmacyService from '../services/PharmacyService'

class ListPharmacyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                pharmacys: []
        }
        this.addPharmacy = this.addPharmacy.bind(this);
        this.editPharmacy = this.editPharmacy.bind(this);
        this.deletePharmacy = this.deletePharmacy.bind(this);
    }

    deletePharmacy(id){
        PharmacyService.deletePharmacy(id).then( res => {
            this.setState({pharmacys: this.state.pharmacys.filter(pharmacy => pharmacy.pharmacyId !== id)});
        });
    }
    viewPharmacy(id){
        this.props.history.push(`/view-pharmacy/${id}`);
    }
    editPharmacy(id){
        this.props.history.push(`/add-pharmacy/${id}`);
    }

    componentDidMount(){
        PharmacyService.getPharmacys().then((res) => {
            this.setState({ pharmacys: res.data});
        });
    }

    addPharmacy(){
        this.props.history.push('/add-pharmacy/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Pharmacy List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPharmacy}> Add Pharmacy</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.pharmacys.map(
                                        pharmacy => 
                                        <tr key = {pharmacy.pharmacyId}>
                                             <td> { pharmacy.name } </td>
                                             <td>
                                                 <button onClick={ () => this.editPharmacy(pharmacy.pharmacyId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePharmacy(pharmacy.pharmacyId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPharmacy(pharmacy.pharmacyId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPharmacyComponent
