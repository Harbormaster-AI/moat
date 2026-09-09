import React, { Component } from 'react'
import ProductVariantService from '../services/ProductVariantService'

class ListProductVariantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                productVariants: []
        }
        this.addProductVariant = this.addProductVariant.bind(this);
        this.editProductVariant = this.editProductVariant.bind(this);
        this.deleteProductVariant = this.deleteProductVariant.bind(this);
    }

    deleteProductVariant(id){
        ProductVariantService.deleteProductVariant(id).then( res => {
            this.setState({productVariants: this.state.productVariants.filter(productVariant => productVariant.productVariantId !== id)});
        });
    }
    viewProductVariant(id){
        this.props.history.push(`/view-productVariant/${id}`);
    }
    editProductVariant(id){
        this.props.history.push(`/add-productVariant/${id}`);
    }

    componentDidMount(){
        ProductVariantService.getProductVariants().then((res) => {
            this.setState({ productVariants: res.data});
        });
    }

    addProductVariant(){
        this.props.history.push('/add-productVariant/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ProductVariant List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addProductVariant}> Add ProductVariant</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Sku </th>
                                    <th> Barcode </th>
                                    <th> Title </th>
                                    <th> Weight </th>
                                    <th> RequiresShipping </th>
                                    <th> WeightUnit </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.productVariants.map(
                                        productVariant => 
                                        <tr key = {productVariant.productVariantId}>
                                             <td> { productVariant.sku } </td>
                                             <td> { productVariant.barcode } </td>
                                             <td> { productVariant.title } </td>
                                             <td> { productVariant.weight } </td>
                                             <td> { productVariant.requiresShipping } </td>
                                             <td> { productVariant.weightUnit } </td>
                                             <td>
                                                 <button onClick={ () => this.editProductVariant(productVariant.productVariantId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteProductVariant(productVariant.productVariantId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewProductVariant(productVariant.productVariantId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListProductVariantComponent
