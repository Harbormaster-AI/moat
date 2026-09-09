import React, { Component } from 'react'
import CategoryService from '../services/CategoryService'

class ListCategoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                categorys: []
        }
        this.addCategory = this.addCategory.bind(this);
        this.editCategory = this.editCategory.bind(this);
        this.deleteCategory = this.deleteCategory.bind(this);
    }

    deleteCategory(id){
        CategoryService.deleteCategory(id).then( res => {
            this.setState({categorys: this.state.categorys.filter(category => category.categoryId !== id)});
        });
    }
    viewCategory(id){
        this.props.history.push(`/view-category/${id}`);
    }
    editCategory(id){
        this.props.history.push(`/add-category/${id}`);
    }

    componentDidMount(){
        CategoryService.getCategorys().then((res) => {
            this.setState({ categorys: res.data});
        });
    }

    addCategory(){
        this.props.history.push('/add-category/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Category List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCategory}> Add Category</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Slug </th>
                                    <th> Position </th>
                                    <th> AsActive </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.categorys.map(
                                        category => 
                                        <tr key = {category.categoryId}>
                                             <td> { category.name } </td>
                                             <td> { category.slug } </td>
                                             <td> { category.position } </td>
                                             <td> { category.asActive } </td>
                                             <td>
                                                 <button onClick={ () => this.editCategory(category.categoryId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCategory(category.categoryId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCategory(category.categoryId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCategoryComponent
