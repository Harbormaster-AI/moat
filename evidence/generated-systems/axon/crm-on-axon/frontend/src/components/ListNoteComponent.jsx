import React, { Component } from 'react'
import NoteService from '../services/NoteService'

class ListNoteComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                notes: []
        }
        this.addNote = this.addNote.bind(this);
        this.editNote = this.editNote.bind(this);
        this.deleteNote = this.deleteNote.bind(this);
    }

    deleteNote(id){
        NoteService.deleteNote(id).then( res => {
            this.setState({notes: this.state.notes.filter(note => note.noteId !== id)});
        });
    }
    viewNote(id){
        this.props.history.push(`/view-note/${id}`);
    }
    editNote(id){
        this.props.history.push(`/add-note/${id}`);
    }

    componentDidMount(){
        NoteService.getNotes().then((res) => {
            this.setState({ notes: res.data});
        });
    }

    addNote(){
        this.props.history.push('/add-note/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Note List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addNote}> Add Note</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Title </th>
                                    <th> Content </th>
                                    <th> CreatedAt </th>
                                    <th> UpdatedAt </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.notes.map(
                                        note => 
                                        <tr key = {note.noteId}>
                                             <td> { note.title } </td>
                                             <td> { note.content } </td>
                                             <td> { note.createdAt } </td>
                                             <td> { note.updatedAt } </td>
                                             <td>
                                                 <button onClick={ () => this.editNote(note.noteId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteNote(note.noteId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewNote(note.noteId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListNoteComponent
