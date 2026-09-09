import axios from 'axios';

const CREDITOR_API_BASE_URL = "/Creditor";

class CreditorService {

    getCreditors(){
        return axios.get(CREDITOR_API_BASE_URL + '/' );
    }

    createCreditor(creditor){
        return axios.post(CREDITOR_API_BASE_URL  + '/create', creditor);
    }

    getCreditorById(creditorId){
        return axios.get(CREDITOR_API_BASE_URL + '/load?creditorId=' + creditorId);
    }

    updateCreditor(creditor){
        return axios.put(CREDITOR_API_BASE_URL + '/update', creditor);
    }

    deleteCreditor(creditorId){
        return axios.delete(CREDITOR_API_BASE_URL + '/delete?creditorId=' + creditorId);
    }
}

export default new CreditorService()