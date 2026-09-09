import React, { Component } from 'react'
import BrandService from '../services/BrandService'

class ListBrandComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                brands: []
        }
        this.addBrand = this.addBrand.bind(this);
        this.editBrand = this.editBrand.bind(this);
        this.deleteBrand = this.deleteBrand.bind(this);
    }

    deleteBrand(id){
        BrandService.deleteBrand(id).then( res => {
            this.setState({brands: this.state.brands.filter(brand => brand.brandId !== id)});
        });
    }
    viewBrand(id){
        this.props.history.push(`/view-brand/${id}`);
    }
    editBrand(id){
        this.props.history.push(`/add-brand/${id}`);
    }

    componentDidMount(){
        BrandService.getBrands().then((res) => {
            this.setState({ brands: res.data});
        });
    }

    addBrand(){
        this.props.history.push('/add-brand/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Brand List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addBrand}> Add Brand</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Description </th>
                                    <th> Website </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.brands.map(
                                        brand => 
                                        <tr key = {brand.brandId}>
                                             <td> { brand.name } </td>
                                             <td> { brand.description } </td>
                                             <td> { brand.website } </td>
                                             <td>
                                                 <button onClick={ () => this.editBrand(brand.brandId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteBrand(brand.brandId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewBrand(brand.brandId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListBrandComponent
