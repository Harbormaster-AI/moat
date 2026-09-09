import React, { Component } from 'react'
import BrandSafetyPolicyService from '../services/BrandSafetyPolicyService'

class ListBrandSafetyPolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                brandSafetyPolicys: []
        }
        this.addBrandSafetyPolicy = this.addBrandSafetyPolicy.bind(this);
        this.editBrandSafetyPolicy = this.editBrandSafetyPolicy.bind(this);
        this.deleteBrandSafetyPolicy = this.deleteBrandSafetyPolicy.bind(this);
    }

    deleteBrandSafetyPolicy(id){
        BrandSafetyPolicyService.deleteBrandSafetyPolicy(id).then( res => {
            this.setState({brandSafetyPolicys: this.state.brandSafetyPolicys.filter(brandSafetyPolicy => brandSafetyPolicy.brandSafetyPolicyId !== id)});
        });
    }
    viewBrandSafetyPolicy(id){
        this.props.history.push(`/view-brandSafetyPolicy/${id}`);
    }
    editBrandSafetyPolicy(id){
        this.props.history.push(`/add-brandSafetyPolicy/${id}`);
    }

    componentDidMount(){
        BrandSafetyPolicyService.getBrandSafetyPolicys().then((res) => {
            this.setState({ brandSafetyPolicys: res.data});
        });
    }

    addBrandSafetyPolicy(){
        this.props.history.push('/add-brandSafetyPolicy/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">BrandSafetyPolicy List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addBrandSafetyPolicy}> Add BrandSafetyPolicy</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Level </th>
                                    <th> ContentRatingThreshold </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.brandSafetyPolicys.map(
                                        brandSafetyPolicy => 
                                        <tr key = {brandSafetyPolicy.brandSafetyPolicyId}>
                                             <td> { brandSafetyPolicy.level } </td>
                                             <td> { brandSafetyPolicy.contentRatingThreshold } </td>
                                             <td>
                                                 <button onClick={ () => this.editBrandSafetyPolicy(brandSafetyPolicy.brandSafetyPolicyId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteBrandSafetyPolicy(brandSafetyPolicy.brandSafetyPolicyId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewBrandSafetyPolicy(brandSafetyPolicy.brandSafetyPolicyId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListBrandSafetyPolicyComponent
