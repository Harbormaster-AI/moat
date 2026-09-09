import axios from 'axios';

const OBLIGATION_API_BASE_URL = "/Obligation";

class ObligationService {

    getObligations(){
        return axios.get(OBLIGATION_API_BASE_URL + '/' );
    }

    createObligation(obligation){
        return axios.post(OBLIGATION_API_BASE_URL  + '/create', obligation);
    }

    getObligationById(obligationId){
        return axios.get(OBLIGATION_API_BASE_URL + '/load?obligationId=' + obligationId);
    }

    updateObligation(obligation){
        return axios.put(OBLIGATION_API_BASE_URL + '/update', obligation);
    }

    deleteObligation(obligationId){
        return axios.delete(OBLIGATION_API_BASE_URL + '/delete?obligationId=' + obligationId);
    }
}

export default new ObligationService()