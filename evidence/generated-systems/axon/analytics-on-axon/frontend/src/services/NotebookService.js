import axios from 'axios';

const NOTEBOOK_API_BASE_URL = "/Notebook";

class NotebookService {

    getNotebooks(){
        return axios.get(NOTEBOOK_API_BASE_URL + '/' );
    }

    createNotebook(notebook){
        return axios.post(NOTEBOOK_API_BASE_URL  + '/create', notebook);
    }

    getNotebookById(notebookId){
        return axios.get(NOTEBOOK_API_BASE_URL + '/load?notebookId=' + notebookId);
    }

    updateNotebook(notebook){
        return axios.put(NOTEBOOK_API_BASE_URL + '/update', notebook);
    }

    deleteNotebook(notebookId){
        return axios.delete(NOTEBOOK_API_BASE_URL + '/delete?notebookId=' + notebookId);
    }
}

export default new NotebookService()