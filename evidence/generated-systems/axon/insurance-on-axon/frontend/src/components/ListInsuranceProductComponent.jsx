import React, { Component } from 'react'
import InsuranceProductService from '../services/InsuranceProductService'

class ListInsuranceProductComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                insuranceProducts: []
        }
        this.addInsuranceProduct = this.addInsuranceProduct.bind(this);
        this.editInsuranceProduct = this.editInsuranceProduct.bind(this);
        this.deleteInsuranceProduct = this.deleteInsuranceProduct.bind(this);
    }

    deleteInsuranceProduct(id){
        InsuranceProductService.deleteInsuranceProduct(id).then( res => {
            this.setState({insuranceProducts: this.state.insuranceProducts.filter(insuranceProduct => insuranceProduct.insuranceProductId !== id)});
        });
    }
    viewInsuranceProduct(id){
        this.props.history.push(`/view-insuranceProduct/${id}`);
    }
    editInsuranceProduct(id){
        this.props.history.push(`/add-insuranceProduct/${id}`);
    }

    componentDidMount(){
        InsuranceProductService.getInsuranceProducts().then((res) => {
            this.setState({ insuranceProducts: res.data});
        });
    }

    addInsuranceProduct(){
        this.props.history.push('/add-insuranceProduct/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">InsuranceProduct List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addInsuranceProduct}> Add InsuranceProduct</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> ProductCode </th>
                                    <th> LineOfBusiness </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.insuranceProducts.map(
                                        insuranceProduct => 
                                        <tr key = {insuranceProduct.insuranceProductId}>
                                             <td> { insuranceProduct.name } </td>
                                             <td> { insuranceProduct.productCode } </td>
                                             <td> { insuranceProduct.lineOfBusiness } </td>
                                             <td>
                                                 <button onClick={ () => this.editInsuranceProduct(insuranceProduct.insuranceProductId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteInsuranceProduct(insuranceProduct.insuranceProductId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewInsuranceProduct(insuranceProduct.insuranceProductId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListInsuranceProductComponent
