import axios from 'axios';

const KYCDOCUMENT_API_BASE_URL = "/KYCDocument";

class KYCDocumentService {

    getKYCDocuments(){
        return axios.get(KYCDOCUMENT_API_BASE_URL + '/' );
    }

    createKYCDocument(kYCDocument){
        return axios.post(KYCDOCUMENT_API_BASE_URL  + '/create', kYCDocument);
    }

    getKYCDocumentById(kYCDocumentId){
        return axios.get(KYCDOCUMENT_API_BASE_URL + '/load?kYCDocumentId=' + kYCDocumentId);
    }

    updateKYCDocument(kYCDocument){
        return axios.put(KYCDOCUMENT_API_BASE_URL + '/update', kYCDocument);
    }

    deleteKYCDocument(kYCDocumentId){
        return axios.delete(KYCDOCUMENT_API_BASE_URL + '/delete?kYCDocumentId=' + kYCDocumentId);
    }
}

export default new KYCDocumentService()