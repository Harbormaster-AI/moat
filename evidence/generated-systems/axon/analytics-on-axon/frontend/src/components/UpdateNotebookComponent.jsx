import React, { Component } from 'react'
import NotebookService from '../services/NotebookService';

class UpdateNotebookComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                title: '',
                repository: '',
                language: ''
        }
        this.updateNotebook = this.updateNotebook.bind(this);

        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changerepositoryHandler = this.changerepositoryHandler.bind(this);
        this.changeLanguageHandler = this.changeLanguageHandler.bind(this);
    }

    componentDidMount(){
        NotebookService.getNotebookById(this.state.id).then( (res) =>{
            let notebook = res.data;
            this.setState({
                title: notebook.title,
                repository: notebook.repository,
                language: notebook.language
            });
        });
    }

    updateNotebook = (e) => {
        e.preventDefault();
        let notebook = {
            notebookId: this.state.id,
            title: this.state.title,
            repository: this.state.repository,
            language: this.state.language
        };
        console.log('notebook => ' + JSON.stringify(notebook));
        console.log('id => ' + JSON.stringify(this.state.id));
        NotebookService.updateNotebook(notebook).then( res => {
            this.props.history.push('/notebooks');
        });
    }

    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changerepositoryHandler= (event) => {
        this.setState({repository: event.target.value});
    }
    changeLanguageHandler= (event) => {
        this.setState({language: event.target.value});
    }

    cancel(){
        this.props.history.push('/notebooks');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Notebook</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> title: </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> repository: </label>
                                                <input placeholder="repository" name="repository" className="form-control" value={this.state.repository} onChange={this.changerepositoryHandler}/>

                                            <label> Language: </label>
                                                <select value={this.state.language} onChange={this.changeLanguageHandler}>
                      <option name="Language" className="form-control" >
                          Python
                      </option>
                      <option name="Language" className="form-control" >
                          R
                      </option>
                      <option name="Language" className="form-control" >
                          SQL
                      </option>
                      <option name="Language" className="form-control" >
                          Julia
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateNotebook}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateNotebookComponent
