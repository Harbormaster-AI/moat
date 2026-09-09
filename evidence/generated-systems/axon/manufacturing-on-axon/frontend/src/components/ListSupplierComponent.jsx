import React, { Component } from 'react'
import SupplierService from '../services/SupplierService'

class ListSupplierComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                suppliers: []
        }
        this.addSupplier = this.addSupplier.bind(this);
        this.editSupplier = this.editSupplier.bind(this);
        this.deleteSupplier = this.deleteSupplier.bind(this);
    }

    deleteSupplier(id){
        SupplierService.deleteSupplier(id).then( res => {
            this.setState({suppliers: this.state.suppliers.filter(supplier => supplier.supplierId !== id)});
        });
    }
    viewSupplier(id){
        this.props.history.push(`/view-supplier/${id}`);
    }
    editSupplier(id){
        this.props.history.push(`/add-supplier/${id}`);
    }

    componentDidMount(){
        SupplierService.getSuppliers().then((res) => {
            this.setState({ suppliers: res.data});
        });
    }

    addSupplier(){
        this.props.history.push('/add-supplier/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Supplier List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addSupplier}> Add Supplier</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> SupplierCode </th>
                                    <th> Address </th>
                                    <th> SupplierTier </th>
                                    <th> PaymentTerms </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.suppliers.map(
                                        supplier => 
                                        <tr key = {supplier.supplierId}>
                                             <td> { supplier.name } </td>
                                             <td> { supplier.supplierCode } </td>
                                             <td> { supplier.address } </td>
                                             <td> { supplier.supplierTier } </td>
                                             <td> { supplier.paymentTerms } </td>
                                             <td>
                                                 <button onClick={ () => this.editSupplier(supplier.supplierId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteSupplier(supplier.supplierId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewSupplier(supplier.supplierId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListSupplierComponent
