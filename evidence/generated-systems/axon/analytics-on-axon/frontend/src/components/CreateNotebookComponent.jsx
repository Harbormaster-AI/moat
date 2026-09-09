import React, { Component } from 'react'
import NotebookService from '../services/NotebookService';

class CreateNotebookComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                title: '',
                repository: '',
                language: ''
        }
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changerepositoryHandler = this.changerepositoryHandler.bind(this);
        this.changeLanguageHandler = this.changeLanguageHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            NotebookService.getNotebookById(this.state.id).then( (res) =>{
                let notebook = res.data;
                this.setState({
                    title: notebook.title,
                    repository: notebook.repository,
                    language: notebook.language
                });
            });
        }        
    }
    saveOrUpdateNotebook = (e) => {
        e.preventDefault();
        let notebook = {
                notebookId: this.state.id,
                title: this.state.title,
                repository: this.state.repository,
                language: this.state.language
            };
        console.log('notebook => ' + JSON.stringify(notebook));

        // step 5
        if(this.state.id === '_add'){
            notebook.notebookId=''
            NotebookService.createNotebook(notebook).then(res =>{
                this.props.history.push('/notebooks');
            });
        }else{
            NotebookService.updateNotebook(notebook).then( res => {
                this.props.history.push('/notebooks');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Notebook</h3>
        }else{
            return <h3 className="text-center">Update Notebook</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> title:&emsp; </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> repository:&emsp; </label>
                                                <input placeholder="repository" name="repository" className="form-control" value={this.state.repository} onChange={this.changerepositoryHandler}/>

                                            <label> Language:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateNotebook}>Save</button>
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

export default CreateNotebookComponent
