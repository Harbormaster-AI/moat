import axios from 'axios';

const EXPIRATIONPOLICY_API_BASE_URL = "/ExpirationPolicy";

class ExpirationPolicyService {

    getExpirationPolicys(){
        return axios.get(EXPIRATIONPOLICY_API_BASE_URL + '/' );
    }

    createExpirationPolicy(expirationPolicy){
        return axios.post(EXPIRATIONPOLICY_API_BASE_URL  + '/create', expirationPolicy);
    }

    getExpirationPolicyById(expirationPolicyId){
        return axios.get(EXPIRATIONPOLICY_API_BASE_URL + '/load?expirationPolicyId=' + expirationPolicyId);
    }

    updateExpirationPolicy(expirationPolicy){
        return axios.put(EXPIRATIONPOLICY_API_BASE_URL + '/update', expirationPolicy);
    }

    deleteExpirationPolicy(expirationPolicyId){
        return axios.delete(EXPIRATIONPOLICY_API_BASE_URL + '/delete?expirationPolicyId=' + expirationPolicyId);
    }
}

export default new ExpirationPolicyService()