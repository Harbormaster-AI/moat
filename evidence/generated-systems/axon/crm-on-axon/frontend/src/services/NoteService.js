import axios from 'axios';

const NOTE_API_BASE_URL = "/Note";

class NoteService {

    getNotes(){
        return axios.get(NOTE_API_BASE_URL + '/' );
    }

    createNote(note){
        return axios.post(NOTE_API_BASE_URL  + '/create', note);
    }

    getNoteById(noteId){
        return axios.get(NOTE_API_BASE_URL + '/load?noteId=' + noteId);
    }

    updateNote(note){
        return axios.put(NOTE_API_BASE_URL + '/update', note);
    }

    deleteNote(noteId){
        return axios.delete(NOTE_API_BASE_URL + '/delete?noteId=' + noteId);
    }
}

export default new NoteService()