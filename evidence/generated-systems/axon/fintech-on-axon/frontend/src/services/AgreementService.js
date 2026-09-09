import axios from 'axios';

const AGREEMENT_API_BASE_URL = "/Agreement";

class AgreementService {

    getAgreements(){
        return axios.get(AGREEMENT_API_BASE_URL + '/' );
    }

    createAgreement(agreement){
        return axios.post(AGREEMENT_API_BASE_URL  + '/create', agreement);
    }

    getAgreementById(agreementId){
        return axios.get(AGREEMENT_API_BASE_URL + '/load?agreementId=' + agreementId);
    }

    updateAgreement(agreement){
        return axios.put(AGREEMENT_API_BASE_URL + '/update', agreement);
    }

    deleteAgreement(agreementId){
        return axios.delete(AGREEMENT_API_BASE_URL + '/delete?agreementId=' + agreementId);
    }
}

export default new AgreementService()