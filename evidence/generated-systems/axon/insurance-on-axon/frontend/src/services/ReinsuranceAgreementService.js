import axios from 'axios';

const REINSURANCEAGREEMENT_API_BASE_URL = "/ReinsuranceAgreement";

class ReinsuranceAgreementService {

    getReinsuranceAgreements(){
        return axios.get(REINSURANCEAGREEMENT_API_BASE_URL + '/' );
    }

    createReinsuranceAgreement(reinsuranceAgreement){
        return axios.post(REINSURANCEAGREEMENT_API_BASE_URL  + '/create', reinsuranceAgreement);
    }

    getReinsuranceAgreementById(reinsuranceAgreementId){
        return axios.get(REINSURANCEAGREEMENT_API_BASE_URL + '/load?reinsuranceAgreementId=' + reinsuranceAgreementId);
    }

    updateReinsuranceAgreement(reinsuranceAgreement){
        return axios.put(REINSURANCEAGREEMENT_API_BASE_URL + '/update', reinsuranceAgreement);
    }

    deleteReinsuranceAgreement(reinsuranceAgreementId){
        return axios.delete(REINSURANCEAGREEMENT_API_BASE_URL + '/delete?reinsuranceAgreementId=' + reinsuranceAgreementId);
    }
}

export default new ReinsuranceAgreementService()