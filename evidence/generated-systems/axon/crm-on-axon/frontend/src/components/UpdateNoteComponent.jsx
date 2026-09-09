import React, { Component } from 'react'
import NoteService from '../services/NoteService';

class UpdateNoteComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                title: '',
                content: '',
                createdAt: '',
                updatedAt: ''
        }
        this.updateNote = this.updateNote.bind(this);

        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changecontentHandler = this.changecontentHandler.bind(this);
        this.changecreatedAtHandler = this.changecreatedAtHandler.bind(this);
        this.changeupdatedAtHandler = this.changeupdatedAtHandler.bind(this);
    }

    componentDidMount(){
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

    updateNote = (e) => {
        e.preventDefault();
        let note = {
            noteId: this.state.id,
            title: this.state.title,
            content: this.state.content,
            createdAt: this.state.createdAt,
            updatedAt: this.state.updatedAt
        };
        console.log('note => ' + JSON.stringify(note));
        console.log('id => ' + JSON.stringify(this.state.id));
        NoteService.updateNote(note).then( res => {
            this.props.history.push('/notes');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Note</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> title: </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> content: </label>
                                                <input placeholder="content" name="content" className="form-control" value={this.state.content} onChange={this.changecontentHandler}/>

                                            <label> createdAt: </label>
                                                <input type="time" placeholder="createdAt" name="createdAt" className="form-control" value={this.state.createdAt} onChange={this.changecreatedAtHandler}/>

                                            <label> updatedAt: </label>
                                                <input type="time" placeholder="updatedAt" name="updatedAt" className="form-control" value={this.state.updatedAt} onChange={this.changeupdatedAtHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateNote}>Save</button>
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

export default UpdateNoteComponent
