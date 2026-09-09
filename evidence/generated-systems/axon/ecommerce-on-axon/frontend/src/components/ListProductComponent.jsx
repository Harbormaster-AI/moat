import React, { Component } from 'react'
import ProductService from '../services/ProductService'

class ListProductComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                products: []
        }
        this.addProduct = this.addProduct.bind(this);
        this.editProduct = this.editProduct.bind(this);
        this.deleteProduct = this.deleteProduct.bind(this);
    }

    deleteProduct(id){
        ProductService.deleteProduct(id).then( res => {
            this.setState({products: this.state.products.filter(product => product.productId !== id)});
        });
    }
    viewProduct(id){
        this.props.history.push(`/view-product/${id}`);
    }
    editProduct(id){
        this.props.history.push(`/add-product/${id}`);
    }

    componentDidMount(){
        ProductService.getProducts().then((res) => {
            this.setState({ products: res.data});
        });
    }

    addProduct(){
        this.props.history.push('/add-product/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Product List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addProduct}> Add Product</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Slug </th>
                                    <th> AsActive </th>
                                    <th> ProductType </th>
                                    <th> DefaultTaxClass </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.products.map(
                                        product => 
                                        <tr key = {product.productId}>
                                             <td> { product.name } </td>
                                             <td> { product.slug } </td>
                                             <td> { product.asActive } </td>
                                             <td> { product.productType } </td>
                                             <td> { product.defaultTaxClass } </td>
                                             <td>
                                                 <button onClick={ () => this.editProduct(product.productId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteProduct(product.productId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewProduct(product.productId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListProductComponent
