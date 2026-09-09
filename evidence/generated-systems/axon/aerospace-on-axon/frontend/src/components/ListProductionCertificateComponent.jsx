import React, { Component } from 'react'
import ProductionCertificateService from '../services/ProductionCertificateService'

class ListProductionCertificateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                productionCertificates: []
        }
        this.addProductionCertificate = this.addProductionCertificate.bind(this);
        this.editProductionCertificate = this.editProductionCertificate.bind(this);
        this.deleteProductionCertificate = this.deleteProductionCertificate.bind(this);
    }

    deleteProductionCertificate(id){
        ProductionCertificateService.deleteProductionCertificate(id).then( res => {
            this.setState({productionCertificates: this.state.productionCertificates.filter(productionCertificate => productionCertificate.productionCertificateId !== id)});
        });
    }
    viewProductionCertificate(id){
        this.props.history.push(`/view-productionCertificate/${id}`);
    }
    editProductionCertificate(id){
        this.props.history.push(`/add-productionCertificate/${id}`);
    }

    componentDidMount(){
        ProductionCertificateService.getProductionCertificates().then((res) => {
            this.setState({ productionCertificates: res.data});
        });
    }

    addProductionCertificate(){
        this.props.history.push('/add-productionCertificate/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ProductionCertificate List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addProductionCertificate}> Add ProductionCertificate</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> CertificateNumber </th>
                                    <th> Authority </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.productionCertificates.map(
                                        productionCertificate => 
                                        <tr key = {productionCertificate.productionCertificateId}>
                                             <td> { productionCertificate.certificateNumber } </td>
                                             <td> { productionCertificate.authority } </td>
                                             <td>
                                                 <button onClick={ () => this.editProductionCertificate(productionCertificate.productionCertificateId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteProductionCertificate(productionCertificate.productionCertificateId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewProductionCertificate(productionCertificate.productionCertificateId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListProductionCertificateComponent
