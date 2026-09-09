import React, { Component } from 'react'
import MedicalSupplierService from '../services/MedicalSupplierService'

class ListMedicalSupplierComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                medicalSuppliers: []
        }
        this.addMedicalSupplier = this.addMedicalSupplier.bind(this);
        this.editMedicalSupplier = this.editMedicalSupplier.bind(this);
        this.deleteMedicalSupplier = this.deleteMedicalSupplier.bind(this);
    }

    deleteMedicalSupplier(id){
        MedicalSupplierService.deleteMedicalSupplier(id).then( res => {
            this.setState({medicalSuppliers: this.state.medicalSuppliers.filter(medicalSupplier => medicalSupplier.medicalSupplierId !== id)});
        });
    }
    viewMedicalSupplier(id){
        this.props.history.push(`/view-medicalSupplier/${id}`);
    }
    editMedicalSupplier(id){
        this.props.history.push(`/add-medicalSupplier/${id}`);
    }

    componentDidMount(){
        MedicalSupplierService.getMedicalSuppliers().then((res) => {
            this.setState({ medicalSuppliers: res.data});
        });
    }

    addMedicalSupplier(){
        this.props.history.push('/add-medicalSupplier/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">MedicalSupplier List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addMedicalSupplier}> Add MedicalSupplier</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Website </th>
                                    <th> SupplierTier </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.medicalSuppliers.map(
                                        medicalSupplier => 
                                        <tr key = {medicalSupplier.medicalSupplierId}>
                                             <td> { medicalSupplier.name } </td>
                                             <td> { medicalSupplier.website } </td>
                                             <td> { medicalSupplier.supplierTier } </td>
                                             <td>
                                                 <button onClick={ () => this.editMedicalSupplier(medicalSupplier.medicalSupplierId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteMedicalSupplier(medicalSupplier.medicalSupplierId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewMedicalSupplier(medicalSupplier.medicalSupplierId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListMedicalSupplierComponent
