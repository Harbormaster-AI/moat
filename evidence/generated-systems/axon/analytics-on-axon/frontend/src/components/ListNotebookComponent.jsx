import React, { Component } from 'react'
import NotebookService from '../services/NotebookService'

class ListNotebookComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                notebooks: []
        }
        this.addNotebook = this.addNotebook.bind(this);
        this.editNotebook = this.editNotebook.bind(this);
        this.deleteNotebook = this.deleteNotebook.bind(this);
    }

    deleteNotebook(id){
        NotebookService.deleteNotebook(id).then( res => {
            this.setState({notebooks: this.state.notebooks.filter(notebook => notebook.notebookId !== id)});
        });
    }
    viewNotebook(id){
        this.props.history.push(`/view-notebook/${id}`);
    }
    editNotebook(id){
        this.props.history.push(`/add-notebook/${id}`);
    }

    componentDidMount(){
        NotebookService.getNotebooks().then((res) => {
            this.setState({ notebooks: res.data});
        });
    }

    addNotebook(){
        this.props.history.push('/add-notebook/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Notebook List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addNotebook}> Add Notebook</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Title </th>
                                    <th> Repository </th>
                                    <th> Language </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.notebooks.map(
                                        notebook => 
                                        <tr key = {notebook.notebookId}>
                                             <td> { notebook.title } </td>
                                             <td> { notebook.repository } </td>
                                             <td> { notebook.language } </td>
                                             <td>
                                                 <button onClick={ () => this.editNotebook(notebook.notebookId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteNotebook(notebook.notebookId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewNotebook(notebook.notebookId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListNotebookComponent
