import React, { Component } from 'react'
import NoteService from '../services/NoteService';

class CreateNoteComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                title: '',
                content: '',
                createdAt: '',
                updatedAt: ''
        }
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changecontentHandler = this.changecontentHandler.bind(this);
        this.changecreatedAtHandler = this.changecreatedAtHandler.bind(this);
        this.changeupdatedAtHandler = this.changeupdatedAtHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            NoteService.getNoteById(this.state.id).then( (res) =>{
                let note = res.data;
                this.setState({
                    title: note.title,
                    content: note.content,
                    createdAt: note.createdAt,
                    updatedAt: note.updatedAt
                });
            });
        }        
    }
    saveOrUpdateNote = (e) => {
        e.preventDefault();
        let note = {
                noteId: this.state.id,
                title: this.state.title,
                content: this.state.content,
                createdAt: this.state.createdAt,
                updatedAt: this.state.updatedAt
            };
        console.log('note => ' + JSON.stringify(note));

        // step 5
        if(this.state.id === '_add'){
            note.noteId=''
            NoteService.createNote(note).then(res =>{
                this.props.history.push('/notes');
            });
        }else{
            NoteService.updateNote(note).then( res => {
                this.props.history.push('/notes');
            });
        }
    }
    
    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changecontentHandler= (event) => {
        this.setState({content: event.target.value});
    }
    changecreatedAtHandler= (event) => {
        this.setState({createdAt: event.target.value});
    }
    changeupdatedAtHandler= (event) => {
        this.setState({updatedAt: event.target.value});
    }

    cancel(){
        this.props.history.push('/notes');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Note</h3>
        }else{
            return <h3 className="text-center">Update Note</h3>
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

                                            <label> content:&emsp; </label>
                                                <input placeholder="content" name="content" className="form-control" value={this.state.content} onChange={this.changecontentHandler}/>

                                            <label> createdAt:&emsp; </label>
                                                <input type="time" placeholder="createdAt" name="createdAt" className="form-control" value={this.state.createdAt} onChange={this.changecreatedAtHandler}/>

                                            <label> updatedAt:&emsp; </label>
                                                <input type="time" placeholder="updatedAt" name="updatedAt" className="form-control" value={this.state.updatedAt} onChange={this.changeupdatedAtHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateNote}>Save</button>
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

export default CreateNoteComponent
