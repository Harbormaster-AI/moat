import axios from 'axios';

const POLICYACKNOWLEDGEMENT_API_BASE_URL = "/PolicyAcknowledgement";

class PolicyAcknowledgementService {

    getPolicyAcknowledgements(){
        return axios.get(POLICYACKNOWLEDGEMENT_API_BASE_URL + '/' );
    }

    createPolicyAcknowledgement(policyAcknowledgement){
        return axios.post(POLICYACKNOWLEDGEMENT_API_BASE_URL  + '/create', policyAcknowledgement);
    }

    getPolicyAcknowledgementById(policyAcknowledgementId){
        return axios.get(POLICYACKNOWLEDGEMENT_API_BASE_URL + '/load?policyAcknowledgementId=' + policyAcknowledgementId);
    }

    updatePolicyAcknowledgement(policyAcknowledgement){
        return axios.put(POLICYACKNOWLEDGEMENT_API_BASE_URL + '/update', policyAcknowledgement);
    }

    deletePolicyAcknowledgement(policyAcknowledgementId){
        return axios.delete(POLICYACKNOWLEDGEMENT_API_BASE_URL + '/delete?policyAcknowledgementId=' + policyAcknowledgementId);
    }
}

export default new PolicyAcknowledgementService()