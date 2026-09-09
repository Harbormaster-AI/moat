import React, { Component } from 'react'
import ContentCategoryService from '../services/ContentCategoryService'

class ListContentCategoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                contentCategorys: []
        }
        this.addContentCategory = this.addContentCategory.bind(this);
        this.editContentCategory = this.editContentCategory.bind(this);
        this.deleteContentCategory = this.deleteContentCategory.bind(this);
    }

    deleteContentCategory(id){
        ContentCategoryService.deleteContentCategory(id).then( res => {
            this.setState({contentCategorys: this.state.contentCategorys.filter(contentCategory => contentCategory.contentCategoryId !== id)});
        });
    }
    viewContentCategory(id){
        this.props.history.push(`/view-contentCategory/${id}`);
    }
    editContentCategory(id){
        this.props.history.push(`/add-contentCategory/${id}`);
    }

    componentDidMount(){
        ContentCategoryService.getContentCategorys().then((res) => {
            this.setState({ contentCategorys: res.data});
        });
    }

    addContentCategory(){
        this.props.history.push('/add-contentCategory/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ContentCategory List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addContentCategory}> Add ContentCategory</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Code </th>
                                    <th> Name </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.contentCategorys.map(
                                        contentCategory => 
                                        <tr key = {contentCategory.contentCategoryId}>
                                             <td> { contentCategory.code } </td>
                                             <td> { contentCategory.name } </td>
                                             <td>
                                                 <button onClick={ () => this.editContentCategory(contentCategory.contentCategoryId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteContentCategory(contentCategory.contentCategoryId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewContentCategory(contentCategory.contentCategoryId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListContentCategoryComponent
