import React, { Component } from 'react'
import ProductOfferingService from '../services/ProductOfferingService'

class ListProductOfferingComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                productOfferings: []
        }
        this.addProductOffering = this.addProductOffering.bind(this);
        this.editProductOffering = this.editProductOffering.bind(this);
        this.deleteProductOffering = this.deleteProductOffering.bind(this);
    }

    deleteProductOffering(id){
        ProductOfferingService.deleteProductOffering(id).then( res => {
            this.setState({productOfferings: this.state.productOfferings.filter(productOffering => productOffering.productOfferingId !== id)});
        });
    }
    viewProductOffering(id){
        this.props.history.push(`/view-productOffering/${id}`);
    }
    editProductOffering(id){
        this.props.history.push(`/add-productOffering/${id}`);
    }

    componentDidMount(){
        ProductOfferingService.getProductOfferings().then((res) => {
            this.setState({ productOfferings: res.data});
        });
    }

    addProductOffering(){
        this.props.history.push('/add-productOffering/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ProductOffering List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addProductOffering}> Add ProductOffering</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> ProductCode </th>
                                    <th> Category </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.productOfferings.map(
                                        productOffering => 
                                        <tr key = {productOffering.productOfferingId}>
                                             <td> { productOffering.name } </td>
                                             <td> { productOffering.productCode } </td>
                                             <td> { productOffering.category } </td>
                                             <td>
                                                 <button onClick={ () => this.editProductOffering(productOffering.productOfferingId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteProductOffering(productOffering.productOfferingId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewProductOffering(productOffering.productOfferingId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListProductOfferingComponent
