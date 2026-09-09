import React, { Component } from 'react'
import ProductPricingService from '../services/ProductPricingService'

class ListProductPricingComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                productPricings: []
        }
        this.addProductPricing = this.addProductPricing.bind(this);
        this.editProductPricing = this.editProductPricing.bind(this);
        this.deleteProductPricing = this.deleteProductPricing.bind(this);
    }

    deleteProductPricing(id){
        ProductPricingService.deleteProductPricing(id).then( res => {
            this.setState({productPricings: this.state.productPricings.filter(productPricing => productPricing.productPricingId !== id)});
        });
    }
    viewProductPricing(id){
        this.props.history.push(`/view-productPricing/${id}`);
    }
    editProductPricing(id){
        this.props.history.push(`/add-productPricing/${id}`);
    }

    componentDidMount(){
        ProductPricingService.getProductPricings().then((res) => {
            this.setState({ productPricings: res.data});
        });
    }

    addProductPricing(){
        this.props.history.push('/add-productPricing/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ProductPricing List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addProductPricing}> Add ProductPricing</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ListPrice </th>
                                    <th> SalePrice </th>
                                    <th> ValidFrom </th>
                                    <th> ValidTo </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.productPricings.map(
                                        productPricing => 
                                        <tr key = {productPricing.productPricingId}>
                                             <td> { productPricing.listPrice } </td>
                                             <td> { productPricing.salePrice } </td>
                                             <td> { productPricing.validFrom } </td>
                                             <td> { productPricing.validTo } </td>
                                             <td>
                                                 <button onClick={ () => this.editProductPricing(productPricing.productPricingId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteProductPricing(productPricing.productPricingId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewProductPricing(productPricing.productPricingId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListProductPricingComponent
