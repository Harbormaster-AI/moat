import axios from 'axios';

const DIRECTDEBITMANDATE_API_BASE_URL = "/DirectDebitMandate";

class DirectDebitMandateService {

    getDirectDebitMandates(){
        return axios.get(DIRECTDEBITMANDATE_API_BASE_URL + '/' );
    }

    createDirectDebitMandate(directDebitMandate){
        return axios.post(DIRECTDEBITMANDATE_API_BASE_URL  + '/create', directDebitMandate);
    }

    getDirectDebitMandateById(directDebitMandateId){
        return axios.get(DIRECTDEBITMANDATE_API_BASE_URL + '/load?directDebitMandateId=' + directDebitMandateId);
    }

    updateDirectDebitMandate(directDebitMandate){
        return axios.put(DIRECTDEBITMANDATE_API_BASE_URL + '/update', directDebitMandate);
    }

    deleteDirectDebitMandate(directDebitMandateId){
        return axios.delete(DIRECTDEBITMANDATE_API_BASE_URL + '/delete?directDebitMandateId=' + directDebitMandateId);
    }
}

export default new DirectDebitMandateService()