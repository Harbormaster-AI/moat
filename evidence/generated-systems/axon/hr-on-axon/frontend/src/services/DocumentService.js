import axios from 'axios';

const DOCUMENT_API_BASE_URL = "/Document";

class DocumentService {

    getDocuments(){
        return axios.get(DOCUMENT_API_BASE_URL + '/' );
    }

    createDocument(document){
        return axios.post(DOCUMENT_API_BASE_URL  + '/create', document);
    }

    getDocumentById(documentId){
        return axios.get(DOCUMENT_API_BASE_URL + '/load?documentId=' + documentId);
    }

    updateDocument(document){
        return axios.put(DOCUMENT_API_BASE_URL + '/update', document);
    }

    deleteDocument(documentId){
        return axios.delete(DOCUMENT_API_BASE_URL + '/delete?documentId=' + documentId);
    }
}

export default new DocumentService()