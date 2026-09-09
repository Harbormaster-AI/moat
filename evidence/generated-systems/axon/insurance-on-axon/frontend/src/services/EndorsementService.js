import axios from 'axios';

const ENDORSEMENT_API_BASE_URL = "/Endorsement";

class EndorsementService {

    getEndorsements(){
        return axios.get(ENDORSEMENT_API_BASE_URL + '/' );
    }

    createEndorsement(endorsement){
        return axios.post(ENDORSEMENT_API_BASE_URL  + '/create', endorsement);
    }

    getEndorsementById(endorsementId){
        return axios.get(ENDORSEMENT_API_BASE_URL + '/load?endorsementId=' + endorsementId);
    }

    updateEndorsement(endorsement){
        return axios.put(ENDORSEMENT_API_BASE_URL + '/update', endorsement);
    }

    deleteEndorsement(endorsementId){
        return axios.delete(ENDORSEMENT_API_BASE_URL + '/delete?endorsementId=' + endorsementId);
    }
}

export default new EndorsementService()