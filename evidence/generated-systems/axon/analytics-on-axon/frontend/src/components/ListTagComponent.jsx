import React, { Component } from 'react'
import TagService from '../services/TagService'

class ListTagComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                tags: []
        }
        this.addTag = this.addTag.bind(this);
        this.editTag = this.editTag.bind(this);
        this.deleteTag = this.deleteTag.bind(this);
    }

    deleteTag(id){
        TagService.deleteTag(id).then( res => {
            this.setState({tags: this.state.tags.filter(tag => tag.tagId !== id)});
        });
    }
    viewTag(id){
        this.props.history.push(`/view-tag/${id}`);
    }
    editTag(id){
        this.props.history.push(`/add-tag/${id}`);
    }

    componentDidMount(){
        TagService.getTags().then((res) => {
            this.setState({ tags: res.data});
        });
    }

    addTag(){
        this.props.history.push('/add-tag/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Tag List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addTag}> Add Tag</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Category </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.tags.map(
                                        tag => 
                                        <tr key = {tag.tagId}>
                                             <td> { tag.name } </td>
                                             <td> { tag.category } </td>
                                             <td>
                                                 <button onClick={ () => this.editTag(tag.tagId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteTag(tag.tagId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewTag(tag.tagId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListTagComponent
