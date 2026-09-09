import axios from 'axios';

const LEGALHOLD_API_BASE_URL = "/LegalHold";

class LegalHoldService {

    getLegalHolds(){
        return axios.get(LEGALHOLD_API_BASE_URL + '/' );
    }

    createLegalHold(legalHold){
        return axios.post(LEGALHOLD_API_BASE_URL  + '/create', legalHold);
    }

    getLegalHoldById(legalHoldId){
        return axios.get(LEGALHOLD_API_BASE_URL + '/load?legalHoldId=' + legalHoldId);
    }

    updateLegalHold(legalHold){
        return axios.put(LEGALHOLD_API_BASE_URL + '/update', legalHold);
    }

    deleteLegalHold(legalHoldId){
        return axios.delete(LEGALHOLD_API_BASE_URL + '/delete?legalHoldId=' + legalHoldId);
    }
}

export default new LegalHoldService()